package handle

import (
	"code-package/pkg/github"
	"code-package/rpc"
	"code-package/utils"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type CodePackage struct{}

func Register(service micro.Service) error {
	err := rpc.RegisterCodePackageHandler(service.Server(), &CodePackage{})
	return err
}

func (h *CodePackage) CheckStatus(ctx context.Context, req *rpc.CpRequest, rsp *rpc.CpResponse) error {
	return nil
}

// CloneCodes 获取代码 使用git clone获取
func (h *CodePackage) CloneCodes(ctx context.Context, req *rpc.CloneCodesRequest, rsp *rpc.CloneCodesResponse) error {

	// 构造 git clone 命令
	cmd := exec.Command("git", "clone", req.GetUrl())

	// 设置标准输出和标准错误到系统标准输出中
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行 git clone 命令
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing git clone: %v\n", err)
		return nil
	}

	fmt.Println("Repository cloned successfully!")
	return nil
}

// GoGitCode 获取代码 打包镜像，推送到私有镜像仓库
func (h *CodePackage) GoGitCode(ctx context.Context, req *rpc.GoGitCodeRequest, rsp *rpc.GoGitCodeResponse) error {

	go func() {
		// 修改任务状态为：拉代码

		// 拉代码
		err := github.CloneCode(req.GetUrl(), utils.Conf.ProjectConfig.Dir)
		if err != nil {
			utils.Tools.LG.Error("Error while cloning repository:" + err.Error())
			return
		}
		// 修改任务状态为：
	}()

	// 生成镜像
	//imageName := "test01"
	//cmd := exec.Command("docker", "build", "-t", imageName, destinationPath)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err = cmd.Run()
	//if err != nil {
	//	utils.Tools.LG.Error("Error building Docker image: " + err.Error())
	//	return nil
	//}
	//utils.Tools.LG.Info("Docker image built successfully")

	// 标记镜像
	//cmd = exec.Command("docker", "tag", imageName, privateRepo+"/"+imageName)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err = cmd.Run()
	//if err != nil {
	//	utils.Tools.LG.Error("Error tagging Docker image:" + err.Error())
	//	return nil
	//}
	//utils.Tools.LG.Info("Docker image tagged successfully")

	// 推送
	//cmd = exec.Command("docker", "push", privateRepo+"/"+imageName)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err = cmd.Run()
	//if err != nil {
	//	utils.Tools.LG.Error("Error pushing Docker image:" + err.Error())
	//	return nil
	//}
	//utils.Tools.LG.Info("Docker image pushed to private registry successfully")
	return nil

}

// 获取用户生成的dockerfile文件
func (h *CodePackage) GetDockerFile() {

}

// 用户一键配置
func (h *CodePackage) ConfigureCI(ctx context.Context, req *rpc.ConfigureCIRequest, rsp *rpc.ConfigureCIResponse) error {

	if len(req.Key) != len(req.Value) {
		utils.Tools.LG.Error("invalid request parameter")
		return fmt.Errorf("invalid request parameter")
	}
	for i, _ := range req.Key {
		err := github.UpdateRepoSecret(req.Key[i], req.Value[i])
		if err != nil {
			return err
		}
	}
	data, err := ioutil.ReadFile("ci.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将内容转换为字符串
	fileContent := string(data)

	// 输出文件内容
	strings.ReplaceAll(fileContent, "version", req.Version)
	// 添加yml文件
	github.UpYml(req.Repository, "./github/workflows", req.CommitMessage, fileContent)
	return err
}
