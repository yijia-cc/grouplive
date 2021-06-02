package tm

import "time"

type Timer interface {
	Now() time.Time
}

var _ Timer = (*LocalTimer)(nil)

type LocalTimer struct {
}

func (l LocalTimer) Now() time.Time {
	return time.Now()
}

func NewLocalTimer() LocalTimer {
	return LocalTimer{}
}
