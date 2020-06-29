package dynamicquery

// Dataset is segment building block
type Dataset struct {
	ID       string `json:"id"`
	EntityID string `json:"entity_id"`
	Owner    string `json:"owner"`
}
