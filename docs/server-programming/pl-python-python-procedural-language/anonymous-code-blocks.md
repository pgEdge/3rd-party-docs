<a id="plpython-do"></a>

## Anonymous Code Blocks


 PL/Python also supports anonymous code blocks called with the [sql-do](../../reference/sql-commands/do.md#sql-do) statement:

```

DO $$
    # PL/Python code
$$ LANGUAGE plpython3u;
```
 An anonymous code block receives no arguments, and whatever value it might return is discarded. Otherwise it behaves just like a function.
