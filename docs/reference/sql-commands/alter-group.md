<a id="sql-altergroup"></a>

# ALTER GROUP

change role name or membership

## Synopsis


```

ALTER GROUP ROLE_SPECIFICATION ADD USER USER_NAME [, ... ]
ALTER GROUP ROLE_SPECIFICATION DROP USER USER_NAME [, ... ]

where ROLE_SPECIFICATION can be:

    ROLE_NAME
  | CURRENT_ROLE
  | CURRENT_USER
  | SESSION_USER

ALTER GROUP GROUP_NAME RENAME TO NEW_NAME
```


## Description


 `ALTER GROUP` changes the attributes of a user group. This is an obsolete command, though still accepted for backwards compatibility, because groups (and users too) have been superseded by the more general concept of roles.


 The first two variants add users to a group or remove them from a group. (Any role can play the part of either a “user” or a “group” for this purpose.) These variants are effectively equivalent to granting or revoking membership in the role named as the “group”; so the preferred way to do this is to use [`GRANT`](grant.md#sql-grant) or [`REVOKE`](revoke.md#sql-revoke). Note that `GRANT` and `REVOKE` have additional options which are not available with this command, such as the ability to grant and revoke `ADMIN OPTION`, and the ability to specify the grantor.


 The third variant changes the name of the group. This is exactly equivalent to renaming the role with [`ALTER ROLE`](alter-role.md#sql-alterrole).


## Parameters


*group_name*
:   The name of the group (role) to modify.

*user_name*
:   Users (roles) that are to be added to or removed from the group. The users must already exist; `ALTER GROUP` does not create or drop users.

*new_name*
:   The new name of the group.


## Examples


 Add users to a group:

```sql

ALTER GROUP staff ADD USER karl, john;
```
 Remove a user from a group:

```sql

ALTER GROUP workers DROP USER beth;
```


## Compatibility


 There is no `ALTER GROUP` statement in the SQL standard.


## See Also
  [sql-grant](grant.md#sql-grant), [sql-revoke](revoke.md#sql-revoke), [sql-alterrole](alter-role.md#sql-alterrole)
