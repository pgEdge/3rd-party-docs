<a id="sql-droprole"></a>

# DROP ROLE

remove a database role

## Synopsis


```

DROP ROLE [ IF EXISTS ] NAME [, ...]
```


## Description


 `DROP ROLE` removes the specified role(s). To drop a superuser role, you must be a superuser yourself; to drop non-superuser roles, you must have `CREATEROLE` privilege and have been granted `ADMIN OPTION` on the role.


 A role cannot be removed if it is still referenced in any database of the cluster; an error will be raised if so. Before dropping the role, you must drop all the objects it owns (or reassign their ownership) and revoke any privileges the role has been granted on other objects. The [`REASSIGN OWNED`](reassign-owned.md#sql-reassign-owned) and [`DROP OWNED`](drop-owned.md#sql-drop-owned) commands can be useful for this purpose; see [Dropping Roles](../../server-administration/database-roles/dropping-roles.md#role-removal) for more discussion.


 However, it is not necessary to remove role memberships involving the role; `DROP ROLE` automatically revokes any memberships of the target role in other roles, and of other roles in the target role. The other roles are not dropped nor otherwise affected.


## Parameters


`IF EXISTS`
:   Do not throw an error if the role does not exist. A notice is issued in this case.

*name*
:   The name of the role to remove.


## Notes


 PostgreSQL includes a program [app-dropuser](../postgresql-client-applications/dropuser.md#app-dropuser) that has the same functionality as this command (in fact, it calls this command) but can be run from the command shell.


## Examples


 To drop a role:

```sql

DROP ROLE jonathan;
```


## Compatibility


 The SQL standard defines `DROP ROLE`, but it allows only one role to be dropped at a time, and it specifies different privilege requirements than PostgreSQL uses.


## See Also
  [sql-createrole](create-role.md#sql-createrole), [sql-alterrole](alter-role.md#sql-alterrole), [sql-set-role](set-role.md#sql-set-role)
