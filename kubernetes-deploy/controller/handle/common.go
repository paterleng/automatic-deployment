package handle

import (
	"context"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
