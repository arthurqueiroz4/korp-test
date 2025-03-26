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

func NewErrBadRequest(message string) ErrBadRequest {
	return &ErrorBase{Message: message, Status: 400}
}

func NewErrNotFound(message string) ErrNotFound {
	return &ErrorBase{Message: message, Status: 404}
}

func NewErrInternalServer(message string) ErrInternalServer {
	return &ErrorBase{Message: message, Status: 500}
}
