//book是单词本
package api

import (
	"log"
	"net/http"
	"strconv"
	"words/words/common"

	"github.com/labstack/echo"
)

type (
	Book struct {
		common.Book
	}
)

func (this *Book) Words(ctx echo.Context) error {
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusBadRequest, "")
	}

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}

	size, err := strconv.Atoi(ctx.QueryParam("size"))
	if err != nil {
		size = 20
	}

	words, err := this.GeWords(bookId, page-1, size)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, words)
	}

	ctx.JSON(http.StatusOK, words)

	return nil
}
