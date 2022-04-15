# Generate service stubs from protos
codegen: update-submodules
	protoc --go_out=. --go-grpc_out=. grpc/protos/accounts/*.proto
	protoc --go_out=. --go-grpc_out=. grpc/protos/notes/*.proto

# Fetch the latest version of the protos submodule.
update-submodules:
	git submodule init
	git submodule update --remote
