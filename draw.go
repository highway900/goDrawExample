package main

import (
    "os"
    "fmt"
    "image"
    "image/png"
    "image/color"
)

func main() {

    // Allocate Image Buffer Memory
    x := 32
    y := 32
    canvasSize := image.Rect(0, 0, x, y)
    img := image.NewRGBA(canvasSize)

    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            v := color.Black
            if i % 2 == 1 && j % 2 == 1 {
                v = color.White
            }
            img.Set(i, j, v)
        }
    }

    // Create Output File
    output := "output.png"

    outputFile, err := os.Create(output)
    if err != nil {
            fmt.Println("ERROR: Can't create ouput file.")
            os.Exit(1)
    }

    if err = png.Encode(outputFile, img); err != nil {
        fmt.Println("ERROR: cannot encode PNG file.")
        os.Exit(1)
    }

    outputFile.Close()

    fmt.Println("Image created: ", output)
}
