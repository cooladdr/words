//book是单词本
package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Book struct {
	}

	BookItem struct {
		Id    int
		Name  string
		CatId int //category id 单词本所在的分类
		Count int
	}

	BookWordItem struct {
		Id          int64
		BookId      int
		Spelling    string
		Meiyin      string
		Yingyin     string
		Type        string
		Definitions string
	}
)

//CatBooks 根据分类id读取所有单词本
func (this *Book) CatBooks(catId int) (bookList []BookItem, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := "select id, showName,catId,wCount from go_cat_book where catId=? and status=? order by showName"

	rows, err := db.Query(sql, catId, "display")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, categoryId, wCount int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name, &categoryId, &wCount)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, BookItem{id, name, categoryId, wCount})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return bookList, nil

}

func (this *Book) GeWords(bookId, start, size int) (wordList []BookWordItem, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := "select id,bookId,spelling,meiyin,yingyin,type,definitions from go_cat_book_words "
	sql += "where bookId = ? limit ?,?"

	rows, err := db.Query(sql, bookId, start*size, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int64
	var bId int
	var spelling, meiyin, yingyin, wtype, definitions string

	for rows.Next() {
		err := rows.Scan(&id, &bId, &spelling, &meiyin, &yingyin, &wtype, &definitions)
		if err != nil {
			return nil, err
		}
		wordList = append(wordList, BookWordItem{id, bId, spelling, meiyin, yingyin, wtype, definitions})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return wordList, nil

}
