# docker-k8s-workshop
Exercises for docker &amp; k8s workshop

## Create Kubernetes cluster

```bash
gcloud container clusters create k8s-workshop  --enable-autoupgrade --enable-autoscaling --min-nodes=3 --max-nodes=10 --num-nodes=5 --zone=us-east1-b
```

## Build microservice docker images & push them to registry 

```bash
REGISTRY_PREFIX="us.gcr.io/devnexus-k8s-workshop"
```

```bash
docker build -t $REGISTRY_PREFIX/conference-website:latest . \
    && docker push $REGISTRY_PREFIX/conference-website:latest
```

```bash
docker build -t $REGISTRY_PREFIX/schedule-service:latest . \
    && docker push $REGISTRY_PREFIX/schedule-service:latest
```

```bash
docker build -t $REGISTRY_PREFIX/conference-details-service:latest . \
    && docker push $REGISTRY_PREFIX/conference-details-service:latest
```

## Deploy microservices to Kubernetes cluster

```bash
kubectl apply -f . 
```
