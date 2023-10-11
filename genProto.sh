#!/bin/bash
function genCommonProto {
    DOMAIN=$1
    PROTO_PATH=./pkg/common/proto/${DOMAIN}
    GO_OUT_PATH=./pkg/common/proto/${DOMAIN}
    PROTO_NAME=${DOMAIN}
    mkdir -p $GO_OUT_PATH
    protoc -I=$PROTO_PATH --proto_path=./pkg/common/proto --go_out=paths=source_relative:$GO_OUT_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
}

genCommonProto page

function genProto {
    DOMAIN=$1
    SKIP_GATEWAY=$2
    PROTO_PATH=./server/${DOMAIN}/api
    GO_OUT_PATH=./server/${DOMAIN}/api/gen/v1
    mkdir -p $GO_OUT_PATH

    #protoc -I=$PROTO_PATH --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    if [ $DOMAIN = "member" ]; then
      for srv in user address;
      do
        OUT_PATH=${GO_OUT_PATH}/${srv}
        mkdir -p $OUT_PATH
        # 这里配置 go 依赖包 路径
        protoc -I=$PROTO_PATH --proto_path=./pkg/common/proto --go_out=paths=source_relative:$OUT_PATH --go-grpc_out=paths=source_relative:$OUT_PATH ${srv}.proto
      done
    else
      protoc -I=$PROTO_PATH --proto_path=./pkg/common/proto --go_out=paths=source_relative:$GO_OUT_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    fi



#    if [ $SKIP_GATEWAY ]; then
#        return
#    fi
#
#    protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml:$GO_OUT_PATH ${DOMAIN}.proto
}

genProto user