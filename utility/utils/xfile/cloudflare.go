package xfile

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"path/filepath"
)

type CloudFlare struct {
	BucketName      string
	AccountId       string
	AccessKeyId     string
	AccessKeySecret string
	MaxSize         float64
}

func NewCloudFlareFromCtx(ctx context.Context) CloudFlare {
	x := CloudFlare{
		BucketName:      g.Cfg().MustGet(ctx, "cloudflare.upload.BucketName").String(),
		AccountId:       g.Cfg().MustGet(ctx, "cloudflare.upload.AccountId").String(),
		AccessKeySecret: g.Cfg().MustGet(ctx, "cloudflare.upload.AccessKeySecret").String(),
		AccessKeyId:     g.Cfg().MustGet(ctx, "cloudflare.upload.AccessKeyId").String(),
		MaxSize:         g.Cfg().MustGet(ctx, "cloudflare.upload.MaxSize").Float64(),
	}
	return x
}

// Upload group 1 系统 2 文件 3 凭证
func (in CloudFlare) Upload(ctx context.Context, group int) (string, error) {
	if group == 0 {
		return "", fmt.Errorf("请输入分组")
	}
	var (
		r = ghttp.RequestFromCtx(ctx)
	)
	file := r.GetUploadFile("file")

	fileSize := float64(file.Size)
	maxFileSize := in.MaxSize * 1024 * 1024 // 0.5MB in bytes
	// Check the file size
	if fileSize > maxFileSize {
		return "", fmt.Errorf("文件大小超过限制")
	}
	// Open the file for read
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	// Get the file size

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			//URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", in.AccountId),
			URL: fmt.Sprintf("https://r2.cloudflarestorage.com/%s", in.AccountId), // 从前端代码获取
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(in.AccessKeyId, in.AccessKeySecret, "")),
	)
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	//判断文件类型 通过后缀名
	ext := filepath.Ext(file.Filename)
	var contentType string
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".mp4":
		contentType = "video/mp4"
	case ".mp3":
		contentType = "audio/mp3"
	default:
		return "", fmt.Errorf("文件类型不支持")
	}

	// 获取当前年和月
	currentYear, currentMonth := gtime.Now().Year(), gtime.Now().Month()

	// Build the object key with year and month
	fileName := fmt.Sprintf("%d/%d/%d/%s%s", group, currentYear, currentMonth, grand.S(6), filepath.Ext(file.Filename))
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &in.BucketName,
		Key:         &fileName,
		Body:        f,
		ContentType: &contentType,
	})

	if err != nil {
		return "", err
	}

	return fileName, nil
}
