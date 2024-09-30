package handle

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/utils"
)

type CornJobHandle struct {
	client *kubernetes.Clientset
}

var cornJobManager CornJobManager

type CornJobInterface struct {
	JobHandle
}

type CornJobManager interface {
	CommentResource
}

func GetCornJobManager() CornJobManager {
	return jobManager
}

func CreateCornJobManager() error {
	var manager CornJobHandle
	client, err := utils.NewKubeConfig()
	if err != nil {
		return err
	}
	manager.client = client
	cornJobManager = &manager
	return nil
}

func (c *CornJobHandle) Before() error {
	return nil
}

func (c *CornJobHandle) CreateResources(r interface{}) error {
	// 定义 CronJob
	cronJob := &v1beta1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-cronjob",
			Namespace: "default", // 设置命名空间
		},
		Spec: v1beta1.CronJobSpec{
			Schedule: "*/5 * * * *", // 设置调度表达式
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:  "example-container",
									Image: "your-image:tag",         // 替换为你的镜像
									Args:  []string{"arg1", "arg2"}, // 可选参数
								},
							},
							RestartPolicy: corev1.RestartPolicyNever, // 设置重启策略
						},
					},
				},
			},
		},
	}

	// 创建 CronJob
	createdCronJob, err := c.client.BatchV1beta1().CronJobs("default").Create(context.TODO(), cronJob, metav1.CreateOptions{})
	if err != nil {
		utils.Tools.LG.Error("创建cornjob资源失败", zap.Error(err))
		return err
	}
	fmt.Printf("CronJob %s created successfully\n", createdCronJob.Name)
	return nil
}

func (c *CornJobHandle) After() error {

	return nil
}
