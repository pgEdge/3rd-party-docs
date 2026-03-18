# Psycopg -- PostgreSQL database adapter for Python

[Psycopg](https://psycopg.org/) is the most popular [PostgreSQL](https://www.postgresql.org/) database adapter for the [Python](https://www.python.org/) programming language.  Its main features are the complete implementation of the Python DBAPI specification and the thread safety (several threads can share the same connection). It was designed for heavily multi-threaded applications that create and destroy lots of cursors and make a large number of concurrent `INSERT`s or `UPDATE`s.

Psycopg 2 is mostly implemented in C as a [libpq](https://www.postgresql.org/docs/current/static/libpq.html) wrapper, resulting in being both efficient and secure. It features client-side and [server-side](usage.md#server-side-cursors) cursors, [asynchronous communication](advanced.md#async-support) and [notifications](advanced.md#async-notify), [COPY](usage.md#copy) support.  Many Python types are supported out-of-the-box and [adapted to matching PostgreSQL data types](usage.md#python-types-adaptation); adaptation can be extended and customized thanks to a flexible [objects adaptation system](advanced.md#adapting-new-types).

Psycopg 2 is both Unicode and Python 3 friendly.
<a id="Psycopg"></a>
<a id="PostgreSQL"></a>
<a id="Python"></a>
<a id="libpq"></a>
