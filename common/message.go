package common

var msgFlags = map[Code]string{
	SUCCESS:        "OK",
	ERROR:          "Fail",
	INVALID_PARAMS: "Request Parameters Error",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token authentication failed",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token has timed out",
	ERROR_AUTH_GEN_TOKEN_FAIL:      "Token generation failed",
}

func GetMsg(code Code) string {
	if msg, ok := msgFlags[code]; ok {
		return msg
	}

	return msgFlags[ERROR]
}
