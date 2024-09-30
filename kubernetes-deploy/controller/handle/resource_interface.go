package handle

type CommentResource interface {
	Before() error
	CreateResources(interface{}) error
	After() error
}
