package dynamicquery

// Aggregate builds aggregate operations
type Aggregate struct {
	Field             string `json:"field"`
	Label             string `json:"label"`
	Operation         string `json:"operation"`
	DecimalAdjustment int    `json:"decimal_adjustment"`
}
