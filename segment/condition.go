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
	PartitionID *string      `json:"partition_id,omitempty"`
	Value       *float64     `json:"value,omitempty"`
	Or          []*Condition `json:"or,omitempty"`
	And         []*Condition `json:"and,omitempty"`
	Gt          *Gt          `json:"gt,omitempty"`
	Gte         *Gte         `json:"gte,omitempty"`
	Lt          *Lt          `json:"lt,omitempty"`
	Lte         *Lte         `json:"lte,omitempty"`
}
