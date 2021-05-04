package pb

func (e *ErrorAlreadyExists) Error() string {
	return e.GetMsg()
}

func (e *Error) Error() string {
	return e.GetMsg()
}

func (e *ErrorNotFound) Error() string {
	return e.GetMsg()
}

func (e *ErrorBadRequest) Error() string {
	return e.GetMsg()
}

func (e *ErrorUnauthorized) Error() string {
	return e.GetMsg()
}
func (e *ErrorAccessDenied) Error() string {
	return e.GetMsg()
}
