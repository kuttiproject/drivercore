package drivercore

// ImageStatus can be Notdownloaded, Downloaded or Unknown
type ImageStatus string

// Image Statuses
const (
	ImageStatusNotDownloaded = ImageStatus("Notdownloaded")
	ImageStatusDownloaded    = ImageStatus("Downloaded")
	ImageStatusUnknown       = ImageStatus("Unknown")
)

// Image describes a template from which Machines are created. Each kutti
// Image should contain binaries for a specific Kubernetes version. A Driver
// is expected to maintain a cache of Images locally, which can be updated
// from a driver-specific source. An Image listed in this cache will have a
// status of Notdownloaded before the actual template is downloaded, or
// Downloaded after it is. An Image in the local cache can be
// downloaded from a driver-specific source using the Fetch method, or
// added to the cache from a local file using the FromFile method.
// The Image can be purged from the local cache (but not from the list)
// using the PurgeLocal method.
type Image interface {
	// K8sVersion returns the version of Kubernetes components in the image.
	K8sVersion() string
	// Status can be Notdownloaded, Downloaded or Unknown.
	Status() ImageStatus
	// Deprecated returns true if the image is no longer supported.
	Deprecated() bool

	// Fetch downloads the image from the driver repository into the local cache.
	Fetch() error
	// FetchWithProgress downloads the image from the driver repository into the
	// local cache, and reports progress via the supplied callback. The callback
	// reports current and total in bytes.
	FetchWithProgress(progress func(current int64, total int64)) error
	// FromFile imports the image from the local filesystem into the local cache.
	FromFile(filepath string) error
	// PurgeLocal removes the image from the local cache.
	PurgeLocal() error
}
