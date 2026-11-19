<a id="functions-logical"></a>

## Logical Operators


 The usual logical operators are available:

```

boolean AND boolean boolean
boolean OR boolean boolean
NOT boolean boolean
```
 SQL uses a three-valued logic system with true, false, and `null`, which represents “unknown”. In the truth tables below, the left operand of a binary operator selects the row, and the right operand the column:

| AND | TRUE | FALSE | NULL |
| --- | --- | --- | --- |
| TRUE | TRUE | FALSE | NULL |
| FALSE | FALSE | FALSE | FALSE |
| NULL | NULL | FALSE | NULL |


| OR | TRUE | FALSE | NULL |
| --- | --- | --- | --- |
| TRUE | TRUE | TRUE | TRUE |
| FALSE | TRUE | FALSE | NULL |
| NULL | TRUE | NULL | NULL |


| NOT | TRUE | FALSE | NULL |
| --- | --- | --- | --- |
|  | FALSE | TRUE | NULL |


 The operators `AND` and `OR` are commutative, that is, you can switch the left and right operands without affecting the result. (However, it is not guaranteed that the left operand is evaluated before the right operand. See [Expression Evaluation Rules](../sql-syntax/value-expressions.md#syntax-express-eval) for more information about the order of evaluation of subexpressions.)
