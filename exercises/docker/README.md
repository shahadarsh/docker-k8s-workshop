# docker exercises

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
