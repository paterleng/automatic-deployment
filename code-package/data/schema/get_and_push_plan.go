package schema

type GetAndPushPlan struct {
	Id          int64
	GetUrl      string
	PushUrl     string
	ImageName   string
	Status      int
	DownloadDir string
	Version     string
	IsSuccess   bool
}

type PlanStatus int

var (
	PlanStatus_CloneCode      PlanStatus = 1 // 拉取代码
	PlanStatus_PackagingImage PlanStatus = 2 // 打包镜像
	PlanStatus_MarkImage      PlanStatus = 3 // 标记镜像
	PlanStatus_PushImage      PlanStatus = 4 // 推送镜像

)
