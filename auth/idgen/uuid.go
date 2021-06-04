package idgen

import (
	"github.com/google/uuid"
	"github.com/yijia-cc/grouplive/auth/entity"
)

var _ IDGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct {
}

func (U UUIDGenerator) NextID() entity.ID {
	return entity.ID(uuid.New().String())
}

func NewUUIDGenerator() UUIDGenerator {
	return UUIDGenerator{}
}
