package wfc

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func saveImage(filepath string, img *image.NRGBA) error {
	imgFile, err := os.Create(filepath)
	defer imgFile.Close()
	if err != nil {
		return fmt.Errorf("Couldn't create file: %w", err)
	}
	return png.Encode(imgFile, img.SubImage(img.Rect))
}

func loadImage(filePath string) (*image.NRGBA, error) {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		return nil, fmt.Errorf("Couldn't open file: %w", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't decode image: %w", err)
	}
	return img.(*image.NRGBA), nil
}
