<a id="ecpg-preproc"></a>

## Preprocessor Directives


 Several preprocessor directives are available that modify how the `ecpg` preprocessor parses and processes a file.
 <a id="ecpg-include"></a>

### Including Files


 To include an external file into your embedded SQL program, use:

```

EXEC SQL INCLUDE FILENAME;
EXEC SQL INCLUDE <FILENAME>;
EXEC SQL INCLUDE "FILENAME";
```
 The embedded SQL preprocessor will look for a file named <em>filename</em><code>.h</code>, preprocess it, and include it in the resulting C output. Thus, embedded SQL statements in the included file are handled correctly.


 The `ecpg` preprocessor will search a file at several directories in following order:

-

  current directory
-

`/usr/local/include`
-

  PostgreSQL include directory, defined at build time (e.g., `/usr/local/pgsql/include`)
-

`/usr/include`
 But when <code>EXEC SQL INCLUDE
    "</code><em>filename</em><code>"</code> is used, only the current directory is searched.


 In each directory, the preprocessor will first look for the file name as given, and if not found will append `.h` to the file name and try again (unless the specified file name already has that suffix).


 Note that `EXEC SQL INCLUDE` is *not* the same as:

```

#include <FILENAME.h>
```
 because this file would not be subject to SQL command preprocessing. Naturally, you can continue to use the C `#include` directive to include other header files.


!!! note

    The include file name is case-sensitive, even though the rest of the `EXEC SQL INCLUDE` command follows the normal SQL case-sensitivity rules.
  <a id="ecpg-define"></a>

### The define and undef Directives


 Similar to the directive `#define` that is known from C, embedded SQL has a similar concept:

```

EXEC SQL DEFINE NAME;
EXEC SQL DEFINE NAME VALUE;
```
 So you can define a name:

```

EXEC SQL DEFINE HAVE_FEATURE;
```
 And you can also define constants:

```

EXEC SQL DEFINE MYNUMBER 12;
EXEC SQL DEFINE MYSTRING 'abc';
```
 Use `undef` to remove a previous definition:

```

EXEC SQL UNDEF MYNUMBER;
```


 Of course you can continue to use the C versions `#define` and `#undef` in your embedded SQL program. The difference is where your defined values get evaluated. If you use `EXEC SQL DEFINE` then the `ecpg` preprocessor evaluates the defines and substitutes the values. For example if you write:

```

EXEC SQL DEFINE MYNUMBER 12;
...
EXEC SQL UPDATE Tbl SET col = MYNUMBER;
```
 then `ecpg` will already do the substitution and your C compiler will never see any name or identifier `MYNUMBER`. Note that you cannot use `#define` for a constant that you are going to use in an embedded SQL query because in this case the embedded SQL precompiler is not able to see this declaration.


 If multiple input files are named on the `ecpg` preprocessor's command line, the effects of `EXEC SQL DEFINE` and `EXEC SQL UNDEF` do not carry across files: each file starts with only the symbols defined by `-D` switches on the command line.
  <a id="ecpg-ifdef"></a>

### ifdef, ifndef, elif, else, and endif Directives


 You can use the following directives to compile code sections conditionally:

<a id="ecpg-ifdef-ifdef"></a>

<code>EXEC SQL ifdef </code><em>name</em><code>;</code>
:   Checks a *name* and processes subsequent lines if *name* has been defined via <code>EXEC SQL define
          </code><em>name</em>.
<a id="ecpg-ifdef-ifndef"></a>

<code>EXEC SQL ifndef </code><em>name</em><code>;</code>
:   Checks a *name* and processes subsequent lines if *name* has *not* been defined via <code>EXEC SQL define </code><em>name</em>.
<a id="ecpg-ifdef-elif"></a>

<code>EXEC SQL elif </code><em>name</em><code>;</code>
:   Begins an optional alternative section after an <code>EXEC SQL ifdef </code><em>name</em> or <code>EXEC SQL ifndef </code><em>name</em> directive. Any number of `elif` sections can appear. Lines following an `elif` will be processed if *name* has been defined *and* no previous section of the same `ifdef`/`ifndef`...`endif` construct has been processed.
<a id="ecpg-ifdef-else"></a>

`EXEC SQL else;`
:   Begins an optional, final alternative section after an <code>EXEC SQL ifdef </code><em>name</em> or <code>EXEC SQL ifndef </code><em>name</em> directive. Subsequent lines will be processed if no previous section of the same `ifdef`/`ifndef`...`endif` construct has been processed.
<a id="ecpg-ifdef-endif"></a>

`EXEC SQL endif;`
:   Ends an `ifdef`/`ifndef`...`endif` construct. Subsequent lines are processed normally.


 `ifdef`/`ifndef`...`endif` constructs can be nested, up to 127 levels deep.


 This example will compile exactly one of the three `SET TIMEZONE` commands:

```

EXEC SQL ifdef TZVAR;
EXEC SQL SET TIMEZONE TO TZVAR;
EXEC SQL elif TZNAME;
EXEC SQL SET TIMEZONE TO TZNAME;
EXEC SQL else;
EXEC SQL SET TIMEZONE TO 'GMT';
EXEC SQL endif;
```
