package i18n

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestNew(t *testing.T) {
	i := New()
	err := i.Load(func() (map[string]map[string]string, error) {
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

		return locales, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "product", i.T("en-US", "product"))
	testify.Equal(t, "test", i.T("en-US", "test"))
	testify.Equal(t, "产品", i.T("zh-CN", "product"))
	testify.Equal(t, "测试", i.T("zh-CN", "test"))

	testify.Equal(t, "who am i", i.T("en-US", "who am i"))
	testify.Equal(t, "我是谁", i.T("zh-CN", "who am i"))

	// support not found key
	testify.Equal(t, "not found", i.T("en-US", "not found"))

	// support args
	testify.Equal(t, "i am zoox", i.T("en-US", "i am {name}", map[string]any{"name": "zoox"}))
	testify.Equal(t, "我是zoox", i.T("zh-CN", "i am {name}", map[string]any{"name": "zoox"}))

	// support nest key on args
	testify.Equal(t, "where is the packing lot ?", i.T("en-US", "where is the {place.name} ?", map[string]any{"place": map[string]any{"name": "packing lot"}}))
	testify.Equal(t, "packing lot在哪里 ?", i.T("zh-CN", "where is the {place.name} ?", map[string]any{"place": map[string]any{"name": "packing lot"}}))
}
