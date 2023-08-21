package syssvc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"star_net/app/admin/api/sys"
	"star_net/app/admin/internal/logic/sysLogic"
	"star_net/consts"
	"star_net/db/model/entity"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
)

type FileStorer interface {
	Upload(ctx context.Context, group int, file *ghttp.UploadFile) (string, error) // 上传文件
	Delete(ctx context.Context, id int) error                                      // 删除文件
	GetFileById(ctx context.Context, id uint64) (*entity.File, error)
	ListFile(ctx context.Context, req *sys.ListFileReq) ([]*entity.File, int, error)
	GetFile(ctx context.Context, id uint64) (*entity.File, error)
}

// 如果是云服务，则传桶名，本地则传路径
func NewFileStorer(storerType string, baseDir string) FileStorer {
	switch storerType {
	case "local":
		return NewLocalFileStorer(baseDir)
	case "cloud":
		return NewS3FileStorer("8429dbc0235a26dda82c90041aa831dd", "fire", "093843554af5bbea77f54ed3ad157599", "d0de875d34365f6c1dc86ab44abb3b9a00131096fe0502a75e61fadc86d5c7e5")
	default:
		return NewLocalFileStorer(baseDir)
	}
}

// 本地操作
type LocalFileStorer struct {
	BaseDir string
}

// 云上传，待实现
type S3FileStorer struct {
	Client     *s3.Client
	BucketName string
}

// 创建新的S3FileStorer
// 创建新的S3FileStorer
func NewS3FileStorer(bucketName, accountId, accessKeyId, accessKeySecret string) *S3FileStorer {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://r2.cloudflarestorage.com/%s", accountId), // 从前端代码获取
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3FileStorer{
		Client:     client,
		BucketName: bucketName,
	}
}

func (s *S3FileStorer) Delete(ctx context.Context, id int) error {
	//获取文件信息
	file, err := sysLogic.File.GetFileById(ctx, uint64(id))

	if err != nil {
		return err
	}
	if file == nil {
		return errors.New("文件不存在")
	}
	// 删除数据库记录
	err = sysLogic.File.DelFile(ctx, uint64(id))
	if err != nil {
		return err
	}
	//删除云端文件
	outPut, err := s.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(file.Url),
	})
	if err != nil {
		return err
	}
	fmt.Println(outPut)
	return nil
}

func (s *S3FileStorer) GetFile(ctx context.Context, id uint64) (*entity.File, error) {
	// 添加获取S3文件的逻辑
	return nil, nil
}

func (s *S3FileStorer) GetFileById(ctx context.Context, id uint64) (*entity.File, error) {
	// 这里你需要添加用于从S3获取文件信息的代码
	return nil, nil
}

func (s *S3FileStorer) ListFile(ctx context.Context, req *sys.ListFileReq) ([]*entity.File, int, error) {
	// 这里你需要添加用于从S3列出文件的代码
	return nil, 0, nil
}

func (s *S3FileStorer) Upload(ctx context.Context, group int, file *ghttp.UploadFile) (string, error) {

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
		return "", errors.New("文件类型不支持")
	}
	if file == nil {
		return "", errors.New("file is nil")
	}

	file.Filename = fmt.Sprint(grand.S(6), filepath.Ext(file.Filename))
	// 这里你需要添加上传文件到S3的代码
	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.BucketName),
		Key:         &file.Filename,
		ContentType: aws.String(contentType),
	}

	reader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()
	var data []byte
	data, err = io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	input.Body = bytes.NewReader(data)
	outPut, err := s.Client.PutObject(ctx, input)
	if err != nil {
		return "", err
	}
	// 检查tag
	if outPut.ETag == nil {
		return "", consts.ErrUploadFile
	}
	f := entity.File{
		Url:    file.Filename,
		Group:  group,
		Status: 1,
	}
	if err = sysLogic.File.AddFile(ctx, &f); err != nil {
		return "", err
	}
	return file.Filename, err
}

func NewLocalFileStorer(baseDir string) *LocalFileStorer {
	return &LocalFileStorer{
		BaseDir: baseDir,
	}
}

// Upload上传文件
func (s *LocalFileStorer) Upload(ctx context.Context, group int, file *ghttp.UploadFile) (string, error) {
	if file == nil {
		return "", consts.ErrImgCannotBeEmpty
	}
	fileName := fmt.Sprint(grand.S(6), filepath.Ext(file.Filename))
	file.Filename = fileName
	datePre := time.Now().Format("2006/01")
	mixPath := fmt.Sprintf("%s/%d/%s/", s.BaseDir, group, datePre)
	_, err := file.Save(mixPath)
	if err != nil {
		return "", err
	}
	dbName := fmt.Sprintf("%d/%s/%s", group, datePre, file.Filename)
	f := entity.File{
		Url:    dbName,
		Group:  group,
		Status: 1,
	}
	if err = sysLogic.File.AddFile(ctx, &f); err != nil {
		return "", err
	}
	return dbName, err
}

// Delete 删除文件
func (l *LocalFileStorer) Delete(ctx context.Context, id int) error {

	//获取文件信息
	file, err := sysLogic.File.GetFileById(context.Background(), uint64(id))
	if err != nil {
		return err
	}
	filePath := filepath.Join(l.BaseDir, file.Url)
	//删除文件
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	//删除数据库记录
	err = sysLogic.File.DelFile(ctx, uint64(id))
	if err != nil {
		return err
	}
	return nil
}

// GetFileById 根据id获取文件
func (l *LocalFileStorer) GetFileById(ctx context.Context, id uint64) (*entity.File, error) {
	return sysLogic.File.GetFileById(ctx, id)
}

// ListFile
func (l *LocalFileStorer) ListFile(ctx context.Context, req *sys.ListFileReq) ([]*entity.File, int, error) {
	return sysLogic.File.ListFile(ctx, req)
}

// GetFile
func (l *LocalFileStorer) GetFile(ctx context.Context, id uint64) (*entity.File, error) {
	return sysLogic.File.GetFileById(ctx, id)
}
