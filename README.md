# docker-k8s-workshop
Exercises for docker &amp; k8s workshop

```bash
gcloud container clusters create k8s-workshop  --enable-autoupgrade --enable-autoscaling --min-nodes=3 --max-nodes=10 --num-nodes=5 --zone=us-east1-b
```

```bash
docker build -t conference-website:latest .
```
```bash
docker tag conference-website:latest us.gcr.io/adarsh-shah-playground/conference-website:latest
```
```bash
docker push us.gcr.io/adarsh-shah-playground/conference-website:latest 
```

```bash
docker build -t schedule-service:latest .
```
```bash
docker tag schedule-service:latest us.gcr.io/adarsh-shah-playground/schedule-service:latest
```
```bash
docker push us.gcr.io/adarsh-shah-playground/schedule-service:latest
```

```bash
docker build -t conference-details-service:latest .
```
```bash
docker tag conference-details-service:latest us.gcr.io/adarsh-shah-playground/conference-details-service:latest
```
```bash
docker push us.gcr.io/adarsh-shah-playground/conference-details-service:latest
```

```bash
kubectl apply -f . 
```
