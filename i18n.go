package i18n

import "github.com/go-zoox/logger"

type I18n interface {
	// Load loads the given locale and translations into the I18n instance.
	Load(fn func() (map[string]Translations, error)) error

	// LoadFromDir loads the given locale and translations into the I18n instance.
	LoadFromDir(dir string) error

	// Translate translates the given key with the given arguments.
	Translate(locale string, key string, data ...map[string]any) (string, error)
	// T is an alias for Translate.
	T(locale string, key string, data ...map[string]any) string
}

type i18n struct {
	// locales is a map of locales to their translations.
	locales map[string]Translations
}

type Translations = map[string]string

// New creates a new I18n instance with the given locale and translations.
func New() I18n {
	return &i18n{}
}

// Translate translates the given key with the given arguments.
func (i *i18n) Translate(locale string, key string, data ...map[string]any) (string, error) {
	return translate(i.locales, locale, key, data...)
}

// T is an alias for Translate.
// if the translation is not found, it returns the key.
func (i *i18n) T(locale string, key string, data ...map[string]any) string {
	return t(i.locales, locale, key, data...)
}

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
		locales, err := loadLocalesFromFile(file)
		if err != nil {
			return nil, err
		}

		return locales, nil
	})
}

// LoadFromDir loads the given locale and translations into the I18n instance.
func (i *i18n) LoadFromDir(dir string) error {
	return i.Load(func() (map[string]Translations, error) {
		locales, err := loadLocalesFromDir(dir)
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
		locales, err := loadLocalesFromURL(url)
		if err != nil {
			return nil, err
		}

		return locales, nil
	})
}
