package github

import (
	"code-package/utils"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v65/github"
	"golang.org/x/crypto/nacl/box"
	"golang.org/x/oauth2"
	"os"
	"time"
)

var ghClient *github.Client

// 初始化Github客户端
func init() {
	// 使用令牌创建一个 OAuth2 客户端
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: utils.Conf.Auth})
	tc := oauth2.NewClient(context.Background(), ts)

	// 创建 GitHub 客户端
	ghClient = github.NewClient(tc)
}

// 仓库添加关键字
func UpdateRepoSecret(secretName, secretValue string) error {

	// 获取仓库的公钥，用于加密 Secret
	publicKey, _, err := ghClient.Actions.GetRepoPublicKey(context.Background(), utils.Conf.GitHub.UserName, utils.Conf.GitHub.Repository)
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("Error fetching repository public key: %v\n", err))
		return err
	}
	// 使用仓库的公钥加密 Secret
	encryptedSecret, err := encryptSecret(publicKey.GetKey(), secretValue)
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("Error encrypting secret: %v\n", err))
		return err
	}

	// 创建 Repository Secret 对象
	secret := &github.EncryptedSecret{
		Name:           secretName,
		KeyID:          publicKey.GetKeyID(),
		EncryptedValue: encryptedSecret,
	}

	// 添加或更新仓库 Secret
	_, err = ghClient.Actions.CreateOrUpdateRepoSecret(context.Background(), utils.Conf.GitHub.UserName, utils.Conf.GitHub.Repository, secret)
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("Error adding repository secret: %v\n", err))
		return err
	}
	utils.Tools.LG.Info(fmt.Sprintf("Secret %s successfully added to repository %s/%s\n", secretName, utils.Conf.GitHub.UserName, utils.Conf.GitHub.Repository))
	return nil
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

// 上传yml文件到 github
func UpYml(repository, filePath, commitMessage, fileContent string) {

	// 上传文件
	err := uploadFile(ghClient, utils.Conf.GitHub.UserName, repository, filePath, commitMessage, fileContent)
	if err != nil {
		utils.Tools.LG.Error(fmt.Sprintf("Error uploading file: %v\n", err))
		return
	}

	utils.Tools.LG.Info("File uploaded successfully!")
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
			// 使用 github.Timestamp 的 New 函数创建时间
			Date: &github.Timestamp{Time: time.Now()},
		},
	}

	// 上传文件到仓库
	_, _, err = client.Repositories.CreateFile(ctx, owner, repo, filePath, options)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	return nil
}

// 拉取代码到指定地址
func CloneCode(url, dir string) error {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,       // 仓库的URL
		Progress: os.Stdout, // 显示克隆进度
		// 如果需要凭据来克隆私有仓库，可以使用 Auth 来传递认证信息
		//Auth: &http.BasicAuth{
		//	Password: auth, // 使用 GitHub Token 作为密码，或其他凭据
		//},
	})
	return err
}
