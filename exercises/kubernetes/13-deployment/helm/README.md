# Helm 3.0 

## Setup Helm 

## Create Helm Chart & install to cluster

```bash
helm create "app"
```

Change docker image in app/values.yaml

```bash
helm install conf-details-svc app 
```

Change service/type from ClusterIP to LoadBalancer

```bash
curl [lb-url]
```
