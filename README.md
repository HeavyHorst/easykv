[![Build Status](https://travis-ci.org/HeavyHorst/easykv.svg?branch=master)](https://travis-ci.org/HeavyHorst/easykv) [![Go Report Card](https://goreportcard.com/badge/github.com/HeavyHorst/easykv)](https://goreportcard.com/report/github.com/HeavyHorst/easykv) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/HeavyHorst/easykv/master/LICENCE)
[![](https://godoc.org/github.com/HeavyHorst/easykv?status.svg)](http://godoc.org/github.com/HeavyHorst/easykv)


# easyKV
easyKV is based on the backends of [confd](https://github.com/kelseyhightower/confd).
easyKV provides a very simple interface to work with some key-value stores.
The goal of easyKV is to abstract these 2 common operations for multiple backends:

  - recursively query the kv-store for key-value pairs.
  - watch a key-prefix for changes.

## Interface
A **storage backend** in `easyKV` should implement (fully or partially) this interface:
```go
type ReadWatcher interface {
	GetValues(keys []string) (map[string]string, error)
	WatchPrefix(ctx context.Context, prefix string, opts ...WatchOption) (uint64, error)
	Close()
}
```

## Compatibility matrix

| Calls                 |   Consul   | Etcdv2 | Etcdv3  |  env  | file |   redis |  vault  |  zookeeper | nats kv |
|-----------------------|:----------:|:------:|:-------:|:-----:|:----:|:-------:|:-------:|:----------:|:-------:|
| GetValues             |     X      |   X    |      X  |    X  |  X   |     X   |   X     |     X      |    X    |
| WatchPrefix           |     X      |   X    |      X  |       |  X   |         |         |     X      |    X    |
| Close                 |     X      |   X    |      X  |    X  |  X   |     X   |   X     |     X      |    X    |
