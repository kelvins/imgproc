package imgproc

import (
	"testing"
)

// Test the load image function
func TestLoadImage(t *testing.T) {
	// Try to load an invalid image
	_, err := LoadImage("123")

	if err == nil {
		t.Error("Expected an error. The path is invalid.")
	}

	// Try to load an valid image
	_, err = LoadImage("./face.png")

	if err != nil {
		t.Error(err)
	}
}

// Test the load multiple images function
func TestLoadMultipleImages(t *testing.T) {
	var paths []string

	// Paths empty
	_, err := LoadMultipleImages(paths)

	if err == nil {
		t.Error("Expected an error. The paths slice is empty.")
	}

	// Invalid path
	paths = append(paths, "123.txt")
	images, err := LoadMultipleImages(paths)

	if images != nil {
		t.Error("Expected images == nil")
	}
	if err == nil {
		t.Error("Expected an error. The path is invalid.")
	}

	// Valid paths
	paths = nil
	paths = append(paths, "./face.png")
	paths = append(paths, "./face.png")

	images, err = LoadMultipleImages(paths)

	if images == nil {
		t.Error("Expected a valid images vector")
	}
	if err != nil {
		t.Error("Expected no error.")
	}
}

// Test the SaveImage function
func TestSaveImage(t *testing.T) {

	img, err := LoadImage("./face.png")

	if err != nil {
		t.Error(err)
	}

	err = SaveImage(img, "./face.png")

	if err != nil {
		t.Error(err)
	}
}

// Test the grayscale function
func TestGrayscale(t *testing.T) {

	img, err := LoadImage("./face.png")

	if err != nil {
		t.Error(err)
	}

	if IsGrayscale(img) {
		t.Error("Image is grayscale")
	}

	newImage := ToGrayscale(img)

	if !IsGrayscale(newImage) {
		t.Error("Image is not grayscale")
	}
}

// Test the ConvertImageToVector function
func TestConvertImageToVector(t *testing.T) {

	img, err := LoadImage("./face.png")

	if err != nil {
		t.Error(err)
	}

	imageSlice := ConvertImageToVector(img)

	expectedImageSlice := [6]uint8{163, 163, 162, 160, 157, 157}

	for i := 0; i < len(expectedImageSlice); i++ {
		if imageSlice[i] != expectedImageSlice[i] {
			t.Error(err)
		}
	}

	if len(imageSlice) != (256 * 256) {
		t.Error("Wrong length")
	}
}

// Test the resize image function
func TestResizeImage(t *testing.T) {
	img, err := LoadImage("./face.png")

	if err != nil {
		t.Error(err)
	}

	b := img.Bounds()
	w, h := b.Max.X, b.Max.Y

	if w != 256 || h != 256 {
		t.Error("Wrong bounds")
	}

	imgResize, err := ResizeImage(img, 100, 0)

	if err == nil {
		t.Error("Error during the resize image")
	}

	imgResize, err = ResizeImage(img, 0, 100)

	if err == nil {
		t.Error("Error during the resize image")
	}

	imgResize, err = ResizeImage(img, 0, 0)

	if err == nil {
		t.Error("Error during the resize image")
	}

	imgResize, err = ResizeImage(img, 9999999, 9999999)

	if err == nil {
		t.Error("Error during the resize image")
	}

	imgResize, _ = ResizeImage(img, 100, 100)

	b = imgResize.Bounds()
	w, h = b.Max.X, b.Max.Y

	if w != 100 || h != 100 {
		t.Error("Wrong bounds")
	}

	imgResize, _ = ResizeImage(img, 10, 10)

	b = imgResize.Bounds()
	w, h = b.Max.X, b.Max.Y

	if w != 10 || h != 10 {
		t.Error("Wrong bounds")
	}
}
