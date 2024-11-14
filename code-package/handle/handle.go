package handle

import (
	"code-package/rpc"
	"context"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

func (h *CodePackage) PullCode(ctx context.Context, req *rpc.PullCodeRequest, rsp *rpc.PullCodeResponse) error {
	//走拉取代码的逻辑
	// 定义目标目录
	targetDir := "./code"

	// 检查目录是否存在，如果存在则删除
	if err := os.Mkdir(targetDir, 0755); !os.IsNotExist(err) {
		fmt.Printf("Directory %s already exists. Deleting...\n", targetDir)
		if err := os.RemoveAll(targetDir); err != nil {
			log.Fatal(err)
		}
	}
	//把克隆的代码放到内存中管理，所以不需要创建额外的目录保存代码
	r, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:           req.Url,
		ReferenceName: plumbing.ReferenceName(req.Branch), // 替换为你要拉取的分支名
		//Auth: &http.BasicAuth{
		//	Username: req.Account,  // 通常是 Git 平台的用户名
		//	Password: req.Password, // 密码或访问令牌
		//},
	})
	fmt.Println(r)
	if err != nil {
		fmt.Println(err)
	}
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
