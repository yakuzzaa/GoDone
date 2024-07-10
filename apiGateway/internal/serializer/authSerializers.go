package serializer

type SignInInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Info SignInInfo `json:"info"`
}

type SignInResponse struct {
	UserID uint64 `json:"id"`
}

type SignUpRequest struct {
	Info SignUpInfo `json:"info"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error error `json:"error"`
}
