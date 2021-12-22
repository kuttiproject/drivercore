package drivermock

import "github.com/kuttiproject/drivercore"

// Image implements the drivercore.Image interface.
type Image struct {
	k8sversion string
	status     drivercore.ImageStatus
	deprecated bool
}

// K8sVersion returns the k8s version of the image.
func (img *Image) K8sVersion() string {
	return img.k8sversion
}

// Status returns the status of the image.
func (img *Image) Status() drivercore.ImageStatus {
	return img.status
}

// Deprecated returns the deprecation status of the image.
func (img *Image) Deprecated() bool {
	return img.deprecated
}

// Fetch fetches an image.
func (img *Image) Fetch() error {
	img.status = drivercore.ImageStatusDownloaded
	return nil
}

// FromFile fetches an image from a file.
func (img *Image) FromFile(filepath string) error {
	return img.Fetch()
}

// PurgeLocal removes the local copy of an image.
func (img *Image) PurgeLocal() error {
	img.status = drivercore.ImageStatusNotDownloaded
	return nil
}
