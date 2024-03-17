package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SanitizeMiddleware is a middleware function to sanitize incoming request data
func SanitizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
				return
			}
			defer c.Request.Body.Close()

			sanitizedBody, err := sanitizeJSONInput(body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error sanitizing JSON data"})
				return
			}

			c.Request.Body = io.NopCloser(bytes.NewBuffer(sanitizedBody))
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
			return
		}

		c.Next()
	}
}

func sanitizeJSONInput(input []byte) ([]byte, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(input, &jsonData)
	if err != nil {
		return nil, err
	}

	for key, value := range jsonData {
		if strVal, ok := value.(string); ok {
			sanitizedValue := SanitizeString(strVal)
			jsonData[key] = sanitizedValue
		}
	}

	sanitizedData, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return sanitizedData, nil
}

func SanitizeString(input string) string {
	// Sanitize for SQL injection by replacing single quotes
	sanitizedString := strings.ReplaceAll(input, "'", "''")

	// Sanitize for CRLF injection by removing newline characters
	sanitizedString = strings.ReplaceAll(sanitizedString, "\r", "")
	sanitizedString = strings.ReplaceAll(sanitizedString, "\n", "")

	return sanitizedString
}

/* package middlewares
//In Case I decide to use http >
import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// SanitizeMiddleware is a middleware function to sanitize incoming request data
func SanitizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/json" {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			sanitizedBody, err := sanitizeJSONInput(body)
			if err != nil {
				http.Error(w, "Error sanitizing JSON data", http.StatusBadRequest)
				return
			}

			r.Body = io.NopCloser(bytes.NewBuffer(sanitizedBody))
		} else {
			http.Error(w, "something went wrong", http.StatusBadRequest) //if content-type != app/json we'll show this vague error (Less information, better)

		}

		next.ServeHTTP(w, r)
	})
}

func sanitizeJSONInput(input []byte) ([]byte, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(input, &jsonData)
	if err != nil {
		return nil, err
	}

	for key, value := range jsonData {
		if strVal, ok := value.(string); ok {
			sanitizedValue := SanitizeString(strVal)
			jsonData[key] = sanitizedValue
		}
	}

	sanitizedData, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return sanitizedData, nil
}

func SanitizeString(input string) string {
	// Sanitize for SQL injection by replacing single quotes
	sanitizedString := strings.ReplaceAll(input, "'", "''")

	// Sanitize for CRLF injection by removing newline characters
	sanitizedString = strings.ReplaceAll(sanitizedString, "\r", "")
	sanitizedString = strings.ReplaceAll(sanitizedString, "\n", "")

	return sanitizedString
}
*/
