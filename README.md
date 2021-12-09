# HTTP API SERVER

## Install choco
### Run administrative powershell.exe, past command line and press ENTER

`@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"`


## Install scoop for migrate
### Run administrative powershell.exe, past command line and press ENTER
1. `Set-ExecutionPolicy RemoteSigned -scope CurrentUser`
2. `Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')`
3. `iwr -useb get.scoop.sh | iex`

Install migrate `scoop install migrate`

## Install MinGW, GCC, G++ and Make
`choco install mingw`

## Create Makefile

`.PHONY: build
build:
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build`

## Install toml-file parser

`go get github.com/BurntSushi/toml`

## In terminal command for create binary file

`make`

### View create binary file, your config-path
`./apiserver -help`

## Library log error
`go get github.com/sirupsen/logrus`

## Testing your code
`go get github.com/stretchr/testify`

## Create migration
`migrate create -ext sql -dir migrations <name-migration>`


## Migrate db up, for connect to postgress
`postgres://user:password@localhost/db_name?sslmode=disable` up