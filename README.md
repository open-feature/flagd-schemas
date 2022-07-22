# schemas
Schemas and spec files pertaining to OpenFeature

<br />

**Golang:**  

generate grpc-gateway, go-grpc and go stubs
```
make gen-go-server
```
generate go-grpc and go stubs
```
make gen-go
```
go package import
```
go get go.buf.build/grpc/go/open-feature/flagd
```  
<br />

**Typescript:**  

generate typescript stubs for grpc client and http/grpc input/output schema
```
make gen-ts
```