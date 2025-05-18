package cmdbuild

import (
	"html/template"
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

	// ck(site.RootGroup, 0)
	// return
	tmpl := template.Must(template.ParseFiles(
		"template/header.html",
		"template/footer.html",
		"template/default.html",
	))

	topNav := sugo.GetTopLevelGroups(site.RootGroup)

	os.RemoveAll("dist")
	if err := sugo.RenderGroupToFiles(site.RootGroup, tmpl, "build", topNav); err != nil {
		log.Fatalf("빌드 실패: %v", err)
	}
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
