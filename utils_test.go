package i18n

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestLoadLocalesFromFile(t *testing.T) {
	locales, err := LoadLocalesFromFile("./tests/locales.json")
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "en-US", locales["en-US"]["en-US"])
	testify.Equal(t, "zh-CN", locales["en-US"]["zh-CN"])
	testify.Equal(t, "name", locales["en-US"]["name"])
	testify.Equal(t, "age", locales["en-US"]["age"])

	testify.Equal(t, "英文", locales["zh-CN"]["en-US"])
	testify.Equal(t, "中文", locales["zh-CN"]["zh-CN"])
	testify.Equal(t, "姓名", locales["zh-CN"]["name"])
	testify.Equal(t, "年龄", locales["zh-CN"]["age"])
}

func TestLoadLocalesFromDir(t *testing.T) {
	locales, err := LoadLocalesFromDir("./tests/lang")
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "en-US", locales["en-US"]["en-US"])
	testify.Equal(t, "zh-CN", locales["en-US"]["zh-CN"])
	testify.Equal(t, "name", locales["en-US"]["name"])
	testify.Equal(t, "age", locales["en-US"]["age"])

	testify.Equal(t, "英文", locales["zh-CN"]["en-US"])
	testify.Equal(t, "中文", locales["zh-CN"]["zh-CN"])
	testify.Equal(t, "姓名", locales["zh-CN"]["name"])
	testify.Equal(t, "年龄", locales["zh-CN"]["age"])
}

func TestLoadLocalesFromURL(t *testing.T) {
	locales, err := LoadLocalesFromURL("https://raw.githubusercontent.com/go-zoox/i18n/master/tests/locales.json")
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "en-US", locales["en-US"]["en-US"])
	testify.Equal(t, "zh-CN", locales["en-US"]["zh-CN"])
	testify.Equal(t, "name", locales["en-US"]["name"])
	testify.Equal(t, "age", locales["en-US"]["age"])

	testify.Equal(t, "英文", locales["zh-CN"]["en-US"])
	testify.Equal(t, "中文", locales["zh-CN"]["zh-CN"])
	testify.Equal(t, "姓名", locales["zh-CN"]["name"])
	testify.Equal(t, "年龄", locales["zh-CN"]["age"])
}
