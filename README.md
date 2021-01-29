## AIPetto OAuth
Service responsible for handling OAuthentication.

User API
```
{
    "email": "email@email.com",
    "password": "123abc"
}

OAuth API:
{
    "grant_type": "password",
    "username": "email@email.com",
    "password": "123abc"
}

{
    "grant_type": "client_credentials",
    "client_id": "id-123",
    "client_secret": "secret-123"
}

```

### Run using Docker and docker-compose
```
sudo docker-compose up --build
sudo docker-compose up -d (daemon mode)
sudo docker-compose up --remove-orphans
```

#### Running tests
```
cd src/domain/access_token
go test 
go test --cover
```

### Cassandra DB

The docker exec command allows you to run commands inside a Docker container. 
The following command line will give you a bash shell inside your cassandra container:

```
docker exec -it go-oauth-cassandra-db bash
cqlsh
```
The Cassandra Server log is available through Docker's container log:
```
docker logs go-oauth-cassandra-db
```

Inside Cassandra:
```
describe keyspaces;
CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1}
describe keyspaces;
USE oauth;
CREATE TABLE access_token(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);
SELECT * FROM access_token where access_token='example';
```


### Troubleshoot & useful commands
```
go mod tidy ---> https://blog.golang.org/using-go-modules
The go mod tidy command cleans up these unused dependencies:

go mod init github.com/aipetto/go-aipetto-users-api
go clean -modcache
```