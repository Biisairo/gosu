package sugo

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func (s *Site) ReadContent(dir string) error {
	s.RootGroup = &Group{
		Name:   "",
		Groups: map[int]*Group{},
		Pages:  map[int]*Page{},
	}

	return recuReadContent(dir, s.RootGroup)
}

func recuReadContent(root string, group *Group) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			readDir(root, &entry, group)
		} else {
			readFile(root, &entry, group)
		}
	}

	return nil
}

func readDir(root string, entry *os.DirEntry, group *Group) error {
	pathName := (*entry).Name()

	index, name, err := ParseIndexedName(pathName)
	if err != nil {
		return err
	}

	url := group.Url
	if url == "" {
		url = name
	} else {
		url = fmt.Sprintf("%v/%v", group.Url, name)
	}

	newGroup := Group{
		Name:   name,
		Groups: map[int]*Group{},
		Pages:  map[int]*Page{},
		Url:    url,
	}

	if _, ok := group.Groups[index]; ok {
		return errors.New("duplicated group index")
	}

	group.Groups[index] = &newGroup

	nextPath := filepath.Join(root, pathName)

	return recuReadContent(nextPath, &newGroup)
}

func readFile(root string, entry *os.DirEntry, group *Group) error {
	pathName := (*entry).Name()

	ext := filepath.Ext(pathName)
	if ext != ".md" && ext != ".html" {
		return errors.New(fmt.Sprintf("file is not md or html : %v", pathName))
	}

	pathStem := strings.TrimSuffix(pathName, ".md")
	pathStem = strings.TrimSuffix(pathStem, ".html")

	if pathName == "index.md" {
		if group.Name == "" {
			group.Name = "Home"
		}

		page := &Page{
			Title:        group.Name,
			OrigFilepath: filepath.Join(root, pathName),
			Template:     "default.html",
			Url:          group.Url,
		}
		parsePage(page)
		group.Index = page
	} else {
		index, name, err := ParseIndexedName(pathStem)
		if err != nil {
			return err
		}

		if _, ok := group.Pages[index]; ok {
			return errors.New("duplicated page index")
		}

		page := &Page{
			Title:        name,
			OrigFilepath: filepath.Join(root, pathName),
			Template:     "default.html",
			Url:          fmt.Sprintf("%v/%v", group.Url, name),
		}
		parsePage(page)
		group.Pages[index] = page
	}

	return nil
}

func parsePage(p *Page) error {
	if err := p.ParseFrontMatter(); err != nil {
		return err
	}

	content, err := p.ParseContent()
	if err != nil {
		return nil
	}

	p.Content = template.HTML(content)

	return nil
}
