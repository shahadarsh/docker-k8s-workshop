# Monitoring

```bash
kubectl create namespace -name monitoring
```

```bash
kubectl apply -f service-monitor.yaml
```

```bash
helm install \
    --name prom \
    --namespace monitoring \
    -f custom-values.yaml \
    stable/prometheus-operator
```
