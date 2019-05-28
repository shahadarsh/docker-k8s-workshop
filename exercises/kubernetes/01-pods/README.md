# Pods
Official docs: https://kubernetes.io/docs/concepts/workloads/pods/pod/

A Pod (as in a pod of whales or pea pod) is a group of one or more containers, with shared storage/network, and a specification for how to run the containers.

## Pod with single container 

```bash
kubectl apply -f single-conference-svc.yaml
```
Verify the state of the pod

```bash
kubectl get pods 
```
```bash
NAME                              READY     STATUS    RESTARTS   AGE
single-conference-svc             1/1       Running   0          69s
```

Get more details of the pod
```bash
kubectl describe pod single-conference-svc 
```

Making pod accessible
```bash
kubectl port-forward single-conference-svc 8080:8080
```
Navigate
```bash
curl http://localhost:8080/conference-details
```

Access Logs
```bash
kubectl logs -f single-conference-svc 
```

## Pod with multiple containers

```bash
kubectl apply -f multi-conference-svc.yaml
```

Get details of the pod
```bash
kubectl describe pod multi-conference-svc
```

Making pod accessible
```bash
kubectl port-forward multi-conference-svc 8080:8080
```
Navigate
```bash
curl http://localhost:8080/conference-details
```

### Cleanup

```bash
kubectl delete -f single-conference-svc.yaml
```
```bash
kubectl delete -f multi-conference-svc.yaml
```
