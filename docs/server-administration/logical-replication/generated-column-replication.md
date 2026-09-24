<a id="logical-replication-gencols"></a>

## Generated Column Replication


 Typically, a table at the subscriber will be defined the same as the publisher table, so if the publisher table has a [`GENERATED column`](../../the-sql-language/data-definition/generated-columns.md#ddl-generated-columns) then the subscriber table will have a matching generated column. In this case, it is always the subscriber table generated column value that is used.


 For example, note below that subscriber table generated column value comes from the subscriber column's calculation.

```

/* pub # */ CREATE TABLE tab_gen_to_gen (a int, b int GENERATED ALWAYS AS (a + 1) STORED);
/* pub # */ INSERT INTO tab_gen_to_gen VALUES (1),(2),(3);
/* pub # */ CREATE PUBLICATION pub1 FOR TABLE tab_gen_to_gen;
/* pub # */ SELECT * FROM tab_gen_to_gen;
 a | b
---+---
 1 | 2
 2 | 3
 3 | 4
(3 rows)

/* sub # */ CREATE TABLE tab_gen_to_gen (a int, b int GENERATED ALWAYS AS (a * 100) STORED);
/* sub # */ CREATE SUBSCRIPTION sub1 CONNECTION 'dbname=test_pub' PUBLICATION pub1;
/* sub # */ SELECT * FROM tab_gen_to_gen;
 a | b
---+----
 1 | 100
 2 | 200
 3 | 300
(3 rows)
```


 In fact, prior to version 18.0, logical replication does not publish `GENERATED` columns at all.


 But, replicating a generated column to a regular column can sometimes be desirable.

!!! tip

    This feature may be useful when replicating data to a non-PostgreSQL database via output plugin, especially if the target database does not support generated columns.


 Generated columns are not published by default, but users can opt to publish stored generated columns just like regular ones.


 There are two ways to do this:

-  Set the `PUBLICATION` parameter [`publish_generated_columns`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-with-publish-generated-columns) to `stored`. This instructs PostgreSQL logical replication to publish current and future stored generated columns of the publication's tables.
-  Specify a table [column list](column-lists.md#logical-replication-col-lists) to explicitly nominate which stored generated columns will be published.

!!! note

    When determining which table columns will be published, a column list takes precedence, overriding the effect of the `publish_generated_columns` parameter.


 The following table summarizes behavior when there are generated columns involved in the logical replication. Results are shown for when publishing generated columns is not enabled, and for when it is enabled.
 <a id="logical-replication-gencols-table-summary"></a>

**Table: Replication Result Summary**

| Publish generated columns? | Publisher table column | Subscriber table column | Result |
| --- | --- | --- | --- |
| No | GENERATED | GENERATED | Publisher table column is not replicated. Use the subscriber table generated column value. |
| No | GENERATED | regular | Publisher table column is not replicated. Use the subscriber table regular column default value. |
| No | GENERATED | --missing-- | Publisher table column is not replicated. Nothing happens. |
| Yes | GENERATED | GENERATED | ERROR. Not supported. |
| Yes | GENERATED | regular | Publisher table column value is replicated to the subscriber table column. |
| Yes | GENERATED | --missing-- | ERROR. The column is reported as missing from the subscriber table. |


!!! warning

    There's currently no support for subscriptions comprising several publications where the same table has been published with different column lists. See [Column Lists](column-lists.md#logical-replication-col-lists).


     This same situation can occur if one publication is publishing generated columns, while another publication in the same subscription is not publishing generated columns for the same table.


!!! note

    If the subscriber is from a release prior to 18, then initial table synchronization won't copy generated columns even if they are defined in the publisher.
