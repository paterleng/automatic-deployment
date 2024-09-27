package controller

import (
	corev1 "k8s.io/api/core/v1"
)

type CommentResource interface {
	Before() error
	Resources() error
	After() error
	ListPods() ([]corev1.Pod, error)
}
