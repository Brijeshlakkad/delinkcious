1. kubectl create -f service-account.yaml

2. kubectl create -f pod-with-secrets.yaml

3. kubectl exec -it trouble bash

# ls /var/run/secrets/kubernetes.io/serviceaccount/
ls: cannot access '/var/run/secrets/kubernetes.io/serviceaccount/': No such file or directory

# cat /etc/generic-secrets2/c
top-secret

# cat /etc/generic-secrets2/d
bottom-secret