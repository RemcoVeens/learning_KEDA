__default__:
	install

format:
	go fmt ./...
	staticcheck ./...

test:
	go test --cover ./...
	gosec ./...

install:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

update_consumer:
	./scripts/buildprod.sh
	eval $(minikube docker-env)
	docker build -t keda-worker:v1 .
	kubectl apply -f k8s/scaledobject.yaml
	kubectl apply -f k8s/deployment.yaml
