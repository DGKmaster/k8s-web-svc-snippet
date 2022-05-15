# k8s-web-svc-snippet

## Application

Develop the small web-service, which:

* Waiting for the data from the customer, receive it, keep it in the DB and respond `Completed`.
* Available for the response for all cities in the database.

## API

Request data: `/add`

```json
{
    "city": "London"
}
```

Response:

```bash
200 OK
Completed
```

Request data: `/all`

Response:

```json
{
    "London",
    "Berlin",
    "Moscow"
}
```

## Infrastructure

Based on Kubernetes provide manifests to deploy the application to any k8s-cluster.
Feel free to use any database or ingress controller.

Try to think about:

* Scaling
* High-availability
* Smooth deployment
* Monitoring
* Documentation

## Project structure

```tree
.
├── k8s      # K8s Manifests
│   ├── db   # Manifests for PostgreSQL database
│   └── svc  # Manifests for Golang web service
└── svc      # Golang web service source code
```

## Tech Stack

* Go 1.18.2
* Minikube 1.25.2
* Kubernetes 1.23.4
* Docker 20.10.16
* PostgreSQL 14.2

## Network

Windows -> VirtualBox -> Ubuntu -> Minikube -> Docker -> Go & PostgreSQL

## How To Run

### Cluster setup

```bash
# Start cluster
minikube start --kubernetes-version=v1.23.3 --cni=cilium --nodes 3

# Enable Ingress controller
minikube addons enable ingress
```

### Local development

```bash
# In Virtualbox VM forward port from Minikube to Windows VS Code
socat tcp-listen:8443,reuseaddr,fork tcp:192.168.49.2:8443

# Debug DB
kubectl exec db-statefulset-0 -c postgres -it -- /bin/bash

# Local dev
kubectl port-forward pod/db-statefulset-0 5432:5432
```

### Build and release application

```bash
k8s-web-svc-snippet/svc$ docker build -f Dockerfile -t svc-image:latest .
docker tag svc-image dgkmaster/ubuntu-pcl
docker push dgkmaster/ubuntu-pcl
```

### Test API

```bash
# From local dev
curl localhost:8081/all
curl localhost:8081/add -H "Content-Type: application/json" -d '{"city":"Chicago"}'

# From ingress
curl -x 192.168.49.2:80 dgk.io/svc/all
curl -x 192.168.49.2:80 dgk.io/svc/add -H "Content-Type: application/json" -d '{"city":"Chicago"}'
```
