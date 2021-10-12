#!/usr/bin/env bash

protoc -Iapi/protobuf/exposure/toutiao --proto_path=/opt/include/ api/protobuf/exposure/toutiao/toutiao_rta.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/tencent api/protobuf/exposure/tencent/tencent_rta.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/tencent api/protobuf/exposure/tencent/tencent_rta_stat.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/baidu api/protobuf/exposure/baidu/baidu_rta.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/kuaishou api/protobuf/exposure/kuaishou/kuaishou_rta.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/duomeng api/protobuf/exposure/duomeng/duomeng_rta.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/qutoutiao api/protobuf/exposure/qutoutiao/qutoutiao_rta.proto --go_out=plugins=grpc:./

protoc -Iapi/protobuf/exposure/vivo api/protobuf/exposure/vivo/vivo_rta.proto --go_out=plugins=grpc:./
