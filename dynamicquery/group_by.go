package dynamicquery

// GroupBy a set of fields
type GroupBy struct {
	Field      string    `json:"field"`
	Label      string    `json:"label"`
	Timebucket string    `json:"timebucket"`
	Gapfill    bool      `json:"gapfill"`
	StartDate  string    `json:"start_date"`
	EndDate    string    `json:"end_date"`
	Datasets   []Dataset `json:"datasets"`
}
