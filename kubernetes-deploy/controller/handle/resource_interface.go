package handle

type CommentResource interface {
	Before() error
	CreateResources() error
	After() error
}
