package i18n

import (
	"fmt"

	"github.com/go-zoox/core-utils/strings"
	"github.com/spf13/cast"
)

// translate translates the given key with the given arguments.
func translate(locales map[string]Translations, locale string, key string, data ...map[string]any) (string, error) {
	if locales == nil {
		return "", fmt.Errorf("locales not loaded")
	}

	// Get the translation for the given key.
	translations, ok := locales[locale]
	if !ok {
		return "", fmt.Errorf("locale(%s) not found", locale)
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

// t is an alias for translate.
func t(locales map[string]Translations, locale string, key string, data ...map[string]any) string {
	translation, err := translate(locales, locale, key, data...)
	if err != nil {
		return key
	}

	return translation
}
