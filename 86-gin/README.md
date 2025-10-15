# Docker commands 

```
docker network create network-demo

docker run -d --name pg -p 5432:5432 --network network-demo -e POSTGRES_USER=admin POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=contactsdb postgres

docker run -d --name pgui -p 8083:8080 --network network-demo adminer
```