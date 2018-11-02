// myEcho project main.go
package main

import (
	"github.com/cooladdr/words/words/api"
	"github.com/cooladdr/words/words/web"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {

	e := echo.New()

	//设置静态文件路径
	e.Static("/static", "assets")

	//api路由
	var apicat api.Category
	var apibook api.Book
	var apiword api.Word
	e.GET("/api/cat", apicat.List)
	e.GET("/api/cat/:id", apicat.Books)
	e.GET("/api/cat/book/:id", apibook.Words)

	e.GET("/api/word/:w", apiword.Info)
	e.GET("/api/word/def/:wid", apiword.Def)
	e.GET("/api/word/expl/:did", apiword.Example)
	e.GET("/api/word/:w/relation", apiword.Relation)

	//web页面路由
	var webcat web.Category
	var webword web.Word
	var webBook web.Book
	e.GET("/cat", webcat.Index)
	e.GET("/word", webword.Index)
	e.GET("/", webword.Index)
	e.GET("/book/:bid", webBook.Index)

	//
	//
	e.Run(standard.New(":1323"))
}
