

.PHONY: bindata.go pcctv.json

generate:
	gen route --openapi -o ./cmd/areacn/router_gen.go github.com/wzshiming/areacn/service/...

bindata.go: pcctv.json
	go-bindata --pkg areacn pcctv.json

pcctv.json:
	cd ./hack/ && go run update.go
	mv ./hack/pcctv.json ./