<a id="plperl"></a>

# PL/Perl — Perl Procedural Language

 PL/Perl is a loadable procedural language that enables you to write PostgreSQL functions and procedures in the [Perl programming language](https://www.perl.org).

 The main advantage to using PL/Perl is that this allows use, within stored functions and procedures, of the manyfold “string munging” operators and functions available for Perl. Parsing complex strings might be easier using Perl than it is with the string functions and control structures provided in PL/pgSQL.

 To install PL/Perl in a particular database, use `CREATE EXTENSION plperl`.

!!! tip

    If a language is installed into `template1`, all subsequently created databases will have the language installed automatically.

!!! note

    Users of source packages must specially enable the build of PL/Perl during the installation process. (Refer to [Installation from Source Code](../../server-administration/installation-from-source-code/index.md#installation) for more information.) Users of binary packages might find PL/Perl in a separate subpackage.

- [PL/Perl Functions and Arguments](pl-perl-functions-and-arguments.md#plperl-funcs)
- [Data Values in PL/Perl](data-values-in-pl-perl.md#plperl-data)
- [Built-in Functions](built-in-functions.md#plperl-builtins)
- [Global Values in PL/Perl](global-values-in-pl-perl.md#plperl-global)
- [Trusted and Untrusted PL/Perl](trusted-and-untrusted-pl-perl.md#plperl-trusted)
- [PL/Perl Triggers](pl-perl-triggers.md#plperl-triggers)
- [PL/Perl Event Triggers](pl-perl-event-triggers.md#plperl-event-triggers)
- [PL/Perl Under the Hood](pl-perl-under-the-hood.md#plperl-under-the-hood)
