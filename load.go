package i18n

import (
	"github.com/go-zoox/logger"
)

// Load loads the given locale and translations into the I18n instance.
func (i *i18n) Load(fn func() (map[string]Translations, error)) error {
	locales, err := fn()
	if err != nil {
		return err
	}

	i.locales = locales
	for locale, _ := range locales {
		logger.Infof("[i18n] loaded locale: %s", locale)
	}

	return nil
}

// LoadFromFile loads the given locale and translations into the I18n instance.
func (i *i18n) LoadFromFile(file string) error {
	return i.Load(func() (map[string]Translations, error) {
		locales, err := LoadLocalesFromFile(file)
		if err != nil {
			return nil, err
		}

		return locales, nil
	})
}

// LoadFromDir loads the given locale and translations into the I18n instance.
func (i *i18n) LoadFromDir(dir string) error {
	return i.Load(func() (map[string]Translations, error) {
		locales, err := LoadLocalesFromDir(dir)
		if err != nil {
			return nil, err
		}

		return locales, nil
	})
}

// LoadFromURL loads the given locale and translations into the I18n instance.
// Only support JSON format.
func (i *i18n) LoadFromURL(url string) error {
	return i.Load(func() (map[string]Translations, error) {
		locales, err := LoadLocalesFromURL(url)
		if err != nil {
			return nil, err
		}

		return locales, nil
	})
}
