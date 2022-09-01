package twirptest

import (
	"context"

	pb "github.com/suraj-swarnapuri/Golang101/rpc/twirptest"
)

type HelloWorld struct{

}

func (hw *HelloWorld) Hello(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hello "}, nil
}