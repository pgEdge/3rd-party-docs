<a id="queries-union"></a>

## Combining Queries (`UNION`, `INTERSECT`, `EXCEPT`)


 The results of two queries can be combined using the set operations union, intersection, and difference. The syntax is

```

QUERY1 UNION [ALL] QUERY2
QUERY1 INTERSECT [ALL] QUERY2
QUERY1 EXCEPT [ALL] QUERY2
```
 where *query1* and *query2* are queries that can use any of the features discussed up to this point.


 `UNION` effectively appends the result of *query2* to the result of *query1* (although there is no guarantee that this is the order in which the rows are actually returned). Furthermore, it eliminates duplicate rows from its result, in the same way as `DISTINCT`, unless `UNION ALL` is used.


 `INTERSECT` returns all rows that are both in the result of *query1* and in the result of *query2*. Duplicate rows are eliminated unless `INTERSECT ALL` is used.


 `EXCEPT` returns all rows that are in the result of *query1* but not in the result of *query2*. (This is sometimes called the *difference* between two queries.) Again, duplicates are eliminated unless `EXCEPT ALL` is used.


 In order to calculate the union, intersection, or difference of two queries, the two queries must be “union compatible”, which means that they return the same number of columns and the corresponding columns have compatible data types, as described in [`UNION`, `CASE`, and Related Constructs](../type-conversion/union-case-and-related-constructs.md#typeconv-union-case).


 Set operations can be combined, for example

```

QUERY1 UNION QUERY2 EXCEPT QUERY3
```
 which is equivalent to

```

(QUERY1 UNION QUERY2) EXCEPT QUERY3
```
 As shown here, you can use parentheses to control the order of evaluation. Without parentheses, `UNION` and `EXCEPT` associate left-to-right, but `INTERSECT` binds more tightly than those two operators. Thus

```

QUERY1 UNION QUERY2 INTERSECT QUERY3
```
 means

```

QUERY1 UNION (QUERY2 INTERSECT QUERY3)
```
 You can also surround an individual *query* with parentheses. This is important if the *query* needs to use any of the clauses discussed in following sections, such as `LIMIT`. Without parentheses, you'll get a syntax error, or else the clause will be understood as applying to the output of the set operation rather than one of its inputs. For example,

```

SELECT a FROM b UNION SELECT x FROM y LIMIT 10
```
 is accepted, but it means

```

(SELECT a FROM b UNION SELECT x FROM y) LIMIT 10
```
 not

```

SELECT a FROM b UNION (SELECT x FROM y LIMIT 10)
```
