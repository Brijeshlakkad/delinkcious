IMAGE=brijeshlakkad/py-kube:0.3
docker build . -t $IMAGE
docker push $IMAGE
