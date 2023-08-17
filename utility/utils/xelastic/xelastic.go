package xelastic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"strings"
	"time"
)

var (
	Cfg = elasticsearch.Config{}
)

func Create(ctx context.Context, index string, body interface{}) {
	es, err := elasticsearch.NewClient(Cfg)
	if err != nil {
		fmt.Println(fmt.Sprintf("new client: %s", err))
		return
	}
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Generate a snowflake ID.
	id := node.Generate()
	docID := id.String()
	marshal, _ := json.Marshal(body)
	// Perform the create operation
	res, err := es.Create(
		index, //  api_operate_log_2023_08_02
		docID, // Document ID
		strings.NewReader(string(marshal)),
		es.Create.WithContext(ctx),
		es.Create.WithRefresh("true"), // Refresh the index immediately for testing purposes
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error creating document: %s", err))

	}
	if res != nil {
		defer res.Body.Close()
		if res.IsError() {
			fmt.Println(fmt.Sprintf("Error response: %s", res.String()))
		}
	}

}

func SaveEsLog(logName string) {
	var logHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
		format := time.Now().Format("2006-01-02")
		sprint := fmt.Sprint(logName, format)
		Create(context.Background(), sprint, in)
		in.Next(ctx)

	}
	// 注册自定义的Handler
	g.Log().SetHandlers(logHandler)
}

func InitCfg(ctx context.Context) {
	url := g.Cfg().MustGet(ctx, "elastic.url").String()
	Cfg = elasticsearch.Config{
		Addresses: []string{
			url, // Replace with the actual Elasticsearch URL
		},
	}
}
