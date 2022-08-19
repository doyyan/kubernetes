KubectlDB: REST APIs for kubectl with datastore
===================================

KubectlDB is a web server which runs kubectl commands and stores the API calls as events in a .postgreSQL database. It can Create, Delete, List and Get deployments and track progress of a Rollout of a deployment.

### Prerequisites to run on a computer

- make (https://www.gnu.org/software/make/)
- docker desktop (https://www.docker.com/products/docker-desktop/) and access to download images
- Connection to a Kubernetes cluster with read/write privileges to create and read objects (https://minikube.sigs.k8s.io/docs/start/)
- PostgreSQL database (optional, https://www.postgresql.org/download/)
- git and access to download from github (https://git-scm.com/downloads)
- download and install Golang 1.18 or higher (https://go.dev/doc/install)
- A REST API Client (e.g Postman https://www.postman.com/downloads/)


## Quick Start
Create and navigate to a directory in the filesystem with write access and

First ensure Docker is running locally and check if connected to the correct k8s cluster by running   

```kubectl config current-context```

    $ git clone https://github.com/doyyan/kubernetes.git
    $ git checkout techtest
    $ cd kubernetes
    $ make rm-postgres
    $ make postgres
    $ sleep 10
    $ make createdb
    $ make migrateup
    $ go build ./...
    $ go run cmd/app/main.go


## REST API calls to process deployments

### Create a deployment (from k8s and DB)
   ```
curl -X POST 'localhost:8080/deployment' -H 'Content-Type: text/plain' --data-raw '{
   "Name":"testdeployment",
   "Namespace":"default",
   "Image":"nginx:1.12",
   "ContainerPort":80,
   "ContainerName":"web",
   "Labels":{
      "app":"demo"
   },
   "Replicas":10
}'
```
```
Returns

{
    "message": "deployment created"
}
```

### Find a deployment (from DB)
  ```
curl -X GET 'localhost:8080/deployment/status?name=testdeployment'
```
```
Returns

{
    "NameSpace": "default",
    "Name": "testdeployment",
    "Kind": "deployment",
    "Image": "nginx:1.12",
    "ContainerName": "web",
    "ContainerPort": 80,
    "Labels": {
        "app": "demo"
    },
    "Replicas": 10,
    "Ready": 0,
    "Current": 0,
    "Available": 0,
    "Status": ""
}
  ```

### List all deployments (from DB)
  ```
curl -X GET 'localhost:8080/deployment/all'
```
```
Returns
[
{
  "NameSpace": "default",
  "Name": "testdeployment",
  "Kind": "deployment",
  "Image": "nginx:1.12",
  "ContainerName": "web",
  "ContainerPort": 80,
  "Labels": {
      "app": "demo"
  },
  "Replicas": 10,
  "Ready": 0,
  "Current": 0,
  "Available": 0,
  "Status": ""
}
]
  ```
### Track status of a deployment (from k8s)
  ```
curl -X GET 'localhost:8080/deployment/status?name=testdeployment'
  ```
```
Returns

Prior starting deployment
{
    "message": "deployment not found"
}

Deployment In Progress
{
    "message": "Waiting for deployment \"testdeployment2\" rollout to finish: 10 of 20 updated replicas are available...\n",
    "rollout Status": "Process in Progress"
}

Finally
{
    "message": "deployment \"testdeployment2\" successfully rolled out\n",
    "rollout Status": "Process Complete"
}
  ```
### Delete a deployment (from k8s and DB)
  ```
curl -X DELETE 'localhost:8080/deployment?name=testdeployment'
```
```
Returns

{
    "message": "deployment deleted"
}
  ```

