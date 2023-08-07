package presenter

type ErrorReponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"msg,omitempty"`
}

func NewErrorReponse(code int, msg string) ErrorReponse {
	return ErrorReponse{
		StatusCode: code,
		Message:    msg,
	}
}
