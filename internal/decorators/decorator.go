package decorators

import "github.com/gin-gonic/gin"

type DecorateFunc func(fn gin.HandlerFunc) gin.HandlerFunc
