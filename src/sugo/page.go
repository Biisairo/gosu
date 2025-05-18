package sugo

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var md = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
	goldmark.WithExtensions(extension.GFM, extension.Footnote))

var frontMatter = []byte("+++")

func (p *Page) Print(layer int) {
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("Title : ", p.Title)
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("Template : ", p.Template)
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("Content : ", p.Content)
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("OrigFilepath : ", p.OrigFilepath)
	for range layer {
		fmt.Print("\t")
	}
	fmt.Println("Url : ", p.Url)
}

func (p *Page) ParseFrontMatter() error {
	fh, err := os.Open(p.OrigFilepath)
	if err != nil {
		return err
	}
	defer fh.Close()

	buf := make([]byte, 1024)
	n, err := fh.Read(buf)
	if err != nil {
		return err
	}

	buf = buf[:n]
	if !bytes.HasPrefix(buf, frontMatter) {
		return nil
	}

	buf = buf[3:]

	pos := bytes.Index(buf, frontMatter)
	if pos == -1 {
		return fmt.Errorf("missing closing front-matter identifier in %s", p.OrigFilepath)
	}

	return toml.Unmarshal(buf[:pos], p)
}

func (p *Page) ParseContent() (string, error) {
	fileContent, err := os.ReadFile(p.OrigFilepath)
	if err != nil {
		return "", err
	}

	if len(fileContent) > 6 {
		pos := bytes.Index(fileContent[3:], frontMatter)
		if pos > -1 {
			fileContent = fileContent[pos+6:]
		}
	}

	if strings.HasSuffix(p.OrigFilepath, ".html") {
		return string(fileContent), nil
	}

	var buf2 strings.Builder
	if err := md.Convert(fileContent, &buf2); err != nil {
		return "", err
	}
	return buf2.String(), nil
}
