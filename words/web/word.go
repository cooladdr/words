package web

import (
	"html/template"
	"io"
	"net/http"
	"github.com/cooladdr/words/words/common"

	"github.com/labstack/echo"
)

type (
	Word struct {
		common.Word
		templates *template.Template
	}
)

func (t *Word) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (this *Word) Index(ctx echo.Context) error {

	t := &Word{
		//templates: template.Must(template.ParseGlob("public/views/*.html")),
		templates: template.Must(template.New("words").Parse(`<!DOCTYPE html>
<html lang="en-US">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Words</title>
    <!--<link href="css/site/index.css" rel="stylesheet">-->
	<link href="/static/assets/36ae7f3/css/bootstrap.css" rel="stylesheet">
	<link href="/static/css/layout/main.css" rel="stylesheet"></head>
	{{range .CssList}}
	<link href="{{.}}" rel="stylesheet">
	{{end}}
<body>

<div class="wrap">
    <nav id="w0" class="navbar-inverse navbar-fixed-top navbar" role="navigation">
		<div class="container">
			<div class="navbar-header">
				<button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#w0-collapse">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
				<a class="navbar-brand" href="/cat">单词酷</a>
			</div>
			<div id="w0-collapse" class="collapse navbar-collapse">
				<ul id="w1" class="navbar-nav navbar-right nav">
					<li class="active"><a href="/word">首页</a></li>
					<li><a href="/cat">单词本分类</a></li>
					<li><a href="">关于我们</a></li>
					<li><a href="">联系我们</a></li>
					<li><a href="">注册</a></li>
					<li><a href="">登陆</a></li>
				</ul>
			</div>
		</div>
	</nav>
    <div class="container">
		<div class="site-index">
			<div class="jumbotron"></div>
    		<div class="body-content ">
        		<div id="container"></div>
    		</div>
		</div>
	</div>
</div>



<script src="/static/assets/a7b1b087/jquery.js"></script>
<script src="/static/js/react/react.min.js"></script>
<script src="/static/js/react/react-dom.min.js"></script>
<script src="/static/js/react/browser.min.js"></script>
<script src="/static/js/react/marked.min.js"></script>
<script src="/static/assets/36ae7f3/js/bootstrap.js"></script>
{{range .JsList}}
<script src="{{.}}"></script>
{{end}}
{{range .LabelJsList}}
<script type="text/babel" src="{{.}}"></script>
{{end}}
</body>
</html>`)),
	}

	ctx.Echo().SetRenderer(t)

	data := common.TplAssets{
		CssList:     []string{"/static/css/words/words.css"},
		LabelJsList: []string{"/static/js/words/words.js"},
	}

	return ctx.Render(http.StatusOK, "words", data)
}

func (this *Word) Def(ctx echo.Context) error {
	return nil
}
