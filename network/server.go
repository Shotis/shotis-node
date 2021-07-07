package network

import (
	"bytes"
	context "context"
	"io"
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
func (s *ServerImpl) UploadImage(stream ShotisService_UploadImageServer) error {
	var header *FileHeader
	var buf bytes.Buffer
	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			stream.SendAndClose(&UploadImageResponse{
				Status:  Status_Failed, // the status failed
				Message: err.Error(),
			})
			return err
		}

		if header == nil {
			header = message.GetHeader()
		}

		buf.Write(message.GetData())
	}

	cf, err := s.GCPService.Upload(header.FileName, &buf)

	if err != nil {
		stream.SendAndClose(&UploadImageResponse{
			Status:  Status_Failed, // the status failed
			Message: err.Error(),
		})
		return err
	}

	return stream.SendAndClose(&UploadImageResponse{
		URL:     cf.AccessURL,
		ImageId: cf.FileName,
		Status:  Status_OK,
		Message: "OK",
	})
}
