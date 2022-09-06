package provider

import (
	article "app/module/articles/handler/grpc"

	"google.golang.org/grpc"
)

func GetGrpcServers() *grpc.Server {
	server := grpc.NewServer()
	article.RegisterArticlesGrpcServer(server, article.NewArticleGrpcServer())
	return server
}
