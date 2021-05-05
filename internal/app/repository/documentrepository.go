package repository

import (
	"context"
	"document-qr/internal/app/model"
	pb "document-qr/proto"
)

type DocumentRepository interface {
	FindById(ctx context.Context, request *pb.PostQRRequest) (*model.Document, error)
	Create(ctx context.Context, request *pb.PostQRRequest, qr string) (*model.Document, error)
}
