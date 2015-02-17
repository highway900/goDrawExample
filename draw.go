package main

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"
    "math/rand"
    "time"
)

const FREE uint32 = 0
const SOLID uint32 = 1

type Level struct {
    Width uint32
    Height uint32
    m [][]uint32
}

func (lvl *Level) Get(i uint32, j uint32) uint32 {
    return lvl.m[i][j]
}

func (lvl *Level) Set(i uint32, j uint32, c uint32) {
    lvl.m[i][j] = c
}

// Simply draw a border around the image
// Fill with white if fill is set to true
func (lvl *Level) drawBorder() {
    for i, row := range lvl.m {
        for j := range row {
            if i == 0 || j == 0 || i == int(lvl.Width)-1 || j == int(lvl.Height)-1 {
                // draw border
                lvl.m[i][j] = SOLID
            } else {
                lvl.m[i][j] = FREE
            }
        }
    }
}

func CreateLevel(XSize uint32, YSize uint32) *Level {
    lvl := &Level{
        Width: XSize,
        Height: YSize,
    }
    // Allocate the top-level slice, the same as before.
    lvl.m = make([][]uint32, XSize) // One row per unit of y.
    // Allocate one large slice to hold all the row.
    for i := range lvl.m {
        lvl.m[i] = make([]uint32, YSize)
    }
    return lvl
}


func (lvl *Level) RandomTrees(n uint32, seed int64) {
    rand.Seed(seed)
    for i := 0; i < int(n); i++  {
        home := true
        for home == true {
            x := uint32(rand.Int31n(int32(lvl.Width)))
            y := uint32(rand.Int31n(int32(lvl.Height)))
            if lvl.Get(x, y) == FREE {
                lvl.Set(x, y, SOLID)
                home = false
            }
        }
    }
}


func (lvl *Level) DrawImage() {
    // Allocate Image Buffer Memory
    x := int(lvl.Width)
    y := int(lvl.Height)
    canvasSize := image.Rect(0, 0, x, y)
    img := image.NewRGBA(canvasSize)

    c := color.Black
    for i, row := range lvl.m {
        for j := range row {
            switch lvl.m[i][j] {
                case FREE:
                    c = color.White
                case SOLID:
                    c = color.Black
            }
            img.Set(i, j, c)
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

func main() {

    level := CreateLevel(16, 32)
    level.drawBorder()
    level.RandomTrees(300, time.Now().Unix())
    level.DrawImage()

}
