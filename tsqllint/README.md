# tsqllint

[`alpine:edge`](https://hub.docker.com/_/alpine/)-based dockerization of [TSQLLint](https://github.com/tsqllint/tsqllint) the util for listing mssql code

As the site says:

> TSQLLint is a tool for describing, identifying, and reporting the presence of anti-patterns in TSQL scripts.

## Usage

### Interactive

The following shell function can assist in running this image interactively:

```sh

tsqllint() {
  docker run \
    --rm \
    --interactive \
    --tty \
    --volume "$(pwd):/work" \
    "galvanist/conex:tsqllint" \
    "$@"
}

```
