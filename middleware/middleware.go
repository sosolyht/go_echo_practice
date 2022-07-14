package middleware

import "github.com/labstack/echo/v4/middleware"

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("dWkrCnyQr2w8e7lmnLkapQtMy0LKDMM4"),
})
