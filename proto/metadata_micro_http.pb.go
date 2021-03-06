// Code generated by protoc-gen-micro
// source: metadata.proto
package pb

import (
	context "context"
	v3 "github.com/unistack-org/micro-client-http/v3"
	api "github.com/unistack-org/micro/v3/api"
	client "github.com/unistack-org/micro/v3/client"
	server "github.com/unistack-org/micro/v3/server"
	http "net/http"
)

type metadataProcessingServiceClient struct {
	c    client.Client
	name string
}

func NewMetadataProcessingServiceClient(name string, c client.Client) MetadataProcessingServiceClient {
	return &metadataProcessingServiceClient{c: c, name: name}
}

func (c *metadataProcessingServiceClient) GetFileMetadata(ctx context.Context, req *GetMetadataRequest, opts ...client.CallOption) (*GetMetadataResponse, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["500"] = &Error{}
	errmap["200"] = &GetMetadataResponse{}
	errmap["401"] = &Error{}
	errmap["404"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodGet),
		v3.Path("/metadata/{metadata_id}"),
	)
	rsp := &GetMetadataResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.GetFileMetadata", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *metadataProcessingServiceClient) UpdateFileMetadata(ctx context.Context, req *UpdateMetadataRequest, opts ...client.CallOption) (*UpdateMetadataResponse, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["200"] = &UpdateMetadataResponse{}
	errmap["400"] = &Error{}
	errmap["401"] = &Error{}
	errmap["404"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodPut),
		v3.Path("/metadata/{update_metadata_id}"),
		v3.Body("*"),
	)
	rsp := &UpdateMetadataResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.UpdateFileMetadata", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *metadataProcessingServiceClient) FindAll(ctx context.Context, req *FindAllRequest, opts ...client.CallOption) (MetadataProcessingService_FindAllClient, error) {
	errmap := make(map[string]interface{}, 3)
	errmap["200"] = &FindAllResponse{}
	errmap["401"] = &Error{}
	errmap["500"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodGet),
		v3.Path("/metadata"),
	)
	stream, err := c.c.Stream(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.FindAll", &FindAllRequest{}), opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(req); err != nil {
		return nil, err
	}
	return &metadataProcessingServiceClientFindAll{stream}, nil
}

type metadataProcessingServiceClientFindAll struct {
	stream client.Stream
}

func (s *metadataProcessingServiceClientFindAll) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceClientFindAll) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceClientFindAll) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceClientFindAll) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceClientFindAll) Recv() (*FindAllResponse, error) {
	msg := &FindAllResponse{}
	if err := s.stream.Recv(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *metadataProcessingServiceClient) FindAllWithConditions(ctx context.Context, req *FindAllWithConditionsRequest, opts ...client.CallOption) (MetadataProcessingService_FindAllWithConditionsClient, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["200"] = &FindAllResponse{}
	errmap["400"] = &Error{}
	errmap["401"] = &Error{}
	errmap["500"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodPost),
		v3.Path("/metadata/search"),
		v3.Body("*"),
	)
	stream, err := c.c.Stream(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.FindAllWithConditions", &FindAllWithConditionsRequest{}), opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(req); err != nil {
		return nil, err
	}
	return &metadataProcessingServiceClientFindAllWithConditions{stream}, nil
}

type metadataProcessingServiceClientFindAllWithConditions struct {
	stream client.Stream
}

func (s *metadataProcessingServiceClientFindAllWithConditions) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceClientFindAllWithConditions) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceClientFindAllWithConditions) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceClientFindAllWithConditions) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceClientFindAllWithConditions) Recv() (*FindAllResponse, error) {
	msg := &FindAllResponse{}
	if err := s.stream.Recv(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *metadataProcessingServiceClient) FindAllWithTextSearching(ctx context.Context, req *FindAllWithTextSearchingRequest, opts ...client.CallOption) (MetadataProcessingService_FindAllWithTextSearchingClient, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["200"] = &FindAllResponse{}
	errmap["400"] = &Error{}
	errmap["401"] = &Error{}
	errmap["500"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodPost),
		v3.Path("/metadata/search/text"),
		v3.Body("*"),
	)
	stream, err := c.c.Stream(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.FindAllWithTextSearching", &FindAllWithTextSearchingRequest{}), opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(req); err != nil {
		return nil, err
	}
	return &metadataProcessingServiceClientFindAllWithTextSearching{stream}, nil
}

type metadataProcessingServiceClientFindAllWithTextSearching struct {
	stream client.Stream
}

