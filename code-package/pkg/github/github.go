package github

import (
	"code-package/utils"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/v65/github"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/oauth2"
	"time"
)

// 获取 GitHub的令牌或者密码
func aaa() {
	// 使用令牌创建一个 OAuth2 客户端
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: utils.Conf.Auth})
	tc := oauth2.NewClient(context.Background(), ts)

	// 创建 GitHub 客户端
	client := github.NewClient(tc)

	// 获取用户的仓库
	repos, _, err := client.Repositories.List(context.Background(), "", nil)
	if err != nil {
		fmt.Printf("Error fetching repositories: %v\n", err)
		return
	}

	// 打印仓库名称
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}

}

func Bbb() {
	// GitHub 令牌
	token := utils.Conf.Auth
	owner := "GuZihang929"
	repo := "draw-together"

	// 要添加的 secret 名称和值
	secretName := "MY_SECRET"
	secretValue := "SuperSecretValue"

	// 设置 OAuth2 客户端
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	// 创建 GitHub 客户端
	client := github.NewClient(tc)
	// 获取仓库的公钥，用于加密 Secret
	publicKey, _, err := client.Actions.GetRepoPublicKey(context.Background(), owner, repo)
	if err != nil {
		fmt.Printf("Error fetching repository public key: %v\n", err)
		return
	}

	// 使用仓库的公钥加密 Secret
	encryptedSecret, err := encryptSecret(publicKey.GetKey(), secretValue)
	if err != nil {
		fmt.Printf("Error encrypting secret: %v\n", err)
		return
	}

	// 创建 Repository Secret 对象
	secret := &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedSecret,
	}

	// 添加或更新仓库 Secret
	_, err = client.Actions.CreateOrUpdateRepoSecret(context.Background(), owner, repo, secret)
	if err != nil {
		fmt.Printf("Error adding repository secret: %v\n", err)
		return
	}

	fmt.Printf("Secret %s successfully added to repository %s/%s\n", secretName, owner, repo)

}

// encryptSecret 使用仓库的公钥加密 secret
func encryptSecret(publicKey string, secretValue string) (string, error) {
	// base64 解码公钥
	decodedKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key: %w", err)
	}

	// 检查解码后的密钥长度是否为 32 字节
	if len(decodedKey) != 32 {
		return "", fmt.Errorf("invalid public key length: expected 32 bytes, got %d", len(decodedKey))
	}

	// 将 decodedKey 转换为 [32]byte 类型
	var key32 [32]byte
	copy(key32[:], decodedKey[:32])

	// 使用 box.SealAnonymous 进行加密
	encryptedBytes, err := box.SealAnonymous(nil, []byte(secretValue), &key32, rand.Reader)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt secret: %w", err)
	}

	// 将加密后的字节进行 Base64 编码
	encryptedValue := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encryptedValue, nil
}

func Ccc() {
	// GitHub 令牌
	// GitHub 令牌
	token := utils.Conf.Auth
	owner := "GuZihang929"
	repo := "draw-together"

	// 要上传的文件路径和内容
	filePath := ".github/workflows/your-workflow.yml"
	commitMessage := "Add workflow file"
	fileContent := `name: Example Workflow

on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout code
      uses: actions/checkout@v2
    - name: Run a one-liner
      run: echo "Hello, World!"
`

	// 设置 OAuth2 客户端
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	// 创建 GitHub 客户端
	client := github.NewClient(tc)

	// 上传文件
	err := uploadFile(client, owner, repo, filePath, commitMessage, fileContent)
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return
	}

	fmt.Println("File uploaded successfully!")
}

// uploadFile 上传文件到 GitHub 仓库
func uploadFile(client *github.Client, owner, repo, filePath, commitMessage, fileContent string) error {
	ctx := context.Background()

	// 获取当前仓库默认分支的最新信息
	ref, _, err := client.Git.GetRef(ctx, owner, repo, "heads/main")
	if err != nil {
		return fmt.Errorf("failed to get ref: %w", err)
	}

	// 获取最新的提交对象
	_, _, err = client.Git.GetCommit(ctx, owner, repo, ref.GetObject().GetSHA())
	if err != nil {
		return fmt.Errorf("failed to get latest commit: %w", err)
	}

	// 将文件内容编码为 base64
	encodedContent := github.String(fileContent)

	// 准备新的文件
	options := &github.RepositoryContentFileOptions{
		Message: github.String(commitMessage),
		Content: []byte(*encodedContent),
		Branch:  github.String("main"),
		Committer: &github.CommitAuthor{
			Name:  github.String("Your Name"),
			Email: github.String("your-email@example.com"),
			Date:  &github.Timestamp{Time: time.Now()},
		},
	}

	// 上传文件到仓库
	_, _, err = client.Repositories.CreateFile(ctx, owner, repo, filePath, options)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}
