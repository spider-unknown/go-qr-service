package impl

import (
	"context"
	"document-qr/internal/app/repository"
	"document-qr/internal/app/service"
	pb "document-qr/proto"
	"encoding/base64"
	"github.com/unistack-org/micro/v3/codec"
	"github.com/unistack-org/micro/v3/metadata"
	"github.com/yeqown/go-qrcode"
	"io/ioutil"
	"os"
	"time"
)

type DocumentServiceImpl struct {
	client             pb.MetadataProcessingServiceClient
	documentRepository repository.DocumentRepository
}

func NewDocumentServiceImpl(client pb.MetadataProcessingServiceClient, documentRepository repository.DocumentRepository) service.DocumentService {
	return &DocumentServiceImpl{client, documentRepository}
}

func (d DocumentServiceImpl) GetQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error {

	document, err := d.documentRepository.FindById(ctx, req)
	if err != nil {
		return &pb.Error{Msg: err.Error()}
	}
	res.Qr = document.QR
	res.Id = document.Id
	res.DocumentId = document.DocumentId
	return nil
}

func (d DocumentServiceImpl) MakeQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error {
	document, err := d.documentRepository.FindById(ctx, req)

	if document != nil {
		return &pb.ErrorAlreadyExists{Msg: "QR документа уже существует!"}
	}

	//metadata, err := d.client.GetFileMetadata(ctx, &pb.GetMetadataRequest{MetadataId: req.DocumentId})
	//if metadata == nil {
	//	return &pb.Error{Msg: "Документ не найден!"}
	//
	//}
	qrc, err := qrcode.New("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", qrcode.WithQRWidth(3))

	if err != nil {
		return &pb.Error{Msg: err.Error()}

	}

	fileName := time.Now().Format("RFC3339") + req.DocumentId + ".png"
	if err := qrc.Save(fileName); err != nil {
		return &pb.Error{Msg: err.Error()}
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return &pb.Error{Msg: err.Error()}
	}
	qr := base64.StdEncoding.EncodeToString(content)
	os.Remove(fileName)
	document, err = d.documentRepository.Create(ctx, req, qr)
	if err != nil {
		return &pb.Error{Msg: err.Error()}
	}

	res.Qr = qr
	res.Id = document.Id
	res.DocumentId = req.DocumentId

	return nil
}

func (d DocumentServiceImpl) GetQRImage(ctx context.Context, req *pb.PostQRRequest, rsp *codec.Frame) error {

	document, err := d.documentRepository.FindById(ctx, req)
	if err != nil {
		return &pb.Error{Msg: err.Error()}
	}

	fileName := time.Now().Format(time.RFC3339) + document.Id + ".png"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	qrBytes, _ := base64.StdEncoding.DecodeString(document.QR)
	_, err = file.Write(qrBytes)
	if err != nil {
		return &pb.Error{Msg: err.Error()}
	}
	data, _ := ioutil.ReadFile(file.Name())
	rsp.Data = data
	md, _ := metadata.FromIncomingContext(ctx)
	md.Set("Content-Disposition", "attachment; filename="+file.Name())
	metadata.SetOutgoingContext(ctx, md)
	os.Remove(fileName)

	return nil
}

func (d DocumentServiceImpl) NewDocument(ctx context.Context, req *pb.PostNewDocumentRequest, res *pb.PostNewDocumentResponse) error {
	document, err := d.documentRepository.CreateEmptyDocument(ctx)
	if err != nil {
		return &pb.Error{Msg: err.Error()}
	}
	res.QrId = document.Id
	return nil
}
