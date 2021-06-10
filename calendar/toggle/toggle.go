package toggle

import "github.com/yijia-cc/grouplive/calendar/entity"

type FeatureToggle interface {
	IsFeatureEnable(featureID string, user *entity.User, properties map[string]string)
}

var _ FeatureToggle = (*InternalToggle)(nil)

type InternalToggle struct {
}

func (i InternalToggle) IsFeatureEnable(featureID string, user *entity.User, properties map[string]string) {
}
