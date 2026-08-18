# Contributing

OctantR is an early community revival. Use GitHub issues and pull requests for project communication.

## Before starting

- Search [existing issues](https://github.com/m8yng/octant/issues).
- Open an issue before a large feature, public API change, or major refactor.
- Keep changes focused and preserve compatibility unless the issue explicitly changes it.

## Development

Follow [HACKING.md](HACKING.md) to configure the development environment.

Before opening a pull request, run the checks relevant to your change:

```sh
go run build.go vet
go run build.go test
go run build.go web-lint
go run build.go web-test
```

For a complete build:

```sh
go run build.go ci
```

## Pull requests

- Explain what changed and why.
- Link the relevant issue when one exists.
- Add or update tests for non-trivial behavior.
- Update documentation when behavior or commands change.
- Use concise commit messages.

For user-visible changes, add a file under `changelogs/unreleased/` named after the issue or pull request and author, for example:

```text
changelogs/unreleased/123-username
```

## Licensing

By contributing, you agree that your original contribution is licensed under the [MIT License](LICENSE) and that you have the right to submit it. Inherited Octant code remains under the [Apache License 2.0](LICENSE-APACHE-2.0).
