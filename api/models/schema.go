package models

import (
	"github.com/uptrace/bun"
)

type (
	UserGroupResults struct {
		UserGroup  UserGroup
		UserGroups []UserGroup
		Count      int
	}

	UserGroup struct {
		bun.BaseModel `bun:"table:user_groups,alias:ug"`

		ID      int64 `bun:"id,pk,autoincrement" json:"id"`
		UserID  int64 `bun:"user_id" json:"user_id"`
		GroupID int64 `bun:"group_id" json:"group_id"`

		AppModel
	}

	GroupPermissionResults struct {
		GroupPermission  GroupPermission
		GroupPermissions []GroupPermission
		Count            int
	}

	GroupPermission struct {
		bun.BaseModel `bun:"table:group_permissions,alias:gp"`

		ID           int64 `bun:"id,pk,autoincrement" json:"id"`
		GroupID      int64 `bun:"group_id" json:"group_id"`
		PermissionID int64 `bun:"permission_id" json:"permission_id"`

		AppModel
	}
)
