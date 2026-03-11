<a id="sql-dropusermapping"></a>

# DROP USER MAPPING

remove a user mapping for a foreign server

## Synopsis


```

DROP USER MAPPING [ IF EXISTS ] FOR { USER_NAME | USER | CURRENT_ROLE | CURRENT_USER | PUBLIC } SERVER SERVER_NAME
```


## Description


 `DROP USER MAPPING` removes an existing user mapping from foreign server.


 The owner of a foreign server can drop user mappings for that server for any user. Also, a user can drop a user mapping for their own user name if `USAGE` privilege on the server has been granted to the user.


## Parameters


`IF EXISTS`
:   Do not throw an error if the user mapping does not exist. A notice is issued in this case.

*user_name*
:   User name of the mapping. `CURRENT_ROLE`, `CURRENT_USER`, and `USER` match the name of the current user. `PUBLIC` is used to match all present and future user names in the system.

*server_name*
:   Server name of the user mapping.


## Examples


 Drop a user mapping `bob`, server `foo` if it exists:

```sql

DROP USER MAPPING IF EXISTS FOR bob SERVER foo;
```


## Compatibility


 `DROP USER MAPPING` conforms to ISO/IEC 9075-9 (SQL/MED). The `IF EXISTS` clause is a PostgreSQL extension.


## See Also
  [sql-createusermapping](create-user-mapping.md#sql-createusermapping), [sql-alterusermapping](alter-user-mapping.md#sql-alterusermapping)
