package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(dto interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if errs := c.ShouldBindJSON(&dto); errs != nil {
			var data = make(map[string]string)
			for _, err := range errs.(validator.ValidationErrors) {
				data[strings.ToLower(err.Field())] = err.Error()
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Validation Errors!.", "errors": data})
			return
		}
		c.Set("data", dto)
		c.Next()
	}
}
