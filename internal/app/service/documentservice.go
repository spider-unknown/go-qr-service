package service

import (
	"context"
	pb "document-qr/proto"
	"github.com/unistack-org/micro/v3/codec"
)

type DocumentService interface {
	MakeQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error
	GetQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error
	GetQRImage(ctx context.Context, req *pb.PostQRRequest, rsp *codec.Frame) error
}
