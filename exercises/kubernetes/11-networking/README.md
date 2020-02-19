# Networking

#### API Gateway (Ambassador)

```bash
kubectl apply -f setup/aes-crds.yaml && \
kubectl wait --for condition=established --timeout=90s crd -lproduct=aes && \
kubectl apply -f setup/aes.yaml && \
kubectl -n ambassador wait --for condition=available --timeout=90s deploy -lproduct=aes
```
Determine DNS/IP address:

```bash
kubectl get -n ambassador service ambassador
```

Replace ambassador DNS/IP address and run curl 

```bash
curl http://[AMBASSADOR_IP]/website/
```
