.PHONY: destroy apply

apply:
	kubectl apply -f deployment.yml
	kubectl apply -f service.yml
	kubectl apply -f ingress.yml

destroy:
	kubectl delete deploy hello-world-deployment
	kubectl delete svc hello-world-svc
	kubectl delete ingress nginx-ingress

