package cmdbuild

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/Biisairo/sugo/src/sugo"
)

func Build(rootPath string, configFile string) {
	site := &sugo.Site{
		RootDir: rootPath,
	}

	if err := site.ParseConfig(filepath.Join(rootPath, configFile)); err != nil {
		log.Fatalf("Error reading configuration file at %s: %s\n", rootPath+configFile, err)
	}

	if err := site.ReadContent(filepath.Join(rootPath, "content")); err != nil {
		log.Fatalf("Error reading content/: %s", err)
	}

	files, err := os.ReadDir("template")
	if err != nil {
		log.Fatalf("Error loading template: %v\n", err)
	}

	var templateName []string
	for _, file := range files {
		name := file.Name()
		ext := filepath.Ext(name)

		if ext != ".html" {
			log.Fatalf("Error wrong template extension: %v\n", err)
		}

		templateFile := fmt.Sprintf("template/%s", name)
		templateName = append(templateName, templateFile)
	}

	sugo.Templates = template.Must(template.ParseFiles(templateName...))

	topNav := sugo.GetTopLevelGroups(site.RootGroup)

	os.RemoveAll("build")
	if err := sugo.RenderGroupToFiles(site, site.RootGroup, topNav); err != nil {
		log.Fatalf("빌드 실패: %v", err)
	}

	copyDir("static", "build/static")
}

func copyDir(src string, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if d.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		}

		return copyFile(path, dstPath)
	})
}

func copyFile(srcFile, dstFile string) error {
	in, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// 권한 복사
	info, err := os.Stat(srcFile)
	if err != nil {
		return err
	}

	return os.Chmod(dstFile, info.Mode())
}

func ck(g *sugo.Group, layer int) {

	g.Print(layer)

	var keys []int
	for k := range g.Pages {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		page := g.Pages[k]
		page.Print(layer + 1)
	}

	keys = []int{}
	for k := range g.Groups {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		group := g.Groups[k]
		ck(group, layer+1)
	}
}
