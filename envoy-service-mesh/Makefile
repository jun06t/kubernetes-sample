.PHONY: configmap destroy apply apply-deployments apply-services

configmap:
	kubectl create configmap envoy-config --from-file=envoy-config

apply: apply-deployments apply-services

apply-deployments:
	kubectl apply -f deployment-alive.yml
	kubectl apply -f deployment-user.yml
	kubectl apply -f deployment-gateway.yml

apply-services:
	kubectl apply -f service-alive.yml
	kubectl apply -f service-alive-admin.yml
	kubectl apply -f service-user.yml
	kubectl apply -f service-user-admin.yml
	kubectl apply -f service-gateway.yml

destroy:
	kubectl delete deployment gateway-deployment
	kubectl delete deployment alive-service-deployment
	kubectl delete deployment user-service-deployment
	kubectl delete service gateway-svc
	kubectl delete service user-svc
	kubectl delete service user-admin-svc
	kubectl delete service alive-admin-svc
	kubectl delete service alive-svc
	kubectl delete configmap envoy-config

