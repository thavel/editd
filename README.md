# editd

`editd` is a lightweight publisher-side client to [etcd](https://github.com/coreos/etcd) focused on:

* publishing new configs.
* keeping key-values stored in etcd up-to-date with the local configurations.

This tool is the missing link between a system and its etcd-[confd](https://github.com/kelseyhightower/confd) cluster.

## Building

Here are the requirements to build editd:

* [go](https://golang.org/) 1.6
* [gb](https://getgb.io/)

Then, all you need is to compile the tool using the sources:

```bash
git clone https://github.com/thavel/editd.git
cd editd
gb build
```

## Getting started

TODO
