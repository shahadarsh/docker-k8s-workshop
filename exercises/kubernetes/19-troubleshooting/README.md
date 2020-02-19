# Troubleshooting Kubernetes

## Pods

#### Verify Pods
`kubectl get pods`

#### Describe Pod
`kubectl describe [pod-name]`

#### Check Logs
`kubectl logs [pod-name]`

#### Check Previous Logs
`kubectl logs [pod-name] --previous`

#### Run commands inside the container
`kubectl exec -it [pod-name] bash`

#### Show Cluster-level events

```bash
# all events sorted by time
kubectl get events --sort-by=.metadata.creationTimestamp
# warnings only
kubectl get events --field-selector type=Warning
# events related to Nodes
kubectl get events --field-selector involvedObject.kind=Node
```

Show kubectx & k9s features
