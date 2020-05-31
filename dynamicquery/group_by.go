package dynamicquery

// GroupBy a set of fields
type GroupBy struct {
	Field      string `json:"field"`
	Label      string `json:"label"`
	Timebucket string `json:"timebucket"`
}
