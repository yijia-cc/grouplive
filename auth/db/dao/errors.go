package dao

var _ error = (*NotFound)(nil)

type NotFound struct{}

func (n NotFound) Error() string {
	return "not found"
}
