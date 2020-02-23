package tools

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/queries"
	"server/db/models"
	"server/initialize/db"
)

func GetUsersFromIds(user_ids []uint) ([]*models.User, error) {
	if len(user_ids) == 0 {
		return nil, nil
	}

	query_str_before := `SELECT id, username, displayname, avatar FROM user WHERE `
	query_str_middle := ``
	query_str_after := `;`

	for _, user_id := range user_ids {
		query_str_middle += fmt.Sprintf("(id = %d) OR ", user_id)
	}
	// "OR " の3文字文を削除
	query_str_middle = query_str_middle[:(len(query_str_middle) - 3)]
	// 文字列を連結
	query_str := query_str_before + query_str_middle + query_str_after
	// userをバルクsearch
	users := []*models.User{}
	err := queries.Raw(query_str).Bind(context.Background(), db.Connection, &users)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return users, nil
}
