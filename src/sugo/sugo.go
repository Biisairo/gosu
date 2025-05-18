package sugo

import (
	"html/template"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var MD = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
	goldmark.WithExtensions(extension.GFM, extension.Footnote))

var FrontMatter = []byte("+++")

var Templates *template.Template

var now = time.Now()

type Site struct {
	RootGroup *Group

	Title   string `toml:"title"`
	SiteUrl string `toml:"url"`
	Port    string `toml:"port"`
	RootDir string
}

type Group struct {
	Name   string
	Groups map[int]*Group

	Index *Page
	Pages map[int]*Page
	Url   string
}

type Page struct {
	Title string `toml:"title"`

	Template string `toml:"template"`

	Content string

	OrigFilepath string
	Url          string
}
