CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:BEKJONS@localhost:5432/collaboration_m?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:123321@localhost:5432/collaboration_m?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database 'postgres://postgres:BEKJONS@localhost:5432/collaboration_m?sslmode=disable' -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_collaboration
