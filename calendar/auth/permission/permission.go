package permission

type ResourceType struct {
	ID   string
	Name string
}

type Permission struct {
	ID             string
	Name           string
	ResourceTypeID string
	ResourceID     string
}

type Factory func() Permission

const allResources = "*"

var (
	ViewAmenityTypes = func() Permission {
		return Permission{
			ID:             "p01",
			ResourceTypeID: "t01",
			ResourceID:     "r01",
		}
	}
)
