package router

import (
	"github.com/VJ-Vijay77/redis_cache_p1/db"
	"github.com/labstack/echo/v4"
)



func InitRouter(database *db.Database) *echo.Echo {
	r := echo.New()
	return r
}