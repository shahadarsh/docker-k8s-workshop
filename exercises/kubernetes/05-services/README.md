# Services 
Official docs: https://kubernetes.io/docs/concepts/services-networking/service/ 

Service is an abstract way to expose an application running on a set of Pods.

## Service   

```bash
kubectl apply -f .
```

Verify the state of deployment

```bash
kubectl get services 
```
```bash
NAME               TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
conf-details-svc   LoadBalancer   10.47.241.57   34.74.25.49   80:31064/TCP   55s
kubernetes         ClusterIP      10.47.240.1    <none>        443/TCP        82m
```

### Cleanup

```bash
kubectl delete -f .
```
