package idgen

import (
	"github.com/yijia-cc/grouplive/auth/entity"
)

type IDGenerator interface {
	NextID() entity.ID
}
