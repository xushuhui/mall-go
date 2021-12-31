package test

import (
	"context"
	"log"
	"mall-go/api/mall"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/assert"
)

var appClient mall.AppHTTPClient

func TestGetBannerById(t *testing.T) {

	reply, err := appClient.GetBannerById(context.Background(), &mall.BannerByIdRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, reply.Id, int64(1))

}
func TestGetBannerByName(t *testing.T) {

	reply, err := appClient.GetBannerByName(context.Background(), &mall.BannerByNameRequest{Name: "b-1"})
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, reply.Name, "b-1")

}
func TestMain(m *testing.M) {
	var err error
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithMiddleware(
			recovery.Recovery(),
		),
		transhttp.WithEndpoint("127.0.0.1:8000"),
	)
	if err != nil {
		panic(err)
	}

	appClient = mall.NewAppHTTPClient(conn)
	m.Run()
}