# Object Audit Logging

Object audit logging logs statements that affect a particular relation. Only `SELECT`, `INSERT`, `UPDATE` and `DELETE` commands are supported. `TRUNCATE` is not included in object audit logging.

Object audit logging is intended to be a finer-grained replacement for `pgaudit.log = 'read, write'`. As such, it may not make sense to use them in conjunction but one possible scenario would be to use session logging to capture each statement and then supplement that with object logging to get more detail about specific relations.

## Configuration

Object-level audit logging is implemented via the roles system. The [pgaudit.role](settings.md#pgauditrole) setting defines the role that will be used for audit logging. A relation (`TABLE`, `VIEW`, etc.) will be audit logged when the audit role has permissions for the command executed or inherits the permissions from another role. This allows you to effectively have multiple audit roles even though there is a single master role in any context.

Set [pgaudit.role](settings.md#pgauditrole) to `auditor` and grant `SELECT` and `DELETE` privileges on the `account` table. Any `SELECT` or `DELETE` statements on the `account` table will now be logged:
```
set pgaudit.role = 'auditor';

grant select, delete
   on public.account
   to auditor;
```

## Example

In this example object audit logging is used to illustrate how a granular approach may be taken towards logging of `SELECT` and `DML` statements. Note that logging on the `account` table is controlled by column-level permissions, while logging on the `account_role_map` table is table-level.

_SQL_:
```
set pgaudit.role = 'auditor';

create table account
(
    id int,
    name text,
    password text,
    description text
);

grant select (password)
   on public.account
   to auditor;

select id, name
  from account;

select password
  from account;

grant update (name, password)
   on public.account
   to auditor;

update account
   set description = 'yada, yada';

update account
   set password = 'HASH2';

create table account_role_map
(
    account_id int,
    role_id int
);

grant select
   on public.account_role_map
   to auditor;

select account.password,
       account_role_map.role_id
  from account
       inner join account_role_map
            on account.id = account_role_map.account_id
```
_Log Output_:
```
AUDIT: OBJECT,1,1,READ,SELECT,TABLE,public.account,select password
  from account,<not logged>
AUDIT: OBJECT,2,1,WRITE,UPDATE,TABLE,public.account,update account
   set password = 'HASH2',<not logged>
AUDIT: OBJECT,3,1,READ,SELECT,TABLE,public.account,select account.password,
       account_role_map.role_id
  from account
       inner join account_role_map
            on account.id = account_role_map.account_id,<not logged>
AUDIT: OBJECT,3,1,READ,SELECT,TABLE,public.account_role_map,select account.password,
       account_role_map.role_id
  from account
       inner join account_role_map
            on account.id = account_role_map.account_id,<not logged>
```

