package entity

type PermissionBinding struct {
	Permission   Permission
	User         User
	Resource     Resource
	ResourceType ResourceType
}
