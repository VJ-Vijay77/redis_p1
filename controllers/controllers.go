package controllers

import (
	"github.com/VJ-Vijay77/redis_cache_p1/db"
	"github.com/labstack/echo/v4"
)

var database *db.Database

func Points(c echo.Context) error {
	var user db.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, err)
	}

	err := database.SaveUser(&user)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, map[string]interface{}{
		"user": user,
	})

}