func (s *metadataProcessingServiceClientFindAllWithTextSearching) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceClientFindAllWithTextSearching) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceClientFindAllWithTextSearching) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceClientFindAllWithTextSearching) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceClientFindAllWithTextSearching) Recv() (*FindAllResponse, error) {
	msg := &FindAllResponse{}
	if err := s.stream.Recv(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *metadataProcessingServiceClient) FindAllWithFullSearching(ctx context.Context, req *FindAllWithTextSearchingRequest, opts ...client.CallOption) (MetadataProcessingService_FindAllWithFullSearchingClient, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["400"] = &Error{}
	errmap["401"] = &Error{}
	errmap["500"] = &Error{}
	errmap["200"] = &FindAllResponse{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodPost),
		v3.Path("/metadata/search/full"),
		v3.Body("*"),
	)
	stream, err := c.c.Stream(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.FindAllWithFullSearching", &FindAllWithTextSearchingRequest{}), opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(req); err != nil {
		return nil, err
	}
	return &metadataProcessingServiceClientFindAllWithFullSearching{stream}, nil
}

type metadataProcessingServiceClientFindAllWithFullSearching struct {
	stream client.Stream
}

func (s *metadataProcessingServiceClientFindAllWithFullSearching) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceClientFindAllWithFullSearching) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceClientFindAllWithFullSearching) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceClientFindAllWithFullSearching) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceClientFindAllWithFullSearching) Recv() (*FindAllResponse, error) {
	msg := &FindAllResponse{}
	if err := s.stream.Recv(msg); err != nil {
		return nil, err
	}
	return msg, nil
}

func (c *metadataProcessingServiceClient) SaveFingerprint(ctx context.Context, req *SaveFingerprintRequest, opts ...client.CallOption) (*SaveFingerprintResponse, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["400"] = &Error{}
	errmap["401"] = &Error{}
	errmap["500"] = &Error{}
	errmap["201"] = &SaveFingerprintRequest{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodPost),
		v3.Path("/metadata/fingerprint"),
		v3.Body("*"),
	)
	rsp := &SaveFingerprintResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.SaveFingerprint", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *metadataProcessingServiceClient) FindAllFingerprintsByID(ctx context.Context, req *FindAllFingerprintsByIDRequest, opts ...client.CallOption) (*FindAllFingerprintsByIDResponse, error) {
	errmap := make(map[string]interface{}, 4)
	errmap["200"] = &FindAllFingerprintsByIDResponse{}
	errmap["400"] = &Error{}
	errmap["401"] = &Error{}
	errmap["500"] = &Error{}
	opts = append(opts,
		v3.ErrorMap(errmap),
	)
	opts = append(opts,
		v3.Method(http.MethodGet),
		v3.Path("/metadata/files/{file_id}/fingerprints"),
	)
	rsp := &FindAllFingerprintsByIDResponse{}
	err := c.c.Call(ctx, c.c.NewRequest(c.name, "MetadataProcessingService.FindAllFingerprintsByID", req), rsp, opts...)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

type metadataProcessingServiceServer struct {
	MetadataProcessingServiceServer
}

func (h *metadataProcessingServiceServer) GetFileMetadata(ctx context.Context, req *GetMetadataRequest, rsp *GetMetadataResponse) error {
	return h.MetadataProcessingServiceServer.GetFileMetadata(ctx, req, rsp)
}

func (h *metadataProcessingServiceServer) UpdateFileMetadata(ctx context.Context, req *UpdateMetadataRequest, rsp *UpdateMetadataResponse) error {
	return h.MetadataProcessingServiceServer.UpdateFileMetadata(ctx, req, rsp)
}

func (h *metadataProcessingServiceServer) FindAll(ctx context.Context, stream server.Stream) error {
	msg := &FindAllRequest{}
	if err := stream.Recv(msg); err != nil {
		return err
	}
	return h.MetadataProcessingServiceServer.FindAll(ctx, msg, &metadataProcessingServiceFindAllStream{stream})
}

type metadataProcessingServiceFindAllStream struct {
	stream server.Stream
}

func (s *metadataProcessingServiceFindAllStream) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceFindAllStream) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceFindAllStream) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceFindAllStream) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceFindAllStream) Send(msg *FindAllResponse) error {
	return s.stream.Send(msg)
}

func (h *metadataProcessingServiceServer) FindAllWithConditions(ctx context.Context, stream server.Stream) error {
	msg := &FindAllWithConditionsRequest{}
	if err := stream.Recv(msg); err != nil {
		return err
	}
	return h.MetadataProcessingServiceServer.FindAllWithConditions(ctx, msg, &metadataProcessingServiceFindAllWithConditionsStream{stream})
}

