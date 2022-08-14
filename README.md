# portdomainservice
Introduction
================
The Port Domain Service can create and maintain data of maritime Ports.
The list of ports must be available in the Home directory, for a sample see the ports.json file

Usage
================
##To RUN the service
```
go run cmd/main.go
```
##And then...
Copy ports.json file to the Home Directory and then...

##To Update Ports
```
curl -X POST  http://localhost:8090/v1/updateports --http0.9
```
##To Create Ports
```
curl -X POST  http://localhost:8090/v1/createports --http0.9
```
##To Build the Service
```
go build ./...
```
##To Test the Service
```
go test ./...
```

Notes
================


Help
================


License
================

Licensed under the New BSD License.

Author
================