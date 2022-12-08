package service

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ResultService struct {
	db *sql.DB
}

func NewResultService(db *sql.DB) *ResultService {
	return &ResultService{
		db: db,
	}
}

func (s *ResultService) GetResult(ctx context.Context, size int64) ([]string, error) {
	const (
		getFmt = `SELECT word FROM Noun WHERE id in (?%s)`
	)
	result := []string{}
	var err error

	str := fmt.Sprintf(getFmt, strings.Repeat(" ,?", int(size-1)))

	var ids_arg = []interface{}{}

	var i int64
	for i = 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		var id = strconv.Itoa(rand.Intn(60476) + 1) //dbの単語idが1~60476
		ids_arg = append(ids_arg, id)
	}

	stmt, err := s.db.PrepareContext(ctx, str)
	if err != nil {
		return nil, err
	} else {
		defer stmt.Close()
		rows, err := stmt.QueryContext(ctx, ids_arg...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var word string
			if err := rows.Scan(&word); err != nil {
				return nil, err
			}
			result = append(result, word)
		}
	}

	return result, err
}
