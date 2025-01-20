# protoc_inject_tag

# 基于开源组件 protoc-go-inject-tag(BSD-2-Clause license) 开发，解决 **swagger-ui注释** 和 json自动生成 **omitempty** 的要求

# 如何使用

## 本地使用

### step 1:

#### Example

```proto
// file: test.proto  
syntax = "proto3";  

package pb;  
option go_package = "/pb";  

message IP {  
  // @gotags: description:"Address"  string Address = 1;  
   // @gotags: description:"Address"  string MAC = 2;}  
```

描述信息只能加在参数上方，因为在 Protocol Buffers（proto3）中，参数右边的注释不会直接传递到生成的代码中。  
proto3 仅保留字段定义上方的注释，并且不支持将注释传递到生成的代码中。

您也可以在生成的.pb.go代码中把注释放在右边，然后再调用protoc_inject_tag方法

Generate your `.pb.go` files with the protoc command as normal:

```console
 protoc --go_out=plugins=grpc:. ./pb/serving/*.proto 
```

#### step 2:

Then run `protoc-go-inject-tag` against the generated files (e.g `test.pb.go`):

执行 pb.protoc_inject_tag.main()方法，自动替换原始的proto.go文件。

## docker内使用

### step 1:

使用Dockerfile文件构建:

```
docker build -t 127.0.0.1:5000/tmp/protogen:with-inject-tag .
```

### step 2:

```
./gen.sh
```
