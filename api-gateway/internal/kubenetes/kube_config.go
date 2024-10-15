package internal

import (
	"api-gateway/model"
	"encoding/base64"
	"fmt"
)

var kubeConfig KubeConfig

type KubeConfig interface {
	CreateConfig(model.ClusterResquest) (string, error)
}

type KubeConfigHandle struct {
}

func GetKubeConfig() KubeConfig {
	kubeConfig = &KubeConfigHandle{}
	return kubeConfig
}

func (k *KubeConfigHandle) CreateConfig(req model.ClusterResquest) (config string, err error) {
	//certificateAuthorityData, err := ioutil.ReadFile(utils.CertificateAuthorityData)
	//if err != nil {
	//	return "", err
	//}
	//clientCertificateData, err := ioutil.ReadFile(utils.ClientCertificateData)
	//if err != nil {
	//	return "", err
	//}
	//clientKeyData, err := ioutil.ReadFile(utils.ClientKeyData)
	//if err != nil {
	//	return "", err
	//}
	certificateAuthorityData := []byte(fmt.Sprintf(`
-----BEGIN CERTIFICATE-----
MIIC4TCCAcmgAwIBAgIBADANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDEwdrdWJl
LWNhMB4XDTI0MDkxODAyMzExMloXDTM0MDkxNjAyMzExMlowEjEQMA4GA1UEAxMH
a3ViZS1jYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAJrQI7SNjYBO
c/ZENeiZEIkJm84MPxeut2Kq+9f2F6qVy4JxYQv5vZ697vjsaHBJGsD3D3qpLtyX
jX17EYlrqRkY7rc59ObPO0HlVm2LIGr7nWhRFdzpLeKY40OTVaTjYLMe7BI8xKcl
8sDnbCtEaj24TsfNVvhTDL0YWfTp6FAClqZI/qCEFE45pOQMsebtWnP1OHqBwX0I
lDXaFArdHKS2zrtY6Tcsjjg4FirqG9vZZDOavSexZAmlLkSaS5Cvfp4qWjLQYUVz
AIZFhGPHyBIzIe+dFDts193fDgAdXlHjdNb78Tu77iZ57ZJ5vxGbEU98Fd68Pqn2
gIu/9ulMsYMCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB/wQFMAMB
Af8wHQYDVR0OBBYEFA89EOU0Z6XPCl5n7ug5qPSQo1AjMA0GCSqGSIb3DQEBCwUA
A4IBAQAu7QUq++XFL4OiFlMr+Q6kl6Jcz1IVIOsBgu3ce8ch/RnLt7FVtchGWRaX
y0inN4zTMZE7YZ1dqt4ZxrPbw1Zg5jIhNdxmCbaqgytes3Nrlrf//8qJAUzFjjmf
Wci8xDSNhMqZN4d76jldavbfekk4FLVIZfroKKjg9u5gmusN14AxYybYTvXFVEn3
BelLTdEDQztFnUPg/+QmIrS5dF3LeeQXUCj3FJPbRZbfcaS5Ek5Df/B8cb6pc393
ATJT8a0fAO22cfbVSX4mYWOKnh8iq9W3HNKQBkG5faBqM9iujzg8XYuFzn3iTh5q
Gn3decKhiRPs0baynBxJZ7iMUm1U
-----END CERTIFICATE-----
`))
	clientCertificateData := []byte(fmt.Sprintf(`
-----BEGIN CERTIFICATE-----
MIIDETCCAfmgAwIBAgIBADANBgkqhkiG9w0BAQsFADAqMSgwJgYDVQQDEx9rdWJl
LWFwaXNlcnZlci1yZXF1ZXN0aGVhZGVyLWNhMB4XDTI0MDkxODAyMzExMloXDTM0
MDkxNjAyMzExMlowKjEoMCYGA1UEAxMfa3ViZS1hcGlzZXJ2ZXItcmVxdWVzdGhl
YWRlci1jYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANWjupeaAEY0
W99DxF7xKBmLtpV4CmmTgeHdImwrxbxeNMeAnat9X+reHFcaxT3j/FLDUJ7Ixr84
WBfwEuw3qrPyyCHqo3mcIQBvAB6L8n6XYajZ+4EtRDJXIgt+zDcDnonuyhzp05wT
Nnk677m5wgl/VjZIDEwNnB10pVYdTzJJzWzRGxhr3x4SDl+VQkt1fmCXN2UCKW81
aPFfdhVk8NVuF6Z6SLDQFMU9SJuYFRmPrrhTsVhoL/Jyhu53zUOLRe3qo9mP35XC
shNGIIMfhRNeZZg2VSlN6MfjjY2ttrX+5WtSZBIBhO1Yry6WUOnL5BBANZXSBpFg
FECphiQcEOsCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB/wQFMAMB
Af8wHQYDVR0OBBYEFNmb2wcjtHYqrhK679o2KfNP7KPXMA0GCSqGSIb3DQEBCwUA
A4IBAQC11FuGR7AOwMdyOzaipQzLSmG41Sx9Bv2fquREzQuPgySIK+hzq7JaWUC+
cMDPkOWWs9NVLDimL3OvlUpqTobxI2lBIcX5JDZ8GyqkPfNQ4HvAQW+obi/k0+Qf
aVLOGsma8GI7axHF5ov3nEUP07+WM4qHOXhck9RN/kD5HsRzIosY9FKlhqYGsqu8
Lj7gZxXikr8LPYkl0ZWygGUwTppqC9NFc7cJq1u0NKBP7U0VTSYRWrzxC3BYsaNF
UqXlROzuCYgthGqsmSD8MWodQ8H8+ijFcMLotSd95s8ftuAV4xeT+oeWVnL7OUrW
f5jtQLHYjPzWT51vEPb2LCIR1aHY
-----END CERTIFICATE-----
`))
	clientKeyData := []byte(fmt.Sprintf(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA1aO6l5oARjRb30PEXvEoGYu2lXgKaZOB4d0ibCvFvF40x4Cd
q31f6t4cVxrFPeP8UsNQnsjGvzhYF/AS7Deqs/LIIeqjeZwhAG8AHovyfpdhqNn7
gS1EMlciC37MNwOeie7KHOnTnBM2eTrvubnCCX9WNkgMTA2cHXSlVh1PMknNbNEb
GGvfHhIOX5VCS3V+YJc3ZQIpbzVo8V92FWTw1W4XpnpIsNAUxT1Im5gVGY+uuFOx
WGgv8nKG7nfNQ4tF7eqj2Y/flcKyE0Yggx+FE15lmDZVKU3ox+ONja22tf7la1Jk
EgGE7VivLpZQ6cvkEEA1ldIGkWAUQKmGJBwQ6wIDAQABAoIBAQDJfydaHVHs5S1R
oqTvtxxPG8sVpqaLdSinUt5Ij/Br/Z5iHBiZyiJRbzPjo8gs2378wGKloFBzwjy0
NyZhzQ8fsfi4pImTBYLK2eaO6QfKBhWCIDZqW7taj3v4w5rwuGUrSyEtrAx9xFE4
fjHkRiOQuZ0ijUmr4ZpA3eejFPnBPsf3rjkhDAg4gjuAgP7JDAAjLcbWpsbjpjg4
R/evaB7teCvvpsQdoUmLVmHgSl195l515CskvH18E/z+Lx+Y2TjGD5xmEyyn+g/m
mvNF9z6wLUvRexgMPRedX7HFuoDvIu6JQObUzPtOc94r7+IE3XhCfJ4BaIXXLBFv
vaSjev3hAoGBAP2/+ckOTB6/vh7mBfESRTxtSL4P/RYM0jCSxvdXTTGzOtAd4cYE
RWzDy1SIdKaYORZFzGf2K/gKWeZ9ltjW/dPovES4uJGj8SHBRvGrNbRAE4vDCiJd
M1REd2rOI4axZWBNjvDaqDJnOFRLx9slvQLgEtPjkhnk11/H840XAwn/AoGBANeI
s2a3puDFIOk3zk43Z4SiKBTSfJy52nclSrdV+VZxtL1e78yw3BEZnaJGivbLrA9T
lpK4IiI8R8Uo2QMaGS0KdtbYrgKBQWddkJx+DAXGGwpYFaUw1tMGTK4EBiSDnVLF
M7hOtVCfAzIVdTm56+/ADXCMROXEXXbYb2HLrcEVAoGBAItiC3QMPXXafzV7xqHA
FnX11bJJGA0np8F0roQo2UdGmzS8ZUsfB3+SITti5SipjxoT5w7oIwjRrsY0EEs1
9SGUwu3AxemCMy8GetC+6fYECHiJ/yQXK93K6gmqB8ux3+zBTZYxlZhyKHftVG+l
UCinKtBnPdapZEDBlZ/XlxzDAoGADtlBGst+OK/8A3Uvxl20yQNu5XhW02lObrTn
/9dxdQ4iWIWI2b45ewgbvwlDG5uOgAPPNM5ws5EZlLCqurb6kwrMgyKsYknLWras
dsuMQn2ScVT+MMI7mpAtijOGxM84cHJbjNAHV8WMr8+goth3M640ftN7D6VGlyB0
E/W3Q70CgYEA2v5obIo/lwqe0eLStXwBDWZtmabT1VmOKLQKyZ1hiZy4YkbvQXIO
gZgh/Dfbu8BTwH0z3IrdAH2Y9sQNwU/KkgHrl3buIy8hxjtuS2gSJsT6COjWeQYJ
LVjisXswcjpzFXio3f6dMUHOm7HPS/aSjS/s0Yn+ORcTDO3GF2rMYgA=
-----END RSA PRIVATE KEY-----
`))
	config = fmt.Sprintf(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    api-version: v1
    certificate-authority-data: %s
    server: "%s"
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
    client-certificate-data: %s
    client-key-data: %s
`, base64.StdEncoding.EncodeToString(certificateAuthorityData), req.ClusterAdr, base64.StdEncoding.EncodeToString(clientCertificateData), base64.StdEncoding.EncodeToString(clientKeyData))
	return config, nil
}
