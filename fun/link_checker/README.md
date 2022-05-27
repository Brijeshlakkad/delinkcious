# Installing Fission on Minikube [https://fission.io/docs/installation/]
```
$ export FISSION_NAMESPACE="fission"
$ kubectl create namespace $FISSION_NAMESPACE
$ kubectl create -k "github.com/fission/fission/crds/v1?ref=v1.16.0"
$ helm repo add fission-charts https://fission.github.io/fission-charts/
$ helm repo update
$ helm install --version v1.16.0 --namespace $FISSION_NAMESPACE fission \
  --set serviceType=NodePort,routerServiceType=NodePort \
  fission-charts/fission-all
```

# Install Fission CLI
```
$ curl -Lo fission https://github.com/fission/fission/releases/download/v1.16.0/fission-v1.16.0-darwin-amd64 \ && chmod +x fission && sudo mv fission /usr/local/bin/
```

# Add the Go environment to your cluster
```
$ fission environment update --name go --image fission/go-env:1.16 --builder fission/go-builder:1.16
```

# Create golang env with builder image to build go plugin
```
$ fission fn create --name link-checker --env go --src link_checker.go --entrypoint Handler
$ fission fn update --name link-checker --env go --src link_checker.go --entrypoint Handler
```

# Before accessing function, need to ensure deploy package of function is in succeeded state.
```
$ fission pkg info --name <pkg-name>
```

# Now, let’s test our first Go function with test command
```
$ fission fn test --name <function-name>
```

# Clusters Only Support NodePort
```
export FISSION_ROUTER=$(minikube ip):$(kubectl -n fission get svc router -o jsonpath='{...nodePort}')
```

curl -X POST http://$FISSION_ROUTER/link-checker -d '{"username": "brijeshlakkad"}'
