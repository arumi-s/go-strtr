# strtr

[![test](https://github.com/arumi-s/go-strtr/actions/workflows/test.yml/badge.svg)](https://github.com/arumi-s/go-strtr/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/arumi-s/go-strtr.svg)](https://pkg.go.dev/github.com/arumi/go-strtr)

This package is a Go implementation of the `strtr` function from [PHP](https://www.php.net/manual/en/function.strtr.php).

## Installation

```bash
go get github.com/arumi-s/go-strtr
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/arumi-s/go-strtr"
)

func main() {
	fmt.Println(strtr.Strtr("Hello World", map[string]string{"Hello": "Hi", "World": "Planet"}))
	// Output: Hi Planet
}
```

## License

[MIT](https://github.com/arumi-s/go-strtr/blob/main/LICENSE)
