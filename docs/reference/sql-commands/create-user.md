<a id="sql-createuser"></a>

# CREATE USER

define a new database role

## Synopsis


```

CREATE USER NAME [ [ WITH ] OPTION [ ... ] ]

where OPTION can be:

      SUPERUSER | NOSUPERUSER
    | CREATEDB | NOCREATEDB
    | CREATEROLE | NOCREATEROLE
    | INHERIT | NOINHERIT
    | LOGIN | NOLOGIN
    | REPLICATION | NOREPLICATION
    | BYPASSRLS | NOBYPASSRLS
    | CONNECTION LIMIT CONNLIMIT
    | [ ENCRYPTED ] PASSWORD 'PASSWORD' | PASSWORD NULL
    | VALID UNTIL 'TIMESTAMP'
    | IN ROLE ROLE_NAME [, ...]
    | ROLE ROLE_NAME [, ...]
    | ADMIN ROLE_NAME [, ...]
    | SYSID UID
```


## Description


 `CREATE USER` is now an alias for [`CREATE ROLE`](create-role.md#sql-createrole). The only difference is that when the command is spelled `CREATE USER`, `LOGIN` is assumed by default, whereas `NOLOGIN` is assumed when the command is spelled `CREATE ROLE`.


## Compatibility


 The `CREATE USER` statement is a PostgreSQL extension. The SQL standard leaves the definition of users to the implementation.


## See Also
  [sql-createrole](create-role.md#sql-createrole)
