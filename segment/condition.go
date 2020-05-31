package segment

// Gte = greater than or equal to
type Gte struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Lte = less than or equal to
type Lte struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Lt = less than
type Lt struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Gt = greater than
type Gt struct {
	PartitionID string  `json:"partition_id"`
	Value       float64 `json:"value"`
}

// Condition is set of logic
type Condition struct {
	PartitionID string       `json:"partition_id"`
	Value       float64      `json:"value"`
	Or          []*Condition `json:"or"`
	And         []*Condition `json:"and"`
	Gt          Gt           `json:"gt"`
	Gte         Gte          `json:"gte"`
	Lt          Lt           `json:"lt"`
	Lte         Lte          `json:"lte"`
}
