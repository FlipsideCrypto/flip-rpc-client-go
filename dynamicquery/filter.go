package dynamicquery

// Filter --
type Filter struct {
	In           In           `json:"in"`
	NotIn        NotIn        `json:"not_in"`
	InSegment    InSegment    `json:"in_segment"`
	NotInSegment NotInSegment `json:"not_in_segment"`
	Gte          Gte          `json:"gte"`
	Lte          Lte          `json:"lte"`
	Gt           Gt           `json:"gt"`
	Eq           Eq           `json:"eq"`
	NotEq        NotEq        `json:"not_eq"`
	Lt           Lt           `json:"lt"`
	And          []*Filter    `json:"and"`
	Or           []*Filter    `json:"or"`
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
