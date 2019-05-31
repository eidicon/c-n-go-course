## preinstallation
minikube(https://kubernetes.io/docs/tasks/tools/install-minikube/)


## build with go local installation
go build -o ./bin/cloud-native-go

## structure
/bin -> compiled binary file


## docker
### build with docker image
docker build -f builder.dockerfile -t cloud-native-go:1.0.2-alpine .

### create image (without builder)
docker build -f app.dockerfile -t cloud-native-go:1.0.1-alpine .

## Minikube commands
 - minikube dashboard (to startup dashboard)
 - minikube start
 - minikube stop
 ### get exposed url (by service, when it's created) 
 - minikube service cloud-native-go --url

## kubectl
### create pod
kubectl create -f k8s-pod.yml
#### create pod o defined namespace
kubectl create -f k8s-pod.yml --namespace cloud-native-go

### delete pod 
kubectl delete pod cloud-native-go

### get list of pods 
kubectl get pods
kubectl get pods --show-labels
kubectl get pods -o wide -L env

### get pod info
kubectl describe pod cloud-native-go

### manage labels
kubectl label pod cloud-native-go env=dev --overwrite

### namespaces 
kubectl get ns
kubectl get pods --namespace cloud-native-go
#### create ns from file 
kubectl create -f k8s-namespace.yml
#### delete namespace and all related pods
kubectl delete -f k8s-namespace.yml

## deployment
### general info
kubectl get deployments,pods,rs
### creates deployment based on config file
kubectl create -f k8s-deployment.yml --record=true
### apply changes in deployment config changes
kubectl apply -f k8s-deployment.yml
### get detailed info
kubectl describe deployment cloud-native-go

## scale
kubectl scale deployment cloud-native-go --replicas=5
kubectl rollout history deployment cloud-native-go

## service
_Service is a "gate" for a pods_
kubectl create -f k8s-service.yml
kubectl get services


## how to update image (this will perform as rollout update)
kubectl set image deployment cloud-native-go cloud-native-go=cloud-native-go:1.0.1-apline
## how to rollback
kubectl rollout undo deployment cloud-native-go


## how to access to the app
kubectl  port-forward cloud-native-go 8080:8080