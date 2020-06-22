package dynamicquery

import "github.com/FlipsideCrypto/flip-rpc-client-go/segment"

// Query user inputs to generate a SQL query
type Query struct {
	Table      string                       `json:"table"`
	Schema     string                       `json:"schema"`
	Aggregates []Aggregate                  `json:"aggregates"`
	GroupBy    []GroupBy                    `json:"group_by"`
	Filter     *Filter                      `json:"filter,omitempty"`
	OrderBy    []OrderBy                    `json:"order_by"`
	Limit      int                          `json:"limit"`
	Latest     *Latest                      `json:"latest,omitempty"`
	Earliest   *Earliest                    `json:"earliest,omitempty"`
	Segments   map[string]segment.Condition `json:"segments"`
}
