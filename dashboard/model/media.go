package model

type MediaType struct {
	MediaTypeId   int    `json:"media_type_id"`
	MediaTypeName string `json:"media_type_name"`
}

type Media struct {
	MediaId     int    `json:"media_id"`
	EventId     int    `json:"event_id"`
	MediaTypeId int    `json:"media_type_id"`
	MediaURL    string `json:"media_url"`
}
