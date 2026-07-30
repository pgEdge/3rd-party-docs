<a id="rules"></a>

# The Rule System

 This chapter discusses the rule system in PostgreSQL. Production rule systems are conceptually simple, but there are many subtle points involved in actually using them.

 Some other database systems define active database rules, which are usually stored procedures and triggers. In PostgreSQL, these can be implemented using functions and triggers as well.

 The rule system (more precisely speaking, the query rewrite rule system) is totally different from stored procedures and triggers. It modifies queries to take rules into consideration, and then passes the modified query to the query planner for planning and execution. It is very powerful, and can be used for many things such as query language procedures, views, and versions. The theoretical foundations and the power of this rule system are also discussed in [ston90b](../../bibliography.md#ston90b) and [ong90](../../bibliography.md#ong90).

- [The Query Tree](the-query-tree.md#querytree)
- [Views and the Rule System](views-and-the-rule-system.md#rules-views)
- [Materialized Views](materialized-views.md#rules-materializedviews)
- [Rules on `INSERT`, `UPDATE`, and `DELETE`](rules-on-insert-update-and-delete.md#rules-update)
- [Rules and Privileges](rules-and-privileges.md#rules-privileges)
- [Rules and Command Status](rules-and-command-status.md#rules-status)
- [Rules Versus Triggers](rules-versus-triggers.md#rules-triggers)
