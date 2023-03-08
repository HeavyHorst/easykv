module github.com/HeavyHorst/easykv

go 1.16

require (
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1
	github.com/garyburd/redigo v1.6.2
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/hashicorp/consul/api v1.11.0
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.0.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.2 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/vault/api v1.3.0
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/nats-io/nats-server/v2 v2.8.4
	github.com/nats-io/nats.go v1.24.0
	github.com/oklog/run v1.1.0 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/tevino/go-zookeeper v0.0.0-20170512024026-c218ec636bef
	go.etcd.io/etcd/client/pkg/v3 v3.5.4
	go.etcd.io/etcd/client/v2 v2.305.4
	go.etcd.io/etcd/client/v3 v3.5.4
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	google.golang.org/genproto v0.0.0-20220819174105-e9f053255caa // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
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
