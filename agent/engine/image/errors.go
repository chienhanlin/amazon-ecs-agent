package image

import (
	"fmt"
)

type ImageError struct {
	ContainerName string
	ImageName     string
	Error         error
}

func (imageError *ImageError) String() string {
	return fmt.Sprintf("ContainerName: %s; ImageName: %s; Error: %v",
		imageError.ContainerName, imageError.ImageName, imageError.Error)
}