type metadataProcessingServiceFindAllWithConditionsStream struct {
	stream server.Stream
}

func (s *metadataProcessingServiceFindAllWithConditionsStream) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceFindAllWithConditionsStream) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceFindAllWithConditionsStream) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceFindAllWithConditionsStream) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceFindAllWithConditionsStream) Send(msg *FindAllResponse) error {
	return s.stream.Send(msg)
}

func (h *metadataProcessingServiceServer) FindAllWithTextSearching(ctx context.Context, stream server.Stream) error {
	msg := &FindAllWithTextSearchingRequest{}
	if err := stream.Recv(msg); err != nil {
		return err
	}
	return h.MetadataProcessingServiceServer.FindAllWithTextSearching(ctx, msg, &metadataProcessingServiceFindAllWithTextSearchingStream{stream})
}

type metadataProcessingServiceFindAllWithTextSearchingStream struct {
	stream server.Stream
}

func (s *metadataProcessingServiceFindAllWithTextSearchingStream) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceFindAllWithTextSearchingStream) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceFindAllWithTextSearchingStream) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceFindAllWithTextSearchingStream) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceFindAllWithTextSearchingStream) Send(msg *FindAllResponse) error {
	return s.stream.Send(msg)
}

func (h *metadataProcessingServiceServer) FindAllWithFullSearching(ctx context.Context, stream server.Stream) error {
	msg := &FindAllWithTextSearchingRequest{}
	if err := stream.Recv(msg); err != nil {
		return err
	}
	return h.MetadataProcessingServiceServer.FindAllWithFullSearching(ctx, msg, &metadataProcessingServiceFindAllWithFullSearchingStream{stream})
}

type metadataProcessingServiceFindAllWithFullSearchingStream struct {
	stream server.Stream
}

func (s *metadataProcessingServiceFindAllWithFullSearchingStream) Close() error {
	return s.stream.Close()
}

func (s *metadataProcessingServiceFindAllWithFullSearchingStream) Context() context.Context {
	return s.stream.Context()
}

func (s *metadataProcessingServiceFindAllWithFullSearchingStream) SendMsg(msg interface{}) error {
	return s.stream.Send(msg)
}

func (s *metadataProcessingServiceFindAllWithFullSearchingStream) RecvMsg(msg interface{}) error {
	return s.stream.Recv(msg)
}

func (s *metadataProcessingServiceFindAllWithFullSearchingStream) Send(msg *FindAllResponse) error {
	return s.stream.Send(msg)
}

func (h *metadataProcessingServiceServer) SaveFingerprint(ctx context.Context, req *SaveFingerprintRequest, rsp *SaveFingerprintResponse) error {
	return h.MetadataProcessingServiceServer.SaveFingerprint(ctx, req, rsp)
}

func (h *metadataProcessingServiceServer) FindAllFingerprintsByID(ctx context.Context, req *FindAllFingerprintsByIDRequest, rsp *FindAllFingerprintsByIDResponse) error {
	return h.MetadataProcessingServiceServer.FindAllFingerprintsByID(ctx, req, rsp)
}

func RegisterMetadataProcessingServiceServer(s server.Server, sh MetadataProcessingServiceServer, opts ...server.HandlerOption) error {
	type metadataProcessingService interface {
		GetFileMetadata(ctx context.Context, req *GetMetadataRequest, rsp *GetMetadataResponse) error
		UpdateFileMetadata(ctx context.Context, req *UpdateMetadataRequest, rsp *UpdateMetadataResponse) error
		FindAll(ctx context.Context, stream server.Stream) error
		FindAllWithConditions(ctx context.Context, stream server.Stream) error
		FindAllWithTextSearching(ctx context.Context, stream server.Stream) error
		FindAllWithFullSearching(ctx context.Context, stream server.Stream) error
		SaveFingerprint(ctx context.Context, req *SaveFingerprintRequest, rsp *SaveFingerprintResponse) error
		FindAllFingerprintsByID(ctx context.Context, req *FindAllFingerprintsByIDRequest, rsp *FindAllFingerprintsByIDResponse) error
	}
	type MetadataProcessingService struct {
		metadataProcessingService
	}
	h := &metadataProcessingServiceServer{sh}
	var nopts []server.HandlerOption
	for _, endpoint := range NewMetadataProcessingServiceEndpoints() {
		nopts = append(nopts, api.WithEndpoint(endpoint))
	}
	return s.Handle(s.NewHandler(&MetadataProcessingService{h}, append(nopts, opts...)...))
}
