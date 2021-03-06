# nebula-go

**IMPORTANT: Code of Nebula go client has been transferred from [nebula-clients](https://github.com/vesoft-inc/nebula-clients) to this repository(nebula-go), and new releases in the future will be published in this repository.
Please update your go.mod and imports correspondingly.**

Official Nebula Go client which communicates with the server using [fbthrift](https://github.com/facebook/fbthrift/).

The code in **master branch** will be updated to accommodate the changes made in Nebula Graph.
To Use the console with a stable release of Nebula Graph, check the branches.
Currently the **[v2.0.0-rc1 branch](https://github.com/vesoft-inc/nebula-go/tree/release-v2.0.0-rc1)** code is compatible with Nebula Graph v2.0.0-rc1.
If you are using a older version, please checkout **[branch v1.0.0](https://github.com/vesoft-inc/nebula-go/tree/v1.0)**.

Please be careful not to modify the files in the nebula directory, these codes were all generated by fbthrift.

## Install

```shell
$ go get -u -v github.com/vesoft-inc/nebula-go@master
```

If you get a message like `cannot use path@version syntax in GOPATH mode`, see the instructions below for [enabling go modules](#enabling-go-modules).

## Usage example

[Code Example](https://github.com/vesoft-inc/nebula-go/tree/master/example/graph_client_example.go)

## Enabling go modules

Dependency management tools are built into go 1.11+ in the form of [go modules](https://github.com/golang/go/wiki/Modules).
If you are using go 1.11 or 1.12 and are working with a project located within `$GOPATH`, you need to do:

```sh
export GO111MODULE=on
```

Ensure your project has a `go.mod` file defined at the root of your project.
If you do not already have one, `go mod init` will create one for you:

```sh
go mod init
```

And then try to get dependencies of `github.com/vesoft-inc/nebula-go` in your go module by simply `go get -u -v github.com/vesoft-inc/nebula-go@master`.
