<a id="sql-dropsubscription"></a>

# DROP SUBSCRIPTION

remove a subscription

## Synopsis


```

DROP SUBSCRIPTION [ IF EXISTS ] NAME [ CASCADE | RESTRICT ]
```


## Description


 `DROP SUBSCRIPTION` removes a subscription from the database cluster.


 To execute this command the user must be the owner of the subscription.


 `DROP SUBSCRIPTION` cannot be executed inside a transaction block if the subscription is associated with a replication slot. (You can use [`ALTER SUBSCRIPTION`](alter-subscription.md#sql-altersubscription) to unset the slot.)


## Parameters


*name*
:   The name of a subscription to be dropped.

`CASCADE`, `RESTRICT`
:   These key words do not have any effect, since there are no dependencies on subscriptions.


## Notes


 When dropping a subscription that is associated with a replication slot on the remote host (the normal state), `DROP SUBSCRIPTION` will connect to the remote host and try to drop the replication slot (and any remaining table synchronization slots) as part of its operation. This is necessary so that the resources allocated for the subscription on the remote host are released. If this fails, either because the remote host is not reachable or because the remote replication slot cannot be dropped or does not exist or never existed, the `DROP SUBSCRIPTION` command will fail. To proceed in this situation, first disable the subscription by executing [`ALTER SUBSCRIPTION ... DISABLE`](alter-subscription.md#sql-altersubscription-params-disable), and then disassociate it from the replication slot by executing [`ALTER SUBSCRIPTION ... SET (slot_name = NONE)`](alter-subscription.md#sql-altersubscription-params-set). After that, `DROP SUBSCRIPTION` will no longer attempt any actions on a remote host. Note that if the remote replication slot still exists, it (and any related table synchronization slots) should then be dropped manually; otherwise it/they will continue to reserve WAL and might eventually cause the disk to fill up. See also [Replication Slot Management](../../server-administration/logical-replication/subscription.md#logical-replication-subscription-slot).


 If a subscription is associated with a replication slot, then `DROP SUBSCRIPTION` cannot be executed inside a transaction block.


## Examples


 Drop a subscription:

```sql

DROP SUBSCRIPTION mysub;
```


## Compatibility


 `DROP SUBSCRIPTION` is a PostgreSQL extension.


## See Also
  [sql-createsubscription](create-subscription.md#sql-createsubscription), [sql-altersubscription](alter-subscription.md#sql-altersubscription)
