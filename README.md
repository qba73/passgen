[![Go](https://github.com/qba73/passgen/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/passgen/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/passgen)](https://goreportcard.com/report/github.com/qba73/passgen)

# passgen

`passgen` is a helper CLI tool for generating passwords.

## Build the passgen

```bash
go build -o passgen cmd/passgen/main.go
```

## Usage

Run passgen:

```bash
$ passgen
{
  "salt": "vo4zUAvV7kcPzttGkzHSL79hvWFJt6AJXT1tEu35ja8=",
  "password": "ch91ipt19fhv0p5d3n20",
  "hash": "d94e77c02f69966cd1a5111a79c572b19208db7e0ffdb6816f8b4ed189f562be"
}
```

Since the `passgen` outputs data in the [JSON](https://www.json.org/json-en.html) format you can pipe the output to [jq](https://stedolan.github.io/jq/):

```bash
$ passgen | jq .password
"ch91jul19fhv18u7m1i0"
```
