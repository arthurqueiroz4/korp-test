package exception

type ErrorBase struct {
	Message string
	Status  int
}

func (e ErrorBase) Error() string {
	return e.Message
}

type (
	ErrBadRequest interface {
		Error() string
	}
	ErrNotFound interface {
		Error() string
	}
	ErrInternalServer interface {
		Error() string
	}
)

func NewErrBadRequest(body any, message string) ErrBadRequest {
	return &ErrorBase{Message: message, Status: 400}
}

func NewErrNotFound(body any, message string) ErrNotFound {
	return &ErrorBase{Message: message, Status: 404}
}

func NewErrInternalServer(body any, message string) ErrInternalServer {
	return &ErrorBase{Message: message, Status: 500}
}
