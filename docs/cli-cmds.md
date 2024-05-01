# util

## Usage

Executa os testes de unidade e gera o relatório de cobertura.

```shell
go test -tags wireinject -failfast -timeout=30s -count=1 -cover -coverprofile="$PWD/test/coverage/coverage.out" -fullpath -race `go list ./... | grep -v ./test`
```
