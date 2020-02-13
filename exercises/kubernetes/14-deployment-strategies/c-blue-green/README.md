# Blue Green Deployment

Deploy Blue: 

```bash
kubectl apply -f conf-details-svc-blue.yaml
```

Deploy Green: 
```bash
kubectl apply -f conf-details-svc-green.yaml
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
