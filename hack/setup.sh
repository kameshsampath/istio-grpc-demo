#!/usr/bin/env bash

set -euxo pipefail

docker network create --opt "com.docker.network.driver.mtu=1450"\
   "${K3D_CLUSTER_NAME}" || true

mkdir -p "$(dirname $KUBECONFIG)"

k3d registry create "${K3D_CLUSTER_NAME}-registry.localhost" --port 5001 --default-network="${K3D_CLUSTER_NAME}"

k3d cluster create "${K3D_CLUSTER_NAME}" --network "${K3D_CLUSTER_NAME}" --registry-use "k3d-${K3D_CLUSTER_NAME}-registry.localhost:5001"

k3d kubeconfig get "${K3D_CLUSTER_NAME}" > "${KUBECONFIG}"
sed -i 's|host.docker.internal|127.0.0.1|' "${KUBECONFIG}"