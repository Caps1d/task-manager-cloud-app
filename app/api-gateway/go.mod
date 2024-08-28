module github.com/Caps1d/task-manager-cloud-app/api-gateway

go 1.22.0

require (
	github.com/Caps1d/task-manager-cloud-app/user v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.19.0
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	golang.org/x/time v0.5.0 // indirect
)

require (
	github.com/Caps1d/task-manager-cloud-app/auth v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.12.0
	github.com/labstack/gommon v0.4.2
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2 // indirect
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-playground/form/v4 v4.2.1
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Caps1d/task-manager-cloud-app/auth => ../auth

replace github.com/Caps1d/task-manager-cloud-app/user => ../user

replace github.com/Caps1d/task-manager-cloud-app/task => ../task

replace github.com/Caps1d/task-manager-cloud-app/notification => ../notification
