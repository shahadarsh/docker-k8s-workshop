### Create cluster if not done already

#### Note: This step will take time. Close to ~20 mins

```
eksctl create cluster \
--name k8s-workshop \
--version 1.13 \
--nodegroup-name standard-workers \
--node-type t3.medium \
--nodes 3 \
--nodes-min 2 \
--nodes-max 3 \
--node-ami auto
```

### Confirm Cluster Connectivity ###
```
kubectl cluster-info
```
