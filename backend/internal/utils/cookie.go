package utils

import (
    "time"
	"github.com/gofiber/fiber/v2"
    "errors"
)

//		function set cookie
func SetCookie(c *fiber.Ctx, token string){
	c.Cookie(&fiber.Cookie{
        Name:     "jwt",
        Value:    token,
        Path:     "/",
        MaxAge:   60 * 20,
        HTTPOnly: true,
        SameSite: "Lax", 
    })
}

func GetCookie(c *fiber.Ctx, name string) (string, error){
	cookie := c.Cookies(name)
	if cookie == "" {
		return "", errors.New("cookie not found")
	}
	return cookie, nil
}

func ClearCookie(c *fiber.Ctx){
    c.Cookie(&fiber.Cookie{
        Name:     "jwt",
        Value:    "",
        Path:     "/",
        Expires: time.Now().Add(-24 * time.Hour),
        MaxAge:   -1,
        HTTPOnly: true,
        SameSite: "Lax", 
    }) 
}
