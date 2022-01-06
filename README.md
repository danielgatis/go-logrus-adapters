# Go - Logrus Adapters

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-logrus-adapters?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-logrus-adapters)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-logrus-adapters/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-logrus-adapters)

A collection of adapters for logrus pkg.

## Install

```bash
go get -u github.com/danielgatis/go-logrus-adapters
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-logrus-adapters"
```

### Example

```go
package main

import (
	"net/http"

	adapters "github.com/danielgatis/go-logrus-adapters"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	e.Logger = adapters.NewEchoLogAdapter(logrus.StandardLogger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
```

### License

Copyright (c) 2021-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)

### Buy me a coffee

Liked some of my work? Buy me a coffee (or more likely a beer)

<a href="https://www.buymeacoffee.com/danielgatis" target="_blank"><img src="https://bmc-cdn.nyc3.digitaloceanspaces.com/BMC-button-images/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;"></a>
