## temporal-tryout

### About
Trying out Temporal and it's Go SDK locally. 

### Local Development

Install the Temporal CLU using `brew` if you don't have it installed already. 

To start the local Temporal Service:
```bash
temporal server start-dev --ui-port 8080 --db-filename local_instance.db
```

Run the client to start a Workflow Execution:
```go
go run client/main.go
```

Run the worker to start a Worker:
```go
go run worker/main.go
``` 
