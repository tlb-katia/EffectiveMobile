package lib

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Resp(status int, message string) Response {
	return Response{
		Status:  status,
		Message: message,
	}
}
