keyring-wincred
===============
[![CI](https://github.com/lox/keyring-wincred/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lox/keyring-wincred/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lox/keyring-wincred.svg)](https://pkg.go.dev/github.com/lox/keyring-wincred)

Windows Credential Manager provider for [`github.com/lox/keyring/v2`](https://github.com/lox/keyring).

## Usage

```bash
go get github.com/lox/keyring-wincred
```

```go
import (
	"context"

	"github.com/lox/keyring/v2"
	wincred "github.com/lox/keyring-wincred"
)

ctx := context.Background()

ring, err := keyring.Open(ctx,
	keyring.WithServiceName("example"),
	keyring.WithProvider(wincred.Provider()),
)
```

`wincred.Provider` accepts the `Prefix` option. On non-Windows platforms, it
returns `keyring.ErrUnavailable` during open.
