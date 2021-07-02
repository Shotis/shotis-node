package network

import (
	"bytes"
	context "context"
	"fmt"
	"runtime"

	"github.com/shotis/shotis-node/storage"
)

type ServerImpl struct {
	UnimplementedShotisServiceServer

	GCPService *storage.GoogleCloudStorageService
}

func (*ServerImpl) Health(context.Context, *HealthReportRequest) (*HealthReport, error) {
	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)

	return &HealthReport{
		MemoryUsage:     float64(stats.Alloc),
		Free:            float64(stats.Sys - stats.Alloc),
		Allocated:       float64(stats.TotalAlloc),
		AwaitingWorkers: 0,
		IdleWorkers:     0,
		UploadedImages:  1337,
	}, nil
}

//UploadImage takes in a response from the client, processes it, and uploads it to a GCP cloud storage
func (s *ServerImpl) UploadImage(ctx context.Context, message *UploadImageMessage) (*UploadImageResponse, error) {
	fmt.Println("Received Upload Request...")
	m, err := s.GCPService.Upload(message.FileName, bytes.NewReader(message.Data))

	if err != nil {
		return nil, err
	}

	fmt.Println("uploaded")

	return &UploadImageResponse{
		URL: m.AccessURL,
	}, err
}
