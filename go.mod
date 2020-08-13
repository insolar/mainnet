module github.com/insolar/mainnet

go 1.12

require (
	github.com/go-pg/pg/v9 v9.1.6
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/google/gops v0.3.6
	github.com/insolar/insolar v1.6.6-0.20200727081822-68505e6410cb
	github.com/insolar/x-crypto v0.0.0-20191031140942-75fab8a325f6
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.6.1
	github.com/tylerb/is v2.1.4+incompatible // indirect
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d // indirect
	golang.org/x/net v0.0.0-20200222033325-078779b8f2d8 // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	golang.org/x/tools v0.0.0-20191108193012-7d206e10da11
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/insolar/mainnet => ./

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
