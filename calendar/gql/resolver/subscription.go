package resolver

type subscription struct{}

func (subscription) ScheduleUpdateSubscribe(args struct {
	Input ScheduleUpdateSubscribeInput
}) *Void {
	return nil
}
