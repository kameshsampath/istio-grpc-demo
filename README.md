# Istio gRPC Demo

## Istio Install

```shell
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
```

```shell
kubectl create namespace istio-system
```

```shell
helm install istio-base istio/base -n istio-system
```

```shell
helm install istiod istio/istiod -n istio-system --wait
```

```shell
kubectl create namespace istio-ingress
kubectl label namespace istio-ingress istio-injection=enabled
helm install istio-ingress istio/gateway -n istio-ingress --wait
```
