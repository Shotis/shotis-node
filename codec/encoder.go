// Images, videos, and gifs can really add up in size
// We'd like to host them all but we also like money and can't go broke
// hosting uncompressed images.
package codec

//EncodeHEVC
// Will take in a video stream and encode it through FFMPEG/libav from whatever
// container format (e.g. H.264) into H.265/HEVC
// Which should keep most quality while producing a smaller file size (like 50% decrease in some cases)
func EncodeHEVC() {}

// EncodeWebP (ideally natively/cgo)
func EncodeWebP() {}
