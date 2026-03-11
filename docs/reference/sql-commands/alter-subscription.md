# ALTER SUBSCRIPTION { #sql-altersubscription }

change the definition of a subscription

## Synopsis


```

ALTER SUBSCRIPTION NAME CONNECTION 'CONNINFO'
ALTER SUBSCRIPTION NAME SET PUBLICATION PUBLICATION_NAME [, ...] [ WITH ( PUBLICATION_OPTION [= VALUE] [, ... ] ) ]
ALTER SUBSCRIPTION NAME ADD PUBLICATION PUBLICATION_NAME [, ...] [ WITH ( PUBLICATION_OPTION [= VALUE] [, ... ] ) ]
ALTER SUBSCRIPTION NAME DROP PUBLICATION PUBLICATION_NAME [, ...] [ WITH ( PUBLICATION_OPTION [= VALUE] [, ... ] ) ]
ALTER SUBSCRIPTION NAME REFRESH PUBLICATION [ WITH ( REFRESH_OPTION [= VALUE] [, ... ] ) ]
ALTER SUBSCRIPTION NAME ENABLE
ALTER SUBSCRIPTION NAME DISABLE
ALTER SUBSCRIPTION NAME SET ( SUBSCRIPTION_PARAMETER [= VALUE] [, ... ] )
ALTER SUBSCRIPTION NAME SKIP ( SKIP_OPTION = VALUE )
ALTER SUBSCRIPTION NAME OWNER TO { NEW_OWNER | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
ALTER SUBSCRIPTION NAME RENAME TO NEW_NAME
```


