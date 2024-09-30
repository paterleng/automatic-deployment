package model

type ResourceReq struct {
	ResourceType string            `json:"resource_type" validate:"required"`
	Name         string            `json:"name"`
	NameSpace    string            `json:"name_space"`
	ImageName    string            `json:"image_name" validate:"required"`
	Replicas     int32             `json:"replicas"`
	Labels       map[string]string `json:"labels"`
	MatchLabels  map[string]string `json:"match_labels"`
}
