package client

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"

	"github.com/craftto/go-tron/pkg/proto/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GrpcClient struct {
	GrpcURL     string
	Conn        *grpc.ClientConn
	Client      api.WalletClient
	grpcTimeout time.Duration
	opts        []grpc.DialOption
	apiKey      string
}

var TronGrpcClient *GrpcClient

func NewGrpcClient(grpcURL string, apiKey string, opts ...grpc.DialOption) (*GrpcClient, error) {
	conn, err := grpc.Dial("grpc.trongrid.io:50051", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		GrpcURL:     grpcURL,
		Conn:        conn,
		Client:      api.NewWalletClient(conn),
		grpcTimeout: 50 * time.Second,
		apiKey:      apiKey,
		opts:        opts,
	}, nil
}
func Start() (*GrpcClient, error) {
	TronGrpcClient, err := NewGrpcClient("grpc.trongrid.io:50051", "bb9ba261-37d3-4731-b2c8-dd56ae5e5abb", grpc.WithInsecure())
	if err != nil {
		g.Log().Error(context.Background(), err)
		return nil, err
	}

	return TronGrpcClient, nil
}
func init() {
	start, err := Start()
	if err != nil {
		panic(err)
	}
	TronGrpcClient = start
}

func (g *GrpcClient) SetAPIKey(apiKey string) error {
	g.apiKey = apiKey
	return nil
}

func (g *GrpcClient) GetContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), g.grpcTimeout)

	if len(g.apiKey) > 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "TRON-PRO-API-KEY", g.apiKey)
	}

	return ctx, cancel
}

func (g *GrpcClient) Close() {
	if g.Conn != nil {
		g.Conn.Close()
	}
}

// GetMessageBytes return grpc message from bytes
func GetMessageBytes(m []byte) *api.BytesMessage {
	return &api.BytesMessage{
		Value: m,
	}
}

// GetPaginatedMessage return grpc message number
func GetPaginatedMessage(offset int64, limit int64) *api.PaginatedMessage {
	return &api.PaginatedMessage{
		Offset: offset,
		Limit:  limit,
	}
}
