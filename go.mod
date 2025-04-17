module github.com/HeavyHorst/easykv

go 1.23.0

toolchain go1.24.0

require (
	github.com/fsnotify/fsnotify v1.5.1
	github.com/garyburd/redigo v1.6.2
	github.com/hashicorp/consul/api v1.11.0
	github.com/hashicorp/vault/api v1.16.0
	github.com/nats-io/nats-server/v2 v2.9.23
	github.com/nats-io/nats.go v1.39.1
	github.com/tevino/go-zookeeper v0.0.0-20170512024026-c218ec636bef
	go.etcd.io/etcd/client/pkg/v3 v3.5.4
	go.etcd.io/etcd/client/v2 v2.305.4
	go.etcd.io/etcd/client/v3 v3.5.4
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/go-jose/go-jose/v4 v4.0.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.9 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.7 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.9.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.5.0 // indirect
	github.com/nats-io/nkeys v0.4.10 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	go.etcd.io/etcd/api/v3 v3.5.4 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.56.3 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace (
	go.etcd.io/etcd/api/v3 v3.5.4 => go.etcd.io/etcd/api/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 => go.etcd.io/etcd/client/pkg/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/client/v3 v3.5.4 => go.etcd.io/etcd/client/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/etcdctl/v3 v3.5.4 => go.etcd.io/etcd/etcdctl/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/pkg/v3 v3.5.4 => go.etcd.io/etcd/pkg/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/raft/v3 v3.5.4 => go.etcd.io/etcd/raft/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/server/v3 v3.5.4 => go.etcd.io/etcd/server/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/tests/v3 v3.5.4 => go.etcd.io/etcd/tests/v3 v3.5.5-0.20220808200321-9e95685d0a6d
	go.etcd.io/etcd/v3 v3.5.4 => go.etcd.io/etcd/v3 v3.5.5-0.20220808200321-9e95685d0a6d
)
