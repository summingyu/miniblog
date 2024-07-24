package errno

var (
	// OK 表示请求成功
	OK = &Errno{HTTP: 200, Code: "", Message: ""}

	// InternalServerError 表示所有未知的服务器端错误
	InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal Server Error."}

	// ErrPageNotFound 表示请求的页面不存在
	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page Not Found."}
)
