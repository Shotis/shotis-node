package worker

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net"

	"github.com/shotis/shotis-node/config"
	"github.com/shotis/shotis-node/network"
	"github.com/shotis/shotis-node/storage"
	"github.com/shotis/shotis-node/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//GRPCWorker is an implementation of the gRPC server
type GRPCWorker struct {
	network.UnimplementedShotisServiceServer
	ctx            context.Context
	storageService *storage.GoogleCloudStorageService
	taskQueue      *tasks.Queue
}

func (*GRPCWorker) Health(context.Context, *network.HealthReportRequest) (*network.HealthReport, error) {
	// TODO
	return nil, nil
}

//UploadImage takes in a response from the client, processes it, and uploads it to a GCP cloud storage
func (s *GRPCWorker) UploadImage(stream network.ShotisService_UploadImageServer) error {

	var header *network.FileHeader
	var buf bytes.Buffer
	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			stream.SendAndClose(&network.UploadImageResponse{
				Status:  network.Status_Failed, // the status failed
				Message: err.Error(),
			})
			return err
		}

		if header == nil {
			header = message.GetHeader()
		}

		buf.Write(message.GetData())
	}

	cf, err := s.storageService.Upload(header.FileName, &buf)

	if err != nil {
		stream.SendAndClose(&network.UploadImageResponse{
			Status:  network.Status_Failed, // the status failed
			Message: err.Error(),
		})
		return err
	}

	return stream.SendAndClose(&network.UploadImageResponse{
		URL:     cf.AccessURL,
		ImageId: cf.FileName,
		Status:  network.Status_OK,
		Message: "OK",
	})
}

//StartTLS starts a GRPC server
func (worker *GRPCWorker) StartTLS(host, cert, key string) error {
	tlsCert, err := tls.LoadX509KeyPair(cert, key)

	if err != nil {
		return err
	}

	server := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&tlsCert)))
	network.RegisterShotisServiceServer(server, worker)

	l, err := net.Listen("tcp", host)

	if err != nil {
		return err
	}

	return server.Serve(l)
}

func (worker *GRPCWorker) Start(host string) error {
	server := grpc.NewServer()
	network.RegisterShotisServiceServer(server, worker)

	l, err := net.Listen("tcp", host)

	if err != nil {
		return err
	}
	return server.Serve(l)
}

func NewGRPCWorker(ctx context.Context, config *config.NodeConfig) (*GRPCWorker, error) {
	service, err := storage.NewGoogleCloudStorage(ctx, config.Cloud.Storage.AuthKey, storage.Bucket(config.Cloud.Storage.Bucket))

	if err != nil {
		return nil, err
	}

	return &GRPCWorker{
		ctx:            ctx,
		storageService: service,
	}, nil
}
