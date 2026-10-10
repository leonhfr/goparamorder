# goparamorder

A Go linter that enforces a configured parameter order in function declarations.

It runs as a [golangci-lint](https://golangci-lint.run/) module plugin, built on [`golang.org/x/tools/go/analysis`](https://pkg.go.dev/golang.org/x/tools/go/analysis).

## Rules

- `context.Context`, if present, must be the first parameter.
- Every other parameter whose type appears in the `order` setting must follow the relative order given there.
- Parameter types absent from `order`, including every predeclared type (`string`, `int`, `error`, `any`, ...), are unchecked: their position is free.
- Matching is by exact declared type, resolved from the file's own import list. A pointer matches the same entry as its value type (`*pkg.Config` matches `pkg.Config`).

The `scope` setting selects which declaration kinds are checked:

- `funcdecl`: function and method declarations. On by default.
- `funclit`: function literals. Off by default.
- `interface`: interface method signatures. Off by default.

A key absent from `scope` keeps its default.

The linter does not suggest fixes. Reordering a signature breaks every call site, since Go arguments are positional; reordering is left to the author.

## Installation

Build a custom golangci-lint binary that includes this plugin. Add it to `.custom-gcl.yml`:

```yaml
version: v2.14.0
name: goparamorder-golangci-lint
destination: bin
plugins:
  - module: github.com/leonhfr/goparamorder
    version: latest
```

Then build it:

```bash
golangci-lint custom
```

See the [module plugin system docs](https://golangci-lint.run/docs/plugins/module-plugins/) for the general mechanism.

## Configuration

Enable the linter in `.golangci.yml` and configure it under `linters.settings.custom`:

```yaml
linters:
  enable:
    - goparamorder
  settings:
    custom:
      goparamorder:
        type: module
        description: Enforces a configured parameter order in function declarations.
        original-url: github.com/leonhfr/goparamorder
        settings:
          # context.Context is handled separately; do not list it here.
          order:
            - github.com/example/project.Config
            - github.com/example/project.Logger
          scope: {} # e.g. {funclit: true} to also check function literals; {funcdecl: false} to turn off the default
```

## Go API

The analyzer is also usable directly, outside golangci-lint, with any `go/analysis` driver such as `singlechecker` or `analysistest`:

```go
import "github.com/leonhfr/goparamorder"

analyzer := goparamorder.New(goparamorder.Settings{
	Order: []string{"github.com/example/project.Config", "github.com/example/project.Logger"},
})
```

## License

MIT, see [LICENSE](LICENSE).
