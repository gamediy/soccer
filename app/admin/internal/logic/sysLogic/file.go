package sysLogic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"star_net/app/admin/api/sys"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
)

type lFile struct {
}

var (
	File = lFile{}
)

// 分片上传
func (l lFile) UploadFileChunk(ctx context.Context, file *ghttp.UploadFile, group, index int, fileName, ext string) (string, error) {
	if file == nil {
		return "", consts.ErrImgCannotBeEmpty
	}
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return "", err
	}
	rootPath := gfile.Pwd() + rootFilePath.String()
	datePre := time.Now().Format("2006/01")
	newFileName := fmt.Sprint(grand.S(6), ext)
	dirName := fmt.Sprint(grand.S(6), ext) + "_" + strconv.Itoa(index)
	filePath := fmt.Sprintf("%s/%d/%s/%s", rootPath, group, datePre, dirName)
	_, err = file.Save(filePath)
	if err != nil {
		return "", err
	}
	return newFileName, nil
}

// 合并分片上传
func (l lFile) MergeFileChunk(ctx context.Context, group int, fileName string, chunks int) (string, error) {
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}

	//创建文件
	datePre := time.Now().Format("2006/01")
	if err != nil {
		return "", err
	}
	rootPath := gfile.Pwd() + rootFilePath.String()
	filePath := fmt.Sprintf("%s/%d/%s/%s", rootPath, group, datePre, fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	//判断chunks是否完整
	if chunks <= 0 {
		return "", consts.ErrImgCannotBeEmpty
	}
	//遍历chunks 从0开始
	for i := 0; i < chunks; i++ {
		//判断分片是否存在
		parts := gstr.Split(fileName, ".")
		indexStr := fmt.Sprintf("%d", i)
		filePath := filepath.Join(rootPath, strconv.Itoa(group), datePre, parts[0]+"_"+indexStr, "blob")
		data, e := os.ReadFile(filePath)
		if e != nil {
			return "", err
		}
		//写入文件
		file.Write(data)
		err := os.RemoveAll(filepath.Dir(filePath))
		if err != nil {
			return "", err
		}
	}
	if err := file.Sync(); err != nil {
		return "", err
	}
	dbName := fmt.Sprintf("%d/%s/%s", group, datePre, fileName)
	f := entity.File{
		Url:    dbName,
		Group:  group,
		Status: 1,
	}
	if err = l.AddFile(ctx, &f); err != nil {
		return "", err
	}
	return dbName, err
}

// 添加文件
func (l lFile) AddFile(ctx context.Context, e *entity.File) error {
	_, err := dao.File.Ctx(ctx).Insert(e)
	return err
}

func (l lFile) GetFileById(ctx context.Context, id uint64) (*entity.File, error) {
	var data entity.File
	one, err := dao.File.Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (l lFile) DelFile(ctx context.Context, id uint64) error {
	if _, err := dao.File.Ctx(ctx).Delete("id", id); err != nil {
		return err
	}
	return nil
}
func (l lFile) ListFile(ctx context.Context, req *sys.ListFileReq) ([]*entity.File, int, error) {
	var data = make([]*entity.File, 0)
	db := dao.File.Ctx(ctx)
	count, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	if err = db.Page(int(req.Page), int(req.Size)).Order("id desc").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, err
	}
	return data, count, nil
}
