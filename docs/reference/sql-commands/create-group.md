<a id="sql-creategroup"></a>

# CREATE GROUP

define a new database role

## Synopsis


```

CREATE GROUP NAME [ [ WITH ] OPTION [ ... ] ]

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
    | IN GROUP ROLE_NAME [, ...]
    | ROLE ROLE_NAME [, ...]
    | ADMIN ROLE_NAME [, ...]
    | USER ROLE_NAME [, ...]
    | SYSID UID
```


## Description


 `CREATE GROUP` is now an alias for [sql-createrole](create-role.md#sql-createrole).


## Compatibility


 There is no `CREATE GROUP` statement in the SQL standard.


## See Also
  [sql-createrole](create-role.md#sql-createrole)
