<a id="sql-createcollation"></a>

# CREATE COLLATION

define a new collation

## Synopsis


```

CREATE COLLATION [ IF NOT EXISTS ] NAME (
    [ LOCALE = LOCALE, ]
    [ LC_COLLATE = LC_COLLATE, ]
    [ LC_CTYPE = LC_CTYPE, ]
    [ PROVIDER = PROVIDER, ]
    [ DETERMINISTIC = BOOLEAN, ]
    [ RULES = RULES, ]
    [ VERSION = VERSION ]
)
CREATE COLLATION [ IF NOT EXISTS ] NAME FROM EXISTING_COLLATION
```
 <a id="sql-createcollation-description"></a>

## Description


 `CREATE COLLATION` defines a new collation using the specified operating system locale settings, or by copying an existing collation.


 To be able to create a collation, you must have `CREATE` privilege on the destination schema.


## Parameters


`IF NOT EXISTS`
:   Do not throw an error if a collation with the same name already exists. A notice is issued in this case. Note that there is no guarantee that the existing collation is anything like the one that would have been created.

*name*
:   The name of the collation. The collation name can be schema-qualified. If it is not, the collation is defined in the current schema. The collation name must be unique within that schema. (The system catalogs can contain collations with the same name for other encodings, but these are ignored if the database encoding does not match.)

*locale*
:   The locale name for this collation. See [libc Collations](../../server-administration/localization/collation-support.md#collation-managing-create-libc) and [ICU Collations](../../server-administration/localization/collation-support.md#collation-managing-create-icu) for details.


     If *provider* is `libc`, this is a shortcut for setting `LC_COLLATE` and `LC_CTYPE` at once. If you specify *locale*, you cannot specify either of those parameters.


     If *provider* is `builtin`, then *locale* must be specified and set to either `C`, `C.UTF-8` or `PG_UNICODE_FAST`.

*lc_collate*
:   If *provider* is `libc`, use the specified operating system locale for the `LC_COLLATE` locale category.

*lc_ctype*
:   If *provider* is `libc`, use the specified operating system locale for the `LC_CTYPE` locale category.

*provider*
:   Specifies the provider to use for locale services associated with this collation. Possible values are `builtin`, `icu` (if the server was built with ICU support) or `libc`. `libc` is the default. See [Locale Providers](../../server-administration/localization/locale-support.md#locale-providers) for details.

`DETERMINISTIC`
:   Specifies whether the collation should use deterministic comparisons. The default is true. A deterministic comparison considers strings that are not byte-wise equal to be unequal even if they are considered logically equal by the comparison. PostgreSQL breaks ties using a byte-wise comparison. Comparison that is not deterministic can make the collation be, say, case- or accent-insensitive. For that, you need to choose an appropriate `LOCALE` setting *and* set the collation to not deterministic here.


     Nondeterministic collations are only supported with the ICU provider.

*rules*
:   Specifies additional collation rules to customize the behavior of the collation. This is supported for ICU only. See [ICU Tailoring Rules](../../server-administration/localization/collation-support.md#icu-tailoring-rules) for details.

*version*
:   Specifies the version string to store with the collation. Normally, this should be omitted, which will cause the version to be computed from the actual version of the collation as provided by the operating system. This option is intended to be used by `pg_upgrade` for copying the version from an existing installation.


     See also [sql-altercollation](alter-collation.md#sql-altercollation) for how to handle collation version mismatches.

*existing_collation*
:   The name of an existing collation to copy. The new collation will have the same properties as the existing one, but it will be an independent object.
 <a id="sql-createcollation-notes"></a>

## Notes


 `CREATE COLLATION` takes a `SHARE ROW EXCLUSIVE` lock, which is self-conflicting, on the `pg_collation` system catalog, so only one `CREATE COLLATION` command can run at a time.


 Use `DROP COLLATION` to remove user-defined collations.


 See [Creating New Collation Objects](../../server-administration/localization/collation-support.md#collation-create) for more information on how to create collations.


 When using the `libc` collation provider, the locale must be applicable to the current database encoding. See [sql-createdatabase](create-database.md#sql-createdatabase) for the precise rules.
 <a id="sql-createcollation-examples"></a>

## Examples


 To create a collation from the operating system locale `fr_FR.utf8` (assuming the current database encoding is `UTF8`):

```sql

CREATE COLLATION french (locale = 'fr_FR.utf8');
```


 To create a collation using the ICU provider using German phone book sort order:

```sql

CREATE COLLATION german_phonebook (provider = icu, locale = 'de-u-co-phonebk');
```


 To create a collation using the ICU provider, based on the root ICU locale, with custom rules:

```sql

CREATE COLLATION custom (provider = icu, locale = 'und', rules = '&V << w <<< W');
```
 See [ICU Tailoring Rules](../../server-administration/localization/collation-support.md#icu-tailoring-rules) for further details and examples on the rules syntax.


 To create a collation from an existing collation:

```sql

CREATE COLLATION german FROM "de_DE";
```
 This can be convenient to be able to use operating-system-independent collation names in applications.
 <a id="sql-createcollation-compat"></a>

## Compatibility


 There is a `CREATE COLLATION` statement in the SQL standard, but it is limited to copying an existing collation. The syntax to create a new collation is a PostgreSQL extension.
 <a id="sql-createcollation-seealso"></a>

## See Also
  [sql-altercollation](alter-collation.md#sql-altercollation), [sql-dropcollation](drop-collation.md#sql-dropcollation)
