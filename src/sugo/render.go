package sugo

import (
	"html/template"
	"os"
	"path/filepath"
	"sort"
)

type templatePage struct {
	IsRoot    bool
	Nav       []*Group
	SubGroups []*Group
	SubPages  []*Page
	Page      *Page
	Title     string
}

// 정렬된 상위 그룹들(nav용)
func GetTopLevelGroups(root *Group) []*Group {
	keys := make([]int, 0, len(root.Groups))
	for k := range root.Groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	result := make([]*Group, 0, len(keys))
	for _, k := range keys {
		result = append(result, root.Groups[k])
	}
	return result
}

// 그룹 렌더링
func RenderGroupToFiles(root *Group, tmpl *template.Template, outputDir string, topNav []*Group) error {
	if root.Index != nil {
		outputPath := filepath.Join(outputDir, root.Url, "index.html")
		if err := renderPageToFile(root.Index, tmpl, outputPath, topNav, groupSubgroups(root), groupSubPages(root)); err != nil {
			return err
		}
	} else {
		outputPath := filepath.Join(outputDir, root.Url, "index.html")
		if err := renderEmptyToFile(root.Name, tmpl, outputPath, topNav, groupSubgroups(root), groupSubPages(root)); err != nil {
			return err
		}
	}

	// Render other pages in the group (not the index)
	for _, page := range root.Pages {
		if page == root.Index {
			continue
		}
		outputPath := filepath.Join(outputDir, page.Url+".html")
		if err := renderPageToFile(page, tmpl, outputPath, topNav, nil, nil); err != nil {
			return err
		}
	}

	// Recursively render subgroups
	for _, sub := range root.Groups {
		if err := RenderGroupToFiles(sub, tmpl, outputDir, topNav); err != nil {
			return err
		}
	}

	return nil
}

// 렌더링에 필요한 데이터 구조 및 실행
func renderPageToFile(page *Page, tmpl *template.Template, path string, nav []*Group, subGroups []*Group, subPages []*Page) error {
	isRoot := false
	if page.Url == "" {
		isRoot = true
	}

	data := templatePage{
		IsRoot:    isRoot,
		Nav:       nav,
		SubGroups: subGroups,
		SubPages:  subPages,
		Page:      page,
		Title:     page.Title,
	}

	return generateHTML(path, tmpl, page.Template, &data)
}

func renderEmptyToFile(title string, tmpl *template.Template, path string, nav []*Group, subGroups []*Group, subPages []*Page) error {
	data := templatePage{
		IsRoot:    false,
		Nav:       nav,
		SubGroups: subGroups,
		SubPages:  subPages,
		Page:      nil,
		Title:     title,
	}

	return generateHTML(path, tmpl, "default.html", &data)
}

func generateHTML(path string, tmpl *template.Template, template string, data *templatePage) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.ExecuteTemplate(f, template, data)
}

// 하위 그룹 정렬된 리스트
func groupSubgroups(g *Group) []*Group {
	keys := make([]int, 0, len(g.Groups))
	for k := range g.Groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	subs := make([]*Group, 0, len(keys))
	for _, k := range keys {
		subs = append(subs, g.Groups[k])
	}
	return subs
}

// 인덱스를 제외한 페이지들 리스트
func groupSubPages(g *Group) []*Page {
	keys := make([]int, 0, len(g.Pages))
	for k := range g.Pages {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	result := make([]*Page, 0, len(keys))
	for _, k := range keys {
		p := g.Pages[k]
		if p != g.Index {
			result = append(result, p)
		}
	}
	return result
}
