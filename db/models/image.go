package models

import "github.com/google/uuid"

type ImagePath string

// Name will returns the base file name
func (s ImagePath) Name() {

}

type Image struct {
	// The user who uploaded it.
	// If the uploader is using a team key
	Uploader uuid.UUID
	// The tag is the short URL path
	// i.e. shot.is/134acqC
	// the `134acqC` is the tag.
	// This should be able to change them
	Tag string
	// The location of the image in the bucket
	Path ImagePath
	// If this image is private only the viewer
	// or team members can view this image
	Private bool
}
