# Controllers 
Official docs: https://kubernetes.io/docs/concepts/workloads/controllers/ 

Controllers make sure that the desired state of the cluster matches the actual state. 

## Deployment controller  

```bash
kubectl apply -f conf-details-svc-1.0.yaml
```

Verify the state of deployment

```bash
kubectl get deployments 
```
```bash
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
conf-details-svc   2/2     2            2           15m
```

## Upgrade container 
```bash
kubectl apply -f conf-details-svc-2.0.yaml
```

Verify upgrade 

```bash
kubectl get deployments 
```
```bash
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
conf-details-svc   2/2     2            2           15m
```
## Scale deployment
```bash
kubectl apply -f conf-details-svc-2.0-scale.yaml
```
Verify scale 

```bash
kubectl get deployments 
```
```bash
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
conf-details-svc   5/5     5            5           15m
```

### Cleanup

```bash
kubectl delete -f conf-details-svc-2.0-scale.yaml
```
