package handler

import (
	"context"
	"document-qr/internal/app/service"
	pb "document-qr/proto"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/v3/codec"
	"net/http"
)

type DocumentQRHandler struct {
	documentService service.DocumentService
}

func NewDocumentQRHandler(documentService service.DocumentService) *DocumentQRHandler {
	return &DocumentQRHandler{documentService}
}

func (h *DocumentQRHandler) GetDocumentQR(ctx context.Context, req *pb.PostQRRequest, rsp *pb.PostQRResponse) error {
	return responseHandler(ctx, h.documentService.GetQR(ctx, req, rsp), http.StatusOK)
}

func (h *DocumentQRHandler) PostDocumentQR(ctx context.Context, req *pb.PostQRRequest, rsp *pb.PostQRResponse) error {
	return responseHandler(ctx, h.documentService.MakeQR(ctx, req, rsp), http.StatusOK)
}

func (h *DocumentQRHandler) PostQRImage(ctx context.Context, req *pb.PostQRRequest, rsp *codec.Frame) error {
	return responseHandler(ctx, h.documentService.GetQRImage(ctx, req, rsp), http.StatusOK)
}

func responseHandler(ctx context.Context, resp interface{}, statusCode int) error {
	_, isInternal := resp.(*pb.Error)
	_, isDenied := resp.(*pb.ErrorAccessDenied)
	_, isUnauthorized := resp.(*pb.ErrorUnauthorized)
	_, isNotFound := resp.(*pb.ErrorNotFound)
	_, isBadRequest := resp.(*pb.ErrorBadRequest)
	_, isExists := resp.(*pb.ErrorAlreadyExists)

	if isInternal {
		httpsrv.SetRspCode(ctx, http.StatusInternalServerError)
	} else if isDenied {
		httpsrv.SetRspCode(ctx, http.StatusForbidden)
	} else if isDenied {
		httpsrv.SetRspCode(ctx, http.StatusForbidden)
	} else if isUnauthorized {
		httpsrv.SetRspCode(ctx, http.StatusUnauthorized)
	} else if isNotFound {
		httpsrv.SetRspCode(ctx, http.StatusNotFound)
	} else if isBadRequest {
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
	} else if isExists {
		httpsrv.SetRspCode(ctx, http.StatusBadRequest)
	} else {
		httpsrv.SetRspCode(ctx, statusCode)
		return nil
	}

	return httpsrv.SetError(resp)
}
