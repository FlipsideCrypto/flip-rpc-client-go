package dynamicquery

// Filter --
type Filter struct {
	In           In           `json:"in,omitempty"`
	NotIn        NotIn        `json:"not_in,omitempty"`
	InSegment    InSegment    `json:"in_segment,omitempty"`
	NotInSegment NotInSegment `json:"not_in_segment,omitempty"`
	Gte          Gte          `json:"gte,omitempty"`
	Lte          Lte          `json:"lte,omitempty"`
	Gt           Gt           `json:"gt,omitempty"`
	Eq           Eq           `json:"eq,omitempty"`
	NotEq        NotEq        `json:"not_eq,omitempty"`
	Lt           Lt           `json:"lt,omitempty"`
	And          []*Filter    `json:"and,omitempty"`
	Or           []*Filter    `json:"or,omitempty"`
}

// Eq equal to filter
type Eq struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// NotEq not equal to filter
type NotEq struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Gte greater than or equal to filter
type Gte struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lte less than or equal to filter
type Lte struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Gt greater than filter
type Gt struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Lt less than filter
type Lt struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// In in filter
type In struct {
	Field string   `json:"field"`
	Value []string `json:"value"`
}

// NotIn not in filter
type NotIn struct {
	Field string   `json:"field"`
	Value []string `json:"value"`
}

// InSegment in filter
type InSegment struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

// NotInSegment not in filter
type NotInSegment struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
