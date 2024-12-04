# tech-challenge

Tech Challenge...

Let's go tech challenge!!!

Links:

<https://miro.com/app/board/uXjVKQtHwOA=/>

## Evidence of the tests carried out

<img width="1419" alt="Screenshot 2024-12-03 at 21 47 03" src="https://github.com/user-attachments/assets/8d7215da-83f9-4e85-bdfc-5913ba308830">


## Run project

To run the application it is necessary to execute the command `make start`

### Aplication

### Migration

All migrations are executed as soon as the `make start` or `make build` command is executed

#### Create

To create a migration, you need to run the `make migrate/create` command passing the file name

example:

```bash
make migrate/create name=add_user
```

to create a migration to add a user

### Swagger

URL to access running Swagger is `/api/v1/swagger/index.html`

## Kubernetes

> [!IMPORTANT]  
> [Minikube](https://minikube.sigs.k8s.io/docs?target=_blank) must be installed.

```bash
minikube start
eval $(minikube docker-env)
minikube addons enable volumesnapshots
minikube addons enable csi-hostpath-driver
docker buildx build -t tech-challenge-fase-4-order .
docker buildx build -t tech-challenge-fase-4-order-migration ./migrations/
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/secrets.yaml
kubectl apply -f k8s/database.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/nodeport.yaml
kubectl apply -f k8s/hpa.yaml
kubectl apply -f k8s/loadbalancer.yaml

#wait for postgres pod to finish
kubectl apply -f k8s/migration-job.yaml
minikube service tech-challenge-fase-4-order-nodeport --url
```
