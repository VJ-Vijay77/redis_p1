package controllers

import (
	"github.com/VJ-Vijay77/redis_cache_p1/db"
	"github.com/labstack/echo/v4"
)

var database *db.Database

func Home(c echo.Context) error {
	return c.JSON(200,"Home Page")
}

func Points(c echo.Context) error {
	var user db.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, "mistake bind")
	}

	err := database.SaveUser(&user)
	if err != nil {
		return c.JSON(500, "500 points")
	}

	return c.JSON(200, map[string]interface{}{
		"user": user,
	})

}

func GetUser(c echo.Context) error {
	username := c.Param("username")
	user, err := database.GetUser(username)
	if err != nil {
		if err == db.ErrNil {
			return c.JSON(404,err)
		}
		return c.JSON(500,err)
	}

	return c.JSON(200, map[string]interface{}{
		"user": user,
	})
}
