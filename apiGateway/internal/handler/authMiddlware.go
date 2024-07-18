package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	signingKey = "d23ud#bGHK54hds#ci5c"
)

func (h *ApiHandler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//authHeader := c.GetHeader("Authorization")
		cookie, err := c.Request.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization is required"})
			c.Abort()
			return

		}
		token := cookie.Value
		//log.Println("authHeader", authHeader)
		//if authHeader == "" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization is required"})
		//	c.Abort()
		//	return
		//}

		//parts := strings.Split(authHeader, " ")
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		//	c.Abort()
		//	return
		//}

		//token := parts[1]
		verifiedToken, err := verifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		userId, err := userIDFromToken(verifiedToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		c.Set("userId", userId)

		c.Next()
	}
}

func verifyToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func userIDFromToken(token *jwt.Token) (uint64, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token claims")
	}

	if userIdFloat, ok := claims["userId"].(float64); ok {
		return uint64(userIdFloat), nil
	}

	return 0, fmt.Errorf("invalid token")
}
