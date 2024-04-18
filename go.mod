module go.indent.com/indent-go

go 1.21.3

toolchain go1.22.2

require (
	github.com/golang/protobuf v1.5.3
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.11.0
	golang.org/x/oauth2 v0.10.0
	golang.org/x/sys v0.10.0
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97
	google.golang.org/genproto/googleapis/api v0.0.0-20231009173412-8bfb1ae86b6c
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97
	google.golang.org/grpc v1.58.3
	google.golang.org/protobuf v1.33.0
	k8s.io/apimachinery v0.26.1
	multiparty.ai v0.0.0-00010101000000-000000000000
)

replace k8s.io/apimachinery v0.26.1 => github.com/indentinc/apimachinery v0.0.0-20220303210610-038787720cef

replace multiparty.ai => github.com/indentapis/multiparty v0.1.11

require (
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/term v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d // indirect
)
