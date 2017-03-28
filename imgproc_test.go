
package imgproc_test

import (
    "testing"
    "github.com/kelvins/imgproc"
)

// Test the load image function
func TestLoadImage(t *testing.T) {
    // Try to load an invalid image
    _, err := imgproc.LoadImage("123")

    if err == nil {
        t.Error(err)
    }

    // Try to load an valid image
    _, err = imgproc.LoadImage("./face.png")

    if err != nil {
        t.Error(err)
    }
}

// Test the SaveImage function
func TestSaveImage(t *testing.T) {

    img, err := imgproc.LoadImage("./face.png")

    if err != nil {
        t.Error(err)
    }

    err = imgproc.SaveImage(img, "./face.png")

    if err != nil {
        t.Error(err)
    }
}

// Test the grayscale function
func TestGrayscale(t *testing.T) {

    img, err := imgproc.LoadImage("./face.png")

    if err != nil {
        t.Error(err)
    }

    if imgproc.IsGrayscale(img) {
        t.Error("Image is grayscale")
    }

    newImage := imgproc.ToGrayscale(img)

    if !imgproc.IsGrayscale(newImage) {
        t.Error("Image is not grayscale")
    }
}

// Test the ConvertImageToVector function
func TestConvertImageToVector(t *testing.T) {

    img, err := imgproc.LoadImage("./face.png")

    if err != nil {
        t.Error(err)
    }

    imageSlice := imgproc.ConvertImageToVector(img)

    expectedImageSlice := [6]uint8{163, 163, 162, 160, 157, 157}

    for i := 0; i < len(expectedImageSlice); i++ {
        if imageSlice[i] != expectedImageSlice[i] {
            t.Error(err)
        }
    }

    if len(imageSlice) != (256*256) {
        t.Error("Wrong length")
    }
}

// Test the resize image function
func TestResizeImage(t *testing.T) {
    img, err := imgproc.LoadImage("./face.png")

    if err != nil {
        t.Error(err)
    }

    b := img.Bounds()
    w, h := b.Max.X, b.Max.Y

    if w != 256 || h != 256 {
        t.Error("Wrong bounds")
    }

    imgResize, err := imgproc.ResizeImage(img, 0, 0)

    if err == nil {
        t.Error("Error during the resize image")
    }

    imgResize, err = imgproc.ResizeImage(img, 9999999, 9999999)

    if err == nil {
        t.Error("Error during the resize image")
    }

    imgResize, err = imgproc.ResizeImage(img, 100, 100)

    b = imgResize.Bounds()
    w, h = b.Max.X, b.Max.Y

    if w != 100 || h != 100 {
        t.Error("Wrong bounds")
    }

    imgResize, err = imgproc.ResizeImage(img, 10, 10)

    b = imgResize.Bounds()
    w, h = b.Max.X, b.Max.Y

    if w != 10 || h != 10 {
        t.Error("Wrong bounds")
    }
}