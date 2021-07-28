package model

type Casbin struct {
	ID          uint   `json:"id"`
	PType       string `json:"p_type"`
	AuthorityId string `json:"authority_id"`
	Path        string `json:"path"`
	Method      string `json:"method"`
}
