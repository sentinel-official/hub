module github.com/sentinel-official/hub

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.33.9
	github.com/tendermint/tm-db v0.5.1
	gopkg.in/yaml.v2 v2.3.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/sentinel-official/cosmos-sdk v0.39.2-sentinel
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
