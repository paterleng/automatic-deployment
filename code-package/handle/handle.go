package handle

import (
	"code-package/data"
	"code-package/data/schema"
	"code-package/pkg/github"
	"code-package/rpc"
	"code-package/utils"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type CodePackage struct {
	repo data.CodePackageRepo
}

func Register(service micro.Service) error {
	repo := data.NewCodePackageRepo(data.NewData(utils.Tools.DB))
	cp := &CodePackage{repo: repo}
	go cp.Start()
	err := rpc.RegisterCodePackageHandler(service.Server(), cp)
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
func (h *CodePackage) StartPlan(ctx context.Context, req *rpc.StartPlanRequest, rsp *rpc.StartPlanResponse) error {

	// 区分这个请求是创建还是任务继续
	if req.IsFailPlan {
		// 查询数据库，获取任务
		d, err1 := h.repo.SelectPlanById(ctx, req.Id)
		if err1 != nil {
			utils.Tools.LG.Error("Error SelectPlanById:" + err1.Error())
			return nil
		}
		SC <- d
	} else {
		// 设置随机种子
		rand.Seed(time.Now().UnixNano())

		// 生成随机的 5 位数
		randomNum := rand.Intn(90000) + 10000 // 生成范围在 10000 到 99999 之间的数
		gapp := schema.GetAndPushPlan{
			GetUrl:      req.Url,
			ImageName:   req.Name,
			Status:      int(schema.PlanStatus_CloneCode),
			DownloadDir: strconv.Itoa(randomNum),
			Version:     req.Version,
			IsSuccess:   false,
		}

		err := h.repo.Create(ctx, gapp)
		if err != nil {
			utils.Tools.LG.Error("Error Create:" + err.Error())
			return nil
		}
		// 发送消息
		SC <- gapp
	}

	return nil

}

// 获取用户生成的dockerfile文件
func (h *CodePackage) GetDockerFile() {

}

// 获取任务状态
func (h *CodePackage) GetPlanStatus(ctx context.Context, req *rpc.GetPlanStatusRequest, rsp *rpc.GetPlanStatusResponse) error {
	// 查询id计划的执行状态
	gapp, err := h.repo.SelectPlanById(ctx, req.Id)
	if err != nil {
		utils.Tools.LG.Error("SelectPlanById err" + err.Error())
		return err
	}
	rsp.Status = int64(gapp.Status)
	rsp.Url = gapp.GetUrl
	rsp.Id = gapp.Id
	return err
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
	d, err := ioutil.ReadFile("ci.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 将内容转换为字符串
	fileContent := string(d)

	// 输出文件内容
	strings.ReplaceAll(fileContent, "version", req.Version)
	// 添加yml文件
	github.UpYml(req.Repository, "./github/workflows", req.CommitMessage, fileContent)
	return err
}
