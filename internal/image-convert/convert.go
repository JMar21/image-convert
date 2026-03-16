package convert

import (
	"fmt"
	"image"
	_ "image/jpeg"
	png "image/png"
	"os"
)

func ConvertJPGToPNG(jpgPath string) error {
	input, err := os.Open(jpgPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer input.Close()
	img, _, err := image.Decode(input)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}
	pngPath := jpgPath[:len(jpgPath)-4] + ".png"
	output, err := os.Create(pngPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer output.Close()
	err = png.Encode(output, img)
	if err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}
	return nil
}
