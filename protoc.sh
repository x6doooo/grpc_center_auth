#!/usr/bin/env bash

protoc3='/Users/yangdongxu/work/protoc-3.0.2-osx-x86_64/bin/protoc'
${protoc3} -I pbs/auth/ pbs/auth/auth.proto --go_out=plugins=grpc:pbs/auth
