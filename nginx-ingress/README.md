# How to use

## Prepare
Run nginx-ingress-controller.
```
$ minikube addons enable ingress
```

## Apply
```
$ make apply
```

## Set virtual host
Get minikube ip
```
$ minikube ip
xxx.xxx.xxx.xxx
```
Add following line to /ets/hosts
```
xxx.xxx.xxx.xxx hello-world.info
```

## Confirm
Now you can access via virtual host.
```
$ curl hello-world.info
<html><head><title>HTTP Hello World</title></head><body><h1>Hello from hello-world-deployment-5c8cd966cd-nvvdk</h1></body></html
```