<a id="tutorial-delete"></a>

## Deletions


 Rows can be removed from a table using the `DELETE` command. Suppose you are no longer interested in the weather of Hayward. Then you can do the following to delete those rows from the table:

```sql

DELETE FROM weather WHERE city = 'Hayward';
```
 All weather records belonging to Hayward are removed.

```sql

SELECT * FROM weather;
```


```

     city      | temp_lo | temp_hi | prcp |    date
---------------+---------+---------+------+------------
 San Francisco |      46 |      50 | 0.25 | 1994-11-27
 San Francisco |      41 |      55 |    0 | 1994-11-29
(2 rows)
```


 One should be wary of statements of the form

```

DELETE FROM TABLENAME;
```
 Without a qualification, `DELETE` will remove *all* rows from the given table, leaving it empty. The system will not request confirmation before doing this!
