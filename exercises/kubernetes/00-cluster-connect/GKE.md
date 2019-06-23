# Create & connect to the cluster

```
gcloud auth login
(follow verification steps)
```

### Create cluster if not done already

```
gcloud container clusters create k8s-workshop  --enable-autoupgrade --enable-autoscaling --min-nodes=2 --max-nodes=3 --num-nodes=2 --zone=us-east1-b
```

### Configure SDK ###
```
gcloud config set project ${PROJECT_ID};
gcloud config set compute/region us-east1;
gcloud config set compute/zone us-east1-b;
```

### Retrieve Cluster Credentials (make sure there's no conflict with a pre-existing KUBECONFIG var) ###
```
gcloud container clusters get-credentials <cluster-name> --region <region>
```

### Confirm Cluster Connectivity ###
```
kubectl cluster-info
```
