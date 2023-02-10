SHELL := /bin/sh

.PHONY: run-adder-api
run-adder-api:
	go run cmd/adder-api/main.go

.PHONY: run-subtractor-api
run-subtractor-api:
	go run cmd/subtractor-api/main.go

.PHONY: run-transformer-api
run-transformer-api:
	go run cmd/transformer-api/main.go

.PHONY: run-transformer-client
run-transformer-client:
	go run cmd/transformer-client/main.go

.PHONY: dep
dep:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/adder/adder.proto proto/subtractor/subtractor.proto proto/transformer/transformer.proto

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up

.PHONY: build-images
build-images:
	docker build --target adder-api --tag adder-api:1.0.0 .
	docker build --target subtractor-api --tag subtractor-api:1.0.0 .
	docker build --target transformer-api --tag transformer-api:1.0.0 .
	docker build --target transformer-client --tag transformer-client:1.0.0 .

.PHONY: minikube-load-images
minikube-load-images:
	minikube image rm adder-api:1.0.0 subtractor-api:1.0.0 transformer-api:1.0.0 transformer-client:1.0.0
	minikube image load adder-api:1.0.0
	minikube image load subtractor-api:1.0.0
	minikube image load transformer-api:1.0.0
	minikube image load transformer-client:1.0.0

.PHONY: kubectl-deploy
kubectl-deploy:
	kubectl apply -f deployment-adder-api.yaml
	kubectl apply -f deployment-subtractor-api.yaml
	kubectl apply -f deployment-transformer-api.yaml

.PHONY: kubectl-deploy-client
kubectl-deploy-client:
	kubectl create -f deployment-transformer-client.yaml

.PHONY: kubectl-deploy-with-linkerd
kubectl-deploy-with-linkerd:
	cat deployment-adder-api.yaml | linkerd inject - | kubectl apply -f -
	cat deployment-subtractor-api.yaml | linkerd inject - | kubectl apply -f -
	cat deployment-transformer-api.yaml | linkerd inject - | kubectl apply -f -

.PHONY: kubectl-deploy-client-with-linkerd
kubectl-deploy-client-with-linkerd:
	cat deployment-transformer-client.yaml | linkerd inject - | kubectl create -f -

.PHONY: kubectl-deploy-with-istio
kubectl-deploy-with-istio:
	istioctl kube-inject -f deployment-adder-api.yaml | kubectl apply -f -
	istioctl kube-inject -f deployment-subtractor-api.yaml | kubectl apply -f -
	istioctl kube-inject -f deployment-transformer-api.yaml | kubectl apply -f -

.PHONY: kubectl-deploy-client-with-istio
kubectl-deploy-client-with-istio:
	istioctl kube-inject -f deployment-transformer-client.yaml | kubectl create -f -

.PHONY: kubectl-cleanup
kubectl-cleanup:
	kubectl get jobs | awk '/transformer-client/{print $$1}' | xargs kubectl delete jobs
	kubectl delete deployments adder-api subtractor-api transformer-api || true
	kubectl delete services adder-api subtractor-api transformer-api || true
