package grpc

type articleGrpc struct {
	UnimplementedArticlesGrpcServer
}

func NewArticleGrpcServer() articleGrpc {
	return articleGrpc{}
}
