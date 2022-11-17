include .env
export

setup:
	@drone exec --trusted --env-file=.env --include="setup cluster" --include="install istio"

build-server:
	@drone exec --trusted --env-file=.env --include="server compile" --include="build server image" --network=$(K3D_CLUSTER_NAME)

protoc-server:
	@drone exec --trusted --include="protoc-server"

deploy:
	kustomize build config | envsubst | kubectl apply -f -

undeploy:
	kustomize build config | envsubst | kubectl delete -f -

all:
	@drone exec --trusted --env-file=.env 

.PHONY: setup	deploy-server protoc-server all