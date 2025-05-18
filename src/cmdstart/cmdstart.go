package cmdstart

import (
	"os"
	"path/filepath"
)

func Start(rootPath string) error {
	for _, dir := range []string{"content", "templates", "public"} {
		if err := os.Mkdir(filepath.Join(rootPath, dir), 0755); err != nil {
			return err
		}
	}

	files := []struct {
		Name    string
		Content []byte
	}{
		{"config.toml", []byte(`url = "http://localhost"
port = "8080"
title = "My website"
`)},
		{"templates/default.html", []byte(`
<!DOCTYPE html>
<head>
	<title>{{ .Title }}</title>
</head>
<body>
{{ .Content }}
</body>
</html>`)},
		{"content/index.md", []byte(`+++
title = "SuGo!"
+++

Welcome to my website.
`)},
	}
	for _, f := range files {
		if err := os.WriteFile(filepath.Join(rootPath, f.Name), f.Content, 0655); err != nil {
			return err
		}
	}

	return nil
}
