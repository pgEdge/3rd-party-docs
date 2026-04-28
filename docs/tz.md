# psycopg2.tz --  tzinfo implementations for Psycopg 2

*module* `psycopg2.tz`

!!! warning "Deprecated since version 2.9"

    The module will be dropped in psycopg 2.10. Use `datetime.timezone`
    instead.

This module holds two different tzinfo implementations that can be used as the `tzinfo` argument to `datetime` constructors, directly passed to Psycopg functions or used to set the `cursor.tzinfo_factory` attribute in cursors.

*class* `psycopg2.tz.FixedOffsetTimezone`

*class* `psycopg2.tz.LocalTimezone`
