package middleware

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PermissionMiddleware interface {
	PermissionMiddleware(permissions ...string) gin.HandlerFunc
	PermissionMiddlewareMap(permissions ...string) gin.HandlerFunc
}

type PermissionMiddlewareImpl struct {
	db         *gorm.DB
	redisCache *redis.Client
}

func NewPermissionMiddleware(db *gorm.DB, redisCache *redis.Client) PermissionMiddleware {
	return &PermissionMiddlewareImpl{
		db:         db,
		redisCache: redisCache,
	}
}

func (p *PermissionMiddlewareImpl) PermissionMiddleware(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user from context
		user := c.MustGet("user").(jwt.MapClaims)

		// get user id
		role_id := int(user["role_id"].(float64))

		// check user permission
		var rolePermissions []models.RolePermission

		err := p.db.
			Joins("Role").
			Joins("Permission").
			Where("role_id = ?", role_id).
			Find(&rolePermissions).Error

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			c.Abort()
			return
		}

		// check if user has permission
		for _, permission := range permissions {
			for _, userPerm := range rolePermissions {
				if permission == userPerm.Permission.Name {
					c.Next()
					return
				}
			}
		}

		c.JSON(403, gin.H{
			"message": "Forbidden",
		})
		c.Abort()
	}

}

func (p *PermissionMiddlewareImpl) PermissionMiddlewareMap(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// get user from context
		user := c.MustGet("user").(jwt.MapClaims)

		// get user id
		role_id := int(user["role_id"].(float64))

		var mapPermissions map[string]bool = make(map[string]bool)

		// check user permission
		var role models.Role

		parseRoleID := fmt.Sprintf("%d", role_id)

		result, err := p.redisCache.Get(ctx, "role:"+parseRoleID).Result()

		if err != nil {

			fmt.Println(err.Error())

			err := p.db.
				Preload("RolePermissions").
				Where("id = ?", role_id).
				Find(&role).Error

			if err != nil {
				c.JSON(500, gin.H{
					"message": "Internal server error",
				})
				c.Abort()
				return
			}

			for _, rolePermission := range role.RolePermissions {
				mapPermissions[rolePermission.Name] = true
			}

			fmt.Printf("Map Permissions: %+v\n", mapPermissions)

			jsonMapPermissions, err := json.Marshal(mapPermissions)

			if err != nil {
				c.JSON(500, gin.H{
					"message": "Internal server error",
				})
				c.Abort()
				return
			}

			err = p.redisCache.Set(ctx, "role:"+parseRoleID, string(jsonMapPermissions), 0).Err()

			if err != nil {
				c.JSON(500, gin.H{
					"message": "Internal server error",
				})
				c.Abort()
				return
			}
		}

		jsonMap := json.Unmarshal([]byte(result), &mapPermissions)

		if jsonMap != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			c.Abort()
			return
		}

		fmt.Printf("Map Permissions from unmarshal: %+v\n", mapPermissions)

		// check if user has permission
		for _, permission := range permissions {
			if mapPermissions[permission] {
				c.Next()
				return
			}
		}

		c.JSON(403, gin.H{
			"message": "Forbidden",
		})
		c.Abort()
	}

}
