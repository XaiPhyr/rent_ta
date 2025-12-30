package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type (
	SpaceResults struct {
		Space  Space
		Spaces []Space
		Count  int
	}

	Space struct {
		bun.BaseModel `bun:"table:space,alias:s"`

		ID            int64           `bun:"id,pk,autoincrement" json:"id"`
		UserID        int64           `bun:"user_id" json:"user_id"`
		Name          string          `bun:"name" json:"name"`
		Description   string          `bun:"description" json:"description"`
		Address       *map[string]any `bun:"address,type:jsonb" json:"address"`
		PricePerHour  float64         `bun:"price_per_hour,nullzero,default:0" json:"price_per_hour"`
		PricePerDay   float64         `bun:"price_per_day,nullzero,default:0" json:"price_per_day"`
		PricePerMonth float64         `bun:"price_per_month,nullzero,default:0" json:"price_per_month"`
		Size          float64         `bun:"size,nullzero,default:0" json:"size"`
		Capacity      int64           `bun:"capacity,nullzero,default:0" json:"capacity"`
		Availability  string          `bun:"space_status,default:A" json:"availability"`

		AppModel
	}
)

func (m Space) Upsert(ctx *gin.Context, item Space) (int, Space, error) {
	var oldData *Space
	httpStatus, action := 201, "POST"

	var setClauseColumns = []string{
		"user_id",
		"name",
		"description",
		"address",
		"price_per_hour",
		"price_per_day",
		"price_per_month",
		"size",
		"capacity",
		"availability",
		"updated_at",
	}

	if item.UUID != "" {
		var tmp Space
		if err := db.NewSelect().Model(&tmp).Where("uuid = ?", item.UUID).Scan(ctx); err == nil {
			httpStatus, action, oldData = 200, "PUT", &tmp
		}
	}

	setClause := parseSetClause(setClauseColumns)
	err := executeTransaction(ctx, func(trx *bun.Tx) error {
		_, err := trx.NewInsert().Model(&item).On("CONFLICT (uuid) DO UPDATE").Set(setClause).Exec(ctx)
		return err
	})

	go auditLog(ctx, oldData, item, item.ID, "space", action, err)
	return httpStatus, item, err
}

func (m Space) Read(qp QueryParams) (res SpaceResults, err error) {
	var coalesceCols = []string{}
	var allowedSortFields = map[string]bool{}

	q := db.NewSelect()

	if qp.UUID != "all" {
		return res, q.Model(&res.Space).Where("uuid = ?", qp.UUID).Scan(qp.Ctx)
	}

	q = sanitizeQuery(q.Model(&res.Spaces), qp, coalesceCols, allowedSortFields)
	res.Count, err = q.ScanAndCount(qp.Ctx)
	return res, err
}

func (m Space) Delete(ctx *gin.Context, uuid string) (deletedAt time.Time, msg string, err error) {
	id, deletedAt, _, msg, err := setStatus(ctx, "spaces", uuid, "deleted_at")

	go auditLog(ctx, nil, map[string]string{"deleted_at": deletedAt.String()}, id, "space", "DELETE", err)
	return
}

func (m Space) UpdateStatus(ctx *gin.Context, uuid string) (status, msg string, err error) {
	id, _, status, msg, err := setStatus(ctx, "spaces", uuid, "status")

	go auditLog(ctx, nil, map[string]string{"status": status}, id, "space", "PATCH", err)
	return
}
