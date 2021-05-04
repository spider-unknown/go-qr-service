package impl

import (
	"bytes"
	"context"
	"document-qr/internal/app/repository"
	"document-qr/internal/app/service"
	pb "document-qr/proto"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/yeqown/go-qrcode"
	"io"
	"io/ioutil"
	"os"
)

type DocumentServiceImpl struct {
	client             pb.MetadataProcessingServiceClient
	documentRepository repository.DocumentRepository
}

func NewDocumentServiceImpl(client pb.MetadataProcessingServiceClient, documentRepository repository.DocumentRepository) service.DocumentService {
	return &DocumentServiceImpl{client, documentRepository}
}

func (d DocumentServiceImpl) GetQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error {
	//md, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	return &pb.Error{Msg: "Системная ошибка!"}
	//}
	//qrType, ok := md.Get("QR-Type")
	//if !ok {
	//	return &pb.Error{Msg: "Не найден QR-Type"}
	//}

	document, err := d.documentRepository.FindById(ctx, req)
	if err != nil {
		return err
	}

	if document.Id != "" {
		res.Qr = document.QR
		res.Id = document.Id
		res.DocumentId = document.DocumentId

		file, err := os.Create(document.Id + ".jpeg")
		if err != nil {
			return err
		}
		defer file.Close()

		//contentType, _ := md.Get("Content-Type")
		//contentLength, _ := md.Get("Content-Length")
		//md.Set("Content-Disposition", "attachment; filename="+file.Name())
		//md.Set("Content-Type", contentType)
		//md.Set("Content-Length", contentLength)

		//if qrType == "1" {
		qrBytes, _ := base64.StdEncoding.DecodeString(document.QR)
		_, err = io.Copy(file, bytes.NewReader(qrBytes))
		if err != nil {
			return err
		}
		//}
		return nil
	} else {
		return &pb.ErrorNotFound{Msg: "Документ не найден!"}
	}
}

func (d DocumentServiceImpl) MakeQR(ctx context.Context, req *pb.PostQRRequest, res *pb.PostQRResponse) error {
	document, err := d.documentRepository.FindById(ctx, req)
	if err != nil {
		return err
	}
	if document.Id != "" {
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

	fileName := uuid.New().String() + "qrrrr.jpeg"
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
		return err
	}

	if document != nil {
		res.Qr = document.QR
		res.Id = document.Id
		res.DocumentId = document.DocumentId
		return nil
	}

	return nil
}
