package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (m Permission) Upsert(ctx *gin.Context, item Permission) (int, Permission, error) {
	var oldData *Permission
	httpStatus, action := 201, "POST"

	var setClauseColumns = []string{"name", "description", "updated_at"}

	if item.UUID != "" {
		var tmp Permission
		if err := db.NewSelect().Model(&tmp).Where("uuid = ?", item.UUID).Scan(ctx); err == nil {
			httpStatus, action, oldData = 200, "PUT", &tmp
		}
	}

	setClause := parseSetClause(setClauseColumns)
	err := executeTransaction(ctx, func(trx *bun.Tx) error {
		_, err := trx.NewInsert().Model(&item).On("CONFLICT (uuid) DO UPDATE").Set(setClause).Exec(ctx)
		return err
	})

	go auditLog(ctx, oldData, item, item.ID, "permission", action, err)
	return httpStatus, item, err
}

func (m Permission) Read(qp QueryParams) (res PermissionResults, err error) {
	var coalesceCols = []string{"name"}
	var allowedSortFields = map[string]bool{"name": true}

	q := db.NewSelect()

	if qp.UUID != "all" {
		return res, q.Model(&res.Permission).Where("uuid = ?", qp.UUID).Scan(qp.Ctx)
	}

	q = sanitizeQuery(q.Model(&res.Permissions), qp, coalesceCols, allowedSortFields)
	res.Count, err = q.ScanAndCount(qp.Ctx)
	return res, err
}

func (m Permission) Delete(ctx *gin.Context, uuid string) (deletedAt time.Time, msg string, err error) {
	id, deletedAt, msg, err := softDelete(ctx, "permissions", uuid)

	go auditLog(ctx, nil, map[string]string{"deleted_at": deletedAt.String()}, id, "permission", "DELETE", err)
	return
}
