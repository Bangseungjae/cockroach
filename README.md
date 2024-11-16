
Install-Package

Install all packages below.

Echo (To serve the HTTP Server)

Viper (To load the config)

GORM + PostgreSQL Driver (To communicate with the database)


```
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/spf13/viper
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

go get github.com/stretchr/testify 
```

Then, drill into the container to create the database line by line.
```text
docker exec -it cockroachdb bash

psql -U postgres

CREATE DATABASE cockroachdb;
```

```sql
cockroachdb=# \c cockroach

cockroach=# \dt
            List of relations
 Schema |    Name     | Type  |  Owner   
--------+-------------+-------+----------
 public | cockroaches | table | postgres
(1 row)

cockroach=# \dt cockroaches
            List of relations
 Schema |    Name     | Type  |  Owner   
--------+-------------+-------+----------
 public | cockroaches | table | postgres
(1 row)

cockroach=# SELECT * FROM cockroaches;
 id | amount |          created_at           
----+--------+-------------------------------
  1 |      1 | 2024-11-16 00:09:02.135528+09
  2 |      2 | 2024-11-16 00:09:02.135528+09
  3 |      2 | 2024-11-16 00:09:02.135528+09
  4 |      5 | 2024-11-16 00:09:02.135528+09
  5 |      3 | 2024-11-16 00:09:02.135528+09
(5 rows)
```

```sql
cockroach=# \d cockroaches
                                       Table "public.cockroaches"
   Column   |           Type           | Collation | Nullable |                 Default                 
------------+--------------------------+-----------+----------+-----------------------------------------
 id         | bigint                   |           | not null | nextval('cockroaches_id_seq'::regclass)
 amount     | bigint                   |           |          | 
 created_at | timestamp with time zone |           |          | 
Indexes:
    "cockroaches_pkey" PRIMARY KEY, btree (id)
```