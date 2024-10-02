package opt

//type OptionFunc func(request *rpc.CreateResourceRequest)
//
//func NewParame(option ...OptionFunc) *rpc.CreateResourceRequest {
//	o := &rpc.CreateResourceRequest{
//		ResourceType: utils.DeploymentResource,
//		NameSpace:    "test",
//		Name:         "test",
//		ImageName:    "test",
//		Labels:       make(map[string]string),
//		MatchLabels:  make(map[string]string),
//		Replicas:     1,
//	}
//	for _, v := range option {
//		v(o)
//	}
//	return o
//}
//
//func Name(n string) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.Name = n
//	}
//}
//
//func ResourceType(r string) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.ResourceType = r
//	}
//}
//
//func NameSpace(n string) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.NameSpace = n
//	}
//}
//
//func ImageName(i string) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.ImageName = i
//	}
//}
//func Replicas(r int32) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.Replicas = r
//	}
//}
//
//func Labels(l map[string]string) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.Labels = l
//	}
//}
//
//func MatchLabels(m map[string]string) OptionFunc {
//	return func(request *rpc.CreateResourceRequest) {
//		request.MatchLabels = m
//	}
//}
