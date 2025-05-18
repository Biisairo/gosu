package sugo

import (
	"strings"

	"github.com/BurntSushi/toml"
)

func (s *Site) ParseConfig(file string) error {
	_, err := toml.DecodeFile(file, s)
	if err != nil {
		return err
	}

	// ensure site url has trailing slash
	if !strings.HasSuffix(s.SiteUrl, "/") {
		s.SiteUrl += "/"
	}

	return nil
}
