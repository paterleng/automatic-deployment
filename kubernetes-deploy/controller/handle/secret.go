package handle

import (
	"context"
	"encoding/base64"
	"fmt"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetes-deploy/model"
	"kubernetes-deploy/utils"
)

func CreateSecret(client *kubernetes.Clientset, nameSpace string, item model.Secret) error {
	//registry是镜像仓库地址，不是镜像地址
	// 4. 创建 base64(username:password) 格式的认证信息
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", item.Account, item.PassWord)))
	// 5. 创建 Docker config JSON 内容
	dockerConfigJSON := fmt.Sprintf(`{
      "auths": {
        "%s": {
          "auth": "%s"
        }
      }
    }`, item.Registry, auth)
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      item.Name,
			Namespace: nameSpace,
		},
		Type: v1.SecretTypeDockerConfigJson,
		Data: map[string][]byte{
			".dockerconfigjson": []byte(dockerConfigJSON),
			"username":          []byte(item.Account),
			"password":          []byte(item.PassWord),
		},
	}
	_, err := client.CoreV1().Secrets(nameSpace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("创建secret%s失败", item.Name), zap.Error(err))
		return err
	}
	return nil
}
