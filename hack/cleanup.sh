#!/usr/bin/env bash

set -euxo pipefail
k3d cluster delete "${K3D_CLUSTER_NAME}"
k3d registry delete "${K3D_CLUSTER_NAME}-registry.localhost"
docker network rm "${DOCKER_NETWORK_NAME}" || true
rm -rf "$KUBECONFIG"
