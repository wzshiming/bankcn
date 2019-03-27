

.PHONY: bindata.go

generate:
	gen route --openapi -o ./cmd/bankcn/router_gen.go github.com/wzshiming/bankcn/service/...

bindata.go:
	cd ./hack/ && go run update.go
	mv ./hack/banks.json ./
	rm -r ./icon && mv ./hack/icon ./icon
	go-bindata --pkg bankcn banks.json icon

