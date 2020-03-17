# conex

This is a repository for utility containers. The latest versions are mirrored to docker hub.

## [`audiosprite`](audiosprite/Dockerfile)

This is an alpine-based dockerization of [audiosprite](https://github.com/tonistiigi/audiosprite).

### interactive use


```sh
audiosprite() {
  docker run \
    --rm \
    --interactive \
    --tty \
    --volume "$(pwd):/data" \
    "galvanist/conex:audiosprite" \
    "$@"
}
```

## [`checkmake`](checkmake/Dockerfile)

Alpine-based containerization of [checkmake](https://github.com/mrtazz/checkmake/), the `Makefile` linter.

### interactive use

This shell function demonstrates using this container in place of having the binary installed.

```sh
checkmake() {
  docker run \
    --rm \
    --interactive \
    --tty \
    --volume "$(pwd):/work" \
    "galvanist/conex:checkmake" \
    "$@"
}
```

Examples sessions using the function:

```sh
$ checkmake
checkmake.

  Usage:
  checkmake [options] <makefile>
  checkmake -h | --help
  checkmake --version
  checkmake --list-rules

  Options:
  -h --help               Show this screen.
  --version               Show version.
  --debug                 Enable debug mode
  --config=<configPath>   Configuration file to read
  --format=<format>       Output format as a Golang text/template template
  --list-rules            List registered rules
```

Checking against this repo's `Makefile`:

```
$ checkmake Makefile 
                                                                
      RULE                 DESCRIPTION             LINE NUMBER  
                                                                
  maxbodylength   Target body for "push" exceeds   25           
                  allowed length of 5 (10).                     
  minphony        Missing required phony target    0            
                  "all"                                         
  minphony        Missing required phony target    0            
                  "test"                                        
```

## [`goenv`](goenv/Dockerfile)

### interactive use

This shell function demonstrates using this container in place of having the actual go installation.

```sh
goenv() {
  docker run \
    --rm \
    --interactive \
    --tty \
    --volume "$(pwd):/home/user/go/src/local" \
    "$@" \
    galvanist/conex:goenv
}
```

Then you just cd to a directory with a go project (or an empty directory) and run `goenv`.

## [`grta`](grta/Dockerfile)

This HTTP endpoint receives webhooks, validates against the PSK, writes the webhook payload to a file (or fifo). Meant to be used behind a load balancer that provides TLS.

## [`hugo`](hugo/Dockerfile)

This is a `debian:stable-slim`-based containerization of hugo-extended. I use it as a builder in multi-stage container builds, I also run it interactively during development.

### interactive use

This shell function demonstrates using this container in place of having the actual hugo binary.

```sh
hugo() {
  insert=""
  if [ "$1" = "server" ] || [ "$1" = "serve" ]; then
    shift;
    insert="server --bind 0.0.0.0 --port 1313"
  fi

  # shellcheck disable=SC2086
  docker run \
    --rm \
    --interactive \
    --tty \
    --volume "$(pwd):/work" \
    --publish "1313:1313" \
    galvanist/conex:hugo \
    hugo \
      $insert \
      "$@"
}
```

You just add the above function to your shell rc file then use hugo normally:

```sh
$ hugo
```

or

```sh
$ hugo serve -D
```

### multistage build usage

```Dockerfile
FROM galvanist/conex:hugo as builder

COPY . .
RUN hugo

FROM nginx:1-alpine as server
COPY --from=builder /work/public/ /usr/share/nginx/html/
```

## [`vueenv`](vueenv/Dockerfile)

This container image is meant to be used as a builder stage for Vue CLI-based apps in a multi-stage build. It is also very useful during development.

### Multistage Build Use

```Dockerfile
FROM galvanist/vueenv:latest as builder

COPY src /app

RUN npm install \
  && npm run build

FROM nginx:1-alpine as server
COPY --from=builder /app/dist/ /usr/share/nginx/html/

# maybe also something like this:
# COPY nginx_conf.d/* /etc/nginx/conf.d/
```

### Interactive Use

Add this to your shell profile:

```sh
vueenv() {
  docker run \
    --rm \
    --interactive \
    --tty \
    --volume "$(pwd):/app" \
    "$@" \
    "galvanist/conex:vueenv"
}
```