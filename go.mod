module github.com/cz-theng/ethspider

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/cz-theng/czkit-go v0.0.0-20210115085902-7a4e7f362a94
	github.com/ethereum/go-ethereum v1.9.25 
	github.com/spf13/cobra v1.1.1
)

replace (
	github.com/ethereum/go-ethereum v1.9.25  => github.com/maticnetwork/bor v0.2.3
)
