package common

type Code int32

const (
	SUCCESS        Code = 200
	ERROR               = 500
	INVALID_PARAMS      = 400

	ERROR_AUTHT_CHECK_TOKEN_FAIL Code = iota + 20001
	ERROR_AUTHT_CHECK_TOKEN_TIMEOUT
	ERROR_AUTHT_GEN_TOKEN_FAIL
	ERROR_AUTHR_CHECK_PERMISSION_FAIL
)

var codeFlags = map[Code]string{
	SUCCESS:        "OK",
	ERROR:          "Fail",
	INVALID_PARAMS: "Request Parameters Error",

	ERROR_AUTHT_CHECK_TOKEN_FAIL:      "Authentication has been denied for this request.",
	ERROR_AUTHT_CHECK_TOKEN_TIMEOUT:   "Authentication Token has timed out",
	ERROR_AUTHT_GEN_TOKEN_FAIL:        "Authentication Token generation failed",
	ERROR_AUTHR_CHECK_PERMISSION_FAIL: "Authorize permission fail",
}

func GetCodeMessage(code Code) string {
	if msg, ok := codeFlags[code]; ok {
		return msg
	}

	return codeFlags[ERROR]
}
