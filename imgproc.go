// imgproc package that provides some simple and useful image processing functions.
package imgproc

import (
	"errors"
	"github.com/disintegration/imaging"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

// IsGrayscale : check if an image is in grayscale.
func IsGrayscale(img image.Image) bool {
	// Gets the width and height of the image
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	if w == 0 || h == 0 {
		return false
	}

	// Verifies each pixel (R,G,B)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r != g && g != b {
				return false
			}
		}
	}

	return true
}

// ToGrayscale : convert an image to grayscale.
func ToGrayscale(img image.Image) image.Image {
	// Get the image bounds
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	// Return the new image in grayscale
	return gray
}

// SaveImage : save an image (Image) to an image file.
func SaveImage(img image.Image, filePath string) error {
	// Creates the file based on the file path
	out, err := os.Create(filePath)

	// Writes the image (img) to out in PNG format
	err = png.Encode(out, img)

	if err != nil {
		return err
	}

	return nil
}

// ConvertImageToSlice : convert the image to grayscale and then convert it to a integer slice.
func ConvertImageToSlice(img image.Image) []uint8 {

	var imageSlice []uint8

	// Get the image bounds
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, img.At(x, y))
			imageSlice = append(imageSlice, gray.GrayAt(x, y).Y)
		}
	}

	return imageSlice
}

// LoadImage : load an image file based on the file path.
func LoadImage(filePath string) (image.Image, error) {
	// Open the file image
	fImage, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	// Ensure that the image file will be closed
	defer fImage.Close()

	// Convert it to an image "object"
	img, _, err := image.Decode(fImage)

	if err != nil {
		return nil, err
	}

	return img, nil
}

// LoadMultipleImages : load multiple images based on a slice of image paths.
func LoadMultipleImages(paths []string) ([]image.Image, error) {

	// If the paths slice is empty
	if len(paths) == 0 {
		return nil, errors.New("The slice is empty")
	}

	// Slice that will store the images loaded
	var images []image.Image

	// For each image path tries to load the image
	for index := 0; index < len(paths); index++ {
		// Try to load the image
		image, err := LoadImage(paths[index])

		// If something went wrong return an error
		if err != nil {
			return nil, err
		}

		// Append the image loaded to the images slice
		images = append(images, image)
	}

	// Return the images slice
	return images, nil
}

// ResizeImage : resize the image based on the parameters (maximum size 10.000x10.000).
func ResizeImage(img image.Image, width int, height int) (image.Image, error) {
	if width == 0 || width > 10000 {
		return nil, errors.New("Invalid width")
	}
	if height == 0 || height > 10000 {
		return nil, errors.New("Invalid height")
	}

	return imaging.Resize(img, width, height, imaging.Lanczos), nil
}
