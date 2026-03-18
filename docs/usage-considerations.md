# Usage Considerations

Depending on settings, it is possible for pgAudit to generate an enormous volume of logging. Be careful to determine exactly what needs to be audit logged in your environment to avoid logging too much.

For example, when working in an OLAP environment it would probably not be wise to audit log inserts into a large fact table. The size of the log file will likely be many times the actual data size of the inserts because the log file is expressed as text. Since logs are generally stored with the OS this may lead to disk space being exhausted very quickly. In cases where it is not possible to limit audit logging to certain tables, be sure to assess the performance impact while testing and allocate plenty of space on the log volume. This may also be true for OLTP environments. Even if the insert volume is not as high, the performance impact of audit logging may still noticeably affect latency.

To limit the number of relations audit logged for `SELECT` and `DML` statements, consider using object audit logging (see [Object Auditing](object-audit-logging.md)). Object audit logging allows selection of the relations to be logged allowing for reduction of the overall log volume. However, when new relations are added they must be explicitly added to object audit logging. A programmatic solution where specified tables are excluded from logging and all others are included may be a good option in this case.

