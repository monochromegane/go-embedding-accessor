# go-embedding-accessor [![Actions Status](https://github.com/monochromegane/go-embedding-accessor/workflows/Go/badge.svg)](https://github.com/monochromegane/go-embedding-accessor/actions)

An accessor generator for files embedded with go:embed.

## Usage

1. Install go-embedding-accessor.
2. Run `go-embedding-accessor` command at your repository.

```sh
$ go-embedding-accessor --pkg PACKAGE_NAME --files 'assets/*.txt' --name asset
```

You can find `asset_embedding.go` in your repository.
So, your Go program get {show,list,restore}-{name} options.
In the following, the option name generated when "asset" is specified as name is used.

### List embedding files

`list-asset` list your embedding files.

```sh
$ my-app --list-asset
assets/a.txt
assets/b.txt
```

### Show embedding file

`show-asset` show embedding file.

```sh
$ my-app --show-asset assets/a.txt
Contents of a.txt
```

### Restore embedding files

`restore-assets` restore embedding files.

```sh
$ my-app --restore-asset
```

You also specify output path by `--restore-path` option.

## Generated code dependency

Generated code (`asset_embedding.go`) doesn't depend on your code, only provides {show,list,restore}-asset options at init() function.


## Installation

```sh
$ go install github.com/monochromegane/go-embedding-accessor/cmd/go-embedding-accessor@latest
```

## Contribution

1. Fork it
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/monochromegane/go-embedding-accessor/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
