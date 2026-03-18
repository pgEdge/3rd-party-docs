<a id="sql_user_management"></a>

# SQL User Management

As mentioned on [JWT Generation](../references/auth.md#jwt_generation), an external service can provide user management and coordinate with the PostgREST server using JWT. It's also possible to support logins entirely through SQL. It's a fair bit of work, so get ready.

## Storing Users and Passwords

The following table, functions, and triggers will live in a `basic_auth` schema that you shouldn't expose publicly in the API. The public views and functions will live in a different schema which internally references this internal information.

First we'll need a table to keep track of our users:

```sql
```

-- We put things inside the basic_auth schema to hide -- them from public view. Certain public procs/views will -- refer to helpers and tables inside.

create table basic_auth.users ( email    text primary key check ( email ~* '^.+@.+..+$' ), pass     text not null check (length(pass) < 512), role     name not null check (length(role) < 512) );

We would like the role to be a foreign key to actual database roles, however PostgreSQL does not support these constraints against the `pg_roles` table. We'll use a trigger to manually enforce it.

```postgres
```

create function basic_auth.check_role_exists() returns trigger as $$ begin if not exists (select 1 from pg_roles as r where r.rolname = new.role) then raise foreign_key_violation using message = 'unknown database role: ' || new.role; return null; end if; return new; end $$ language plpgsql;

create constraint trigger ensure_user_role_exists after insert or update on basic_auth.users for each row execute procedure basic_auth.check_role_exists();

Next we'll use the pgcrypto extension and a trigger to keep passwords safe in the `users` table.

```postgres
```

create extension pgcrypto;

create function basic_auth.encrypt_pass() returns trigger as $$ begin if tg_op = 'INSERT' or new.pass <> old.pass then new.pass = crypt(new.pass, gen_salt('bf')); end if; return new; end $$ language plpgsql;

create trigger encrypt_pass before insert or update on basic_auth.users for each row execute procedure basic_auth.encrypt_pass();

With the table in place we can make a helper to check a password against the encrypted column. It returns the database role for a user if the email and password are correct.

```postgres
```

create function basic_auth.user_role(email text, pass text) returns name language plpgsql as $$ begin return ( select role from basic_auth.users where users.email = user_role.email and users.pass = crypt(user_role.pass, users.pass) ); end; $$;
<a id="public_ui"></a>

## Public User Interface

In the previous section we created an internal table to store user information. Here we create a login function which takes an email address and password and returns JWT if the credentials match a user in the internal table.

### Permissions

Your database roles need access to the schema, tables, views and functions in order to service HTTP requests. Recall from the [Overview of role system](../references/auth.md#roles) that PostgREST uses special roles to process requests, namely the authenticator and anonymous roles. Below is an example of permissions that allow anonymous users to create accounts and attempt to log in.

```postgres
```

create role anon noinherit; create role authenticator noinherit; grant anon to authenticator;

Then, add `db-anon-role` to the configuration file to allow anonymous requests.

```ini
```

db-anon-role = "anon"
<a id="jwt-from-sql"></a>

### JWT from SQL

You can create JWT tokens in SQL using the [pgjwt extension](https://github.com/michelp/pgjwt). It's simple and requires only pgcrypto. If you're on an environment like Amazon RDS which doesn't support installing new extensions, you can still manually run the [SQL inside pgjwt](https://github.com/michelp/pgjwt/blob/master/pgjwt--0.1.1.sql) (you'll need to replace `@extschema@` with another schema or just delete it) which creates the functions you will need.

Next write a function that returns the token. The one below returns a token with a hard-coded role, which expires five minutes after it was issued. Note this function has a hard-coded secret as well.

```postgres
```

CREATE FUNCTION jwt_test(OUT token text) AS $$ SELECT public.sign( row_to_json(r), 'reallyreallyreallyreallyverysafe' ) AS token FROM ( SELECT 'my_role'::text as role, extract(epoch from now())::integer + 300 AS exp ) r; $$ LANGUAGE sql;

PostgREST exposes this function to clients via a POST request to `/rpc/jwt_test`.

!!! note


To avoid hard-coding the secret in functions, save it as a property of the database.

```postgres
 -- run this once
 ALTER DATABASE mydb SET "app.jwt_secret" TO 'reallyreallyreallyreallyverysafe';

 -- then all functions can refer to app.jwt_secret
 SELECT sign(
   row_to_json(r), current_setting('app.jwt_secret')
 ) AS token
 FROM ...
```

### Logins

As described in JWT from SQL we'll create a JWT inside our login function. Note that you'll need to adjust the secret key which is hard-coded in this example to a secure (at least thirty-two character) secret of your choosing.

```postgres
```

-- login should be on your exposed schema create function login(email text, pass text, out token text) as $$ declare _role name; begin -- check email and password select basic_auth.user_role(email, pass) into _role; if _role is null then raise invalid_password using message = 'invalid user or password'; end if;

select sign( row_to_json(r), 'reallyreallyreallyreallyverysafe' ) as token from ( select _role as role, login.email as email, extract(epoch from now())::integer + 60*60 as exp ) r into token; end; $$ language plpgsql security definer;

grant execute on function login(text,text) to anon;

Since the above [login` function is defined as `security definer](https://www.postgresql.org/docs/current/sql-createfunction.html#id-1.9.3.67.10.2), the anonymous user `anon` doesn't need permission to read the `basic_auth.users` table. It doesn't even need permission to access the `basic_auth` schema. `grant execute on function` is included for clarity but it might not be needed, see [Functions](../explanations/db_authz.md#func_privs) for more details.

An API request to call this function would look like:

```bash
```

curl "http://localhost:3000/rpc/login"  -X POST -H "Content-Type: application/json"  -d '{ "email": "foo@bar.com", "pass": "foobar" }'

The response would look like the snippet below. Try decoding the token at [jwt.io](https://jwt.io/). (It was encoded with a secret of `reallyreallyreallyreallyverysafe` as specified in the SQL code above. You'll want to change this secret in your app!)

```json
```

{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZvb0BiYXIuY29tIiwicGFzcyI6ImZvb2JhciJ9.37066TTRlh-1hXhnA9oO9Pj6lgL6zFuJU0iCHhuCFno" }

### Alternatives

See the how-to [SQL User Management using postgres' users and passwords](sql-user-management-using-postgres-users-and-passwords.md#sql-user-management-using-postgres-users-and-passwords) for a similar way that completely avoids the table `basic_auth.users`.
