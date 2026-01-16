package pingRequest

type Start struct {
	IP     string `json:"ip" binding:"required,ip4_addr"`
	Target int    `json:"target" binding:"required"`
}
