package utils

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"os"
)

func UploadToCloudinary(filePath string) (string, error) {
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	uploadResult, err := cld.Upload.Upload(context.Background(), filePath, uploader.UploadParams{
		Folder: "belio",
	})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}
