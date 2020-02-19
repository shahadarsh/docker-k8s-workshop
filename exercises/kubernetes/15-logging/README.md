# Logging

```bash
helm repo add stable https://kubernetes-charts.storage.googleapis.com
```

```bash
kubectl create namespace -name logging
```

```bash
helm repo add akomljen-charts \
    https://raw.githubusercontent.com/komljen/helm-charts/master/charts/
```

```bash
helm install --name es-operator \
    --namespace logging \
    akomljen-charts/elasticsearch-operator
```

```bash
helm install --name efk \
    --namespace logging \
    akomljen-charts/efk
```

```bash
kubectl get pods -n logging
```

```bash
kubectl port-forward [efk-container-name] 5601 -n logging
```
