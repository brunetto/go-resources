package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
)

const (
	tileSize = 100
	imgSize  = 1e3
	nWkrs    = 1e2
	chanSize = nWkrs + 100
)

func main() {
	t0 := time.Now()
	defer func() { fmt.Println("done in", time.Since(t0)) }()

	// create final image container
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	// create channels and waitgroups
	in := make(chan work, chanSize)
	out := make(chan *image.RGBA, chanSize)

	wkrs, rdc := &sync.WaitGroup{}, &sync.WaitGroup{}

	// launch "reduce" closure
	rdc.Add(1)

	go func() {
		defer rdc.Done() // signal we are done on exit

		// until tiles arrive, add them to the final image
		for tile := range out {
			draw.Draw(img, tile.Bounds(), tile, tile.Bounds().Min, draw.Src)
		}
	}()

	// start workers
	for i := 0; i < nWkrs; i++ {
		wkrs.Add(1)

		go worker(wkrs, in, out)
	}

	// send work:
	// it would be not efficient to create a worker per pixel,
	// much better use a pool of workers, each for a tile containing
	// a certain number of pixels
	for x := 0; x < imgSize; x += tileSize {
		for y := 0; y < imgSize; y += tileSize {
			in <- work{
				x: x, y: y, dx: tileSize, dy: tileSize,
				// here you can change the type of resulting image by choosing
				// how to color each tile
				colorFunc: func() cf {
					if x >= y {
						return newCF(imgSize) // single image
					} else {
						return newCF(tileSize) // identical tiles
					}
				}(),
			}
		}
	}

	// close channel and wait for the goroutines to complete
	close(in)
	wkrs.Wait()

	close(out)
	rdc.Wait()

	// save final image
	err := save(img)
	if err != nil {
		log.Fatal(errors.Wrap(err, "can't save final image"))
	}
}

type cf = func(x, y int) color.RGBA

func newCF(sideSize int) cf {
	return func(x, y int) color.RGBA {
		return color.RGBA{uint8(x * 255 / sideSize), uint8(y * 255 / sideSize), 100, 255}
	}
}

type work struct {
	x, y, dx, dy int
	colorFunc    cf //func(x, y int) color.RGBA
}

func worker(wkrs *sync.WaitGroup, in chan work, out chan *image.RGBA) {
	defer wkrs.Done() // signal the goroutine is over

	for w := range in { // until we have work to do
		tile := image.NewRGBA(image.Rect(w.x, w.y, w.x+w.dx, w.y+w.dy)) // create empty tile
		// loop over the tile pixels
		for xx := w.x; xx < w.x+w.dx; xx++ {
			for yy := w.y; yy < w.y+w.dy; yy++ {
				// set pixel color using the chosen color function
				tile.SetRGBA(xx, yy, w.colorFunc(xx, yy))
			}
		}

		// send the result to the reducer
		out <- tile
	}
}

func save(img image.Image) error {
	f, err := os.Create("example.png")
	if err != nil {
		return errors.Wrap(err, "can't create final image file")
	}

	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return errors.Wrap(err, "can't encode image to png file")
	}

	return nil
}
