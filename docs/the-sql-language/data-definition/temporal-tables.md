<a id="ddl-temporal-tables"></a>

## Temporal Tables


 *Temporal tables* allow users to track different dimensions of history. *Application time* tracks the history of a thing out in the world, and *system time* tracks the history of the database itself. (A database that does both is also called *bitemporal*.) This section describes how to express and manage such histories in temporal tables.
 <a id="ddl-application-time"></a>

### Application Time


 *Application time* refers to a history of the entity described by a table. In a typical non-temporal table, there is a single row for each entity. In a temporal table, an entity may have multiple rows, as long as those rows describe non-overlapping periods from its history. Application time requires each row to have a start and end time, expressing when the row is applicable.


 The following SQL creates a temporal table that can store application time:

```sql

CREATE TABLE products (
    product_no integer,
    price numeric,
    valid_at daterange
);
```


 Records in a temporal table can be imagined on a timeline, as in [Application Time Example](#temporal-entities-figure). Here we show three records describing two products. Each record is a tuple with three attributes: the product number, the price, and the application time. So product 5 was first offered for a price of 5.00 starting January 1, 2020, but then became 8.00 starting January 1, 2022. Its second record has no specified end time, indicating that it is true indefinitely, or for all future time. The last record shows that product 6 was introduced January 1, 2021 for 9.00, then canceled January 1, 2024.
 <a id="temporal-entities-figure"></a>

**Application Time Example**


![image](images/temporal-entities.svg)


 In a table, these records would be:

```

 product_no | price |        valid_at
------------+-------+-------------------------
          5 |  5.00 | [2020-01-01,2022-01-01)
          5 |  8.00 | [2022-01-01,)
          6 |  9.00 | [2021-01-01,2024-01-01)
```


 We show the application time using range-type notation, because it is stored as a single column (either a range or multirange). Ranges include their start point but exclude their end point. That way two adjacent ranges cover all points without overlapping. See [Range Types](../data-types/range-types.md#rangetypes) for more information about range types.


 In principle, a table with application-time ranges/multiranges is equivalent to a table that stores application-time “instants”: one for each second, millisecond, nanosecond, or whatever finest granularity is available. But such a table would contain far too many rows, so ranges/multiranges offer an optimization to represent the same information in a compact form. In addition, ranges and multiranges offer a more convenient interface for typical temporal operations, where records change infrequently enough that separate “versions” persist for extended periods of time.
 <a id="ddl-application-time-primary-keys"></a>

#### Temporal Primary Keys and Unique Constraints


 A table with application time has a different concept of entity uniqueness than a non-temporal table. Temporal entity uniqueness can be enforced with a temporal primary key. A regular primary key has at least one column, all columns are `NOT NULL`, and the combined value of all columns is unique. A temporal primary key also has at least one such column, but in addition it has a final column that is of a range type or multirange type that shows when the row is applicable. The regular parts of the key must be unique for any moment in time, but non-unique rows are allowed if their application time does not overlap.


 The syntax to create a temporal primary key is as follows:

```sql

CREATE TABLE products (
    product_no integer,
    price numeric,
    valid_at daterange,
    PRIMARY KEY (product_no, valid_at WITHOUT OVERLAPS)
);
```
 In this example, `product_no` is the non-temporal part of the key, and `valid_at` is a range column containing the application time.


 The `WITHOUT OVERLAPS` column is implicitly `NOT NULL` (like the other parts of the key). In addition it may not contain empty values, that is, a range of `'empty'` or a multirange of `{}`. An empty application time would have no meaning.


 It is also possible to create a temporal unique constraint that is not a primary key. The syntax is similar:

```sql

CREATE TABLE products (
    product_no integer,
    price numeric,
    valid_at daterange,
    UNIQUE (product_no, valid_at WITHOUT OVERLAPS)
);
```
 Temporal unique constraints also forbid empty ranges/multiranges for their application time, but that column is permitted to be null (like the other columns of the unique constraint).


 Temporal primary keys and unique constraints are backed by GiST indexes (see [GiST Indexes](../../internals/built-in-index-access-methods/gist-indexes.md#gist)) rather than B-Tree indexes. In practice, creating a temporal primary key or constraint requires installing the [btree_gist](../../appendixes/additional-supplied-modules-and-extensions/btree_gist-gist-operator-classes-with-b-tree-behavior.md#btree-gist) extension, so that the database has GiST operator classes for the non-temporal parts of the key.


 Temporal primary keys and unique constraints have the same behavior as exclusion constraints (see [Exclusion Constraints](constraints.md#ddl-constraints-exclusion)), where each regular key part is compared with equality, and the application time is compared with overlaps, for example `EXCLUDE USING gist (id WITH =, valid_at WITH &&)`. The only difference is that they also forbid an empty application time.
  <a id="ddl-application-time-foreign-keys"></a>

#### Temporal Foreign Keys


 A temporal foreign key is a reference from one application-time table to another application-time table. Just as a non-temporal reference requires a referenced key to exist, so a temporal reference requires a referenced key to exist, but during whatever history the reference exists (at least). So if the `products` table is referenced by a `variants` table, and a variant of product 5 has an application-time of `[2020-01-01,2026-01-01)`, then product 5 must exist throughout that period.


 We can create the `variants` table with the following schema (without a foreign key yet):

```sql

CREATE TABLE variants (
  id         integer,
  product_no integer,
  name       text,
  valid_at   daterange,
  PRIMARY KEY (id, valid_at WITHOUT OVERLAPS)
);
```
 We have included a temporal primary key as a best practice, but it is not strictly required by foreign keys.


 [Temporal Foreign Key Example](#temporal-references-figure) plots product 5 (in green) and two variants referencing it (in yellow) on the same timeline. Variant 8 (Medium) was introduced first, then variant 9 (XXL). Both satisfy the foreign key constraint, because the referenced product exists throughout their entire history.
 <a id="temporal-references-figure"></a>

**Temporal Foreign Key Example**


![image](images/temporal-references.svg)


 In a table, these records would be:

```

 id | product_no |  name  |        valid_at
----+------------+--------+-------------------------
  8 |          5 | Medium | [2021-01-01,2023-06-01)
  9 |          5 | XXL    | [2022-03-01,2024-06-01)
```


 Note that a temporal reference need not be fulfilled by a single row in the referenced table. Product 5 had a price change in the middle of variant 8's history, but the reference is still valid. The combination of all matching rows is used to test whether the referenced history contains the referencing row.


 The syntax to add a temporal foreign key to our table is:

```sql

CREATE TABLE variants (
  id         integer,
  product_no integer,
  name       text,
  valid_at   daterange,
  PRIMARY KEY (id, valid_at WITHOUT OVERLAPS),
  FOREIGN KEY (product_no, PERIOD valid_at) REFERENCES products (product_no, PERIOD valid_at)
);
```
 Note that the keyword `PERIOD` must be used for the application-time column in both the referencing and referenced table.


 A temporal primary key or unique constraint matching the referenced columns must exist on the referenced table.


 PostgreSQL supports temporal foreign keys with action `NO ACTION`, but not `RESTRICT`, `CASCADE`, `SET NULL`, or `SET DEFAULT`.
   <a id="ddl-system-time"></a>

### System Time


 *System time* refers to the history of the database table, not the entity it describes. It captures when each row was inserted/updated/deleted.


 PostgreSQL does not currently support system time, but it could be emulated using triggers, and there are external extensions that provide such functionality.
