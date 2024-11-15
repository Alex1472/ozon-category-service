#!/bin/sh

GRPC_HOST="localhost:7002"
GRPC_METHOD="alex1472.ozon_film_service.film_service.v1.FilmService/GetCategoryById" # package in .proto

payload=$(
  cat <<EOF
{
  "id": 4
}
EOF
)

grpcurl -plaintext -emit-defaults \
  -rpc-header 'x-app-name:dev' \
  -rpc-header 'x-app-version:1' \
  -d "${payload}" ${GRPC_HOST} ${GRPC_METHOD}
