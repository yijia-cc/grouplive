package obs

import "github.com/yijia-cc/grouplive/calendar/entity"

type Analytics interface {
	Track(eventName string, user *entity.User, properties map[string]string)
}

