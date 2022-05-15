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

## Tech Stack

* Go 1.18.2
* Minikube 1.25.2
* Kubernetes 1.23.4
* Docker 20.10.16
* PostgreSQL 14.2

## Network

Windows -> VirtualBox -> Ubuntu -> Minikube -> Docker -> Containerd -> Go & PostgreSQL

## How To Run

```bash
# Start cluster
minikube start --kubernetes-version=v1.23.3 --cni=cilium --nodes 3

# Enable Ingress controller
minikube addons enable ingress

# In Virtualbox VM forward port from Minikube to Windows VS Code
socat tcp-listen:8443,reuseaddr,fork tcp:192.168.49.2:8443

kubectl exec db-statefulset-0 -c postgres -it -- /bin/bash

curl -x 192.168.49.2:80 dgk.io/db

k8s-web-svc-snippet$ docker build -f docker/svc.dockerfile -t svc-image:latest .

kubectl port-forward pod/db-statefulset-0 5432:5432
```
