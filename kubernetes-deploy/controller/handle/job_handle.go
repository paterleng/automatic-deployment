package handle

import (
	"context"
	"fmt"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/utils"
)

type JobHandle struct {
	client *kubernetes.Clientset
}

var jobManager JobManager

type JobInterface struct {
	JobHandle
}

type JobManager interface {
	CommentResource
}

func GetJobManager() JobManager {
	return jobManager
}

func CreateJobManager() error {
	var manager JobHandle
	client, err := utils.NewKubeConfig()
	if err != nil {
		return err
	}
	manager.client = client
	jobManager = &manager
	return nil
}

func (d *JobHandle) Before() error {
	return nil
}

func (d *JobHandle) CreateResources(r interface{}) error {
	// 定义 Job
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-job",
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "example-container",
							Image: "busybox",
							Command: []string{
								"sh",
								"-c",
								"echo Hello, Kubernetes! && sleep 30",
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyOnFailure,
				},
			},
		},
	}

	// 创建 Job
	result, err := d.client.BatchV1().Jobs("default").Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created Job %q.\n", result.GetObjectMeta().GetName())

	return nil
}

func (d *JobHandle) After() error {
	return nil
}
