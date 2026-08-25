package puregen

import "flag"

type OptionsInfo struct {
	StatshouseURL    string
	PartitioningMode int
}

func (opt *OptionsInfo) Bind(f *flag.FlagSet) {
	f.StringVar(&opt.StatshouseURL, "info-statshouse-url", "",
		`url to statshouse (required)`)
	f.IntVar(&opt.PartitioningMode, "info-partitioning-mode", 1,
		`mode how to partitioning namespaces for queries (1 - by chunks of 200 methods, 2 - by namespace)`)
}
