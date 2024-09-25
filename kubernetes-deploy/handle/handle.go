package handle

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"kubernetes-deploy/rpc"
)

type KubernetesDeploy struct{}

func Register(service micro.Service) error {
	err := rpc.RegisterKubernetesDeployHandler(service.Server(), &KubernetesDeploy{})
	return err
}

func (h *KubernetesDeploy) GetKubernetesConfig(ctx context.Context, req *rpc.ConfigRequest, resp *rpc.ConfigResponse) (err error) {
	config := fmt.Sprintf(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    api-version: v1
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM0VENDQWNtZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFTTVJBd0RnWURWUVFERXdkcmRXSmwKTFdOaE1CNFhEVEkwTURreE9EQXlNekV4TWxvWERUTTBNRGt4TmpBeU16RXhNbG93RWpFUU1BNEdBMVVFQXhNSAphM1ZpWlMxallUQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUpyUUk3U05qWUJPCmMvWkVOZWlaRUlrSm04NE1QeGV1dDJLcSs5ZjJGNnFWeTRKeFlRdjV2WjY5N3Zqc2FIQkpHc0QzRDNxcEx0eVgKalgxN0VZbHJxUmtZN3JjNTlPYlBPMEhsVm0yTElHcjduV2hSRmR6cExlS1k0ME9UVmFUallMTWU3Qkk4eEtjbAo4c0RuYkN0RWFqMjRUc2ZOVnZoVERMMFlXZlRwNkZBQ2xxWkkvcUNFRkU0NXBPUU1zZWJ0V25QMU9IcUJ3WDBJCmxEWGFGQXJkSEtTMnpydFk2VGNzampnNEZpcnFHOXZaWkRPYXZTZXhaQW1sTGtTYVM1Q3ZmcDRxV2pMUVlVVnoKQUlaRmhHUEh5Qkl6SWUrZEZEdHMxOTNmRGdBZFhsSGpkTmI3OFR1NzdpWjU3Wko1dnhHYkVVOThGZDY4UHFuMgpnSXUvOXVsTXNZTUNBd0VBQWFOQ01FQXdEZ1lEVlIwUEFRSC9CQVFEQWdLa01BOEdBMVVkRXdFQi93UUZNQU1CCkFmOHdIUVlEVlIwT0JCWUVGQTg5RU9VMFo2WFBDbDVuN3VnNXFQU1FvMUFqTUEwR0NTcUdTSWIzRFFFQkN3VUEKQTRJQkFRQXU3UVVxKytYRkw0T2lGbE1yK1E2a2w2SmN6MUlWSU9zQmd1M2NlOGNoL1JuTHQ3RlZ0Y2hHV1JhWAp5MGluTjR6VE1aRTdZWjFkcXQ0WnhyUGJ3MVpnNWpJaE5keG1DYmFxZ3l0ZXMzTnJscmYvLzhxSkFVekZqam1mCldjaTh4RFNOaE1xWk40ZDc2amxkYXZiZmVrazRGTFZJWmZyb0tLamc5dTVnbXVzTjE0QXhZeWJZVHZYRlZFbjMKQmVsTFRkRURRenRGblVQZy8rUW1JclM1ZEYzTGVlUVhVQ2ozRkpQYlJaYmZjYVM1RWs1RGYvQjhjYjZwYzM5MwpBVEpUOGEwZkFPMjJjZmJWU1g0bVlXT0tuaDhpcTlXM0hOS1FCa0c1ZmFCcU05aXVqemc4WFl1RnpuM2lUaDVxCkduM2RlY0toaVJQczBiYXluQnhKWjdpTVVtMVUKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: "https://192.168.10.8:6443"
  name: "local"
contexts:
- context:
    cluster: "local"
    user: "kube-admin-local"
  name: "local"
current-context: "local"
users:
- name: "kube-admin-local"
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDakNDQWZLZ0F3SUJBZ0lJTmE2enJxNW5FSVF3RFFZSktvWklodmNOQVFFTEJRQXdFakVRTUE0R0ExVUUKQXhNSGEzVmlaUzFqWVRBZUZ3MHlOREE1TVRnd01qTXhNVEphRncwek5EQTVNVFl3TWpNeE1USmFNQzR4RnpBVgpCZ05WQkFvVERuTjVjM1JsYlRwdFlYTjBaWEp6TVJNd0VRWURWUVFERXdwcmRXSmxMV0ZrYldsdU1JSUJJakFOCkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQS9MSHBpbnMwbWt5cVRGa0ozS2pyY2xPU1laaFEKUEs0SkZpU2s3UmoxUUZDTXJxT0pVeW1aMkVWL3Vsb2tXWkFNV1BMYjlXRDlvNUpjMklkUWpOMTZraTF0WkdpVwpFbUVHUno2OVBiRktFKy8yb3Njc3pLa2dFQ1F4SlZwcmJzNW1Ib045aW5hL2NFNklKb1kwVWlrcFhjSXQ0Q1NyCnEvWW1wTjVaT094czI0ZWxiZDB6TVMvbDNKMDZyd0JNcUNrU2VRYkNBZTlTRWI0bjRjNHJvazlIUkNzd2JtV3YKZGxHMFZTNUdQdkxiT1Z4V0FEaFhPdm9FdU1abUwxYXhySTZkUHBONTBoUkJvbVdMUWtQYTlydGxDNVB4VWNFcgp1Rk5JWmNxM1Y4SG5CeUhCR0Y1ZlhIMTcyTjkyMm8rN0VZSkJ0dm13bXd1OXV5eTlzZlJYQjNobVF3SURBUUFCCm8wZ3dSakFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUhBd0l3SHdZRFZSMGoKQkJnd0ZvQVVEejBRNVRSbnBjOEtYbWZ1NkRtbzlKQ2pVQ013RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUc1WgovZmx6RWFnQXJQekQ1UzhlamcrQXJOaWd4NzRQUys4Y1I0NzVvaGVycDFuOThKdFRWRG5qcXdURWZrNjAvbUw4CmtRcHcyWWlXQWhzcHRKYlRPQk1jcTVSVk42K2FzVVliQ0Z1VndkN0gxM2FhZDZoYjlMOWhDUllqN0VxTkhvU3AKMlA3Z1hjdzVWS0YvbzhqcURpcklicE5YMHMzWmt6VDdrYjQ5UFhWLzRZdUlzZWJ2ZHFreEFoZWtxSWMwa0lVRgpodkNSbkQ5RHYwdFdLbnYySHd6SDVuZlhIeitLT2paVDh2L2FJWnFLYmJMV0xzUE8yMzgvU0VJRXRJdU1QcHFYCk10OE1WdGlFOU1za1NFeUZvUzg1UTJBUFgvbFdMU0FBSm1DQ2NFWFc3YVd6aXR6ZDA2STN6QTlmQ3BFcTlBK3AKRzRWcXphbEU4bjh2Z1hkMGJyUT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBL0xIcGluczBta3lxVEZrSjNLanJjbE9TWVpoUVBLNEpGaVNrN1JqMVFGQ01ycU9KClV5bVoyRVYvdWxva1daQU1XUExiOVdEOW81SmMySWRRak4xNmtpMXRaR2lXRW1FR1J6NjlQYkZLRSsvMm9zY3MKektrZ0VDUXhKVnByYnM1bUhvTjlpbmEvY0U2SUpvWTBVaWtwWGNJdDRDU3JxL1ltcE41Wk9PeHMyNGVsYmQwegpNUy9sM0owNnJ3Qk1xQ2tTZVFiQ0FlOVNFYjRuNGM0cm9rOUhSQ3N3Ym1XdmRsRzBWUzVHUHZMYk9WeFdBRGhYCk92b0V1TVptTDFheHJJNmRQcE41MGhSQm9tV0xRa1BhOXJ0bEM1UHhVY0VydUZOSVpjcTNWOEhuQnlIQkdGNWYKWEgxNzJOOTIybys3RVlKQnR2bXdtd3U5dXl5OXNmUlhCM2htUXdJREFRQUJBb0lCQVFEM09VRFpvVXlIb1hZRAptQ0VUNTVmaWxHemtXdFkxWjdld0dFM2JRYzNBREk1MjFieW14azZqenVyMThQY2JRTmt6dFBjS0RNaFozZnBFClJhN0RBblRHeXZiNWczRS9yZmtzZnU4NmlsZUZSZ0l3bzFqcktwUVo5cTYyY3M0azdIM3dZSWpValBVaTlWSXoKRHAzbHZDdWlkVFhZMDBlNXgrdHE1YmE1QUdaSUoyNCtoNE5kbU1nb0FJck1VM0NKQ2U3L2xwVVVXS2dmME84WAo0QkNoTkJUdXc1a1YvRmRyODhneFJuUHdDZHpGKytKMFJJSjhCU1hLY2lIb3RXQThJUVNwSFRHOHhHbko2eTlPCkdsdkd0UG9GbTAvbkhraU1BNUxmQUFNazZhZ2VLbEhhdkl2SzFmUEFXZWZBSjFodmJlcXZVUW1ad0lFczkyM2EKb3Jjb0h0YmhBb0dCQVA0Yzh0RENlVjViVm53dFJWS1M4NWo5YzhjcitEeXlHeVk1M003MDdnN011a1ErSFRFMQo2NGR2L3JVTlQ2VzNReXh0SWRtNklYUE5UTms3ZzlMSzBzTkhIVktueW41Zm54YTdJbzB5Q1Z6c0hEY3dMN3lLCnVSWm1GUk1FbS9lcU1hWUVOT1UxbmtEL0VZY210UU5acTFNSjNwR2tjQzYvQ2sybFRub2pRNERyQW9HQkFQNlMKUko1WmZOQU9XNEN3YzRkMzM0Tk1iRG1BSS9WNGMyVUZocENkemo0MmJFQjNoYW1xbFNEVFJiS1ArZlUwRUVQYQpmdy83clUxNFJZSldMbVRraGcwcytyeTBvL1JIK0VFbE8yQk9UNG5WVytaODFNSnVjMFdET1RrRDA4VTNxenVSCkR1SlNmeC91WkZDT1JYSk9sNXZPM0Y1MWhuVXR1SDVtRUpRMHd4b0pBb0dBYktNUG8vUWdVeGlWWnMwcjM5dDMKVFNhK0FwNW5wL3MwNUNqRW42M1N3SDBCL1A3WkdCckhNTVhPakxTK0lraXc3U05TMzNSVnJONU9SbWpOT0tjQgpaLzdWS0dzWGhPTjRiZzFlL2lJd0lvdkduRWw3ZncySGgyM3BBSkp0NDNuZGMrNUVkdUZ2WUxZclZpNVVJTFNMCkxCc0NEaDNRNEF0SWFsMkxxV2UrQ2pVQ2dZRUE0N2p5T3Zod0J5bWd2dGlaTUFpU1g4Q0FXMXQrVHpwbGQydmUKK1FUQkZyUFdXRURTeDNoNW1IaUZzM3JKcllmYU9PaCttMHlXdFdNdVFxNVhLQmVvdHFhUzBUV0NLS1lzdEJIYgpNSVk4N2w1MnJCTGt1OXpUcnMzNDVVREFNbjVlNVpVWVRHcGhuNjIvL0xPWCtlM0YvekdudmJQQ0NKWlNvSGNxCncza0RRZ2tDZ1lFQXlSRGJrM3RGVzRNWVI1R3dTSXcyNm9RbkFMZWtKUG04MkZNVEJKdzJCOFowcURYUHBlTWMKbm82MDZiZzF1dXFHeTRDK1hpQUNEUXY4emxBckxGZnFhRTJJdVBUTEZ2Sy9GbFV4aGFWT1p1a2F0M2oxYXVGUgphWmJtMzd3TGNZVWJoa3Y5cE9nNExIWmJGTkprVmdMOGhQaFpka0ZhVG1FRVRyM2s3U2FlWGRRPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
`)
	resp.Config = config
	return nil

}

func (h *KubernetesDeploy) CheckStatus(ctx context.Context, req *rpc.KsRequest, rsp *rpc.KsResponse) error {
	return nil
}
