package handle

import (
	"code-package/data/schema"
	"code-package/pkg/cmd"
	"code-package/pkg/github"
	"code-package/utils"
	"context"
)

var SC chan schema.GetAndPushPlan

// Start 任务开启线程
func (h *CodePackage) Start() {

	SC = make(chan schema.GetAndPushPlan, 10)

	for {
		select {

		case d := <-SC:

			switch d.Status {

			case int(schema.PlanStatus_CloneCode):
				err := github.CloneCode(d.GetUrl, d.DownloadDir, utils.Conf.GitHub.Auth)
				if err != nil {
					utils.Tools.LG.Error("Error while Packaging image:" + err.Error())
				} else {
					dd := d
					dd.Status = int(schema.PlanStatus_PackagingImage)
					// 数据库修改任务状态为：schema.PlanStatus_PackagingImage
					if err = h.repo.UpdatePlanStatus(context.Background(), d.Id, int64(schema.PlanStatus_PackagingImage)); err != nil {
						utils.Tools.LG.Error("Error while Packaging image:" + err.Error())
						continue
					}
					SC <- dd
				}
			case int(schema.PlanStatus_PackagingImage):
				err := cmd.PackagingImage(d.ImageName, d.GetUrl, d.DownloadDir)
				if err != nil {
					utils.Tools.LG.Error("Error while Packaging image:" + err.Error())
				} else {
					dd := d
					dd.Status = int(schema.PlanStatus_MarkImage)
					// 数据库修改任务状态为：schema.PlanStatus_MarkImage
					if err = h.repo.UpdatePlanStatus(context.Background(), d.Id, int64(schema.PlanStatus_MarkImage)); err != nil {
						utils.Tools.LG.Error("Error while Packaging image:" + err.Error())
						continue
					}
					SC <- dd
				}
			case int(schema.PlanStatus_MarkImage):
				err := cmd.MarkImage(d.ImageName, d.Version)
				if err != nil {
					utils.Tools.LG.Error("Error while Mark image:" + err.Error())
				} else {
					dd := d
					dd.Status = int(schema.PlanStatus_PushImage)
					// 数据库修改任务状态为：schema.PlanStatus_PushImage
					if err = h.repo.UpdatePlanStatus(context.Background(), d.Id, int64(schema.PlanStatus_PushImage)); err != nil {
						utils.Tools.LG.Error("Error while Packaging image:" + err.Error())
						continue
					}
					SC <- dd
				}
			case int(schema.PlanStatus_PushImage):
				err := cmd.PushImage(d.ImageName, d.Version)
				if err != nil {
					utils.Tools.LG.Error("Error while Push image:" + err.Error())
				}
			}

		}
	}
}
