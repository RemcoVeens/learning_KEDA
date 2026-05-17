#!/bin/bash

case "$1" in
    start)
    	minikube start
    	minikube start --drivers=docker
    	minikube dashboard --port=63840 &
    	kubectl proxy &
    	minikube tunnel --bind-address="127.0.0.1" -c &
        disown
        ;;
    stop)
        echo "Stopping minikube"
    	minikube stop
        sudo pkill -f "minikube tunnel"
        sudo pkill -f "minikube port-forward"
        ;;
    update)
        ./scripts/buildprod.sh
        eval $(minikube docker-env)
        docker build -t keda-worker:v1 .
        kubectl apply -f k8s/scaledobject.yaml
        kubectl apply -f k8s/deployment.yaml
        ;;
    port)
        kubectl port-forward svc/redis-master 6379:6379 &
        drivers
        ;;
    instal)
    	helm repo add kedacore https://kedacore.github.io/charts
    	helm repo add bitnami https://charts.bitnami.com/bitnami
    	helm repo update
    	helm install keda kedacore/keda --namespace keda --create-namespace
    	helm install redis bitnami/redis --set auth.enabled=false
     ;;
    *)
        echo "Usage: $0 {start|stop|instal}"
        exit 1
        ;;
esac
