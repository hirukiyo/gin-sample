package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hirukiyo/gin-sample/apiserver/applog"
	"github.com/rs/xid"
	"github.com/samber/lo"
)

const maxLogValueLength = 256

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := xid.New().String()

		// extracting request body
		body := []byte{}
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			body, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

		applog.Info(c, "REQ",
			"id", id,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", c.Request.URL.RawQuery,
			"client_ip", c.ClientIP(),
			"real_ip", c.GetHeader("X-Real-IP"),
			"forward-for", c.GetHeader("X-Forwarded-For"),
			"host", c.Request.Host,
			"user_agent", c.GetHeader("User-Agent"),
			"referer", c.GetHeader("Referer"),
			"body", string(body),
		)

		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)

		applog.Info(c, "RES",
			"id", id,
			"status", c.Writer.Status(),
			"size", c.Writer.Size(),
			"latency", slog.AnyValue(latency),
		)
	}
}

// Middleware to log requests
func RequestLoggingMiddleware(maskKeys []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody any
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore the body for the handler

		contentType := c.GetHeader("Content-Type")
		if strings.Contains(contentType, "application/json") {
			if err := json.Unmarshal(bodyBytes, &requestBody); err == nil {
				requestBody = maskAndTruncateJSON(requestBody, maskKeys)
			} else {
				requestBody = string(bodyBytes)
			}
		} else {
			requestBody = string(bodyBytes)
		}

		log := fmt.Sprintf("Method: %s, Path: %s, Query: %s, Body: %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.URL.RawQuery,
			requestBody,
		)
		fmt.Println(log)

		c.Next()
	}
}

func maskAndTruncateJSON(data any, maskKeys []string) any {
	switch v := data.(type) {
	case map[string]any:
		result := make(map[string]any)
		for key, value := range v {
			if lo.Contains(maskKeys, key) {
				result[key] = "***MASKED***"
			} else {
				result[key] = maskAndTruncateJSON(value, maskKeys)
			}
		}
		return result
	case []any:
		for i, value := range v {
			v[i] = maskAndTruncateJSON(value, maskKeys)
		}
		return v
	case string:
		if len(v) > maxLogValueLength {
			return v[:maxLogValueLength]
		}
		return v
	default:
		return v
	}
}
