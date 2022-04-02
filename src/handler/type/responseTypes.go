package _type

type QuickResponseType struct {
	Result bool `json:"result"`
}

type IdResponseType struct {
	Id int `json:"id"`
}

type ResponseType struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func NewResponseType(code int, message interface{}) *ResponseType {
	return &ResponseType{
		Code:    code,
		Message: message,
	}
}
