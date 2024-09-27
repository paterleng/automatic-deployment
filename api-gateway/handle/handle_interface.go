package handle

import "context"

type KubernetesHandlerInterface interface {
	CreateResource(ctx context.Context) error
}
