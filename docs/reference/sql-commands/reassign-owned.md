<a id="sql-reassign-owned"></a>

# REASSIGN OWNED

change the ownership of database objects owned by a database role

## Synopsis


```

REASSIGN OWNED BY { OLD_ROLE | CURRENT_ROLE | CURRENT_USER | SESSION_USER } [, ...]
               TO { NEW_ROLE | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
```


## Description


 `REASSIGN OWNED` instructs the system to change the ownership of database objects owned by any of the *old_roles* to *new_role*.


## Parameters


*old_role*
:   The name of a role. The ownership of all the objects within the current database, and of all shared objects (databases, tablespaces), owned by this role will be reassigned to *new_role*.

*new_role*
:   The name of the role that will be made the new owner of the affected objects.


## Notes


 `REASSIGN OWNED` is often used to prepare for the removal of one or more roles. Because `REASSIGN OWNED` does not affect objects within other databases, it is usually necessary to execute this command in each database that contains objects owned by a role that is to be removed.


 `REASSIGN OWNED` requires membership on both the source role(s) and the target role.


 The [`DROP OWNED`](drop-owned.md#sql-drop-owned) command is an alternative that simply drops all the database objects owned by one or more roles.


 The `REASSIGN OWNED` command does not affect any privileges granted to the *old_roles* on objects that are not owned by them. Likewise, it does not affect default privileges created with `ALTER DEFAULT PRIVILEGES`. Use `DROP OWNED` to revoke such privileges.


 See [Dropping Roles](../../server-administration/database-roles/dropping-roles.md#role-removal) for more discussion.


## Compatibility


 The `REASSIGN OWNED` command is a PostgreSQL extension.


## See Also
  [sql-drop-owned](drop-owned.md#sql-drop-owned), [sql-droprole](drop-role.md#sql-droprole), [sql-alterdatabase](alter-database.md#sql-alterdatabase)
