.PHONY: destroy apply

apply:
	kubectl apply -f sa.yaml
	kubectl apply -f clusterrolebinding.yaml
	kubectl apply -f pod.yaml

destroy:
	kubectl delete po sample-kubectl
	kubectl delete sa sample-serviceaccount
	kubectl delete clusterrolebinding sample-clusterrolebinding

test:
	kubectl exec -it sample-kubectl -- kubectl get pods
