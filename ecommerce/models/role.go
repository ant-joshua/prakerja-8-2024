package models

type Role struct {
	ID              int          `json:"id" gorm:"primaryKey"`
	Name            string       `json:"name"`
	RolePermissions []Permission `json:"role_permissions" gorm:"many2many:role_permissions;"`
}

type CreateRoleRequest struct {
	Name string `json:"name" name:"name" binding:"required"`
}

type Permission struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type RolePermission struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	RoleID       int        `json:"role_id"`
	Role         Role       `json:"role" gorm:"foreignKey:RoleID"`
	PermissionID int        `json:"permission_id"`
	Permission   Permission `json:"permission" gorm:"foreignKey:PermissionID"`
}
