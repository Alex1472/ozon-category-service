module github.com/Alex1472/ozon-category-service

go 1.22.7

toolchain go1.23.3

require (
	github.com/Alex1472/ozon-category-service/pkg/category-service v0.0.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/lib/pq v1.10.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.24.0
	golang.org/x/net v0.29.0
	google.golang.org/grpc v1.68.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)

replace github.com/Alex1472/ozon-category-service/pkg/category-service => ./pkg/category-service
