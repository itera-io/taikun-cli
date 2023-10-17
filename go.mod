module github.com/itera-io/taikun-cli

go 1.17

require (
	github.com/Smidra/taikungoclient v0.0.0-20231017062658-a9e57a75534d
	github.com/Smidra/taikungoclient/client v0.0.0-20231017062658-a9e57a75534d
	github.com/go-openapi/strfmt v0.21.7
	github.com/jedib0t/go-pretty/v6 v6.3.1
	github.com/spf13/cobra v1.4.0
)

//replace github.com/Smidra/taikungoclient => /home/radek/taikun/taikun-cli/taikungoclient-radek-repo
//replace github.com/itera-io/taikungoclient => /home/radek/taikun/taikun-cli/taikungoclient-repo

require (
	github.com/Smidra/taikungoclient/showbackclient v0.0.0-20231006131231-05bf431ea403 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/go-openapi/errors v0.20.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	golang.org/x/sys v0.5.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
