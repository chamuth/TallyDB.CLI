package core

type Request struct {
	RequestId   string          `json:"requestId"`
	Credentials CredentialsType `json:"credentials"`
}

type CredentialsType struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
