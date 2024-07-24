package errno

var ErrUserAlreayExist = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "user already exist."}
