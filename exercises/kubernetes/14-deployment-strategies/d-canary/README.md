# Blue Green Deployment

Deploy Blue: 

```bash
kubectl apply -f conf-details-svc-v1.yaml
```

Deploy Green: 
```bash
kubectl apply -f conf-details-svc-v2.yaml
```

Install Istio:
```bash
helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.1.7/charts/
helm repo list
kubectl create namespace -n istio-system
helm install istio-init --namespace istio-system istio.io/istio-init
helm install istio --namespace istio-system --set grafana.enabled=true istio.io/istio
kubectl get svc -n istio-system
kubectl label namespace default istio-injection=enabled
kubectl get namespace -L istio-injection
```

Test Green:
```bash
kubectl port-forward [pod-name] 8080:8080
curl http://localhost:8080/conference-details
```

Switch from Blue to Green:
```bash
kubectl apply -f conf-details-svc-blue-to-green.yaml
```
