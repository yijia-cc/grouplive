package entity

type PermissionBinding struct {
	Permission   Permission
	User         User
	ResourceType ResourceType
	Resource     Resource
}
