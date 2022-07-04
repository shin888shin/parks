## Parks

A basic Go API

## Description

A simple Go based API using a grab bag of technologies.

## Run it

Start containers
```
docker-compose up --build
docker exec -it <db-container-id> bash
```

Add fun data
```
mysql -u root -p
  - rootrootroot
show databases;
use parks;
show tables;
INSERT INTO parks (name, description, location, created_at) VALUES ('grand canyon', 'hot as hell', 'arizona', CURRENT_TIMESTAMP);
```

See it in action
```
curl http://localhost:8082/ping
=> {"message":"pongtuesday"}

curl http://localhost:8082/parks
=> [{"id":1,"name":"grand canyon","description":"hot as hell","location":"arizona","created_at":"2021-07-04T02:48:16Z"}]
```
