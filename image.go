package drivercore

// ImageStatus can be notdownloaded, downloaded or deprecated
type ImageStatus string

// Image Statuses
const (
	ImageStatusNotDownloaded = ImageStatus("notdownloaded")
	ImageStatusDownloaded    = ImageStatus("downloaded")
	ImageStatusDeprecated    = ImageStatus("deprecated")
)

// Image describes a template from which Machines are created. A Driver
// is expected to maintain a cache of Images locally, which can be updated
// from a driver-specific source. A Image listed in this cache may have a
// status of NotDownloaded before the actual template is downloaded, or
// Downloaded after it is. The actual template can be
// downloaded from a driver-specific source using the Fetch method, or
// added to the cache from a local file using the FromFile method.
// The template (but not the image) can be purged from the local cache
// using the PurgeLocal method.
type Image interface {
	K8sVersion() string
	Type() string
	Status() ImageStatus

	Fetch() error
	FromFile(filepath string) error
	PurgeLocal() error
}
