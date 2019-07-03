package common

type Status int

const (
	SUCCESS        Status = 200
	ERROR          Status = 500
	INVALID_PARAMS Status = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    Status = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT Status = 20002
	ERROR_AUTH_GEN_TOKEN_FAIL      Status = 20003
)
