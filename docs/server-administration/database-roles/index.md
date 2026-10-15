<a id="user-manag"></a>

# Database Roles

 PostgreSQL manages database access permissions using the concept of *roles*. A role can be thought of as either a database user, or a group of database users, depending on how the role is set up. Roles can own database objects (for example, tables and functions) and can assign privileges on those objects to other roles to control who has access to which objects. Furthermore, it is possible to grant *membership* in a role to another role, thus allowing the member role to use privileges assigned to another role.

 The concept of roles subsumes the concepts of “users” and “groups”. In PostgreSQL versions before 8.1, users and groups were distinct kinds of entities, but now there are only roles. Any role can act as a user, a group, or both.

 This chapter describes how to create and manage roles. More information about the effects of role privileges on various database objects can be found in [Privileges](../../the-sql-language/data-definition/privileges.md#ddl-priv).

- [Database Roles](database-roles.md#database-roles)
- [Role Attributes](role-attributes.md#role-attributes)
- [Role Membership](role-membership.md#role-membership)
- [Dropping Roles](dropping-roles.md#role-removal)
- [Predefined Roles](predefined-roles.md#predefined-roles)
- [Function Security](function-security.md#perm-functions)
