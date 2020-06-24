package dynamicquery

// Field selects a field to return
type Field struct {
	Field             string `json:"field"`
	Label             string `json:"label"`
	DecimalAdjustment int    `json:"decimal_adjustment"`
	ToNegative        bool   `json:"to_negative"`
}
