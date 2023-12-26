package router

import (
	"github.com/gin-gonic/gin"
	order_status_bar "github.com/shulst/go-kraken-safe-trader/order-status-bar"
	progress_bar "github.com/shulst/go-kraken-safe-trader/progress-bar"
	"image/png"
	"strconv"
	"strings"
)

func Router() *gin.Engine {
	router := gin.New()
	router.GET("/:type/:percentOff/:percentFilled/:dimensions", Handle)
	return router
}

func DecodeDimensions(d string) (int, int) {
	s := strings.Split(strings.ToLower(string(d)), "x")
	width, _ := strconv.ParseInt(s[0], 10, 32)
	height, _ := strconv.ParseInt(s[1], 10, 32)
	return int(width), int(height)
}

func Handle(c *gin.Context) {
	orderType := order_status_bar.OrderType(c.Param("type"))

	off, err1 := strconv.ParseInt(c.Param("percentOff"), 10, 32)
	filled, err2 := strconv.ParseInt(c.Param("percentFilled"), 10, 32)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			c.AbortWithError(400, err1)
		} else {
			c.AbortWithError(400, err2)
		}
		return
	}

	progressBar := progress_bar.Create(100)
	progressBar = progressBar.Fill(progress_bar.Fill(off), progress_bar.Fill(filled))

	orderStatusBar := order_status_bar.Create(progressBar, orderType)

	img := orderStatusBar.Draw(DecodeDimensions(c.Param("dimensions")))

	c.Writer.Header().Set("Content-Type", "image/png")

	err := png.Encode(c.Writer, img)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

}
