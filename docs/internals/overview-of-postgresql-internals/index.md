# Overview of PostgreSQL Internals { #overview }

!!! note "Author"

    This chapter originated as part of [Enhancement of the ANSI SQL Implementation of PostgreSQL](../../bibliography.md#sim98) Stefan Simkovics' Master's Thesis prepared at Vienna University of Technology under the direction of O.Univ.Prof.Dr. Georg Gottlob and Univ.Ass. Mag. Katrin Seyr.

 This chapter gives an overview of the internal structure of the backend of PostgreSQL. After having read the following sections you should have an idea of how a query is processed. This chapter is intended to help the reader understand the general sequence of operations that occur within the backend from the point at which a query is received, to the point at which the results are returned to the client.

- [The Path of a Query](the-path-of-a-query.md#query-path)
- [How Connections Are Established](how-connections-are-established.md#connect-estab)
- [The Parser Stage](the-parser-stage.md#parser-stage)
- [The PostgreSQL Rule System](the-postgresql-rule-system.md#rule-system)
- [Planner/Optimizer](planner-optimizer.md#planner-optimizer)
- [Executor](executor.md#executor)
