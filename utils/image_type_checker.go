package utils

import (
	"bytes"
	"errors"
	"io"
	"log"
)

type MediaType string

const (
	ImageJPEG MediaType = "image/jpeg"
	ImagePNG  MediaType = "image/png"
	ImageGIF  MediaType = "image/gif"
	TextPlain MediaType = "text/plain"
)

func CheckImageType(imageBytes []byte) MediaType {
	reader := bytes.NewReader(imageBytes)

	header := make([]byte, 8)
	_, err := io.ReadFull(reader, header)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		log.Println("Error Check Image =>", err)
		return TextPlain
	}

	switch {
	case len(header) >= 2 && header[0] == 0xFF && header[1] == 0xD8:
		return ImageJPEG
	case len(header) >= 8 && header[0] == 0x89 && header[1] == 'P' && header[2] == 'N' &&
		header[3] == 'G' && header[4] == 0x0D && header[5] == 0x0A && header[6] == 0x1A && header[7] == 0x0A:
		return ImagePNG
	case len(header) >= 6 && header[0] == 'G' && header[1] == 'I' && header[2] == 'F' &&
		header[3] == '8' && (header[4] == '7' || header[4] == '9') && header[5] == 'a':
		return ImageGIF
	default:
		return ImageJPEG
	}
}
