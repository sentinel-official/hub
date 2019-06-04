module github.com/ironman0x7b2/sentinel-sdk

go 1.12

require (
	github.com/cosmos/cosmos-sdk v0.34.7
	github.com/etcd-io/bbolt v1.3.2 // indirect
	github.com/gorilla/mux v1.7.2
	github.com/rakyll/statik v0.1.6
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.31.7
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
