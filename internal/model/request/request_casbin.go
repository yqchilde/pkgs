package request

type CasbinInReceive struct {
	AuthorityID string       `json:"authority_id"`
	CasbinInfos []CasbinInfo `json:"casbin_infos"`
}

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}
