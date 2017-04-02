imgproc Package
==========================

[![Build Status](https://circleci.com/gh/kelvins/imgproc.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/kelvins/imgproc)


GoLang package that provides simple image processing functions.
Use: `go get github.com/kelvins/imgproc`

Functions:
----

- [X] IsGrayscale
- [X] ToGrayscale
- [X] ConvertImageToVector
- [X] ResizeImage
- [X] SaveImage
- [X] LoadImage

Example:
----

``` go
import (
    "fmt"
    "log"
    "github.com/kelvins/imgproc"
)

img, err := imgproc.LoadImage("./face.png")
if err != nil {
    log.Fatal(err)
}

grayscaleImage := imgproc.ToGrayscale(img)

err = imgproc.SaveImage(grayscaleImage, "./face.png")
if err != nil {
    log.Fatal(err)
}

if !imgproc.IsGrayscale(grayscaleImage) {
    fmt.Println("It is not grayscale")
}

imageSlice := imgproc.ConvertImageToVector(grayscaleImage)

resizedImage, err = imgproc.ResizeImage(grayscaleImage, 100, 100)
if err != nil {
    log.Fatal(err)
}
```

**Note**: this package uses the `github.com/disintegration/imaging` package to resize the image
