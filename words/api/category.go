//Category是单词分类
package api

import (
	"log"
	"net/http"
	"strconv"
	"words/words/common"

	"github.com/labstack/echo"
)

type (
	Category struct {
		common.Category
	}
)

func (this *Category) List(ctx echo.Context) error {
	list, err := this.GetList()
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, list)
	}

	ctx.JSON(http.StatusOK, list)

	return nil
}

func (this *Category) Books(ctx echo.Context) error {
	catId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusBadRequest, "")
	}

	var book common.Book

	books, err := book.CatBooks(catId)
	if err != nil {
		log.Printf("%s", err)
		ctx.JSON(http.StatusServiceUnavailable, books)
	}

	ctx.JSON(http.StatusOK, books)

	return nil
}
