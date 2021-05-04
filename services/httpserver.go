package services

import (
	"context"
	serverconfig "document-qr/configs"
	servicehandler "document-qr/internal/app/handler"
	"document-qr/internal/app/repository"
	"document-qr/internal/app/service/impl"
	pb "document-qr/proto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	httpcli "github.com/unistack-org/micro-client-http/v3"
	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/v3"
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/server"
)

func StartDocumentQRService(ctx context.Context, errCh chan<- error) {
	cfg := serverconfig.NewConfig("document-service", "1.0", "127.0.0.1:8080")

	svc := micro.NewService()

	if err := svc.Init(
		micro.Server(httpsrv.NewServer(
			server.Name(cfg.Server.Name),
			server.Version(cfg.Server.Version),
			server.Address(cfg.Server.Addr),
			server.Context(ctx),
			server.Codec("application/json", jsoncodec.NewCodec()))),
		micro.Client(httpcli.NewClient(
			client.ContentType("application/json"),
			client.Codec("application/json", jsoncodec.NewCodec()),
			client.Context(ctx),
		)),
	); err != nil {
		logger.Errorf(ctx, "Error initializing service %v", err)
		errCh <- err
		return
	}

	db := <-initDB("postgres", "postgres://postgres:password@localhost:5432/document_qrs?sslmode=disable", errCh)

	documentRepository := repository.NewDocumentRepository(db, jsoncodec.NewCodec())
	cli := client.NewClientCallOptions(svc.Client(), client.WithAddress(""))
	processCli := pb.NewMetadataProcessingServiceClient("metadata-client", cli)
	//processCli.GetFileMetadata(ctx, &pb.GetMetadataRequest{MetadataId: "123123"})

	documentService := impl.NewDocumentServiceImpl(processCli, documentRepository)
	handler := servicehandler.NewDocumentQRHandler(documentService)
	if err := pb.RegisterDocumentQRProcessingServiceServer(svc.Server(), handler); err != nil {
		logger.Errorf(ctx, "Error registering server %v", err)
		errCh <- err
		return
	}

	if err := svc.Run(); err != nil {
		logger.Errorf(ctx, "Error runnig service %v", err)
		errCh <- err
		return
	}
}

func initDB(name, url string, errCh chan<- error) <-chan *sqlx.DB {
	dbCh := make(chan *sqlx.DB, 1)
	go func() {
		defer close(dbCh)
		db, err := sqlx.Connect(name, url)
		if err != nil {
			errCh <- err
		}
		dbCh <- db
	}()
	return dbCh
}
