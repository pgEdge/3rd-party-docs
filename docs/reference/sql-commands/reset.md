<a id="sql-reset"></a>

# RESET

restore the value of a run-time parameter to the default value

## Synopsis


```

RESET CONFIGURATION_PARAMETER
RESET ALL
```


## Description


 `RESET` restores run-time parameters to their default values. `RESET` is an alternative spelling for

```

SET CONFIGURATION_PARAMETER TO DEFAULT
```
 Refer to [sql-set](set.md#sql-set) for details.


 The default value is defined as the value that the parameter would have had, if no `SET` had ever been issued for it in the current session. The actual source of this value might be a compiled-in default, the configuration file, command-line options, or per-database or per-user default settings. This is subtly different from defining it as “the value that the parameter had at session start”, because if the value came from the configuration file, it will be reset to whatever is specified by the configuration file now. See [Server Configuration](../../server-administration/server-configuration/index.md#runtime-config) for details.


 The transactional behavior of `RESET` is the same as `SET`: its effects will be undone by transaction rollback.


## Parameters


*configuration_parameter*
:   Name of a settable run-time parameter. Available parameters are documented in [Server Configuration](../../server-administration/server-configuration/index.md#runtime-config) and on the [sql-set](set.md#sql-set) reference page.

`ALL`
:   Resets all settable run-time parameters to default values.


## Examples


 Set the `timezone` configuration variable to its default value:

```

RESET timezone;
```


## Compatibility


 `RESET` is a PostgreSQL extension.


## See Also
  [sql-set](set.md#sql-set), [sql-show](show.md#sql-show)
