package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type (
	AppModel struct {
		Active    bool      `bun:"active,default:true" json:"active"`
		Status    string    `bun:"status,default:O" json:"status"`
		Flag      string    `bun:"flag,nullzero,default:null" json:"flag,omitempty"`
		UUID      string    `bun:"uuid,default:gen_random_uuid()" json:"uuid,omitempty"`
		CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at,omitzero"`
		CreatedBy int64     `bun:"created_by,default:0" json:"created_by,omitzero"`
		UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at,omitzero"`
		UpdatedBy int64     `bun:"updated_by,default:0" json:"updated_by,omitzero"`
		DeletedAt time.Time `bun:"deleted_at,nullzero,default:null" json:"deleted_at,omitzero"`
		DeletedBy int64     `bun:"deleted_by,default:0" json:"deleted_by,omitzero"`
	}

	QueryParams struct {
		UUID        string
		Filter      string
		FilterExtOp string
		FilterExt   string
		Sort        string
		Status      string
		Deleted     bool
		Limit       int
		Page        int
		Ctx         *gin.Context
	}

	AuditLog struct {
		bun.BaseModel `bun:"table:audit_logs,alias:au"`

		ID               int64     `bun:"id,pk,autoincrement" json:"id"`
		UserID           int64     `bun:"user_id" json:"user_id"`
		Token            string    `bun:"token,default:null" json:"token,omitempty"`
		Path             string    `bun:"path" json:"path"`
		Action           string    `bun:"action" json:"action"`
		ResponseStatus   int       `bun:"response_status" json:"response_status"`
		ModuleID         int64     `bun:"module_id" json:"module_id"`
		Module           string    `bun:"module" json:"module"`
		BeforeDataChange any       `bun:"before_data_change" json:"before_data_change"`
		AfterDataChange  any       `bun:"after_data_change" json:"after_data_change"`
		Description      string    `bun:"description,default:null" json:"description"`
		IPAddress        string    `bun:"ip_address" json:"ip_address"`
		UserAgent        string    `bun:"user_agent" json:"user_agent"`
		CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at,omitzero"`
	}

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
