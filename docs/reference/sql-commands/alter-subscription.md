<a id="sql-altersubscription"></a>

# ALTER SUBSCRIPTION

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


 Commands `ALTER SUBSCRIPTION ... REFRESH PUBLICATION`, `ALTER SUBSCRIPTION ... {SET|ADD|DROP} PUBLICATION ...` with `refresh` option as `true`, `ALTER SUBSCRIPTION ... SET (failover = true|false)` and `ALTER SUBSCRIPTION ... SET (two_phase = false)` cannot be executed inside a transaction block.


 Commands `ALTER SUBSCRIPTION ... REFRESH PUBLICATION` and `ALTER SUBSCRIPTION ... {SET|ADD|DROP} PUBLICATION ...` with `refresh` option as `true` also cannot be executed when the subscription has [`two_phase`](create-subscription.md#sql-createsubscription-params-with-two-phase) commit enabled, unless [`copy_data`](create-subscription.md#sql-createsubscription-params-with-copy-data) is `false`. See column `subtwophasestate` of [`pg_subscription`](../../internals/system-catalogs/pg_subscription.md#catalog-pg-subscription) to know the actual two-phase state.


## Parameters


<a id="sql-altersubscription-params-name"></a>

*name*
:   The name of a subscription whose properties are to be altered.
<a id="sql-altersubscription-params-connection"></a>

<code>CONNECTION '</code><em>conninfo</em><code>'</code>
:   This clause replaces the connection string originally set by [sql-createsubscription](create-subscription.md#sql-createsubscription). See there for more information.
<a id="sql-altersubscription-params-setadddrop-publication"></a>

<code>SET PUBLICATION </code><em>publication_name</em>, <code>ADD PUBLICATION </code><em>publication_name</em>, <code>DROP PUBLICATION </code><em>publication_name</em>
:   These forms change the list of subscribed publications. `SET` replaces the entire list of publications with a new list, `ADD` adds additional publications to the list of publications, and `DROP` removes the publications from the list of publications. We allow non-existent publications to be specified in `ADD` and `SET` variants so that users can add those later. See [sql-createsubscription](create-subscription.md#sql-createsubscription) for more information. By default, this command will also act like `REFRESH PUBLICATION`.


     *publication_option* specifies additional options for this operation. The supported options are:

    `refresh` (`boolean`)
    :   When false, the command will not try to refresh table information. `REFRESH PUBLICATION` should then be executed separately. The default is `true`.
     Additionally, the options described under `REFRESH PUBLICATION` may be specified, to control the implicit refresh operation.
<a id="sql-altersubscription-params-refresh-publication"></a>

`REFRESH PUBLICATION`
:   Fetch missing table information from publisher. This will start replication of tables that were added to the subscribed-to publications since [`CREATE SUBSCRIPTION`](create-subscription.md#sql-createsubscription) or the last invocation of `REFRESH PUBLICATION`.


     *refresh_option* specifies additional options for the refresh operation. The supported options are:

    `copy_data` (`boolean`)
    :   Specifies whether to copy pre-existing data in the publications that are being subscribed to when the replication starts. The default is `true`.


         Previously subscribed tables are not copied, even if a table's row filter `WHERE` clause has since been modified.


         See [Notes](create-subscription.md#sql-createsubscription-notes) for details of how `copy_data = true` can interact with the [`origin`](create-subscription.md#sql-createsubscription-params-with-origin) parameter.


         See the [`binary`](create-subscription.md#sql-createsubscription-params-with-binary) parameter of `CREATE SUBSCRIPTION` for details about copying pre-existing data in binary format.
<a id="sql-altersubscription-params-enable"></a>

`ENABLE`
:   Enables a previously disabled subscription, starting the logical replication worker at the end of the transaction.
<a id="sql-altersubscription-params-disable"></a>

`DISABLE`
:   Disables a running subscription, stopping the logical replication worker at the end of the transaction.
<a id="sql-altersubscription-params-set"></a>

<code>SET ( </code><em>subscription_parameter</em><code> [= </code><em>value</em><code>] [, ... ] )</code>
:   This clause alters parameters originally set by [sql-createsubscription](create-subscription.md#sql-createsubscription). See there for more information. The parameters that can be altered are [`slot_name`](create-subscription.md#sql-createsubscription-params-with-slot-name), [`synchronous_commit`](create-subscription.md#sql-createsubscription-params-with-synchronous-commit), [`binary`](create-subscription.md#sql-createsubscription-params-with-binary), [`streaming`](create-subscription.md#sql-createsubscription-params-with-streaming), [`disable_on_error`](create-subscription.md#sql-createsubscription-params-with-disable-on-error), [`password_required`](create-subscription.md#sql-createsubscription-params-with-password-required), [`run_as_owner`](create-subscription.md#sql-createsubscription-params-with-run-as-owner), [`origin`](create-subscription.md#sql-createsubscription-params-with-origin), [`failover`](create-subscription.md#sql-createsubscription-params-with-failover), and [`two_phase`](create-subscription.md#sql-createsubscription-params-with-two-phase). Only a superuser can set `password_required = false`.


     When altering the [`slot_name`](create-subscription.md#sql-createsubscription-params-with-slot-name), the `failover` and `two_phase` property values of the named slot may differ from the counterpart [`failover`](create-subscription.md#sql-createsubscription-params-with-failover) and [`two_phase`](create-subscription.md#sql-createsubscription-params-with-two-phase) parameters specified in the subscription. When creating the slot, ensure the slot properties `failover` and `two_phase` match their counterpart parameters of the subscription. Otherwise, the slot on the publisher may behave differently from what these subscription options say: for example, the slot on the publisher could either be synced to the standbys even when the subscription's [`failover`](create-subscription.md#sql-createsubscription-params-with-failover) option is disabled or could be disabled for sync even when the subscription's [`failover`](create-subscription.md#sql-createsubscription-params-with-failover) option is enabled.


     The [`failover`](create-subscription.md#sql-createsubscription-params-with-failover) and [`two_phase`](create-subscription.md#sql-createsubscription-params-with-two-phase) parameters can only be altered when the subscription is disabled.


     When altering [`two_phase`](create-subscription.md#sql-createsubscription-params-with-two-phase) from `true` to `false`, the backend process reports an error if any prepared transactions done by the logical replication worker (from when `two_phase` parameter was still `true`) are found. You can resolve prepared transactions on the publisher node, or manually roll back them on the subscriber, and then try again. The transactions prepared by logical replication worker corresponding to a particular subscription have the following pattern: “`pg_gid_%u_%u`” (parameters: subscription `oid`, remote transaction id `xid`). To resolve such transactions manually, you need to roll back all the prepared transactions with corresponding subscription IDs in their names. Applications can check [`pg_prepared_xacts`](../../internals/system-views/pg_prepared_xacts.md#view-pg-prepared-xacts) to find the required prepared transactions. After the `two_phase` option is changed from `true` to `false`, the publisher will replicate the transactions again when they are committed.
<a id="sql-altersubscription-params-skip"></a>

<code>SKIP ( </code><em>skip_option</em><code> = </code><em>value</em><code> )</code>
:   Skips applying all changes of the remote transaction. If incoming data violates any constraints, logical replication will stop until it is resolved. By using the `ALTER SUBSCRIPTION ... SKIP` command, the logical replication worker skips all data modification changes within the transaction. This option has no effect on the transactions that are already prepared by enabling [`two_phase`](create-subscription.md#sql-createsubscription-params-with-two-phase) on the subscriber. After the logical replication worker successfully skips the transaction or finishes a transaction, the LSN (stored in `pg_subscription`.`subskiplsn`) is cleared. See [Conflicts](../../server-administration/logical-replication/conflicts.md#logical-replication-conflicts) for the details of logical replication conflicts.


     *skip_option* specifies options for this operation. The supported option is:

    `lsn` (`pg_lsn`)
    :   Specifies the finish LSN of the remote transaction whose changes are to be skipped by the logical replication worker. The finish LSN is the LSN at which the transaction is either committed or prepared. Skipping individual subtransactions is not supported. Setting `NONE` resets the LSN.
<a id="sql-altersubscription-params-new-owner"></a>

*new_owner*
:   The user name of the new owner of the subscription.
<a id="sql-altersubscription-params-new-name"></a>

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
