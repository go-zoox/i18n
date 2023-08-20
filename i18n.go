package i18n

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
