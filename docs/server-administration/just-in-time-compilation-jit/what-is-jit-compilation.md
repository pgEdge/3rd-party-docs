<a id="jit-reason"></a>

## What Is JIT compilation?


 Just-in-Time (JIT) compilation is the process of turning some form of interpreted program evaluation into a native program, and doing so at run time. For example, instead of using general-purpose code that can evaluate arbitrary SQL expressions to evaluate a particular SQL predicate like `WHERE a.col = 3`, it is possible to generate a function that is specific to that expression and can be natively executed by the CPU, yielding a speedup.


 PostgreSQL has builtin support to perform JIT compilation using [LLVM](https://llvm.org/) when PostgreSQL is built with [`--with-llvm`](../installation-from-source-code/building-and-installation-with-autoconf-and-make.md#configure-with-llvm).


 See `src/backend/jit/README` for further details.
 <a id="jit-accelerated-operations"></a>

### JIT Accelerated Operations


 Currently PostgreSQL's JIT implementation has support for accelerating expression evaluation and tuple deforming. Several other operations could be accelerated in the future.


 Expression evaluation is used to evaluate `WHERE` clauses, target lists, aggregates and projections. It can be accelerated by generating code specific to each case.


 Tuple deforming is the process of transforming an on-disk tuple (see [Table Row Layout](../../internals/database-physical-storage/database-page-layout.md#storage-tuple-layout)) into its in-memory representation. It can be accelerated by creating a function specific to the table layout and the number of columns to be extracted.
  <a id="jit-inlining"></a>

### Inlining


 PostgreSQL is very extensible and allows new data types, functions, operators and other database objects to be defined; see [Extending SQL](../../server-programming/extending-sql/index.md#extend). In fact the built-in objects are implemented using nearly the same mechanisms. This extensibility implies some overhead, for example due to function calls (see [User-Defined Functions](../../server-programming/extending-sql/user-defined-functions.md#xfunc)). To reduce that overhead, JIT compilation can inline the bodies of small functions into the expressions using them. That allows a significant percentage of the overhead to be optimized away.
  <a id="jit-optimization"></a>

### Optimization


 LLVM has support for optimizing generated code. Some of the optimizations are cheap enough to be performed whenever JIT is used, while others are only beneficial for longer-running queries. See [https://llvm.org/docs/Passes.html#transform-passes](https://llvm.org/docs/Passes.html#transform-passes) for more details about optimizations.
