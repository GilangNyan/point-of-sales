package middleware

import "github.com/gin-gonic/gin"

func GetUserID(ctx *gin.Context) string {
	if userID, exists := ctx.Get("userID"); exists {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}

func GetUsername(ctx *gin.Context) string {
	if username, exists := ctx.Get("username"); exists {
		if name, ok := username.(string); ok {
			return name
		}
	}
	return ""
}

func GetUserEmail(ctx *gin.Context) string {
	if userMail, exists := ctx.Get("email"); exists {
		if mail, ok := userMail.(string); ok {
			return mail
		}
	}
	return ""
}

func GetUserRoles(ctx *gin.Context) []string {
	if roles, exists := ctx.Get("roles"); exists {
		if r, ok := roles.([]string); ok {
			return r
		}
	}
	return []string{}
}

func HasRole(ctx *gin.Context, role string) bool {
	roles := GetUserRoles(ctx)
	for _, userRole := range roles {
		if userRole == role {
			return true
		}
	}
	return false
}
