package handle

import (
	"code-package/pkg/github"
	"code-package/rpc"
	"code-package/utils"
	"context"
	"fmt"
	"github.com/go-git/go-git/v5"
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
func (h *CodePackage) CloneCodes(ctx context.Context, req *rpc.CpRequest, rsp *rpc.CpResponse) error {
	// Git 仓库地址和目标目录
	repoURL := "https://github.com/paterleng/evaluation-of-teaching.git"
	//targetDir := ""

	// 构造 git clone 命令
	cmd := exec.Command("git", "clone", repoURL)

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

// GoGitCode 获取代码 使用gogit库
func (h *CodePackage) GoGitCode(ctx context.Context, req *rpc.CpRequest, rsp *rpc.CpResponse) error {
	// 要克隆的远程仓库地址
	repoURL := "https://github.com/paterleng/evaluation-of-teaching.git"

	// 本地目标路径
	destinationPath := "./cloned-repo"

	// 克隆仓库到本地路径
	fmt.Println("Cloning repository...")

	// 克隆仓库
	_, err := git.PlainClone(destinationPath, false, &git.CloneOptions{
		URL:      repoURL,   // 仓库的URL
		Progress: os.Stdout, // 显示克隆进度
		// 如果需要凭据来克隆私有仓库，可以使用 Auth 来传递认证信息
		//Auth: &http.BasicAuth{
		//	Username: "your-username", // 用户名，这里可以使用 "any" 字符串，GitHub不使用这个字段
		//	Password: "your-token",    // 使用 GitHub Token 作为密码，或其他凭据
		//},
	})
	if err != nil {
		log.Fatalf("Error while cloning repository: %v", err)
	}

	fmt.Println("Repository cloned successfully!")
	return nil
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
	fmt.Println("File Content:")
	fmt.Println(fileContent)
	strings.ReplaceAll(fileContent, "version", req.Version)
	fmt.Println(fileContent)
	// 添加yml文件
	github.UpYml(req.Repository, "./github/workflows", req.CommitMessage, fileContent)
	return err
}
