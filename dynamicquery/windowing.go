package dynamicquery

// WindowField returns the following in the parent select
type WindowField struct {
	Field string `json:"field"`
	Label string `json:"label"`
}

// Latest builds a window operation over the latest data from a partition
type Latest struct {
	PartitionBy []string      `json:"partition_by"`
	OrderBy     string        `json:"order_by"`
	Fields      []WindowField `json:"fields"`
}

// Earliest builds a window operation over the earliest data from a partition
type Earliest struct {
	PartitionBy []string      `json:"partition_by"`
	OrderBy     string        `json:"order_by"`
	Fields      []WindowField `json:"fields"`
}
