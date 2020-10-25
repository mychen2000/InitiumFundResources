# Artis-Go

This is the backend repo for Artis

## Configuration

Alpaca API credentials

```shell script
cp ./config/config.example.yaml ./config/config.yaml
vim ./config/config.yaml # 或者其他任何你惯用的编辑器
```

## Dependencies

Go 的版本会跟随上游最新版本。因此最好使用 [官网推荐的方式安装](https://golang.org/doc/install) ，而不是 apt 或 brew 源 (which often gives you outdated packages)

所有的项目依赖都在 [./go.mod](./go.mod) 文件中，而 [./go.sum](./go.sum) 则列出了各个依赖的 checksum。

go 会在执行 `go build`, `go run` 命令时自动安装所需要的依赖，但你也可以手动使用 `go mod download` 来安装所有依赖。这些依赖默认会被安装到你的 GOPATH 目录中，你也可以使用 `go mod vendor` 来将所有依赖的代码存储在 `/vendor` 目录下。

## Run

你可以先将程序编译成二进制文件再运行，也可以直接通过 `go run` 运行程序

```shell script
# 测试 API 是否可用
go build ./cmd/test && ./test
```

```shell script
go run ./cmd/test/test.go
```

## 代码风格

Go 有内建的 formatter，你可以通过运行以下命令来 format 所有 `.go` 的代码文件

```
go fmt -x ./...
```
# Artis-Go-Duplicate
# Artis-Go-Duplicate
