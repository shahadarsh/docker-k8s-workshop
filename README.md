# docker-k8s-workshop
Exercises for docker &amp; k8s workshop

## Create Kubernetes cluster

```bash
gcloud container clusters create k8s-workshop  --enable-autoupgrade --enable-autoscaling --min-nodes=3 --max-nodes=10 --num-nodes=5 --zone=us-east1-b
```

## Build microservice docker images & push them to registry 

```bash
docker build -t us.gcr.io/adarsh-shah-playground/conference-website:latest .
    && docker push us.gcr.io/adarsh-shah-playground/conference-website:latest
```

```bash
docker build -t us.gcr.io/adarsh-shah-playground/schedule-service:latest .
    && docker push us.gcr.io/adarsh-shah-playground/schedule-service:latest
```

```bash
docker build -t us.gcr.io/adarsh-shah-playground/conference-details-service:latest .
    && docker push us.gcr.io/adarsh-shah-playground/conference-details-service:latest
```

## Deploy microservices to Kubernetes cluster

```bash
kubectl apply -f . 
```
