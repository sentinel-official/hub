module github.com/sentinel-official/hub

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.40.0
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.2
	github.com/tendermint/tm-db v0.6.3
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
