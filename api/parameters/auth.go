package parameters

//TokenAuthentication representa o token
type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}
