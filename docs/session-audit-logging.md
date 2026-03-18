# Session Audit Logging

Session audit logging provides detailed logs of all statements executed by a user in the backend.

## Configuration

Session logging is enabled with the [pgaudit.log](settings.md#pgauditlog) setting.

Enable session logging for all `DML` and `DDL` and log all relations in `DML` statements:
```
set pgaudit.log = 'write, ddl';
set pgaudit.log_relation = on;
```
Enable session logging for all commands except `MISC` and raise audit log messages as `NOTICE`:
```
set pgaudit.log = 'all, -misc';
set pgaudit.log_level = notice;
```

## Example

In this example session audit logging is used for logging `DDL` and `SELECT` statements. Note that the insert statement is not logged since the `WRITE` class is not enabled

_SQL_:
```
set pgaudit.log = 'read, ddl';

create table account
(
    id int,
    name text,
    password text,
    description text
);

insert into account (id, name, password, description)
             values (1, 'user1', 'HASH1', 'blah, blah');

select *
    from account;
```
_Log Output_:
```
AUDIT: SESSION,1,1,DDL,CREATE TABLE,TABLE,public.account,create table account
(
    id int,
    name text,
    password text,
    description text
);,<not logged>
AUDIT: SESSION,2,1,READ,SELECT,,,select *
    from account,,<not logged>
```

