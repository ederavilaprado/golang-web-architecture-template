# golang-web-architecture-template

[![Build Status](https://travis-ci.org/ederavilaprado/golang-web-architecture-template.svg?branch=master)](https://travis-ci.org/ederavilaprado/golang-web-architecture-template)

## TODO

- [ ] request scope + reques id
- [ ] log for request/response
- [ ] log to file
- [ ] struct, params and querystring validation
- [x] change httprouter to "echo"
- [ ] handling errors
- [ ] graceful shutdown with https://github.com/facebookgo/grace
- [ ] swagger
- [ ] migrations
- [ ] metrics (with influxdb or prometheus)
- [ ] Change the api json for query from snake case to camel case.
- [ ] Helper to return json ?!?
- [ ] improve readme


```bash
$ docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=go_restful -p 5432:5432 -d postgres
```

```bash
$ psql -h localhost -p 5432 -d apidb -U postgres -W
mysecretpassword
```

```sql
CREATE TABLE customers (
  id integer,
  full_name text,
  email text NULL
);
```

CompileDaemon for dev env
```
$ go get github.com/githubnemo/CompileDaemon
```

```
docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=go_restful -p 5432:5432 -d postgres

psql -h localhost -p 5432 -d go_restful -U postgres -W < db.sql



eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTEyNTUyNzQsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.90mklqO3anYDjps-h7bY7GQORHnSwGP2a3_P7EqxusM


curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTEyNTUyNzQsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.90mklqO3anYDjps-h7bY7GQORHnSwGP2a3_P7EqxusM" http://localhost:8080/v1/artists
curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTEyNTUyNzQsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.90mklqO3anYDjps-h7bY7GQORHnSwGP2a3_P7EqxusM" http://localhost:8080/v1/artists/10
```



Following the good patterns like S.O.L.I.D. and Clean Architecture, with many thanks to Uncle Bob (https://twitter.com/unclebobmartin) and also Dave Cheney (https://twitter.com/davecheney)
- https://8thlight.com/blog/uncle-bob/
- https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html
- https://8thlight.com/blog/uncle-bob/2014/05/08/SingleReponsibilityPrinciple.html
- https://www.youtube.com/watch?v=zzAdEt3xZ1M

Hardly inspired by https://github.com/qiangxue/golang-restful-starter-kit
