## Replicating Sequences { #logical-replication-sequences }


 To synchronize sequences from a publisher to a subscriber, first publish them using [`CREATE PUBLICATION ... FOR ALL SEQUENCES`](../../reference/sql-commands/create-publication.md#sql-createpublication-params-for-all-sequences) and then on the subscriber:


-  use [`CREATE SUBSCRIPTION`](../../reference/sql-commands/create-subscription.md#sql-createsubscription) to initially synchronize the published sequences.
-  use [`ALTER SUBSCRIPTION ... REFRESH PUBLICATION`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-refresh-publication) to synchronize only newly added sequences.
-  use [`ALTER SUBSCRIPTION ... REFRESH SEQUENCES`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-refresh-sequences) to re-synchronize all sequences currently known to the subscription.


 A *sequence synchronization worker* will be started after executing any of the above subscriber commands, and will exit once the sequences are synchronized.


 The ability to launch a sequence synchronization worker is limited by the [`max_sync_workers_per_subscription`](../server-configuration/replication.md#guc-max-sync-workers-per-subscription) configuration.


### Sequence Definition Mismatches { #sequence-definition-mismatches }


 The sequence synchronization worker validates that sequence definitions match between publisher and subscriber. If mismatches exist, the worker logs an error identifying them and exits. The apply worker continues respawning the sequence synchronization worker until synchronization succeeds. See also [`wal_retrieve_retry_interval`](../server-configuration/replication.md#guc-wal-retrieve-retry-interval).


 To resolve this, use [`ALTER SEQUENCE`](../../reference/sql-commands/alter-sequence.md#sql-altersequence) to align the subscriber's sequence parameters with those of the publisher.


### Refreshing Out-of-Sync Sequences { #sequences-out-of-sync }


 Subscriber sequence values will become out of sync as the publisher advances them.


 To detect this, compare the [pg_subscription_rel](../../internals/system-catalogs/pg_subscription_rel.md#catalog-pg-subscription-rel).`srsublsn` on the subscriber with the `page_lsn` obtained from the [`pg_get_sequence_data`](../../the-sql-language/functions-and-operators/sequence-manipulation-functions.md#func-pg-get-sequence-data) function for the sequence on the publisher. Then run [`ALTER SUBSCRIPTION ... REFRESH SEQUENCES`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-refresh-sequences) to re-synchronize if necessary.


!!! warning

    Each sequence caches a block of values (typically 32) in memory before generating a new WAL record, so its LSN advances only after the entire cached batch has been consumed. As a result, sequence value drift cannot be detected by LSN comparison when sequence increments fall within the same cached block (typically 32 values).


### Examples { #logical-replication-sequences-examples }


 Create some sequences on the publisher.

```

/* pub # */ CREATE SEQUENCE s1 START WITH 10 INCREMENT BY 1;
/* pub # */ CREATE SEQUENCE s2 START WITH 100 INCREMENT BY 10;
```


 Create the same sequences on the subscriber.

```

/* sub # */ CREATE SEQUENCE s1 START WITH 10 INCREMENT BY 1;
/* sub # */ CREATE SEQUENCE s2 START WITH 100 INCREMENT BY 10;
```


 Advance the sequences on the publisher a few times.

```

/* pub # */ SELECT nextval('s1');
 nextval
---------
      10
(1 row)
/* pub # */ SELECT nextval('s1');
 nextval
---------
      11
(1 row)
/* pub # */ SELECT nextval('s2');
 nextval
---------
     100
(1 row)
/* pub # */ SELECT nextval('s2');
 nextval
---------
     110
(1 row)
```


 Check the sequence page LSNs on the publisher.

```

/* pub # */ SELECT * FROM pg_get_sequence_data('s1');
 last_value | is_called |  page_lsn
------------+-----------+------------
         11 | t         | 0/0178F9E0
(1 row)
/* pub # */ SELECT * FROM pg_get_sequence_data('s2');
 last_value | is_called |  page_lsn
------------+-----------+------------
        110 | t         | 0/0178FAB0
(1 row)
```


 Create a publication for the sequences.

```

/* pub # */ CREATE PUBLICATION pub1 FOR ALL SEQUENCES;
```


 Subscribe to the publication.

```

/* sub # */ CREATE SUBSCRIPTION sub1
/* sub - */ CONNECTION 'host=localhost dbname=test_pub application_name=sub1'
/* sub - */ PUBLICATION pub1;
```


 Verify that the initial sequence values are synchronized.

```

/* sub # */ SELECT last_value, is_called FROM s1;
 last_value | is_called
------------+-----------
         11 | t
(1 row)

/* sub # */ SELECT last_value, is_called FROM s2;
 last_value | is_called
------------+-----------
        110 | t
(1 row)
```


 Confirm that the sequence page LSNs on the publisher have been recorded on the subscriber.

```

/* sub # */ SELECT srrelid::regclass, srsublsn FROM pg_subscription_rel;
 srrelid |  srsublsn
---------+------------
 s1      | 0/0178F9E0
 s2      | 0/0178FAB0
(2 rows)
```


 Advance the sequences on the publisher 50 more times.

```

/* pub # */  SELECT nextval('s1') FROM generate_series(1,50);
/* pub # */  SELECT nextval('s2') FROM generate_series(1,50);
```


 Check the sequence page LSNs on the publisher.

```

/* pub # */ SELECT * FROM pg_get_sequence_data('s1');
 last_value | is_called |  page_lsn
------------+-----------+------------
         61 | t         | 0/017CED28
(1 row)

/* pub # */ SELECT * FROM pg_get_sequence_data('s2');
 last_value | is_called |  page_lsn
------------+-----------+------------
        610 | t         | 0/017CEDF8
(1 row)
```


 The difference between the sequence page LSNs on the publisher and the sequence page LSNs on the subscriber indicates that the sequences are out of sync. Re-synchronize all sequences known to the subscriber using [`ALTER SUBSCRIPTION ... REFRESH SEQUENCES`](../../reference/sql-commands/alter-subscription.md#sql-altersubscription-params-refresh-sequences).

```

/* sub # */ ALTER SUBSCRIPTION sub1 REFRESH SEQUENCES;
```


 Recheck the sequences on the subscriber.

```

/* sub # */ SELECT last_value, is_called FROM s1;
 last_value | is_called
------------+-----------
         61 | t
(1 row)

/* sub # */ SELECT last_value, is_called FROM s2;
 last_value | is_called
------------+-----------
        610 | t
(1 row)
```
