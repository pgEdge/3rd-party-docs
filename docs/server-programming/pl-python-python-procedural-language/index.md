<a id="plpython"></a>

# PL/Python — Python Procedural Language

 The PL/Python procedural language allows PostgreSQL functions and procedures to be written in the [Python language](https://www.python.org).

 To install PL/Python in a particular database, use `CREATE EXTENSION plpython3u`.

!!! tip

    If a language is installed into `template1`, all subsequently created databases will have the language installed automatically.

 PL/Python is only available as an “untrusted” language, meaning it does not offer any way of restricting what users can do in it and is therefore named `plpython3u`. A trusted variant `plpython` might become available in the future if a secure execution mechanism is developed in Python. The writer of a function in untrusted PL/Python must take care that the function cannot be used to do anything unwanted, since it will be able to do anything that could be done by a user logged in as the database administrator. Only superusers can create functions in untrusted languages such as `plpython3u`.

!!! note

    Users of source packages must specially enable the build of PL/Python during the installation process. (Refer to the installation instructions for more information.) Users of binary packages might find PL/Python in a separate subpackage.

- [PL/Python Functions](pl-python-functions.md#plpython-funcs)
- [Data Values](data-values.md#plpython-data)
- [Sharing Data](sharing-data.md#plpython-sharing)
- [Anonymous Code Blocks](anonymous-code-blocks.md#plpython-do)
- [Trigger Functions](trigger-functions.md#plpython-trigger)
- [Database Access](database-access.md#plpython-database)
- [Explicit Subtransactions](explicit-subtransactions.md#plpython-subtransaction)
- [Transaction Management](transaction-management.md#plpython-transactions)
- [Utility Functions](utility-functions.md#plpython-util)
- [Python 2 vs. Python 3](python-2-vs-python-3.md#plpython-python23)
- [Environment Variables](environment-variables.md#plpython-envar)
