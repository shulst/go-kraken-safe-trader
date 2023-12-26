package order_status_bar

import (
	"fmt"
	"github.com/shulst/go-kraken-safe-trader/order-status-bar/config"
	"testing"
)

func TestCombineColors(t *testing.T) {
	cfg := config.FromEnv("../.env")
	buyBg := cfg.BuyBgColor.RGBA()
	buyFg := cfg.BuyFgColor.RGBA()
	sellBg := cfg.SellBgColor.RGBA()
	sellFg := cfg.SellFgColor.RGBA()
	overlay := cfg.OverlayColor.RGBA()

	fmt.Println("======= Buying colors =======")
	fmt.Printf("Buy BG: %v\nBuy FG: %v\nOverlay: %v\n", buyBg, buyFg, overlay)
	fmt.Println("Blending ...")
	fmt.Printf("Bg + Fg: %v\n", CombineColors(buyBg, buyFg))
	fmt.Printf("Bg + Overlay: %v\n", CombineColors(buyBg, overlay))
	fmt.Printf("Bg + Fg + Overlay: %v\n", CombineColors(buyBg, buyFg, overlay))

	fmt.Println("======= Selling colors =======")
	fmt.Printf("Sell Bg: %v\nSell Fg: %v\nOverlay: %v\n", sellBg, sellFg, overlay)
	fmt.Println("Blending ...")
	fmt.Printf("Bg + Fg: %v\n", CombineColors(sellBg, sellFg))
	fmt.Printf("Bg + Overlay: %v\n", CombineColors(sellBg, overlay))
	fmt.Printf("Bg + Fg + Overlay: %v\n", CombineColors(sellBg, sellFg, overlay))
}
