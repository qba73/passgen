[![Go](https://github.com/qba73/passgen/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/passgen/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/passgen)](https://goreportcard.com/report/github.com/qba73/passgen)

# passgen

`passgen` is a Go library and CLI tool for generating passwords.

## Installing the command-line client

```bash
go install github.com/qba73/passgen/cmd/passgen@latest
```

## Using the command-line client

```bash
$ passgen
{
  "salt": "bPDkbaVZljXhWJ5tMDY6fwzAO9keqZ3woJNYfd_ZcZw=",
  "password": "chatr1519fhph7hf1igg",
  "hash": "66d47fed894fd89cfd94983b95dbb986565fec1e41a8a17b7634e1e4da1f8247"
}
```

Since the `passgen` outputs data in the [JSON](https://www.json.org/json-en.html) format you can pipe the output to [jq](https://stedolan.github.io/jq/):

```bash
$ passgen | jq .password
"ch91jul19fhv18u7m1i0"
```
