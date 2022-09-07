package api

import (
	"github.com/labstack/echo/v4"
	"github.com/VJ-Vijay77/redis_cache_p1/controllers"
)


func API(e *echo.Echo) {

	e.POST("/points",controllers.Points)

}