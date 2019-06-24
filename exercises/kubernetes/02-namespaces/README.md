# Pods
Official docs: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/

Namespaces are a way to divide cluster resources between multiple users (via resource quota).

## Create Namespace 

```bash
kubectl apply -f conf-website-namespace.yaml 
```
Verify namespace was created

```bash
kubectl get namespaces 
```
```bash
NAME                 STATUS   AGE
conference-website   Active   5s
default              Active   6d6h
kube-public          Active   6d6h
kube-system          Active   6d6h
```
### Cleanup

```bash
kubectl delete -f conf-website-namespace.yaml 
```
