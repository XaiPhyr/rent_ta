package models

import (
	"api/utils"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type (
	AppModel struct {
		ID        int64     `bun:"id,pk,autoincrement" json:"id"`
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
)

var db = utils.InitDB()

func sanitizeQuery(q *bun.SelectQuery, cols []string, filter, sortFields, status string, validateField func(string) bool) *bun.SelectQuery {
	if len(cols) > 0 {
		var cls []string
		for _, col := range cols {
			cls = append(cls, "coalesce("+col+",'')")
		}

		filter = strings.ToLower(filter)
		cl := "(" + strings.Join(cls, " || ") + ")"
		if filter != "" {
			q = q.Where(cl+" ~* ?", filter)
		}
	}

	if sortFields != "" {
		sortFieldList := strings.SplitSeq(sortFields, ",")
		for sortField := range sortFieldList {
			if validateField(sortField) {
				if after, ok := strings.CutPrefix(sortField, "-"); ok {
					q = q.Order(after + " DESC")
				} else {
					q = q.Order(sortField + " ASC")
				}
			}
		}
	} else {
		q = q.Order("id ASC")
	}

	if status != "" {
		q = q.Where("status = ?", status)
	}

	return q
}

func executeTransaction(ctx context.Context, trxFunc func(*bun.Tx) error) error {
	trx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := trxFunc(&trx); err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit()
}

func auditLog(ctx *gin.Context, beforeDataChange, afterDataChange any, moduleId int64, module, action, description string) {
	v, _ := ctx.Get("userId")

	var userID int64
	id, ok := v.(int)
	if ok {
		userID = int64(id)
	}

	auditLog := &AuditLog{
		UserID:           userID,
		Token:            "",
		Path:             ctx.FullPath(),
		Action:           action,
		ResponseStatus:   ctx.Writer.Status(),
		ModuleID:         moduleId,
		Module:           module,
		BeforeDataChange: beforeDataChange,
		AfterDataChange:  afterDataChange,
		Description:      description,
		IPAddress:        ctx.ClientIP(),
		UserAgent:        ctx.Request.UserAgent(),
	}

	_, err := db.NewInsert().Model(auditLog).Exec(ctx)

	if err != nil {
		fmt.Println("AUDIT LOG ERR: ", err)
	}
}
