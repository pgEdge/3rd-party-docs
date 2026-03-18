<a id="queries-graph"></a>

## Graph Queries


 This section describes the sublanguage for querying property graphs, defined as described in [Property Graphs](../data-definition/property-graphs.md#ddl-property-graphs).
 <a id="queries-graph-overview"></a>

### Overview


 Consider this example from [Property Graphs](../data-definition/property-graphs.md#ddl-property-graphs):

```

-- get list of customers active today
SELECT customer_name FROM GRAPH_TABLE (myshop MATCH (c IS customers)-[IS customer_orders]->(o IS orders WHERE o.ordered_when = current_date) COLUMNS (c.name AS customer_name));
```
 The graph query part happens inside the `GRAPH_TABLE` construct. As far as the rest of the query is concerned, this acts like a table function in that it produces a computed table as output. Like other `FROM` clause elements, table alias and column alias names can be assigned to the result, and the result can be joined with other tables, subsequently filtered, and so on, for example:

```sql

SELECT ... FROM GRAPH_TABLE (mygraph MATCH ... COLUMNS (...)) AS myresult (a, b, c) JOIN othertable USING (a) WHERE b > 0 ORDER BY c;
```


 The `GRAPH_TABLE` clause consists of the graph name, followed by the keyword `MATCH`, followed by a graph pattern expression (see below), followed by the keyword `COLUMNS` and a column list.
  <a id="queries-graph-patterns"></a>

### Graph Patterns


 The core of the graph querying functionality is the graph pattern, which appears after the keyword `MATCH`. Formally, a graph pattern consists of one or more path patterns. A path is a sequence of graph elements, starting and ending with a vertex and alternating between vertices and edges. A path pattern is a syntactic expressions that matches paths.


 A path pattern thus matches a sequence of vertices and edges. The simplest possible path pattern is

```

()
```
 which matches a single vertex. The next simplest pattern would be

```

()-[]->()
```
 which matches a vertex followed by an edge followed by a vertex. The characters `()` are a vertex pattern and the characters `-[]->` are an edge pattern.


 These characters can also be separated by whitespace, for example:

```

( ) - [ ] - > ( )
```


!!! tip

    A way to remember these symbols is that in visual representations of property graphs, vertices are usually circles (like `()`) and edges have rectangular labels (like `[]`).


 The above patterns would match any vertex, or any two vertices connected by any edge, which isn't very interesting. Normally, we want to search for elements (vertices and edges) that have certain characteristics. These characteristics are written in between the parentheses or brackets. (This is also called an element pattern filler.) Typically, we would search for elements with a certain label. This is written by <code>IS
    </code><em>labelname</em>. For example, this would match all vertices with the label `person`:

```

(IS person)
```
 The next example would match a vertex with the label `person` connected to a vertex with the label `account` connected by an edge with the label `has`.

```

(IS person)-[IS has]->(IS account)
```
 Multiple labels can also be matched, using “or” semantics:

```

(IS person)-[IS has]->(IS account|creditcard)
```


 Recall that edges are directed. The other direction is also possible in a path pattern, for example:

```

(IS account)<-[IS has]-(IS person)
```
 It is also possible to match both directions:

```

(IS person)-[IS is_friend_of]-(IS person)
```
 This has a meaning of “or”: An edge in either direction would match.


 In many cases, the edge patterns don't need a filler. (All the filtering then happens on the vertices.) For these cases, an abbreviated edge pattern syntax is available that omits the brackets, for example:

```

(IS person)->(IS account)
(IS account)<-(IS person)
(IS person)-(IS person)
```
 As is often the case, abbreviated syntax can make expressions more compact but also sometimes harder to understand.


 Furthermore, it is possible to define graph pattern variables in the path pattern expressions. These are bound to the matched elements and can be used to refer to the property values from those elements. The most important use is to use them in the `COLUMNS` clause to define the tabular result of the `GRAPH_TABLE` clause. For example (assuming appropriate definitions of the property graph as well as the underlying tables):

```

GRAPH_TABLE (mygraph MATCH (p IS person)-[h IS has]->(a IS account)
             COLUMNS (p.name AS person_name, h.since AS has_account_since, a.num AS account_number)
```
 `WHERE` clauses can be used inside element patterns to filter matches:

```

(IS person)-[IS has]->(a IS account WHERE a.type = 'savings')
```
