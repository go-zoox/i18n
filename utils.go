package i18n

import (
	"fmt"

	"github.com/go-zoox/fetch"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/fs/type/json"
	"github.com/go-zoox/fs/type/yaml"
)

// LoadLocalesFromFile loads locales from file
// Only support JSON, YAML, TOML and INI format.
func LoadLocalesFromFile(filepath string) (map[string]Translations, error) {
	ext := fs.ExtName(filepath)
	locales := make(map[string]Translations)
	switch ext {
	case ".json":
		if err := json.Read(filepath, &locales); err != nil {
			return nil, err
		}
	case ".yaml", ".yml":
		if err := yaml.Read(filepath, &locales); err != nil {
			return nil, err
		}
	case ".toml":
		if err := yaml.Read(filepath, &locales); err != nil {
			return nil, err
		}
	case ".ini":
		if err := yaml.Read(filepath, &locales); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported locale file format: %s", ext)
	}

	return locales, nil
}

// LoadLocalesFromDir loads locales from dir
// The directory structure should be like this:
// lang/
//
//		en-US.json
//		zh-CN.json
//		en-US.yaml
//		en-US.toml
//	 	en-US.ini
//		...
//
// Only support JSON, YAML, TOML and INI format.
func LoadLocalesFromDir(dir string) (map[string]Translations, error) {
	files, err := fs.ListDir(dir)
	if err != nil {
		return nil, err
	}

	locales := make(map[string]Translations, len(files))
	for _, file := range files {
		fileName := file.Name()
		ext := fs.ExtName(fileName)
		localeName := fileName[:len(fileName)-len(ext)]
		filepath := fs.JoinPath(dir, fileName)
		translations := Translations{}

		switch ext {
		case ".json":
			if err := json.Read(filepath, &translations); err != nil {
				return nil, err
			}
		case ".yaml", ".yml":
			if err := yaml.Read(filepath, &translations); err != nil {
				return nil, err
			}
		case ".toml":
			if err := yaml.Read(filepath, &translations); err != nil {
				return nil, err
			}
		case ".ini":
			if err := yaml.Read(filepath, &translations); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unsupported locale file format: %s", ext)
		}

		locales[localeName] = translations
	}

	return locales, nil
}

// LoadLocalesFromURL loads locales from url with JSON format.
func LoadLocalesFromURL(url string) (map[string]Translations, error) {
	response, err := fetch.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locale file from %s: %s", url, err)
	}

	locales := map[string]Translations{}
	if err := response.UnmarshalJSON(&locales); err != nil {
		return nil, err
	}

	return locales, nil
}
