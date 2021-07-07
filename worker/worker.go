package worker

import (
	"bytes"
	"context"
	"io"
	"runtime"

	"github.com/shotis/shotis-node/config"
	"github.com/shotis/shotis-node/network"
	"github.com/shotis/shotis-node/storage"
	"github.com/shotis/shotis-node/tasks"
)

//GRPCWorker is an implementation of the gRPC server
type GRPCWorker struct {
	network.UnimplementedShotisServiceServer
	ctx            context.Context
	storageService *storage.GoogleCloudStorageService
	taskQueue      *tasks.Queue
}

func (*GRPCWorker) Health(context.Context, *network.HealthReportRequest) (*network.HealthReport, error) {
	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)

	return &network.HealthReport{
		MemoryUsage:     float64(stats.Alloc),
		Free:            float64(stats.Sys - stats.Alloc),
		Allocated:       float64(stats.TotalAlloc),
		AwaitingWorkers: 0,
		IdleWorkers:     0,
		UploadedImages:  1337,
	}, nil
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

func (worker *GRPCWorker) Start() {}

func NewGRPCWorker(ctx context.Context, config *config.NodeConfig) *GRPCWorker {

}
