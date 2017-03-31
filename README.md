# golang-web-architecture-template

```bash
$ docker run -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=postgres -e POSTGRES_DB=apidb -p 5437:5437 -p 5432:5432 -d postgres
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

```
docker run -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=go_restful -p 5432:5432 -d postgres

psql -h localhost -p 5432 -d go_restful -U postgres -W < db.sql



eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTEyNTUyNzQsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.90mklqO3anYDjps-h7bY7GQORHnSwGP2a3_P7EqxusM


curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTEyNTUyNzQsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.90mklqO3anYDjps-h7bY7GQORHnSwGP2a3_P7EqxusM" http://localhost:8080/v1/artists
curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTEyNTUyNzQsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.90mklqO3anYDjps-h7bY7GQORHnSwGP2a3_P7EqxusM" http://localhost:8080/v1/artists/10
```
