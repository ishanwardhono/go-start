package grpc

import context "context"

type articleGrpc struct {
	ArticlesGrpcServer
}

func NewArticleGrpcServer() articleGrpc {
	return articleGrpc{}
}

func (g articleGrpc) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return &GetAllResponse{
		Articles: []*Article{
			{
				Title:   "Hallo",
				Content: "testt",
				Author:  "tess2",
			},
		},
	}, nil
}
