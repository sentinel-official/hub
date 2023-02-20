module github.com/sentinel-official/hub

go 1.20

require (
	github.com/CosmWasm/wasmd v0.29.2
	github.com/cosmos/cosmos-sdk v0.46.10
	github.com/cosmos/ibc-go/v3 v3.4.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/prometheus/client_golang v1.14.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.34.26
	github.com/tendermint/tm-db v0.6.7
	google.golang.org/genproto v0.0.0-20221014213838-99cd37c6964a
	google.golang.org/grpc v1.50.1
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/informalsystems/tendermint v0.34.26
)
