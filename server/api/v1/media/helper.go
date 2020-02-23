package media

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"
)

func DecodeBase64ImageToBytes(base64_image string) ([]byte, error) {
	base64_header_index := strings.Index(base64_image, ",")
	if base64_header_index <= 0 {
		return nil, errors.New("base64 has no header")
	}

	// decode
	base64_image_without_header := base64_image[(base64_header_index + 1):]
	return base64.StdEncoding.DecodeString(base64_image_without_header)
}

func DecodeBytesToImage(image_bytes []byte) (img image.Image, mime string, err error) {
	mime = http.DetectContentType(image_bytes) // image/png, image/jpeg
	if mime == "image/png" {
		png_image, err := png.Decode(bytes.NewReader(image_bytes))
		if err != nil {
			return nil, mime, err
		}
		img = png_image
	} else if mime == "image/jpeg" {
		jpeg_image, err := jpeg.Decode(bytes.NewReader(image_bytes))
		if err != nil {
			return nil, mime, err
		}
		img = jpeg_image
	} else {
		return nil, mime, errors.New("nor png/jpeg")
	}
	return img, mime, nil
}

func EncodeImageToBuffer(img image.Image, mime string) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	if mime == "image/png" {
		err := png.Encode(buffer, img)
		if err != nil {
			return buffer, err
		}
	} else if mime == "image/jpeg" {
		err := jpeg.Encode(buffer, img, nil)
		if err != nil {
			return buffer, err
		}
	} else {
		return buffer, errors.New("nor png/jpeg")
	}
	return buffer, nil
}

func ResizeImage(img image.Image, width uint, height uint) (image.Image, error) {
	resized_img := resize.Resize(width, height, img, resize.Lanczos3)
	return resized_img, nil
}
