<a id="wal-for-extensions"></a>

# Write Ahead Logging for Extensions

 Certain extensions, principally extensions that implement custom access methods, may need to perform write-ahead logging in order to ensure crash-safety. PostgreSQL provides two ways for extensions to achieve this goal.

 First, extensions can choose to use [generic WAL](generic-wal-records.md#generic-wal), a special type of WAL record which describes changes to pages in a generic way. This method is simple to implement and does not require that an extension library be loaded in order to apply the records. However, generic WAL records will be ignored when performing logical decoding.

 Second, extensions can choose to use a [custom resource manager](custom-wal-resource-managers.md#custom-rmgr). This method is more flexible, supports logical decoding, and can sometimes generate much smaller write-ahead log records than would be possible with generic WAL. However, it is more complex for an extension to implement.

- [Generic WAL Records](generic-wal-records.md#generic-wal)
- [Custom WAL Resource Managers](custom-wal-resource-managers.md#custom-rmgr)
