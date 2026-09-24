<a id="logical-replication-col-lists"></a>

## Column Lists


 Each publication can optionally specify which columns of each table are replicated to subscribers. The table on the subscriber side must have at least all the columns that are published. If no column list is specified, then all columns on the publisher are replicated. See [sql-createpublication](../../reference/sql-commands/create-publication.md#sql-createpublication) for details on the syntax.


 The choice of columns can be based on behavioral or performance reasons. However, do not rely on this feature for security: a malicious subscriber is able to obtain data from columns that are not specifically published. If security is a consideration, protections can be applied at the publisher side.


 If no column list is specified, any columns added to the table later are automatically replicated. This means that having a column list which names all columns is not the same as having no column list at all.


 A column list can contain only simple column references. The order of columns in the list is not preserved.


 Generated columns can also be specified in a column list. This allows generated columns to be published, regardless of the publication parameter [`publish_generated_columns`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-with-publish-generated-columns). See [Generated Column Replication](generated-column-replication.md#logical-replication-gencols) for details.


 Specifying a column list when the publication also publishes [`FOR TABLES IN SCHEMA`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-for-tables-in-schema) is not supported.


 For partitioned tables, the publication parameter [`publish_via_partition_root`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-with-publish-via-partition-root) determines which column list is used. If `publish_via_partition_root` is `true`, the root partitioned table's column list is used. Otherwise, if `publish_via_partition_root` is `false` (the default), each partition's column list is used.


 If a publication publishes `UPDATE` or `DELETE` operations, any column list must include the table's replica identity columns (see [REPLICA IDENTITY](../../reference/sql-commands/alter-table.md#sql-altertable-replica-identity)). If a publication publishes only `INSERT` operations, then the column list may omit replica identity columns.


 Column lists have no effect for the `TRUNCATE` command.


 During initial data synchronization, only the published columns are copied. However, if the subscriber is from a release prior to 15, then all the columns in the table are copied during initial data synchronization, ignoring any column lists. If the subscriber is from a release prior to 18, then initial table synchronization won't copy generated columns even if they are defined in the publisher.
 <a id="logical-replication-col-list-combining"></a>

!!! warning "Warning: Combining Column Lists from Multiple Publications"

    There's currently no support for subscriptions comprising several publications where the same table has been published with different column lists. [sql-createsubscription](../../reference/sql-commands/create-subscription.md#sql-createsubscription) disallows creating such subscriptions, but it is still possible to get into that situation by adding or altering column lists on the publication side after a subscription has been created.


     This means changing the column lists of tables on publications that are already subscribed could lead to errors being thrown on the subscriber side.


     If a subscription is affected by this problem, the only way to resume replication is to adjust one of the column lists on the publication side so that they all match; and then either recreate the subscription, or use [`ALTER SUBSCRIPTION ... DROP PUBLICATION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-setadddrop-publication) to remove one of the offending publications and add it again.
 <a id="logical-replication-col-list-examples"></a>

### Examples


 Create a table `t1` to be used in the following example.

```

/* pub # */ CREATE TABLE t1(id int, a text, b text, c text, d text, e text, PRIMARY KEY(id));
```


 Create a publication `p1`. A column list is defined for table `t1` to reduce the number of columns that will be replicated. Notice that the order of column names in the column list does not matter.

```

/* pub # */ CREATE PUBLICATION p1 FOR TABLE t1 (id, b, a, d);
```


 `psql` can be used to show the column lists (if defined) for each publication.

```

/* pub # */ \dRp+
                                                Publication p1
  Owner   | All tables | All sequences | Inserts | Updates | Deletes | Truncates | Generated columns | Via root
----------+------------+---------------+---------+---------+---------+-----------+-------------------+----------
 postgres | f          | f             | t       | t       | t       | t         | none              | f
Tables:
    "public.t1" (id, a, b, d)
```


 `psql` can be used to show the column lists (if defined) for each table.

```

/* pub # */ \d t1
                 Table "public.t1"
 Column |  Type   | Collation | Nullable | Default
--------+---------+-----------+----------+---------
 id     | integer |           | not null |
 a      | text    |           |          |
 b      | text    |           |          |
 c      | text    |           |          |
 d      | text    |           |          |
 e      | text    |           |          |
Indexes:
    "t1_pkey" PRIMARY KEY, btree (id)
Included in publications:
    "p1" (id, a, b, d)
```


 On the subscriber node, create a table `t1` which now only needs a subset of the columns that were on the publisher table `t1`, and also create the subscription `s1` that subscribes to the publication `p1`.

```

/* sub # */ CREATE TABLE t1(id int, b text, a text, d text, PRIMARY KEY(id));
/* sub # */ CREATE SUBSCRIPTION s1
/* sub - */ CONNECTION 'host=localhost dbname=test_pub application_name=s1'
/* sub - */ PUBLICATION p1;
```


 On the publisher node, insert some rows to table `t1`.

```

/* pub # */ INSERT INTO t1 VALUES(1, 'a-1', 'b-1', 'c-1', 'd-1', 'e-1');
/* pub # */ INSERT INTO t1 VALUES(2, 'a-2', 'b-2', 'c-2', 'd-2', 'e-2');
/* pub # */ INSERT INTO t1 VALUES(3, 'a-3', 'b-3', 'c-3', 'd-3', 'e-3');
/* pub # */ SELECT * FROM t1 ORDER BY id;
 id |  a  |  b  |  c  |  d  |  e
----+-----+-----+-----+-----+-----
  1 | a-1 | b-1 | c-1 | d-1 | e-1
  2 | a-2 | b-2 | c-2 | d-2 | e-2
  3 | a-3 | b-3 | c-3 | d-3 | e-3
(3 rows)
```


 Only data from the column list of publication `p1` is replicated.

```

/* sub # */ SELECT * FROM t1 ORDER BY id;
 id |  b  |  a  |  d
----+-----+-----+-----
  1 | b-1 | a-1 | d-1
  2 | b-2 | a-2 | d-2
  3 | b-3 | a-3 | d-3
(3 rows)
```
