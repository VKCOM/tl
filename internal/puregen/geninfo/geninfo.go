// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package geninfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/puregen"
	"github.com/VKCOM/tl/internal/tlast"
	"github.com/VKCOM/tl/internal/utils"
)

const (
	statshouseViewBase  = "https://statshouse.mvk.com/view"
	statshouseQueryBase = "https://statshouse.mvk.com/api/query"

	cookieEnvVar        = "STATSHOUSE_VKUTH_DATA"
	cookieKey           = "vkuth_data"
	cookieSessionEnvVar = "STATSHOUSE_VKUTH_SESSION"
	cookieSessionKey    = "vkuth_session"

	sectionLengthInSeconds = "60"
)

func statshouseViewLink(queries []tlast.Name) string {
	v := url.Values{}
	v.Set("t", "0")
	v.Set("f", "-7200")
	v.Set("s", "vkcom_rpc_get_time")
	v.Set("g", sectionLengthInSeconds)
	v.Add("qb", "15")
	v.Add("qb", "16")
	v.Set("qf", "15-kphp")
	v.Set("dg", "0")
	v.Set("dl", "0.0.12.12")
	for _, entry := range queries {
		v.Add("qf", "11-"+entry.String())
	}
	return statshouseViewBase + "?" + v.Encode()
}

func statshouseQueryLink(queries []tlast.Name) string {
	v := url.Values{}
	v.Set("qw", "count_norm")
	v.Set("t", "0")
	v.Set("f", "-7200")
	v.Set("w", sectionLengthInSeconds+"s")
	v.Set("v", "2")
	v.Add("qb", "15")
	v.Add("qb", "16")
	v.Set("ep", "1")
	v.Set("qv", "0")
	v.Set("s", "vkcom_rpc_get_time")
	v.Set("n", "5")
	v.Set("priority", "1")
	for _, entry := range queries {
		v.Add("qf", "11-"+entry.String())
	}
	return statshouseQueryBase + "?" + v.Encode()
}

type reportPart struct {
	Index      int    `json:"index"`
	QueryCount int    `json:"query_count"`
	ViewURL    string `json:"view_url"`
	QueryURL   string `json:"query_url"`
}

type report struct {
	Table  string             `json:"table"`
	Values map[string]float64 `json:"values"`
	Parts  []reportPart       `json:"parts"`
}

type queryResponse struct {
	Data struct {
		Series struct {
			Time       []int64 `json:"time"`
			SeriesMeta []struct {
				Tags map[string]struct {
					Value string `json:"value"`
				} `json:"tags"`
			} `json:"series_meta"`
			SeriesData [][]*float64 `json:"series_data"`
		} `json:"series"`
	} `json:"data"`
}

func fetchCountNorm(link string, dataCookie, sessionCookie string) (map[string]float64, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	cookieHeader := cookieKey + "=" + dataCookie
	if sessionCookie != "" {
		cookieHeader += "; " + cookieSessionKey + "=" + sessionCookie
	}
	req.Header.Set("Cookie", cookieHeader)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var result queryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	seriesMap := make(map[string]float64)
	for i, meta := range result.Data.Series.SeriesMeta {
		if i >= len(result.Data.Series.SeriesData) {
			break
		}
		// строим ключ из тегов (отсортированные ключи для стабильности)
		tagKeys := make([]string, 0, len(meta.Tags))
		for k := range meta.Tags {
			tagKeys = append(tagKeys, k)
		}
		sort.Strings(tagKeys)

		var keyParts []string
		for _, k := range tagKeys {
			val := meta.Tags[k].Value
			if val == " 0" {
				val = "?"
			}
			keyParts = append(keyParts, val)
		}
		key := ""
		if len(keyParts) > 0 {
			key = keyParts[0]
			for i := 1; i < len(keyParts); i++ {
				key += "," + keyParts[i]
			}
		}

		var sum float64
		var count int
		for _, v := range result.Data.Series.SeriesData[i] {
			if v != nil {
				sum += *v
				count++
			}
		}
		if count > 0 {
			seriesMap[key] = sum / float64(count)
		} else {
			seriesMap[key] = 0
		}
	}
	return seriesMap, nil
}

func Generate(kernel *pure.Kernel, options *puregen.Options) error {
	namespaces := map[string][]tlast.Name{}
	for _, comb := range kernel.TL1() {
		if comb.IsFunction {
			namespaces[comb.Construct.Name.Namespace] = append(namespaces[comb.Construct.Name.Namespace], comb.Construct.Name)
		}
	}
	ns := utils.Keys(namespaces)
	sort.Strings(ns)
	sort.Slice(ns, func(i, j int) bool {
		return len(namespaces[ns[i]])-len(namespaces[ns[j]]) > 0
	})

	sum := 0
	for _, n := range ns {
		//fmt.Println(n, len(namespaces[n]))
		sum += len(namespaces[n])
	}
	//fmt.Println(sum)

	const PartSize = 200

	parts := make([][]tlast.Name, 1)
	for _, n := range ns {
		i := len(parts) - 1
		parts[i] = append(parts[i], namespaces[n]...)
		if len(parts[i]) >= PartSize {
			parts = append(parts, []tlast.Name{})
		}
	}

	cookieValue := os.Getenv(cookieEnvVar)
	sessionCookie := os.Getenv(cookieSessionEnvVar)
	if cookieValue == "" {
		fmt.Fprintf(os.Stderr, "Warning: %s not set, printing view links only\n", cookieEnvVar)
		for _, part := range parts {
			fmt.Println(statshouseViewLink(part))
		}
		return nil
	}

	totalSeries := make(map[string]float64)
	var reportParts []reportPart
	for i, part := range parts {
		queryLink := statshouseQueryLink(part)
		viewLink := statshouseViewLink(part)
		reportParts = append(reportParts, reportPart{
			Index:      i,
			QueryCount: len(part),
			ViewURL:    viewLink,
			QueryURL:   queryLink,
		})

		seriesMap, err := fetchCountNorm(queryLink, cookieValue, sessionCookie)
		if err != nil {
			return fmt.Errorf("failed to fetch part %d (%d queries): %w", i, len(part), err)
		}
		fmt.Printf("Part %d (%d queries): series = %d\n", i, len(part), len(seriesMap))
		for k, v := range seriesMap {
			totalSeries[k] += v
		}
	}

	type item struct {
		key   string
		value float64
	}
	var items []item
	maxKeyLen := 3 // "Key"
	for k, v := range totalSeries {
		items = append(items, item{key: k, value: v})
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	totalSum := 0.0
	for _, v := range totalSeries {
		totalSum += v
	}

	var tableBuilder strings.Builder
	fmt.Fprintf(&tableBuilder, "%-*s  %10s  %10s\n", maxKeyLen, "Key", "Avg", "Percent")
	tableBuilder.WriteString(strings.Repeat("-", maxKeyLen+2+10+2+10) + "\n")
	for _, it := range items {
		percent := 0.0
		if totalSum > 0 {
			percent = it.value / totalSum * 100
		}
		fmt.Fprintf(&tableBuilder, "%-*s  %10.2f  %9.2f%%\n", maxKeyLen, it.key, it.value, percent)
	}
	tableStr := tableBuilder.String()
	fmt.Print(tableStr)

	rep := report{
		Table:  tableStr,
		Values: totalSeries,
		Parts:  reportParts,
	}

	outPath := filepath.Join(options.Outdir, "info.json")
	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create report file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(rep); err != nil {
		return fmt.Errorf("failed to encode report: %w", err)
	}

	return nil
}
