include .envrc
export

setup:
	@drone exec --trusted --env-file=.env --include="setup cluster" --include="install istio"

deploy-server:
	@drone exec --trusted --env-file=.env --include="deploy server" --network=$(K3D_CLUSTER_NAME)

protoc-server:
	@drone exec --trusted --include="protoc-server"

clean:
	kubectl delete -f config/v1
	kubectl delete -f config/v2

deploy:
	ko create --insecure-registry  -f config/v1
	ko create --insecure-registry  -f config/v2

all:
	@drone exec --trusted --env-file=.env 

.PHONY: setup	deploy-server protoc-server all