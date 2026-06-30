# supautils

[![Coverage Status](https://coveralls.io/repos/github/supabase/supautils/badge.svg?branch=master)](https://coveralls.io/github/supabase/supautils?branch=master)
![PostgreSQL version](https://img.shields.io/badge/postgresql-13+-blue.svg)

Supautils is an extension that unlocks advanced Postgres features without granting SUPERUSER access.

It's a loadable library that securely allows creating event triggers, publications, extensions to non-superusers. Built for cloud deployments where giving SUPERUSER rights to end users isn’t an option.

Completely managed by configuration — no tables, functions, or security labels are added to your database. This makes upgrades effortless and lets you apply settings cluster-wide solely via `postgresql.conf`.
