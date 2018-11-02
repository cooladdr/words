package api

import (
	"log"
	"net/http"
	"strconv"
	"github.com/cooladdr/words/words/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type (
	Word struct {
		common.Word
	}
)

func (this *Word) Info(ctx echo.Context) error {
	word := ctx.Param("w")

	wordInfo, err := this.GetInfo(word)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, "")
	} else {
		ctx.JSON(http.StatusOK, wordInfo)
	}

	return nil
}

func (this *Word) Def(ctx echo.Context) error {
	wordId, err := strconv.Atoi(ctx.Param("wid"))
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusBadRequest, "wid")
		return nil
	}

	defs, err := this.GetDef(wordId, "")
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, "")
	} else {
		ctx.JSON(http.StatusOK, defs)
	}

	return nil
}

func (this *Word) Example(ctx echo.Context) error {
	defId, err := strconv.Atoi(ctx.Param("did"))
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusBadRequest, "did")
		return nil
	}

	examples, err := this.GetExample(defId)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, "")
	} else {
		ctx.JSON(http.StatusOK, examples)
	}

	return nil
}

func (this *Word) Relation(ctx echo.Context) error {
	word := ctx.Param("w")
	relation := ctx.FormValue("r")

	relatedWords, err := this.GetRelation(word, relation)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, "")
	} else {
		ctx.JSON(http.StatusOK, relatedWords)
	}

	return nil
}
