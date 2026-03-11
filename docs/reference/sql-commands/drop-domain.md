# DROP DOMAIN { #sql-dropdomain }

remove a domain

## Synopsis


```

DROP DOMAIN [ IF EXISTS ] NAME [, ...] [ CASCADE | RESTRICT ]
```


## Description


 `DROP DOMAIN` removes a domain. Only the owner of a domain can remove it.


## Parameters


`IF EXISTS`
:   Do not throw an error if the domain does not exist. A notice is issued in this case.

*name*
:   The name (optionally schema-qualified) of an existing domain.

`CASCADE`
:   Automatically drop objects that depend on the domain (such as table columns), and in turn all objects that depend on those objects (see [Dependency Tracking](../../the-sql-language/data-definition/dependency-tracking.md#ddl-depend)).

`RESTRICT`
:   Refuse to drop the domain if any objects depend on it. This is the default.


## Examples { #sql-dropdomain-examples }


 To remove the domain `box`:

```sql

DROP DOMAIN box;
```


## Compatibility { #sql-dropdomain-compatibility }


 This command conforms to the SQL standard, except for the `IF EXISTS` option, which is a PostgreSQL extension.


## See Also { #sql-dropdomain-see-also }
  [sql-createdomain](create-domain.md#sql-createdomain), [sql-alterdomain](alter-domain.md#sql-alterdomain)
