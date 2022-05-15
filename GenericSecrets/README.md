1. kubectl create secret generic generic-secrets --from-env-file=generic-secrets.txt -o yaml

2. echo -n $(kubectl get secret generic-secrets -o jsonpath="{.data.a}") | base64 -D


1. base64-encode the values
   1. echo -n top-secret | base64
   2. echo -n bottom-secret | base64
2. kubectl create -f generic-secrets2.yaml


# Create a mutual secret for link-manager and graph-manager:
1. echo -n "social-graph-manager: 123" | base64
2. kubectl create -f link-mutual-auth.yaml
3. kubectl get secret link-mutual-auth -o "jsonpath={.data['mutual-auth\.yaml']}" | base64 -D

1. echo -n "link-manager: 123" | base64
2. kubectl create -f social-mutual-auth.yaml
3. kubectl get secret social-mutual-auth -o "jsonpath={.data['mutual-auth\.yaml']}" | base64 -D