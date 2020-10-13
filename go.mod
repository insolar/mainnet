module github.com/insolar/mainnet

go 1.15

require (
	github.com/google/gops v0.3.6
	github.com/insolar/insolar v1.9.0
	github.com/insolar/x-crypto v0.0.0-20191031140942-75fab8a325f6
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0
	golang.org/x/tools v0.0.0-20191108193012-7d206e10da11
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/insolar/mainnet => ./

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
