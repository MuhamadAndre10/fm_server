package utils

type Err struct {
	msg string
}

func (err Err) Error() string {
	return err.msg
}
