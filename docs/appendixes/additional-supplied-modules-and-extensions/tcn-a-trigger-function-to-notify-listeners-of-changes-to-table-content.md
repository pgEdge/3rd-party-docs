<a id="tcn"></a>

## tcn — a trigger function to notify listeners of changes to table content


 The `tcn` module provides a trigger function that notifies listeners of changes to any table on which it is attached. It must be used as an `AFTER` trigger `FOR EACH ROW`.


 This module is considered “trusted”, that is, it can be installed by non-superusers who have `CREATE` privilege on the current database.


 Only one parameter may be supplied to the function in a `CREATE TRIGGER` statement, and that is optional. If supplied it will be used for the channel name for the notifications. If omitted `tcn` will be used for the channel name.


 The payload of the notifications consists of the table name, a letter to indicate which type of operation was performed, and column name/value pairs for primary key columns. Each part is separated from the next by a comma. For ease of parsing using regular expressions, table and column names are always wrapped in double quotes, and data values are always wrapped in single quotes. Embedded quotes are doubled.


 A brief example of using the extension follows.

```

test=# CREATE TABLE tcndata
test-#   (
test(#     a int NOT NULL,
test(#     b date NOT NULL,
test(#     c text,
test(#     PRIMARY KEY (a, b)
test(#   );
CREATE TABLE
test=# CREATE TRIGGER tcndata_tcn_trigger
test-#   AFTER INSERT OR UPDATE OR DELETE ON tcndata
test-#   FOR EACH ROW EXECUTE FUNCTION triggered_change_notification();
CREATE TRIGGER
test=# LISTEN tcn;
LISTEN
test=# INSERT INTO tcndata VALUES (1, date '2012-12-22', 'one'),
test-#                            (1, date '2012-12-23', 'another'),
test-#                            (2, date '2012-12-23', 'two');
INSERT 0 3
Asynchronous notification "tcn" with payload ""tcndata",I,"a"='1',"b"='2012-12-22'" received from server process with PID 22770.
Asynchronous notification "tcn" with payload ""tcndata",I,"a"='1',"b"='2012-12-23'" received from server process with PID 22770.
Asynchronous notification "tcn" with payload ""tcndata",I,"a"='2',"b"='2012-12-23'" received from server process with PID 22770.
test=# UPDATE tcndata SET c = 'uno' WHERE a = 1;
UPDATE 2
Asynchronous notification "tcn" with payload ""tcndata",U,"a"='1',"b"='2012-12-22'" received from server process with PID 22770.
Asynchronous notification "tcn" with payload ""tcndata",U,"a"='1',"b"='2012-12-23'" received from server process with PID 22770.
test=# DELETE FROM tcndata WHERE a = 1 AND b = date '2012-12-22';
DELETE 1
Asynchronous notification "tcn" with payload ""tcndata",D,"a"='1',"b"='2012-12-22'" received from server process with PID 22770.
```
