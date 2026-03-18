<a id="ddl-property-graphs"></a>

## Property Graphs


 A property graph is a way to represent database contents, as an alternative to the usual (in SQL) approach of representing database contents using relational structures such as tables. A property graph can then be queried using graph pattern matching syntax, instead of join queries typical of relational databases. PostgreSQL implements SQL/PGQ (Here, PGQ stands for “property graph query”. In the jargon of graph databases, “property graph” is normally abbreviated as PG, which is clearly confusing for practioners of PostgreSQL, also usually abbreviated as PG.), which is part of the SQL standard, where a property graph is defined as a kind of read-only view over relational tables. So the actual data is still in tables or table-like objects, but is exposed as a graph for graph querying operations. (This is in contrast to native graph databases, where the data is stored directly in a graph structure.) Underneath, both relational queries and graph queries use the same query planning and execution infrastructure, and in fact relational and graph queries can be combined and mixed in single queries.


 A graph is a set of vertices and edges. Each edge has two distinguishable associated vertices called the source and destination vertices. (So in this model, all edges are directed.) Vertices and edges together are called the elements of the graph. A property graph extends this well-known mathematical structure with a way to represent user data. In a property graph, each vertex or edge has one or more associated labels, and each label has zero or more properties. The labels are similar to table row types in that they define the kind of the contained data and its structure. The properties are similar to columns in that they contain the actual data. In fact, by default, a property graph definition exposes the underlying tables and columns as labels and properties, but more complicated definitions are possible.


 Consider the following table definitions:

```sql

CREATE TABLE products (
    product_no integer PRIMARY KEY,
    name varchar,
    price numeric
);

CREATE TABLE customers (
    customer_id integer PRIMARY KEY,
    name varchar,
    address varchar
);

CREATE TABLE orders (
    order_id integer PRIMARY KEY,
    ordered_when date
);

CREATE TABLE order_items (
    order_items_id integer PRIMARY KEY,
    order_id integer REFERENCES orders (order_id),
    product_no integer REFERENCES products (product_no),
    quantity integer
);

CREATE TABLE customer_orders (
    customer_orders_id integer PRIMARY KEY,
    customer_id integer REFERENCES customers (customer_id),
    order_id integer REFERENCES orders (order_id)
);
```
 When mapping this to a graph, the first three tables would be the vertices and the last two tables would be the edges. The foreign-key definitions correspond to the fact that edges link two vertices. (Graph definitions work more naturally with many-to-many relationships, so this example is organized like that, even though one-to-many relationships might be used here in a pure relational approach.)


 Here is an example how a property graph could be defined on top of these tables:

```sql

CREATE PROPERTY GRAPH myshop
    VERTEX TABLES (
        products,
        customers,
        orders
    )
    EDGE TABLES (
        order_items SOURCE orders DESTINATION products,
        customer_orders SOURCE customers DESTINATION orders
    );
```


 This graph could then be queried like this:

```

-- get list of customers active today
SELECT customer_name FROM GRAPH_TABLE (myshop MATCH (c IS customers)-[IS customer_orders]->(o IS orders WHERE o.ordered_when = current_date) COLUMNS (c.name AS customer_name));
```
 corresponding approximately to this relational query:

```

-- get list of customers active today
SELECT customers.name FROM customers JOIN customer_orders USING (customer_id) JOIN orders USING (order_id) WHERE orders.ordered_when = current_date;
```


 The above definition requires that all tables have primary keys and that for each edge there is an appropriate foreign key. Otherwise, additional clauses have to be specified to identify the key columns. For example, this would be the fully verbose definition that does not rely on primary and foreign keys:

```sql

CREATE PROPERTY GRAPH myshop
    VERTEX TABLES (
        products KEY (product_no),
        customers KEY (customer_id),
        orders KEY (order_id)
    )
    EDGE TABLES (
        order_items KEY (order_items_id)
            SOURCE KEY (order_id) REFERENCES orders (order_id)
            DESTINATION KEY (product_no) REFERENCES products (product_no),
        customer_orders KEY (customer_orders_id)
            SOURCE KEY (customer_id) REFERENCES customers (customer_id)
            DESTINATION KEY (order_id) REFERENCES orders (order_id)
    );
```


 As mentioned above, by default, the names of the tables and columns are exposed as labels and properties, respectively. The clauses `IS customer`, `IS order`, etc. in the `MATCH` clause in fact refer to labels, not table names.


 One use of labels is to expose a table through a different name in the graph. For example, in graphs, vertices typically have singular nouns as labels and edges typically have verbs or phrases derived from verbs as labels, such as “has”, “contains”, or something specific like “approved_by”. We can introduce such labels into our example like this:

```sql

CREATE PROPERTY GRAPH myshop
    VERTEX TABLES (
        products LABEL product,
        customers LABEL customer,
        orders LABEL "order"
    )
    EDGE TABLES (
        order_items SOURCE orders DESTINATION products LABEL contains,
        customer_orders SOURCE customers DESTINATION orders LABEL has_placed
    );
```


 With this definition, we can write a query like this:

```sql

SELECT customer_name FROM GRAPH_TABLE (myshop MATCH (c IS customer)-[IS has_placed]->(o IS "order" WHERE o.ordered_when = current_date) COLUMNS (c.name AS customer_name));
```
 With the new labels the `MATCH` clause is now more intuitive.


 Notice that the label `order` is quoted. If we run above statements without adding quotes around `order`, we will get a syntax error since `order` is a keyword.


 Another use is to apply the same label to multiple element tables. For example, consider this additional table:

```sql

CREATE TABLE employees (
    employee_id integer PRIMARY KEY,
    employee_name varchar,
    ...
);
```
 and the following graph definition:

```sql

CREATE PROPERTY GRAPH myshop
    VERTEX TABLES (
        products LABEL product,
        customers LABEL customer LABEL person PROPERTIES (name),
        orders LABEL order,
        employees LABEL employee LABEL person PROPERTIES (employee_name AS name)
    )
    EDGE TABLES (
        order_items SOURCE orders DESTINATION products LABEL contains,
        customer_orders SOURCE customers DESTINATION orders LABEL has
    );
```
 (In practice, there ought to be an edge linking the `employees` table to something, but it is allowed like this.) Then we can run a query like this (incomplete):

```sql

SELECT ... FROM GRAPH_TABLE (myshop MATCH (IS person WHERE name = '...')-[]->... COLUMNS (...));
```
 This would automatically consider both the `customers` and the `employees` tables when looking for an edge with the `person` label.


 When more than one element table has the same label, it is required that the properties match in number, name, and type. In the example, we specify an explicit property list and in one case override the name of the column to achieve this.


 Using more than one label associated with an element table and each label exposing a different set of properties, the same relational data, and the graph structure contained therein, can be exposed through multiple co-existing logical views, which can be queried using graph pattern matching constructs.


 For more details on the syntax for creating property graphs, see [`CREATE PROPERTY GRAPH`](../../reference/sql-commands/create-property-graph.md#sql-create-property-graph). More details about the graph query syntax is in [Graph Queries](../queries/graph-queries.md#queries-graph).
