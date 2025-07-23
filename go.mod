module github.com/freitasgouvea/pq-signer

go 1.21

toolchain go1.21.10

require (
	github.com/open-quantum-safe/liboqs-go v0.2.0
	github.com/spf13/cobra v1.9.1
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)

replace github.com/open-quantum-safe/liboqs-go => ../liboqs-go
