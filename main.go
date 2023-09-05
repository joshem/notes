package notes

import (
	"github.com/labstack/echo"
	jww "github.com/spf13/jwalterweatherman"
)

func main() {
	e := echo.New()

	err := e.Start(":8080")
	if err != nil {
		jww.FATAL.Panicf("Failed to start server: %+v", err)
	}
}
