# Workshop Prerequisites

Introduction to Kubernetes & Advanced Kubernetes workshop consists of hands-on exercises. In order to be prepared for that please follow below instructions prior to arriving. You can email me on adarsh@shahadarsh.com if you have any questions.

Install following in advance but it will also be good to have admin rights to your laptop in case we need to install something else. 

## kubectl 
Follow instructions [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Docker

* [Mac](https://docs.docker.com/docker-for-mac/install/)
* [Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/)  
* [Centos](https://docs.docker.com/install/linux/docker-ce/cens/)
* [Windows](https://docs.docker.com/docker-for-windows/install/)

## Git 
Follow instructions [here](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Kubernetes Cluster

There are various options for running k8s cluster that can be used with this workshop:

#### Note: Charges will be incurred if you use option 1 or 2.

### Option 1: Using GKE (Recommended)

* Setup GCP account (if you don't have one): Instructions [here](https://console.cloud.google.com/getting-started)
    * You might get a free $300 credit depending on the region
* [GKE](https://cloud.google.com/run/docs/gke/setup)
* [Google Cloud SDK](https://cloud.google.com/sdk/install)

### Option 2: Using EKS

* Setup AWS account (if you don't have one): Instructions [here](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/)
* Setup [eksctl](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html)

### Option 3: Using Minikube

Follow instructions [here](https://kubernetes.io/docs/tasks/tools/install-minikube/)

**Note:** This will only create single node cluster so you won't be able to try out few of the features like scaling during hands-on exercises.

## Clone repo and verify setup
Clone repo
```bash
git@github.com:shahadarsh/docker-k8s-workshop.git
```
Create cluster and connect using instructions [here](https://github.com/shahadarsh/docker-k8s-workshop/tree/master/exercises/kubernetes/00-cluster-connect)

Cleanup after to avoid charges(we will recreate the cluster the day off workshop). Follow steps [here](https://github.com/shahadarsh/docker-k8s-workshop/tree/master/exercises/kubernetes/999-cleanup)

## Other aspects 
* If you have a MiFi device bring it with you in-case venue WiFi has issues
* Its recommended to have Admin rights to your laptop in case we need to install something else to try out
* Few additional things that are helpful but not needed: 
    * [k9s](https://github.com/derailed/k9s)
    * [kubectx](https://github.com/ahmetb/kubectx)

