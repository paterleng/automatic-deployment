package handle

import (
	"context"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/utils"
)

func CheckNameSpace(client *kubernetes.Clientset, nameSpace string) error {
	_, err := client.CoreV1().Namespaces().Get(context.TODO(), nameSpace, metav1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			namespace := &v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: nameSpace,
				},
			}
			_, err = client.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func GetSecret(client *kubernetes.Clientset, nameSpace string) (*v1.SecretList, error) {
	secret, err := client.CoreV1().Secrets(nameSpace).List(context.TODO(), metav1.ListOptions{})
	return secret, err
}

func imagePullSecret(client *kubernetes.Clientset, nameSpace string) []v1.LocalObjectReference {
	secretList, err := GetSecret(client, nameSpace)
	if err != nil {
		utils.Tools.LG.Error("获取secret失败", zap.Error(err))
		return nil
	}
	references := make([]corev1.LocalObjectReference, 0)
	for _, item := range secretList.Items {
		references = append(references, corev1.LocalObjectReference{
			Name: item.Name,
		})
	}
	return references
}
