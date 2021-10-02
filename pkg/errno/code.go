package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrDecode            = &Errno{Code: 20105, Message: "Error occurred while decoding the error code."}

	// checkpoint3 errors
	ErrWrongMethod = &Errno{Code: 20301,Message: "你因为使用\"错误的方法\"写入病毒，被银行安全系统识别。很遗憾，你被逮捕了。（请尝试使用\"其他方法\"重新访问)"}
	ErrMatch = &Errno{Code: 20305, Message: "the decoded error code is not match."}
)
