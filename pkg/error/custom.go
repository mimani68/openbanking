package error

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	return e
}
