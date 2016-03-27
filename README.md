# editd

[![Circle CI](https://img.shields.io/circleci/project/thavel/editd/master.svg)](https://circleci.com/gh/thavel/editd)

`editd` is a lightweight publisher-side client to [etcd](https://github.com/coreos/etcd) focused on:

* publishing new configs.
* keeping key-values stored in etcd up-to-date with the local configurations.

This tool is the missing link between a system and its etcd-[confd](https://github.com/kelseyhightower/confd) cluster.


## Building

Here are the requirements to build editd:

* [go](https://golang.org/) 1.6
* [gb](https://getgb.io/)

Then, all you need is to compile the tool from the sources:

```bash
git clone https://github.com/thavel/editd.git
cd editd
gb vendor restore
gb build
```

You should now have an `editd` binary in the `bin/`.

```bash
$ ls bin/
editd
```


## Getting started

All you need to know is available in the binary help (using `-h`).

### Quick start

Here are the minimal arguments to use editd:

```bash
editd -node localhost:4001 -key /hello -value world
```

### Optional arguments

| Flag        | Default | Description                   |
|-------------|---------|-------------------------------|
| `-interval` | 5000    | Synchronization interval      |
| `-onetime`  | No      | Run once and exit             |
| `-ttl`      | 10000   | TTL duration for keys         |
| `-nottl`    | No      | Disable TTL duration for keys |
| `-safe`     | No      | Exit upon errors              |
| `-fvalue`   | n/a     | Use file content as value     |


## Upcoming features

* Handle different backends (for now, only _etcd_ is supported).
* Improve signal handling.
* Support secured requests.
