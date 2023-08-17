package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"star_net/consts"
	"star_net/db/dao"
	"star_net/db/model/entity"
	"star_net/pure/get"
	"star_net/utility/utils/xcrud"
	"strings"
)

func HandleGenApi(ctx context.Context, dirPath string, apiPrefix ...string) error {
	if err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return fmt.Errorf("I can't find sorry~")
		}

		// 只处理Go源文件
		prefix := ""
		if len(apiPrefix) > 0 {
			prefix = apiPrefix[0]
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			if err = processFile(ctx, path, prefix); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func processFile(ctx context.Context, filePath string, apiPrefix string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// 解析Go源代码，生成语法树
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", file, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return nil
	}

	// 遍历语法树，查找结构体字段中的tags标签
	for _, decl := range f.Decls {
		// 判断节点类型是否为结构体
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				// 获取结构体信息
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						// 遍历结构体字段
						for _, field := range structType.Fields.List {
							// 判断字段是否有tags标签
							if field.Tag != nil {
								tags := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
								tagValue := tags.Get("tags")
								if tagValue != "" {
									var api entity.Api
									api.Group = tags.Get("tags")
									api.Method = strings.ToUpper(tags.Get("method"))
									api.Url = fmt.Sprintf("%s%s", apiPrefix, tags.Get("path"))
									api.Desc = fmt.Sprintf("%s %s", tags.Get("summery"), tags.Get("dc"))
									if err = genOneApi(ctx, api); err != nil {
										return err
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func genOneApi(ctx context.Context, api entity.Api) error {
	g.Log().Info(ctx, api)
	a, err := get.ApiByGroupAndUrl(ctx, api.Group, api.Url, api.Method)
	if err != nil {
		if errors.Is(err, consts.ErrDataNotFound) {
			x := xcrud.Create{Ctx: ctx, Table: "s_api", Data: api}
			if err = x.Exec(); err != nil {
				return err
			}
		}
		return nil
	}
	a.Desc = api.Desc
	a.Method = api.Method
	if _, err = dao.Api.Ctx(ctx).Save(a); err != nil {
		return err
	}
	return nil
}
