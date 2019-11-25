### 安装

####安装grpc
```
go get -u google.golang.org/grpc
```

####安装Protoc Plugin
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

####安装Protocol Buffers v3
```
wget https://github.com/google/protobuf/releases/download/v3.5.1/protobuf-all-3.5.1.zip
unzip protobuf-all-3.5.1.zip
cd protobuf-3.5.1/
./configure
make
make install
ldconfig
protoc --version
```

####protoc使用
```
protoc --go_out=plugins=grpc,import_path=mypackage:. *.proto
```

* `--go_out=.：`设置 Go 代码输出的目录,该指令会加载 protoc-gen-go 插件达到生成 Go 代码的目的，生成的文件以 .pb.go 为文件后缀
* `import_prefix=xxx：`将指定前缀添加到所有import路径的开头
* `import_path=foo/bar：`如果文件没有声明go_package，则用作包。如果它包含斜杠，那么最右边的斜杠将被忽略。
* `plugins=plugin1+plugin2：`指定要加载的子插件列表（我们所下载的repo中唯一的插件是grpc）
* `Mfoo/bar.proto=quux/shme：` M参数，指定.proto文件编译后的包名（foo/bar.proto编译后为包名为quux/shme）

