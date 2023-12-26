package order_status_bar

import (
	"fmt"
	"github.com/shulst/go-kraken-safe-trader/order-status-bar/config"
	progress_bar "github.com/shulst/go-kraken-safe-trader/progress-bar"
	"image/png"
	"os"
	"testing"
)

func TestOrderStatusBar_Buy(t *testing.T) {
	cfg := config.FromEnv("../.env")
	buy := BuyColors{cfg: cfg}
	progressBar := progress_bar.Create(100)
	progressBar = progressBar.Fill(17, 18)

	orderStatusBar := Create(buy, progressBar, Buy)
	fmt.Printf("Order status bar: %v\n", orderStatusBar)

	img := orderStatusBar.Draw(200, 20)

	newPngFile := "./two_rectangles.png" // output image will live here
	myfile, err := os.Create(newPngFile) // ... now lets save output image
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	png.Encode(myfile, img) // output file /tmp/two_rectangles.png
}
