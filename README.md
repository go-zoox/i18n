# i18n - 🏳️An simple i18n messages manage implementation for Go.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/i18n)](https://pkg.go.dev/github.com/go-zoox/i18n)
[![Build Status](https://github.com/go-zoox/i18n/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/i18n/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/i18n)](https://goreportcard.com/report/github.com/go-zoox/i18n)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/i18n/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/i18n?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/i18n.svg)](https://github.com/go-zoox/i18n/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/i18n.svg?label=Release)](https://github.com/go-zoox/i18n/releases)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/i18n
```

## Getting Started

```go
package main

import (
	"context"

	"github.com/go-zoox/i18n"
)

func main() {
	i := i18n.New()
	err := i.Load(func (lang string) (map[string][string]string, error) {
		// loads locales
		return map[string][string]string{
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
		}, nil
	})
	if err != nil {
		panic(err)
	}

	translation := i.T("en-US", "product")

	fmt.Println(translation) // product
}
```

### Load From Directory
```go
package main

import (
	"context"

	"github.com/go-zoox/i18n"
)

// The directory structure should be like this:
// lang/
//
//		en-US.json
//		zh-CN.json
//		en-US.yaml
//		en-US.toml
//	 	en-US.ini
//		...

func main() {
	i := i18n.New()
	err := i.LoadFromDir("./lang")
		if err != nil {
		panic(err)
	}

	translation := i.T("en-US", "product")

	fmt.Println(translation) // product
}
```

### Load From URL
```go
package main

import (
	"context"

	"github.com/go-zoox/i18n"
)

func main() {
	i := i18n.New()
	err := i.LoadFromURL("https://raw.githubusercontent.com/go-zoox/i18n/master/tests/locales.json")
		if err != nil {
		panic(err)
	}

	translation := i.T("en-US", "product")

	fmt.Println(translation) // product
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
