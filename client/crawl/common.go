package crawl

import (
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/goodguy-project/goodguy-core/idl"
)

var (
	Client idl.CrawlServiceClient
)

func MustInitClient() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", viper.GetString("client.crawler.host"), viper.GetInt("client.crawler.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	Client = idl.NewCrawlServiceClient(conn)
}
