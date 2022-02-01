# DEPRECATED

**Archiving this repository in favour of https://github.com/nats-io/natscli**

# nats-cli

simple cli for nats pub/sub

## Install

```sh
$ go get -u github.com/moonwalker/nats-cli
```

## Usage

Subscribe:

```sh
$ nats-cli sub 'FOO'
```

Publish:

```sh
$ nats-cli pub 'FOO' '{ "baz": "bar" }'
```
