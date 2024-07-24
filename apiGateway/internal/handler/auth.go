package handler

import (
	"context"
	"net/http"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/yakuzzaa/GoDone/apiGateway/internal/serializer"
	auth "github.com/yakuzzaa/GoDone/backendService/grpc/pkg/auth_v1"
)

const tokenTTL = 24 * time.Hour

// sign-in handles the user sign-in process.
// @Summary User sign-in
// @Description Sign in a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   request body  serializer.SignInRequest true "SignInRequest"
// @Success 200 {object} serializer.SignInResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /auth/sign-in [post]
func (h *ApiHandler) signIn(c *gin.Context) {
	var req auth.SignInRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("signIn request received", slog.String("username", req.Info.Username))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authClient.SignIn(ctx, &req)
	if err != nil {
		h.logger.Error("failed to sign in", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("signIn request successful", slog.Uint64("userID", resp.Id))

	c.JSON(http.StatusOK, serializer.SignInResponse{
		UserID: resp.Id})
}

// sign-up handles the user sign-up process.
// @Summary User sign-up
// @Description Sign up a user and get token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   request body  serializer.SignUpRequest true "SignInRequest"
// @Success 200 {object} serializer.SignUpResponse
// @Failure 400 {object} serializer.ErrorResponse
// @Failure 500 {object} serializer.ErrorResponse
// @Router /auth/sign-up [post]
func (h *ApiHandler) signUp(c *gin.Context) {
	var req auth.SignUpRequest

	if err := c.BindJSON(&req); err != nil {
		h.logger.Error("failed to bind JSON", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("signUp request received", slog.String("username", req.Info.Username))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := h.authClient.SignUp(ctx, &req)
	if err != nil {
		h.logger.Error("failed to sign up", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, serializer.ErrorResponse{Error: err.Error()})
		return
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    resp.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   int(tokenTTL.Seconds()),
	}

	http.SetCookie(c.Writer, cookie)

	h.logger.Info("signUp request successful", slog.String("Token", resp.Token))

	c.JSON(http.StatusOK, serializer.SignUpResponse{Token: resp.Token})
}
