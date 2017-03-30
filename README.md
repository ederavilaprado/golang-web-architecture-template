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
