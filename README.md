KubectlDB: REST APIs for kubectl with datastore
===================================

KubectlDB is a web server which runs kubectl commands and stores the API calls as events in a .postgreSQL database. It can Create, Delete, List and Get deployments and track progress of a Rollout of a deployment.

### Prerequisites to run on a computer

- make (https://www.gnu.org/software/make/)
- docker desktop (https://www.docker.com/products/docker-desktop/) and access to download images
- Connection to a Kubernetes cluster with read/write privileges to create and read objects (https://minikube.sigs.k8s.io/docs/start/)
- PostgreSQL database (optional, https://www.postgresql.org/download/)
- git and access to download from github
- download and install Golang 1.16 or higher (https://go.dev/doc/install)
- A REST API Client (e.g Postman https://www.postman.com/downloads/)


## Quick Start
Create and navigate to a directory in the filesystem with write access and

    $ git clone https://github.com/doyyan/kubernetes.git
    $ cd kubernetes
    $ make postgres
    $ make createdb
    $ make migrateup
    $ go build ./...
    $ go run cmd/app/main.go


## REST API calls to process deployments

### Create a deployment
   ```
curl -X POST 'localhost:8080/deployment' -H 'Content-Type: text/plain' --data-raw '{
   "Name":"testDeployment",
   "Namespace":"default",
   "Image":"nginx:1.12",
   "ContainerPort":80,
   "ContainerName":"web",
   "Labels":{
      "app":"demo"
   },
   "Replicas":10
}'

Returns

```

### Find a deployment
  ```
curl -X GET 'localhost:8080/deployment/status?name=testDeployment'
  ```

### List all deployments


### Track status of a deployment

### Delete a deployment
  ```
curl -X DELETE 'localhost:8080/deployment?name=testDeployment'
  ```

