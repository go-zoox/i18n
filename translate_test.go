package i18n

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestTranslate(tx *testing.T) {
	locales := map[string]map[string]string{
		"en-US": {
			"product":                     "product",
			"design":                      "design",
			"frontend":                    "frontend",
			"backend":                     "backend",
			"test":                        "test",
			"who am i":                    "who am i",
			"i am {name}":                 "i am {name}",
			"where is the {place.name} ?": "where is the {place.name} ?",
		},
		"zh-CN": {
			"product":                     "产品",
			"design":                      "设计",
			"frontend":                    "前端",
			"backend":                     "后端",
			"test":                        "测试",
			"who am i":                    "我是谁",
			"i am {name}":                 "我是{name}",
			"where is the {place.name} ?": "{place.name}在哪里 ?",
		},
	}

	testify.Equal(tx, "product", t(locales, "en-US", "product"))
	testify.Equal(tx, "test", t(locales, "en-US", "test"))
	testify.Equal(tx, "产品", t(locales, "zh-CN", "product"))
	testify.Equal(tx, "测试", t(locales, "zh-CN", "test"))

	testify.Equal(tx, "who am i", t(locales, "en-US", "who am i"))
	testify.Equal(tx, "我是谁", t(locales, "zh-CN", "who am i"))

	// support not found key
	testify.Equal(tx, "not found", t(locales, "en-US", "not found"))

	// support args
	testify.Equal(tx, "i am zoox", t(locales, "en-US", "i am {name}", map[string]any{"name": "zoox"}))
	testify.Equal(tx, "我是zoox", t(locales, "zh-CN", "i am {name}", map[string]any{"name": "zoox"}))

	// support nest key on args
	testify.Equal(tx, "where is the packing lot ?", t(locales, "en-US", "where is the {place.name} ?", map[string]any{"place": map[string]any{"name": "packing lot"}}))
	testify.Equal(tx, "packing lot在哪里 ?", t(locales, "zh-CN", "where is the {place.name} ?", map[string]any{"place": map[string]any{"name": "packing lot"}}))

}
