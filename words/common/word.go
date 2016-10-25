package common

import (
	"database/sql"
	_ "log"
	_ "net/http"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Word struct {
	}

	WordItem struct {
		Id       int
		Spelling string
		Meiyin   string
		Yingyin  string
		Type     string
	}

	RelatedWordsItem struct {
		Spelling string
		Relation string
	}

	Definition struct {
		Id       int
		WId      int
		Part     string
		From     string
		Order    int
		SubOrder int
		Level    string
		EnDef    string
		ChDef    string
	}

	Example struct {
		Id         int
		DefId      int
		Order      int
		From       string
		EnSentence string
		ChSentence string
	}
)

func (this *Word) dbQuery(sqlStr string, params ...interface{}) {

}

func (this *Word) GetInfo(word string) (wordItem *WordItem, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sql := "select id,spelling, meiyin,yingyin,type from go_words where spelling=?"

	rows, err := db.Query(sql, word)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id int
	var spelling, meiyin, yingyin, wtype string

	for rows.Next() {
		err := rows.Scan(&id, &spelling, &meiyin, &yingyin, &wtype)
		if err != nil {
			return nil, err
		}
		wordItem = &WordItem{id, spelling, meiyin, yingyin, wtype}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return wordItem, nil
}

func (this *Word) GetDef(wordId int, part string) (wordDefs []Definition, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlStr := "select id,wordId,part,editFrom,showOrder,subOrder,level,enDef,chDef from go_words_definitions where wordId=? "

	var rows *sql.Rows

	if len(part) > 0 {
		sqlStr += " and part=?"
		rows, err = db.Query(sqlStr, wordId, part)
	} else {
		rows, err = db.Query(sqlStr, wordId)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, wId, order, subOrder int
	var p, efrom, level, enDef, chDef string

	for rows.Next() {
		err := rows.Scan(&id, &wId, &p, &efrom, &order, &subOrder, &level, &enDef, &chDef)
		if err != nil {
			return nil, err
		}

		wordDefs = append(wordDefs, Definition{id, wId, p, efrom, order, subOrder, level, enDef, chDef})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return wordDefs, nil
}

func (this *Word) GetExample(defId int) (examples []Example, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var rows *sql.Rows
	sqlStr := "select id,defId,showOrder,efitFrom,enSentence,chSentence from go_words_def_examples where defId=?"

	rows, err = db.Query(sqlStr, defId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var id, deffinitionId, order int
	var efitFrom, enSentence, chSentence string

	for rows.Next() {
		err := rows.Scan(&id, &deffinitionId, &order, &efitFrom, &enSentence, &chSentence)
		if err != nil {
			return nil, err
		}
		examples = append(examples, Example{id, deffinitionId, order, efitFrom, enSentence, chSentence})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return examples, nil
}

func (this *Word) GetRelation(Word, relation string) (relatedWords []RelatedWordsItem, err error) {
	db, err := sql.Open("mysql", "root:@/wordnew")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var rows *sql.Rows
	sqlStr := "select minor,relation from go_words_relation where major=?"

	if len(relation) > 0 {
		sqlStr += "  and relation=? order by relation"
		rows, err = db.Query(sqlStr, Word, relation)
	} else {
		sqlStr += " order by relation"
		rows, err = db.Query(sqlStr, Word)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tmpWord, rel string

	for rows.Next() {
		err := rows.Scan(&tmpWord, &rel)
		if err != nil {
			return nil, err
		}
		relatedWords = append(relatedWords, RelatedWordsItem{tmpWord, rel})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return relatedWords, nil
}
