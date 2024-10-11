package cmd

import (
	"fmt"
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

	// 执行命令
	fmt.Printf("Running command: docker build -t %s %s\n", name, destinationPath)
	return cmd.Run()

}
