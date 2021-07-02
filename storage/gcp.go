package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type CloudFile struct {
	// The name of the file
	FileName string
	// The access URL
	AccessURL string
	// The time it was created  at
	CreatedAt int64
}

type GCPOption func(*GoogleCloudStorageService)

func Bucket(name string) func(*GoogleCloudStorageService) {
	return func(gcss *GoogleCloudStorageService) {
		gcss.Bucket = name
	}
}

type GoogleCloudStorageService struct {
	internalClient *storage.Client
	Bucket         string
}

func (gcp *GoogleCloudStorageService) Upload(fileName string, file io.Reader) (*CloudFile, error) {

	/*
		When we upload a file it could be in the form of like a path.

		Say for example when we have teams

		The path could be like

		team/{teamUUID}/{imageID}

		or

		user/{userUUID}/{imageID}

		and that would just be behind the scenes

		front facing it would be like

		struct Image {
			Owner uuid   // the actual owner of the image
			Path string // the "path" in the Google Cloud Storage Bucket
			Name string // the custom file name
			URLExt string // for those with custom urls it can be like shot.is/my_image
			Public bool // whether or not this is public. For a team, if it is private, only team members should be able to view it
		};

		which would be stored in a sharded MongoDB
	*/

	bucket := gcp.internalClient.Bucket(gcp.Bucket)
	obj := bucket.Object(fileName)

	w := obj.NewWriter(context.Background())

	br, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	w.Write(br)
	w.Close()

	attr, err := obj.Attrs(context.Background())

	if err != nil {
		return nil, err
	}

	// TODO: return some sort of url
	return &CloudFile{
		FileName:  attr.Name,
		AccessURL: fmt.Sprintf("https://storage.cloud.google.com/%s/%s", attr.Bucket, attr.Name),
		CreatedAt: w.Created.UnixNano(),
	}, nil
}

func NewGoogleCloudStorage(ctx context.Context, credentialsFile string, options ...GCPOption) (*GoogleCloudStorageService, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))

	if err != nil {
		return nil, err
	}

	service := &GoogleCloudStorageService{
		internalClient: client,
	}

	for _, option := range options {
		option(service)
	}

	return service, nil
}
