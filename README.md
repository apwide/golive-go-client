# Golive Client Generation

* [openapi-client-generator deepmap](https://github.com/deepmap/oapi-codegen)
* [doc-openapi-generator](https://medium.com/@kyodo-tech/go-client-code-generation-from-swagger-and-openapi-a0576831836c)

## Procedure
* go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
* oapi-codegen -config gen-config.yaml golive-api.yaml
* go mod tidy

On linux mint, don't forget to set GOPATH to ~/go and add GOPATH to your PATH