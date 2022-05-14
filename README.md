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
