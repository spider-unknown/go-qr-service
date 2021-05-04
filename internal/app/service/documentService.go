package service

import (
	"context"
	pb "document-qr/proto"
)

type DocumentService interface {
	MakeQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error
	GetQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error
}
