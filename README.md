# learning KEDA

this is my first experience with kubernetes event-driven autoscaling. using redis in this case

`producer` creates job_queue object and waits for completion-queue object

`consumer` waits for job_queue objects, 'proccesses' it and creates completion-queue object

using minukube to create a local kubernetes cluster

## install

1. run minikube cluster

```Bash
./dashboard.sh start
```

2. Install KEDA via Helm

Add the KEDA Helm repository and install it into your cluster:

```Bash
helm repo add kedacore https://kedacore.github.io/charts
helm repo update
helm install keda kedacore/keda --namespace keda --create-namespace
```

Verify installation: Run `kubectl get pods -n keda` and ensure the following are running:

- admission webhook
- KEDA operator
- metrics server

3. Install Redis via Helm

We will use Bitnami's Redis chart for a quick setup:

```Bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install redis bitnami/redis --set auth.enabled=false
```

4. create image and upload minikube

run `make update_consumer`
