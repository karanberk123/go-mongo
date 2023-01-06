package exception

type ValidationErr struct {
	Message string
}

func (validateErr ValidationErr) Error() string {
	return validateErr.Message
}
