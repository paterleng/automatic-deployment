package handle

import (
	"code-package/rpc"
	"context"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
	"os/exec"
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

// 获取用户 github 添加yml
