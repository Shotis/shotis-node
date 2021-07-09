package models

import "github.com/google/uuid"

type ContentPath string

type Content struct {
	// The user who uploaded it.
	// If the uploader is using a team key
	Uploader uuid.UUID
	// The tag is the short URL path
	// i.e. shot.is/134acqC
	// the `134acqC` is the tag.
	// This should be able to change them
	Tag string
	// The location of the image in the bucket
	Path ContentPath
	// The amount of space this image is currently using
	SpaceUsed float64
	// If this image is private only the viewer
	// or team members can view this image
	Private bool
}
