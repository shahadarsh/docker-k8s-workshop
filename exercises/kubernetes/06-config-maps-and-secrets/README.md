#  Config Maps & Secrets  
Official docs: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/  

## Config Maps   

```bash
kubectl apply -f .
```
```bash
kubectl get services
```
Note: External IP assignment takes some time so wait for it
```bash
NAME               TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
conf-details-svc   LoadBalancer   10.47.241.57   34.74.25.49   80:31064/TCP   6d9h
kubernetes         ClusterIP      10.47.240.1    <none>        443/TCP        6d11h
```
```bash
curl http://34.74.25.49/conference-details
```
### Cleanup

```bash
kubectl delete -f .
```
