# User Service

The user microservice. It uses a multi-stage [Dockerfile](Dockerfile) to generate a lean and mean image from SCRATCH that just includes the Go binary. The system has a CI/CD pipeline, but you can also build and deploy it yourself.


## Build Docker image

```
$ docker build -t brijeshlakkad/delinkcious-user:${VERSION} .
```

## Push to Registry

This will go by default to DockerHub. Make sure you're logged in:

```
$ docker login
```

Then push your image:

```
$ docker push brijeshlakkad/delinkcious-user:${VERSION}
```

## Deploy to active Kubernetes cluster

If you want to push a local minikube make sure your kubectl is pointed to the right cluster and type:

```
$ kubectl apply -f k8s
```