## Description


 `ALTER SUBSCRIPTION` can change most of the subscription properties that can be specified in [sql-createsubscription](create-subscription.md#sql-createsubscription).


 You must own the subscription to use `ALTER SUBSCRIPTION`. To rename a subscription or alter the owner, you must have `CREATE` permission on the database. In addition, to alter the owner, you must be able to `SET ROLE` to the new owning role. If the subscription has `password_required=false`, only superusers can modify it.


 When refreshing a publication we remove the relations that are no longer part of the publication and we also remove the table synchronization slots if there are any. It is necessary to remove these slots so that the resources allocated for the subscription on the remote host are released. If due to network breakdown or some other error, PostgreSQL is unable to remove the slots, an error will be reported. To proceed in this situation, the user either needs to retry the operation or disassociate the slot from the subscription and drop the subscription as explained in [sql-dropsubscription](drop-subscription.md#sql-dropsubscription).


 Commands `ALTER SUBSCRIPTION ... REFRESH PUBLICATION` and `ALTER SUBSCRIPTION ... {SET|ADD|DROP} PUBLICATION ...` with `refresh` option as `true` cannot be executed inside a transaction block. These commands also cannot be executed when the subscription has [`two_phase`](create-subscription.md#sql-createsubscription-with-two-phase) commit enabled, unless [`copy_data`](create-subscription.md#sql-createsubscription-with-copy-data) is `false`. See column `subtwophasestate` of [`pg_subscription`](../../internals/system-catalogs/pg_subscription.md#catalog-pg-subscription) to know the actual two-phase state.


## Parameters


*name*
:   The name of a subscription whose properties are to be altered.

<code>CONNECTION '</code><em>conninfo</em><code>'</code>
:   This clause replaces the connection string originally set by [sql-createsubscription](create-subscription.md#sql-createsubscription). See there for more information.

<code>SET PUBLICATION </code><em>publication_name</em>, <code>ADD PUBLICATION </code><em>publication_name</em>, <code>DROP PUBLICATION </code><em>publication_name</em>
:   These forms change the list of subscribed publications. `SET` replaces the entire list of publications with a new list, `ADD` adds additional publications to the list of publications, and `DROP` removes the publications from the list of publications. We allow non-existent publications to be specified in `ADD` and `SET` variants so that users can add those later. See [sql-createsubscription](create-subscription.md#sql-createsubscription) for more information. By default, this command will also act like `REFRESH PUBLICATION`.


     *publication_option* specifies additional options for this operation. The supported options are:

    `refresh` (`boolean`)
    :   When false, the command will not try to refresh table information. `REFRESH PUBLICATION` should then be executed separately. The default is `true`.
     Additionally, the options described under `REFRESH PUBLICATION` may be specified, to control the implicit refresh operation.

`REFRESH PUBLICATION`
:   Fetch missing table information from publisher. This will start replication of tables that were added to the subscribed-to publications since `CREATE SUBSCRIPTION` or the last invocation of `REFRESH PUBLICATION`.


     *refresh_option* specifies additional options for the refresh operation. The supported options are:

    `copy_data` (`boolean`)
    :   Specifies whether to copy pre-existing data in the publications that are being subscribed to when the replication starts. The default is `true`.


         Previously subscribed tables are not copied, even if a table's row filter `WHERE` clause has since been modified.


         See [Notes](create-subscription.md#sql-createsubscription-notes) for details of how `copy_data = true` can interact with the [`origin`](create-subscription.md#sql-createsubscription-with-origin) parameter.


         See the [`binary`](create-subscription.md#sql-createsubscription-with-binary) parameter of `CREATE SUBSCRIPTION` for details about copying pre-existing data in binary format.

`ENABLE`
:   Enables a previously disabled subscription, starting the logical replication worker at the end of the transaction.

`DISABLE`
:   Disables a running subscription, stopping the logical replication worker at the end of the transaction.

<code>SET ( </code><em>subscription_parameter</em><code> [= </code><em>value</em><code>] [, ... ] )</code>
:   This clause alters parameters originally set by [sql-createsubscription](create-subscription.md#sql-createsubscription). See there for more information. The parameters that can be altered are [`slot_name`](create-subscription.md#sql-createsubscription-with-slot-name), [`synchronous_commit`](create-subscription.md#sql-createsubscription-with-synchronous-commit), [`binary`](create-subscription.md#sql-createsubscription-with-binary), [`streaming`](create-subscription.md#sql-createsubscription-with-streaming), [`disable_on_error`](create-subscription.md#sql-createsubscription-with-disable-on-error), [`password_required`](create-subscription.md#sql-createsubscription-with-password-required), [`run_as_owner`](create-subscription.md#sql-createsubscription-with-run-as-owner), and [`origin`](create-subscription.md#sql-createsubscription-with-origin). Only a superuser can set `password_required = false`.

<code>SKIP ( </code><em>skip_option</em><code> = </code><em>value</em><code> )</code>
:   Skips applying all changes of the remote transaction. If incoming data violates any constraints, logical replication will stop until it is resolved. By using the `ALTER SUBSCRIPTION ... SKIP` command, the logical replication worker skips all data modification changes within the transaction. This option has no effect on the transactions that are already prepared by enabling [`two_phase`](create-subscription.md#sql-createsubscription-with-two-phase) on the subscriber. After the logical replication worker successfully skips the transaction or finishes a transaction, the LSN (stored in `pg_subscription`.`subskiplsn`) is cleared. See [Conflicts](../../server-administration/logical-replication/conflicts.md#logical-replication-conflicts) for the details of logical replication conflicts.


     *skip_option* specifies options for this operation. The supported option is:

    `lsn` (`pg_lsn`)
    :   Specifies the finish LSN of the remote transaction whose changes are to be skipped by the logical replication worker. The finish LSN is the LSN at which the transaction is either committed or prepared. Skipping individual subtransactions is not supported. Setting `NONE` resets the LSN.

*new_owner*
:   The user name of the new owner of the subscription.

*new_name*
:   The new name for the subscription.


 When specifying a parameter of type `boolean`, the `=` *value* part can be omitted, which is equivalent to specifying `TRUE`.


## Examples


 Change the publication subscribed by a subscription to `insert_only`:

```sql

ALTER SUBSCRIPTION mysub SET PUBLICATION insert_only;
```


 Disable (stop) the subscription:

```sql

ALTER SUBSCRIPTION mysub DISABLE;
```


## Compatibility


 `ALTER SUBSCRIPTION` is a PostgreSQL extension.


## See Also
  [sql-createsubscription](create-subscription.md#sql-createsubscription), [sql-dropsubscription](drop-subscription.md#sql-dropsubscription), [sql-createpublication](create-publication.md#sql-createpublication), [sql-alterpublication](alter-publication.md#sql-alterpublication)
