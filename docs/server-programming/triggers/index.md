<a id="triggers"></a>

# Triggers

 This chapter provides general information about writing trigger functions. Trigger functions can be written in most of the available procedural languages, including PL/pgSQL ([PL/pgSQL — SQL Procedural Language](../pl-pgsql-sql-procedural-language/index.md#plpgsql)), PL/Tcl ([PL/Tcl — Tcl Procedural Language](../pl-tcl-tcl-procedural-language/index.md#pltcl)), PL/Perl ([PL/Perl — Perl Procedural Language](../pl-perl-perl-procedural-language/index.md#plperl)), and PL/Python ([PL/Python — Python Procedural Language](../pl-python-python-procedural-language/index.md#plpython)). After reading this chapter, you should consult the chapter for your favorite procedural language to find out the language-specific details of writing a trigger in it.

 It is also possible to write a trigger function in C, although most people find it easier to use one of the procedural languages. It is not currently possible to write a trigger function in the plain SQL function language.

- [Overview of Trigger Behavior](overview-of-trigger-behavior.md#trigger-definition)
- [Visibility of Data Changes](visibility-of-data-changes.md#trigger-datachanges)
- [Writing Trigger Functions in C](writing-trigger-functions-in-c.md#trigger-interface)
- [A Complete Trigger Example](a-complete-trigger-example.md#trigger-example)
