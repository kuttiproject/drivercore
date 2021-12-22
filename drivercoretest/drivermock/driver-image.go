package drivermock

import (
	"errors"

	"github.com/kuttiproject/drivercore"
)

var (
	localimages  = map[string]*Image{}
	remoteimages = map[string]*Image{}
)

// UpdateRemoteImage adds or updates a mock remote image.
func (vd *Driver) UpdateRemoteImage(k8sversion string, deprecated bool) {
	remoteimages[k8sversion] = &Image{
		k8sversion: k8sversion,
		status:     drivercore.ImageStatusNotDownloaded,
		deprecated: deprecated,
	}
}

// DeleteRemoteImage deletes a mock remote image.
func (vd *Driver) DeleteRemoteImage(k8sversion string) {
	delete(remoteimages, k8sversion)
}

// UpdateImageList updates the image list
func (vd *Driver) UpdateImageList() error {
	localimages = make(map[string]*Image, len(remoteimages))
	for key, value := range remoteimages {
		localimages[key] = value
	}

	return nil
}

// ValidK8sVersion returns true if the specified k8s version
// is available in localimages, false otherwise.
func (vd *Driver) ValidK8sVersion(k8sversion string) bool {
	_, ok := localimages[k8sversion]
	return ok
}

// K8sVersions lists locally available k8s versions.
func (vd *Driver) K8sVersions() []string {
	result := make([]string, len(localimages))
	index := 0
	for _, value := range localimages {
		result[index] = value.k8sversion
		index++
	}
	return result
}

// ListImages returns the local images.
func (vd *Driver) ListImages() ([]drivercore.Image, error) {
	result := make([]drivercore.Image, len(localimages))
	index := 0
	for _, value := range localimages {
		result[index] = value
		index++
	}

	return result, nil
}

// GetImage returns the image for the specified k8s version,
// or nil.
func (vd *Driver) GetImage(k8sversion string) (drivercore.Image, error) {
	result, ok := localimages[k8sversion]
	if !ok {
		return nil, errors.New("image not available locally")
	}

	return result, nil
}
