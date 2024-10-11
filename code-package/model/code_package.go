package model

type GetAndPushPlan struct {
	GetUrl    string
	PushUrl   string
	ImageName string
	Status    int8
	Context   string
}
