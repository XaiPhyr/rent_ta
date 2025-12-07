package models

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type (
	User struct {
		bun.BaseModel `bun:"table:users,alias:u"`
		AppModel

		Username   string          `bun:"username" json:"username"`
		Password   string          `bun:"password" json:"-"`
		Email      string          `bun:"email,unique" json:"email"`
		FirstName  *string         `bun:"first_name,nullzero,default:null" json:"first_name"`
		MiddleName *string         `bun:"middle_name,nullzero,default:null" json:"middle_name"`
		LastName   *string         `bun:"last_name,nullzero,default:null" json:"last_name"`
		Mobile     *string         `bun:"mobile,unique,nullzero,default:null" json:"mobile"`
		Address    *map[string]any `bun:"address,type:jsonb" json:"address"`
		Optin      bool            `bun:"optin,default:false" json:"optin"`
		LastLogin  time.Time       `bun:"last_login,nullzero,default:null" json:"last_login,omitzero"`
		Metadata   *map[string]any `bun:"metadata,type:jsonb,default:null" json:"metadata"`
		IsAdmin    bool            `bun:"is_admin" json:"is_admin"`
	}

	UserResults struct {
		User  User   `json:"user,omitempty"`
		Users []User `json:"users,omitempty"`
		Count int    `json:"count,omitempty"`
	}
)

func (m User) ParseUser(ctx *gin.Context) *User {
	return &User{}
}

func (m User) Upsert(ctx *gin.Context, user User) (User, error) {
	var newData, oldData User
	var parseOldData *User
	action := "POST"

	setClause := m.parseSetClause()

	if user.UUID != "" {
		db.NewSelect().Model(&oldData).Where("uuid = ?", user.UUID).Scan(ctx, &oldData)
		parseOldData = &oldData
		action = "PUT"
	}

	err := executeTransaction(ctx, func(trx *bun.Tx) error {
		q := trx.NewInsert().Model(&user)
		q = q.On("CONFLICT (uuid) DO UPDATE")
		q = q.Set(setClause)
		_, err := q.Exec(ctx)
		return err
	})

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	newData = user
	go auditLog(ctx, parseOldData, newData, user.ID, "user", action, errMsg)
	return user, err
}

func (m User) Read(ctx *gin.Context, uuid, filter, filterExt, sort, status string, limit, page int) (res UserResults, err error) {
	cols := []string{}
	q := db.NewSelect()

	if uuid != "all" {
		return res, q.Model(&res.User).Where("uuid = ?", uuid).Scan(ctx, &res.User)
	}

	q = q.Model(&res.Users)
	q = sanitizeQuery(q, cols, filter, sort, status, m.validateField)
	q = m.applyFilter(q, filterExt)

	if limit > 0 {
		q = q.Limit(limit)

		if page > 0 {
			offset := (page - 1) * limit
			q = q.Offset(offset)
		}
	}

	res.Count, err = q.ScanAndCount(ctx)
	return res, err
}

func (m User) Delete(ctx *gin.Context, uuid string) (user User, err error) {
	err = executeTransaction(ctx, func(trx *bun.Tx) error {
		q := trx.NewUpdate().Model(&user).Where("uuid = ?", uuid)
		q = q.Set("deleted_at = CASE WHEN deleted_at IS NULL THEN NOW() ELSE NULL END")
		q = q.Returning("deleted_at, id")
		_, err = q.Exec(ctx)
		return err
	})

	var deletedAt any
	if !user.DeletedAt.IsZero() {
		deletedAt = map[string]string{"deleted_at": user.DeletedAt.String()}
	}

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	go auditLog(ctx, nil, deletedAt, user.ID, "user", "DELETE", errMsg)
	return
}

func (m User) applyFilter(q *bun.SelectQuery, filterExt string) *bun.SelectQuery {
	if filterExt == "" {
		return q
	}

	filters := strings.SplitSeq(filterExt, ",")
	for f := range filters {
		kv := strings.Split(f, "=")

		if len(kv) == 2 {
			k, v := kv[0], kv[1]

			q = q.Where(k+" = ?", v)
		}
	}

	return q
}

func (m User) parseSetClause() string {
	cols := []string{
		"username",
		"email",
		"first_name",
		"middle_name",
		"last_name",
		"mobile",
		"address",
		"optin",
		"metadata",
		"is_admin",
		"updated_at",
	}

	setClauses := make([]string, 0, len(cols))
	for _, col := range cols {
		setClauses = append(setClauses, col+" = EXCLUDED."+col)
	}

	return strings.Join(setClauses, ", ")
}

func (m User) validateField(sortField string) bool {
	after, _ := strings.CutPrefix(sortField, "-")

	allowedSortFields := map[string]bool{
		"id":         true,
		"username":   true,
		"email":      true,
		"created_at": true,
	}

	return allowedSortFields[after]
}
