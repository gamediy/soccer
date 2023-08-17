package syssvc

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

// 测试upload
func TestFileStorer_Upload(t *testing.T) {
	storer := NewFileStorer("cloud", "D:\\")
	//强转 S3FileStorer

	storer.Upload(nil, 1, nil)
}

// 测试 Delete
func TestFileStorer_Delete(t *testing.T) {
	storer := NewFileStorer("cloud", "D:\\")
	//强转 S3FileStorer
	ctx := context.Background()
	storer.Delete(ctx, 239)
}
