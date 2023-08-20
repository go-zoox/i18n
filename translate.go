package i18n

import (
	"fmt"

	"github.com/go-zoox/core-utils/strings"
	"github.com/spf13/cast"
)

// Translate translates the given key with the given arguments.
func (i *i18n) Translate(locale string, key string, data ...map[string]any) (string, error) {
	// Get the translation for the given key.
	translations, ok := i.locales[locale]
	if !ok {
		return "", fmt.Errorf("invalid locale: %s", locale)
	}

	translation, ok := translations[key]
	if !ok {
		return "", fmt.Errorf("translation key(%s) not found", key)
	}

	// If no data was given, return the translation.
	if len(data) == 0 || data[0] == nil {
		return cast.ToString(translation), nil
	}

	return strings.Format(translation, data[0]), nil
}

// T is an alias for Translate.
// if the translation is not found, it returns the key.
func (i *i18n) T(locale string, key string, data ...map[string]any) string {
	translation, err := i.Translate(locale, key, data...)
	if err != nil {
		return key
	}

	return translation
}
