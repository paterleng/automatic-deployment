package cmd

import (
	"code-package/utils"
	"os"
	"os/exec"
)

func PackagingImage(name, destinationPath, tempDir string) error {

	// 检查目录是否存在
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		return err
	}
	// 构建 docker build 命令
	cmd := exec.Command("docker", "build", "-t", name, destinationPath)
	// 将标准输出和标准错误连接到当前进程的输出中
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func MarkImage(name, version string) error {
	cmd := exec.Command("docker", "tag", name, utils.Conf.ImagesHub.Repo+"/"+utils.Conf.ImagesHub.NameSpace+"/"+name+":"+version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func PushImage(name, version string) error {
	cmd := exec.Command("docker", "push", utils.Conf.ImagesHub.Repo+"/"+utils.Conf.ImagesHub.NameSpace+"/"+name+":"+version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func DockerLogin() error {
	cmd := exec.Command("docker", "login", "-u", utils.Conf.ImagesHub.UserName, "-p", utils.Conf.ImagesHub.Password)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func DockerOut() error {
	cmd := exec.Command("docker", "logout")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
