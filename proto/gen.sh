
# 根据proto文件生成go源码
protoc --go_out=plugins=grpc:. *.proto

# 生成API文档的swagger.json文件
protoc -I /usr/local/include -I . -I $HOME/workspace/gocode/helloGRPC/proto/google/api --swagger_out=logtostderr=true:. ./hello.proto

if [ $? -eq 0 ]; then
    docker run --rm -d -p 8888:8080 -e SWAGGER_JSON=/foo/hello.swagger.json -v $(pwd):/foo swaggerapi/swagger-ui
fi

if [ $? -eq 0 ]; then
    echo "succeed"
else
    echo "failed"
fi

