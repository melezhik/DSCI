package main

import (
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
	"regexp"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
	e.Use(middleware.Recover())       // recover panics as errors for proper error handling

	// Routes
	e.GET("/", home)
	e.GET("/doc/:thing", thing)
	e.Static("/js", "js")
	e.Static("/css", "css")

	// Start server
	if err := e.Start("0.0.0.0:9292"); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}

// Handlers

func home(c *echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/doc/README")
}

func thing(c *echo.Context) error {

	thing := c.Param("thing")

	path := fmt.Sprintf("%s.md",thing)

	data, err := os.ReadFile(path)

	re := regexp.MustCompile(`\/(\S+?)\.md`)

	data = re.ReplaceAll(data, []byte("/doc/" + "$1" ))

	re = regexp.MustCompile(`\.\<(\S+?)\>`)

	data = re.ReplaceAll(data, []byte(".&lt;" + "$1" + "&gt;" ))

	if err != nil {
		log.Printf("thing: Error reading file: %s: %s", path, err)
		return echo.ErrNotFound
	}

	return c.HTML(
		http.StatusOK,
		fmt.Sprintf(
			`<html data-theme="dark">
  <head>
    <meta charset="utf-8">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.4/css/bulma.min.css">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/default.min.css">
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.15.0/dist/katex.min.css">
	<link rel="stylesheet" href="https://unpkg.com/@wcj/markdown-to-html/dist/marked.css">
	<link href="/css/main.css" rel="stylesheet">
	<script src="https://unpkg.com/@wcj/markdown-to-html/dist/markdown.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/perl.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/bash.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/yaml.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/languages/json.min.js"></script>
	<script>hljs.initHighlightingOnLoad();</script>
	<script src="/js/fix_navbar.js" type="text/javascript"></script>
	<title>DSCI - [D]ead [S]imple CI</title>
  </head>
  <body>
	<nav class="navbar is-transparent" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
       <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
          <span aria-hidden="true"></span>
        </a>
      </div>
      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item" href="/">
            Home
          </a>
          <a class="navbar-item" href="/doc/cicd">
            CICD
          </a>
          <a class="navbar-item" href="/doc/pipeline">
            Pipelines
          </a>
          <a class="navbar-item" href="/doc/job">
            Jobs
          </a>
          <a class="navbar-item" href="/doc/task">
            Tasks
          </a>
          <a class="navbar-item" href="/doc/forgejo-setup">
            Forgejo
          </a>
         <a class="navbar-item" href="/doc/cookbook">
            Cookbook
          </a>
          <a class="navbar-item" href="/doc/discord">
            Discord
          </a>
          <a class="navbar-item" href="/doc/demo">
            Demo Server
          </a>
          <a class="navbar-item" href="/doc/screenshots">
            Screenshots
          </a>
         <a class="navbar-item" href="https://github.com/melezhik/DSCI">
            GitHub
          </a>
        </div>
      </div>
    </nav>
    <section class="hero">
      <div class="hero-body">
        <p class="title">%s</p>
		<hr>
		<p class="content" id="data">%s</p>
      </div>
    </section>
  <script type="text/javascript">
	;(() => {
		const str = document.getElementById("data").textContent;
		var p = document.getElementById("data");
		const div = document.createElement('div');
		//div.className = 'markdown-body';
		p.innerHTML = markdown.default(str)
		//document.body.appendChild(div)
	})()
  </script>
  </body>
</html>`,
			thing,
			data,
		),
	)

}
