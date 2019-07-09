package common

type Code int32

const (
	SUCCESS        Code = 200
	ERROR               = 500
	INVALID_PARAMS      = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_GEN_TOKEN_FAIL      = 20003
)

var codeFlags = map[Code]string{
	SUCCESS:        "OK",
	ERROR:          "Fail",
	INVALID_PARAMS: "Request Parameters Error",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token has timed out",
	ERROR_AUTH_GEN_TOKEN_FAIL:      "Token generation failed",
}

func GetCodeMessage(code Code) string {
	if msg, ok := codeFlags[code]; ok {
		return msg
	}

	return codeFlags[ERROR]
}
