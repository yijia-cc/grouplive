package obs

import "context"

type Analytics interface {
	Track(ctx context.Context, eventName string, properties map[string]string)
}
