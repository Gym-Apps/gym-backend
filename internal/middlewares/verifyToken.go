package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Gym-Apps/gym-backend/internal/config/db"
	jwtPackage "github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/Gym-Apps/gym-backend/internal/utils/response"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenUser := c.Get("user").(*jwt.Token)
		if !tokenUser.Valid {
			return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, "Oturum Bulunamadı.1"))
		}
		claims, ok := tokenUser.Claims.(*jwtPackage.JwtCustomClaims)
		if !ok {
			return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, "Oturum Bulunamadı."))
		}
		verifiedUser := &models.User{}

		db.DB.Debug().
			Where("id=?", claims.ID).
			First(&verifiedUser)
		if verifiedUser.ID == 0 {
			fmt.Println("girdi 3")
			return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, "Oturum Bulunamadı."))
		}

		c.Set("verifiedUser", verifiedUser)
		return next(c)
	}
}
