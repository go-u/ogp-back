package media

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"server/initialize/gcs"
)

func UploadImageGCS(BucketName string, img_bytes []byte, path string, add_extension bool, width uint, height uint, cache_time int) (*string, error) {
	// decode bytes to image
	img, mime, err := DecodeBytesToImage(img_bytes)
	if err != nil {
		return nil, err
	}
	// resize
	if width != 0 && height != 0 {
		img, err = ResizeImage(img, width, height)
		if err != nil {
			return nil, err
		}
	}
	// make write buffer
	buffer, err := EncodeImageToBuffer(img, mime)
	if err != nil {
		return nil, err
	}
	// build path
	var extension string
	if mime == "image/png" {
		extension = ".png"
	} else if mime == "image/jpeg" {
		extension = ".jpg"
	} else {
		return nil, errors.New("nor png/jpeg")
	}
	if add_extension {
		path = path + extension
	}
	// upload
	Handler := gcs.Client.Bucket(BucketName)
	gcs_writer := Handler.Object(path).NewWriter(context.Background())
	gcs_writer.ContentType = mime
	gcs_writer.CacheControl = fmt.Sprintf("public, max-age=%d", cache_time)
	if cache_time == 0 {
		gcs_writer.CacheControl = "no-cache, max-age=0"
	}
	gcs_writer.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}
	if _, err = io.Copy(gcs_writer, buffer); err != nil {
		return nil, err
	}
	if err := gcs_writer.Close(); err != nil {
		return nil, err
	}
	const publicURL = "https://storage.googleapis.com/%s/%s"
	url := fmt.Sprintf(publicURL, BucketName, path)
	return &url, nil
}

func DeleteObjectGCS(BucketName string, path string) error {
	Handler := gcs.Client.Bucket(BucketName)
	object := Handler.Object(path)
	if err := object.Delete(context.Background()); err != nil {
		return err
	}
	return nil
}

func DeleteDirGCS(BucketName string, dir_path string) error {
	if len(dir_path) != 0 {
		Handler := gcs.Client.Bucket(BucketName)
		objects := Handler.Objects(context.Background(), &storage.Query{Prefix: dir_path})
		for {
			object, err := objects.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			o := Handler.Object(object.Name)
			err = o.Delete(context.Background())
			if err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("wrong dir_path")
}
