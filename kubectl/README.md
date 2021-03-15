# Kubernetes

- You need a pod where you containter will run. PODs IP will change over time.
- You need a service with fixed IP to load balance your POD.
- You need to ingress object to expose your service to outside world.
- If your application is stateful then you also need a volume to mount your application to like PVC.