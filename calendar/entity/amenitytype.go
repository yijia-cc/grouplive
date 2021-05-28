package entity

type AmenityType struct {
	ID              ID
	Title           string
	Description     string
	ThumbnailURL    string
	AmenityInfoList []AmenityInfo
}
