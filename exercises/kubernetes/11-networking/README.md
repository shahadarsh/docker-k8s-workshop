# Networking

1. Highly-coupled container-to-container communications 
- Pods and localhost communications

2. Pod-to-Pod communications

3. Pod-to-Service communications: this is covered by services.

4. External-to-Service communications: this is covered by services.
- Load Balancer exercise in #05-Services

#### API Gateway (Ambassador)

```bash
kubectl apply -f setup/aes-crds.yaml && \
kubectl wait --for condition=established --timeout=90s crd -lproduct=aes && \
kubectl apply -f setup/aes.yaml && \
kubectl -n ambassador wait --for condition=available --timeout=90s deploy -lproduct=aes
```

Determine IP address:

```bash
kubectl get -n ambassador service ambassador -o 'go-template={{range .status.loadBalancer.ingress}}{{print .ip "\n"}}{{end}}'
```

- Ingress


Calico/Flannel
