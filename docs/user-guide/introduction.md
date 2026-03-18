# Introduction
<a name="introduction"></a>

This user guide is intended to be followed sequentially from beginning to end &mdash; each section depends on the last. For example, the [Restore](#/restore) section relies on setup that is performed in the [Quick Start](#/quickstart) section. Once pgBackRest is up and running then skipping around is possible but following the user guide in order is recommended the first time through.

Although the examples in this guide are targeted at Debian/Ubuntu and PostgreSQL 16, it should be fairly easy to apply the examples to any Unix distribution and PostgreSQL version. The only OS-specific commands are those to create, start, stop, and drop PostgreSQL clusters. The pgBackRest commands will be the same on any Unix system though the location of the executable may vary. While pgBackRest strives to operate consistently across versions of PostgreSQL, there are subtle differences between versions of PostgreSQL that may show up in this guide when illustrating certain examples, e.g. PostgreSQL path/file names and settings.

Configuration information and documentation for PostgreSQL can be found in the PostgreSQL [Manual](http://www.postgresql.org/docs/16/static/index.html).

A somewhat novel approach is taken to documentation in this user guide. Each command is run on a virtual machine when the documentation is built from the XML source. This means you can have a high confidence that the commands work correctly in the order presented. Output is captured and displayed below the command when appropriate. If the output is not included it is because it was deemed not relevant or was considered a distraction from the narrative.

All commands are intended to be run as an unprivileged user that has sudo privileges for both the `root` and `postgres` users. It's also possible to run the commands directly as their respective users without modification and in that case the `sudo` commands can be stripped off.
