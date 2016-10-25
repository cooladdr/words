//Category是单词分类
package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Category struct {
	}

	CatItem struct {
		Id    int
		Name  string
		Count int
	}
)

func (this *Category) GetList() (catList []CatItem, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("select id, showName, bCount from go_cat where status=? order by showName", "display")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, bcount int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name, &bcount)
		if err != nil {
			return nil, err
		}
		catList = append(catList, CatItem{id, name, bcount})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return catList, nil
}
