<a id="plpython-sharing"></a>

## Sharing Data


 The global dictionary `SD` is available to store private data between repeated calls to the same function within a session. Each PL/Python function has its own instance of `SD`, except that for trigger functions, there is one instance of `SD` per function per triggered-on relation.


 The global dictionary `GD` is public data, that is available to all PL/Python functions within a session; use with care.


 Each function gets its own execution environment in the Python interpreter, so that global data and function arguments from `myfunc` are not available to `myfunc2`. The exception is the data in the `GD` dictionary, as mentioned above.
