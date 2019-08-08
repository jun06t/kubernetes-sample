# How to use

## Before start
### Set virtual host
Get minikube ip
```
$ minikube ip
xxx.xxx.xxx.xxx
```
Add following line to /ets/hosts
```
xxx.xxx.xxx.xxx hello-world.info
```

## Case 1: with minikube addons
### Run nginx-ingress-controller.
```
$ minikube addons enable ingress
```

### Apply
```
$ make apply
```

### Confirm
Now you can access via virtual host.
```
$ curl hello-world.info
<html><head><title>HTTP Hello World</title></head><body><h1>Hello from hello-world-deployment-5c8cd966cd-nvvdk</h1></body></html
```

## Case 2: with helm charts
### Run nginx-ingress-controller.
```
$ helm install --name nginx-ingress stable/nginx-ingress 
```

Check nginx-ingress-controller's Service port.

```
$ kubectl get svc
NAME                            TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
kubernetes                      ClusterIP      10.96.0.1        <none>        443/TCP                      17h
nginx-ingress-controller        LoadBalancer   10.106.110.157   <pending>     80:32123/TCP,443:31527/TCP   5m51s
nginx-ingress-default-backend   ClusterIP      10.106.9.142     <none>        80/TCP                       5m51s
```

### Apply
```
$ make apply
```

### Confirm
Now you can access via virtual host.
```
$ curl hello-world.info:32123
<html><head><title>HTTP Hello World</title></head><body><h1>Hello from hello-world-deployment-5c8cd966cd-nvvdk</h1></body></html
```
