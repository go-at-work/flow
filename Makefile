mock:
    mockery --all --keeptree

migrate:
    migrate -source file://postgres/migrations \
            -database postgres://postgres:postgres@127.0.0.1:5432/flow_dev?sslmode=disable up

rollback:
    migrate -source file://postgres/migrations \
            -database postgres://postgres:postgres@127.0.0.1:5432/flow_dev?sslmode=disable down

drop:
    migrate -source file://postgres/migrations \
            -database postgres://postgres:postgres@127.0.0.1:5432/flow_dev?sslmode=disable drop

migration:
@read -p "Enter migration name: " name; \
    migrate create -ext sql -dir postgres/migrations -seq $$name

run:
   go run cmd/graphqlserver/*.go

generate:
    go generate ./...   
