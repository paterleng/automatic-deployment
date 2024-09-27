package handle

import corev1 "k8s.io/api/core/v1"

type DeployHandle struct{}

func (d *DeployHandle) Before() error {

	return nil
}

func (d *DeployHandle) Resources() error {
	return nil
}

func (d *DeployHandle) After() error {
	return nil
}

func (d *DeployHandle) ListPods() ([]corev1.Pod, error) {
	return nil, nil

}
