// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func JsonPrettyPrint(in []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, in, "", "  ")
	if err != nil {
		return in
	}
	return out.Bytes()
}

// https://github.com/pkg/errors/issues/102.
func ErrorStringWithStack(err error) string {
	if err == nil {
		return ""
	}
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	cause := errors.Cause(err)
	if stackTrace, ok := cause.(stackTracer); ok {
		buf := bytes.Buffer{}
		for _, frame := range stackTrace.StackTrace() {
			buf.WriteString(fmt.Sprintf("\n%+v", frame))
		}
		return err.Error() + buf.String()
	}
	return err.Error()
}
