<a id="sql-createusermapping"></a>

# CREATE USER MAPPING

define a new mapping of a user to a foreign server

## Synopsis


```

CREATE USER MAPPING [ IF NOT EXISTS ] FOR { USER_NAME | USER | CURRENT_ROLE | CURRENT_USER | PUBLIC }
    SERVER SERVER_NAME
    [ OPTIONS ( OPTION 'VALUE' [ , ... ] ) ]
```


## Description


 `CREATE USER MAPPING` defines a mapping of a user to a foreign server. A user mapping typically encapsulates connection information that a foreign-data wrapper uses together with the information encapsulated by a foreign server to access an external data resource.


 The owner of a foreign server can create user mappings for that server for any user. Also, a user can create a user mapping for their own user name if `USAGE` privilege on the server has been granted to the user.


## Parameters


`IF NOT EXISTS`
:   Do not throw an error if a mapping of the given user to the given foreign server already exists. A notice is issued in this case. Note that there is no guarantee that the existing user mapping is anything like the one that would have been created.

*user_name*
:   The name of an existing user that is mapped to foreign server. `CURRENT_ROLE`, `CURRENT_USER`, and `USER` match the name of the current user. When `PUBLIC` is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.

*server_name*
:   The name of an existing server for which the user mapping is to be created.

<code>OPTIONS ( </code><em>option</em><code> '</code><em>value</em><code>' [, ... ] )</code>
:   This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.


## Examples


 Create a user mapping for user `bob`, server `foo`:

```sql

CREATE USER MAPPING FOR bob SERVER foo OPTIONS (user 'bob', password 'secret');
```


## Compatibility


 `CREATE USER MAPPING` conforms to ISO/IEC 9075-9 (SQL/MED).


## See Also
  [sql-alterusermapping](alter-user-mapping.md#sql-alterusermapping), [sql-dropusermapping](drop-user-mapping.md#sql-dropusermapping), [sql-createforeigndatawrapper](create-foreign-data-wrapper.md#sql-createforeigndatawrapper), [sql-createserver](create-server.md#sql-createserver)
