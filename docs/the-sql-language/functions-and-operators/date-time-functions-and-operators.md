<a id="functions-datetime"></a>

## Date/Time Functions and Operators


 [Date/Time Functions](#functions-datetime-table) shows the available functions for date/time value processing, with details appearing in the following subsections. [Date/Time Operators](#operators-datetime-table) illustrates the behaviors of the basic arithmetic operators (`+`, `*`, etc.). For formatting functions, refer to [Data Type Formatting Functions](data-type-formatting-functions.md#functions-formatting). You should be familiar with the background information on date/time data types from [Date/Time Types](../data-types/date-time-types.md#datatype-datetime).


 In addition, the usual comparison operators shown in [Comparison Operators](comparison-functions-and-operators.md#functions-comparison-op-table) are available for the date/time types. Dates and timestamps (with or without time zone) are all comparable, while times (with or without time zone) and intervals can only be compared to other values of the same data type. When comparing a timestamp without time zone to a timestamp with time zone, the former value is assumed to be given in the time zone specified by the [TimeZone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) configuration parameter, and is rotated to UTC for comparison to the latter value (which is already in UTC internally). Similarly, a date value is assumed to represent midnight in the `TimeZone` zone when comparing it to a timestamp.


 All the functions and operators described below that take `time` or `timestamp` inputs actually come in two variants: one that takes `time with time zone` or `timestamp with time zone`, and one that takes `time without time zone` or `timestamp without time zone`. For brevity, these variants are not shown separately. Also, the `+` and `*` operators come in commutative pairs (for example both `date` `+` `integer` and `integer` `+` `date`); we show only one of each such pair.
 <a id="operators-datetime-table"></a>

**Table: Date/Time Operators**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>date</code> <code>+</code> <code>integer</code> <code>date</code></td>
<td>Add a number of days to a date</td>
<td><code>date '2001-09-28' + 7</code> <code>2001-10-05</code></td>
</tr>
<tr>
<td><code>date</code> <code>+</code> <code>interval</code> <code>timestamp</code></td>
<td>Add an interval to a date</td>
<td><code>date '2001-09-28' + interval '1 hour'</code> <code>2001-09-28 01:00:00</code></td>
</tr>
<tr>
<td><code>date</code> <code>+</code> <code>time</code> <code>timestamp</code></td>
<td>Add a time-of-day to a date</td>
<td><code>date '2001-09-28' + time '03:00'</code> <code>2001-09-28 03:00:00</code></td>
</tr>
<tr>
<td><code>interval</code> <code>+</code> <code>interval</code> <code>interval</code></td>
<td>Add intervals</td>
<td><code>interval '1 day' + interval '1 hour'</code> <code>1 day 01:00:00</code></td>
</tr>
<tr>
<td><code>timestamp</code> <code>+</code> <code>interval</code> <code>timestamp</code></td>
<td>Add an interval to a timestamp</td>
<td><code>timestamp '2001-09-28 01:00' + interval '23 hours'</code> <code>2001-09-29 00:00:00</code></td>
</tr>
<tr>
<td><code>time</code> <code>+</code> <code>interval</code> <code>time</code></td>
<td>Add an interval to a time</td>
<td><code>time '01:00' + interval '3 hours'</code> <code>04:00:00</code></td>
</tr>
<tr>
<td><code>-</code> <code>interval</code> <code>interval</code></td>
<td>Negate an interval</td>
<td><code>- interval '23 hours'</code> <code>-23:00:00</code></td>
</tr>
<tr>
<td><code>date</code> <code>-</code> <code>date</code> <code>integer</code></td>
<td>Subtract dates, producing the number of days elapsed</td>
<td><code>date '2001-10-01' - date '2001-09-28'</code> <code>3</code></td>
</tr>
<tr>
<td><code>date</code> <code>-</code> <code>integer</code> <code>date</code></td>
<td>Subtract a number of days from a date</td>
<td><code>date '2001-10-01' - 7</code> <code>2001-09-24</code></td>
</tr>
<tr>
<td><code>date</code> <code>-</code> <code>interval</code> <code>timestamp</code></td>
<td>Subtract an interval from a date</td>
<td><code>date '2001-09-28' - interval '1 hour'</code> <code>2001-09-27 23:00:00</code></td>
</tr>
<tr>
<td><code>time</code> <code>-</code> <code>time</code> <code>interval</code></td>
<td>Subtract times</td>
<td><code>time '05:00' - time '03:00'</code> <code>02:00:00</code></td>
</tr>
<tr>
<td><code>time</code> <code>-</code> <code>interval</code> <code>time</code></td>
<td>Subtract an interval from a time</td>
<td><code>time '05:00' - interval '2 hours'</code> <code>03:00:00</code></td>
</tr>
<tr>
<td><code>timestamp</code> <code>-</code> <code>interval</code> <code>timestamp</code></td>
<td>Subtract an interval from a timestamp</td>
<td><code>timestamp '2001-09-28 23:00' - interval '23 hours'</code> <code>2001-09-28 00:00:00</code></td>
</tr>
<tr>
<td><code>interval</code> <code>-</code> <code>interval</code> <code>interval</code></td>
<td>Subtract intervals</td>
<td><code>interval '1 day' - interval '1 hour'</code> <code>1 day -01:00:00</code></td>
</tr>
<tr>
<td><code>timestamp</code> <code>-</code> <code>timestamp</code> <code>interval</code></td>
<td>Subtract timestamps (converting 24-hour intervals into days, similarly to <a href="#function-justify-hours"><code>justify_hours()</code></a>)</td>
<td><code>timestamp '2001-09-29 03:00' - timestamp '2001-07-27 12:00'</code> <code>63 days 15:00:00</code></td>
</tr>
<tr>
<td><code>interval</code> <code>*</code> <code>double precision</code> <code>interval</code></td>
<td>Multiply an interval by a scalar</td>
<td><code>interval '1 second' * 900</code> <code>00:15:00</code><br><code>interval '1 day' * 21</code> <code>21 days</code><br><code>interval '1 hour' * 3.5</code> <code>03:30:00</code></td>
</tr>
<tr>
<td><code>interval</code> <code>/</code> <code>double precision</code> <code>interval</code></td>
<td>Divide an interval by a scalar</td>
<td><code>interval '1 hour' / 1.5</code> <code>00:40:00</code></td>
</tr>
</tbody>
</table>
 <a id="functions-datetime-table"></a>

**Table: Date/Time Functions**

<table>
<thead>
<tr>
<th>Function</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>age</code> ( <code>timestamp</code>, <code>timestamp</code> ) <code>interval</code></td>
<td>Subtract arguments, producing a “symbolic” result that uses years and months, rather than just days</td>
<td><code>age(timestamp '2001-04-10', timestamp '1957-06-13')</code> <code>43 years 9 mons 27 days</code></td>
</tr>
<tr>
<td><code>age</code> ( <code>timestamp</code> ) <code>interval</code></td>
<td>Subtract argument from <code>current_date</code> (at midnight)</td>
<td><code>age(timestamp '1957-06-13')</code> <code>62 years 6 mons 10 days</code></td>
</tr>
<tr>
<td><code>clock_timestamp</code> ( ) <code>timestamp with time zone</code></td>
<td>Current date and time (changes during statement execution); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>clock_timestamp()</code> <code>2019-12-23 14:39:53.662522-05</code></td>
</tr>
<tr>
<td><code>current_date</code> <code>date</code></td>
<td>Current date; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>current_date</code> <code>2019-12-23</code></td>
</tr>
<tr>
<td><code>current_time</code> <code>time with time zone</code></td>
<td>Current time of day; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>current_time</code> <code>14:39:53.662522-05</code></td>
</tr>
<tr>
<td><code>current_time</code> ( <code>integer</code> ) <code>time with time zone</code></td>
<td>Current time of day, with limited precision; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>current_time(2)</code> <code>14:39:53.66-05</code></td>
</tr>
<tr>
<td><code>current_timestamp</code> <code>timestamp with time zone</code></td>
<td>Current date and time (start of current transaction); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>current_timestamp</code> <code>2019-12-23 14:39:53.662522-05</code></td>
</tr>
<tr>
<td><code>current_timestamp</code> ( <code>integer</code> ) <code>timestamp with time zone</code></td>
<td>Current date and time (start of current transaction), with limited precision; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>current_timestamp(0)</code> <code>2019-12-23 14:39:53-05</code></td>
</tr>
<tr>
<td><code>date_add</code> ( <code>timestamp with time zone</code>, <code>interval</code> [, <code>text</code> ] ) <code>timestamp with time zone</code></td>
<td>Add an <code>interval</code> to a <code>timestamp with time zone</code>, computing times of day and daylight-savings adjustments according to the time zone named by the third argument, or the current <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone">TimeZone</a> setting if that is omitted. The form with two arguments is equivalent to the <code>timestamp with time zone</code> <code>+</code> <code>interval</code> operator.</td>
<td><code>date_add('2021-10-31 00:00:00+02'::timestamptz, '1 day'::interval, 'Europe/Warsaw')</code> <code>2021-10-31 23:00:00+00</code></td>
</tr>
<tr>
<td><code>date_bin</code> ( <code>interval</code>, <code>timestamp</code>, <code>timestamp</code> ) <code>timestamp</code></td>
<td>Bin input into specified interval aligned with specified origin; see <a href="#functions-datetime-bin"><code>date_bin</code></a></td>
<td><code>date_bin('15 minutes', timestamp '2001-02-16 20:38:40', timestamp '2001-02-16 20:05:00')</code> <code>2001-02-16 20:35:00</code></td>
</tr>
<tr>
<td><code>date_part</code> ( <code>text</code>, <code>timestamp</code> ) <code>double precision</code></td>
<td>Get timestamp subfield (equivalent to <code>extract</code>); see <a href="#functions-datetime-extract"><code>EXTRACT</code>, <code>date_part</code></a></td>
<td><code>date_part('hour', timestamp '2001-02-16 20:38:40')</code> <code>20</code></td>
</tr>
<tr>
<td><code>date_part</code> ( <code>text</code>, <code>interval</code> ) <code>double precision</code></td>
<td>Get interval subfield (equivalent to <code>extract</code>); see <a href="#functions-datetime-extract"><code>EXTRACT</code>, <code>date_part</code></a></td>
<td><code>date_part('month', interval '2 years 3 months')</code> <code>3</code></td>
</tr>
<tr>
<td><code>date_subtract</code> ( <code>timestamp with time zone</code>, <code>interval</code> [, <code>text</code> ] ) <code>timestamp with time zone</code></td>
<td>Subtract an <code>interval</code> from a <code>timestamp with time zone</code>, computing times of day and daylight-savings adjustments according to the time zone named by the third argument, or the current <a href="../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone">TimeZone</a> setting if that is omitted. The form with two arguments is equivalent to the <code>timestamp with time zone</code> <code>-</code> <code>interval</code> operator.</td>
<td><code>date_subtract('2021-11-01 00:00:00+01'::timestamptz, '1 day'::interval, 'Europe/Warsaw')</code> <code>2021-10-30 22:00:00+00</code></td>
</tr>
<tr>
<td><code>date_trunc</code> ( <code>text</code>, <code>timestamp</code> ) <code>timestamp</code></td>
<td>Truncate to specified precision; see <a href="#functions-datetime-trunc"><code>date_trunc</code></a></td>
<td><code>date_trunc('hour', timestamp '2001-02-16 20:38:40')</code> <code>2001-02-16 20:00:00</code></td>
</tr>
<tr>
<td><code>date_trunc</code> ( <code>text</code>, <code>timestamp with time zone</code>, <code>text</code> ) <code>timestamp with time zone</code></td>
<td>Truncate to specified precision in the specified time zone; see <a href="#functions-datetime-trunc"><code>date_trunc</code></a></td>
<td><code>date_trunc('day', timestamptz '2001-02-16 20:38:40+00', 'Australia/Sydney')</code> <code>2001-02-16 13:00:00+00</code></td>
</tr>
<tr>
<td><code>date_trunc</code> ( <code>text</code>, <code>interval</code> ) <code>interval</code></td>
<td>Truncate to specified precision; see <a href="#functions-datetime-trunc"><code>date_trunc</code></a></td>
<td><code>date_trunc('hour', interval '2 days 3 hours 40 minutes')</code> <code>2 days 03:00:00</code></td>
</tr>
<tr>
<td><code>extract</code> ( <code>field</code> <code>FROM</code> <code>timestamp</code> ) <code>numeric</code></td>
<td>Get timestamp subfield; see <a href="#functions-datetime-extract"><code>EXTRACT</code>, <code>date_part</code></a></td>
<td><code>extract(hour FROM timestamp '2001-02-16 20:38:40')</code> <code>20</code></td>
</tr>
<tr>
<td><code>extract</code> ( <code>field</code> <code>FROM</code> <code>interval</code> ) <code>numeric</code></td>
<td>Get interval subfield; see <a href="#functions-datetime-extract"><code>EXTRACT</code>, <code>date_part</code></a></td>
<td><code>extract(month FROM interval '2 years 3 months')</code> <code>3</code></td>
</tr>
<tr>
<td><code>isfinite</code> ( <code>date</code> ) <code>boolean</code></td>
<td>Test for finite date (not +/-infinity)</td>
<td><code>isfinite(date '2001-02-16')</code> <code>true</code></td>
</tr>
<tr>
<td><code>isfinite</code> ( <code>timestamp</code> ) <code>boolean</code></td>
<td>Test for finite timestamp (not +/-infinity)</td>
<td><code>isfinite(timestamp 'infinity')</code> <code>false</code></td>
</tr>
<tr>
<td><code>isfinite</code> ( <code>interval</code> ) <code>boolean</code></td>
<td>Test for finite interval (not +/-infinity)</td>
<td><code>isfinite(interval '4 hours')</code> <code>true</code></td>
</tr>
<tr>
<td><a id="function-justify-days"></a>
 `justify_days` ( `interval` ) `interval`</td>
<td>Adjust interval, converting 30-day time periods to months</td>
<td><code>justify_days(interval '1 year 65 days')</code> <code>1 year 2 mons 5 days</code></td>
</tr>
<tr>
<td><a id="function-justify-hours"></a>
 `justify_hours` ( `interval` ) `interval`</td>
<td>Adjust interval, converting 24-hour time periods to days</td>
<td><code>justify_hours(interval '50 hours 10 minutes')</code> <code>2 days 02:10:00</code></td>
</tr>
<tr>
<td><code>justify_interval</code> ( <code>interval</code> ) <code>interval</code></td>
<td>Adjust interval using <code>justify_days</code> and <code>justify_hours</code>, with additional sign adjustments</td>
<td><code>justify_interval(interval '1 mon -1 hour')</code> <code>29 days 23:00:00</code></td>
</tr>
<tr>
<td><code>localtime</code> <code>time</code></td>
<td>Current time of day; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>localtime</code> <code>14:39:53.662522</code></td>
</tr>
<tr>
<td><code>localtime</code> ( <code>integer</code> ) <code>time</code></td>
<td>Current time of day, with limited precision; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>localtime(0)</code> <code>14:39:53</code></td>
</tr>
<tr>
<td><code>localtimestamp</code> <code>timestamp</code></td>
<td>Current date and time (start of current transaction); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>localtimestamp</code> <code>2019-12-23 14:39:53.662522</code></td>
</tr>
<tr>
<td><code>localtimestamp</code> ( <code>integer</code> ) <code>timestamp</code></td>
<td>Current date and time (start of current transaction), with limited precision; see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>localtimestamp(2)</code> <code>2019-12-23 14:39:53.66</code></td>
</tr>
<tr>
<td><code>make_date</code> ( <code>year</code> <code>int</code>, <code>month</code> <code>int</code>, <code>day</code> <code>int</code> ) <code>date</code></td>
<td>Create date from year, month and day fields (negative years signify BC)</td>
<td><code>make_date(2013, 7, 15)</code> <code>2013-07-15</code></td>
</tr>
<tr>
<td><code>make_interval</code> ( [ <code>years</code> <code>int</code> [, <code>months</code> <code>int</code> [, <code>weeks</code> <code>int</code> [, <code>days</code> <code>int</code> [, <code>hours</code> <code>int</code> [, <code>mins</code> <code>int</code> [, <code>secs</code> <code>double precision</code> ]]]]]]] ) <code>interval</code></td>
<td>Create interval from years, months, weeks, days, hours, minutes and seconds fields, each of which can default to zero</td>
<td><code>make_interval(days =&gt; 10)</code> <code>10 days</code></td>
</tr>
<tr>
<td><code>make_time</code> ( <code>hour</code> <code>int</code>, <code>min</code> <code>int</code>, <code>sec</code> <code>double precision</code> ) <code>time</code></td>
<td>Create time from hour, minute and seconds fields</td>
<td><code>make_time(8, 15, 23.5)</code> <code>08:15:23.5</code></td>
</tr>
<tr>
<td><code>make_timestamp</code> ( <code>year</code> <code>int</code>, <code>month</code> <code>int</code>, <code>day</code> <code>int</code>, <code>hour</code> <code>int</code>, <code>min</code> <code>int</code>, <code>sec</code> <code>double precision</code> ) <code>timestamp</code></td>
<td>Create timestamp from year, month, day, hour, minute and seconds fields (negative years signify BC)</td>
<td><code>make_timestamp(2013, 7, 15, 8, 15, 23.5)</code> <code>2013-07-15 08:15:23.5</code></td>
</tr>
<tr>
<td><code>make_timestamptz</code> ( <code>year</code> <code>int</code>, <code>month</code> <code>int</code>, <code>day</code> <code>int</code>, <code>hour</code> <code>int</code>, <code>min</code> <code>int</code>, <code>sec</code> <code>double precision</code> [, <code>timezone</code> <code>text</code> ] ) <code>timestamp with time zone</code></td>
<td>Create timestamp with time zone from year, month, day, hour, minute and seconds fields (negative years signify BC). If <code>timezone</code> is not specified, the current time zone is used; the examples assume the session time zone is <code>Europe/London</code></td>
<td><code>make_timestamptz(2013, 7, 15, 8, 15, 23.5)</code> <code>2013-07-15 08:15:23.5+01</code><br><code>make_timestamptz(2013, 7, 15, 8, 15, 23.5, 'America/New_York')</code> <code>2013-07-15 13:15:23.5+01</code></td>
</tr>
<tr>
<td><code>now</code> ( ) <code>timestamp with time zone</code></td>
<td>Current date and time (start of current transaction); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>now()</code> <code>2019-12-23 14:39:53.662522-05</code></td>
</tr>
<tr>
<td><code>random</code> ( <code>min</code> <code>date</code>, <code>max</code> <code>date</code> ) <code>date</code></td>
<td><code>random</code> ( <code>min</code> <code>timestamp</code>, <code>max</code> <code>timestamp</code> ) <code>timestamp</code></td>
<td><code>random</code> ( <code>min</code> <code>timestamptz</code>, <code>max</code> <code>timestamptz</code> ) <code>timestamptz</code><br>Returns a random value in the range <code>min</code> <= x <= <code>max</code>.<br>Note that these functions use the same pseudo-random number generator as the functions listed in <a href="mathematical-functions-and-operators.md#functions-math-random-table">Random Functions</a>, and respond in the same way to calling <a href="mathematical-functions-and-operators.md#function-setseed"><code>setseed()</code></a>.<br><code>random('1979-02-08'::date,'2025-07-03'::date)</code> <code>1983-04-21</code><br><code>random('2000-01-01'::timestamptz, now())</code> <code>2015-09-27 09:11:33.732707+00</code></td>
</tr>
<tr>
<td><code>statement_timestamp</code> ( ) <code>timestamp with time zone</code></td>
<td>Current date and time (start of current statement); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>statement_timestamp()</code> <code>2019-12-23 14:39:53.662522-05</code></td>
</tr>
<tr>
<td><code>timeofday</code> ( ) <code>text</code></td>
<td>Current date and time (like <code>clock_timestamp</code>, but as a <code>text</code> string); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>timeofday()</code> <code>Mon Dec 23 14:39:53.662522 2019 EST</code></td>
</tr>
<tr>
<td><code>transaction_timestamp</code> ( ) <code>timestamp with time zone</code></td>
<td>Current date and time (start of current transaction); see <a href="#functions-datetime-current">Current Date/Time</a></td>
<td><code>transaction_timestamp()</code> <code>2019-12-23 14:39:53.662522-05</code></td>
</tr>
<tr>
<td><code>to_timestamp</code> ( <code>double precision</code> ) <code>timestamp with time zone</code></td>
<td>Convert Unix epoch (seconds since 1970-01-01 00:00:00+00) to timestamp with time zone</td>
<td><code>to_timestamp(1284352323)</code> <code>2010-09-13 04:32:03+00</code></td>
</tr>
</tbody>
</table>


  In addition to these functions, the SQL `OVERLAPS` operator is supported:

```

(START1, END1) OVERLAPS (START2, END2)
(START1, LENGTH1) OVERLAPS (START2, LENGTH2)
```
 This expression yields true when two time periods (defined by their endpoints) overlap, false when they do not overlap. The endpoints can be specified as pairs of dates, times, or time stamps; or as a date, time, or time stamp followed by an interval. When a pair of values is provided, either the start or the end can be written first; `OVERLAPS` automatically takes the earlier value of the pair as the start. Each time period is considered to represent the half-open interval *start* `<=` *time* `<` *end*, unless *start* and *end* are equal in which case it represents that single time instant. This means for instance that two time periods with only an endpoint in common do not overlap.


```

SELECT (DATE '2001-02-16', DATE '2001-12-21') OVERLAPS
       (DATE '2001-10-30', DATE '2002-10-30');
Result: true
SELECT (DATE '2001-02-16', INTERVAL '100 days') OVERLAPS
       (DATE '2001-10-30', DATE '2002-10-30');
Result: false
SELECT (DATE '2001-10-29', DATE '2001-10-30') OVERLAPS
       (DATE '2001-10-30', DATE '2001-10-31');
Result: false
SELECT (DATE '2001-10-30', DATE '2001-10-30') OVERLAPS
       (DATE '2001-10-30', DATE '2001-10-31');
Result: true
```


 When adding an `interval` value to (or subtracting an `interval` value from) a `timestamp` or `timestamp with time zone` value, the months, days, and microseconds fields of the `interval` value are handled in turn. First, a nonzero months field advances or decrements the date of the timestamp by the indicated number of months, keeping the day of month the same unless it would be past the end of the new month, in which case the last day of that month is used. (For example, March 31 plus 1 month becomes April 30, but March 31 plus 2 months becomes May 31.) Then the days field advances or decrements the date of the timestamp by the indicated number of days. In both these steps the local time of day is kept the same. Finally, if there is a nonzero microseconds field, it is added or subtracted literally. When doing arithmetic on a `timestamp with time zone` value in a time zone that recognizes DST, this means that adding or subtracting (say) `interval '1 day'` does not necessarily have the same result as adding or subtracting `interval '24 hours'`. For example, with the session time zone set to `America/Denver`:

```

SELECT timestamp with time zone '2005-04-02 12:00:00-07' + interval '1 day';
Result: 2005-04-03 12:00:00-06
SELECT timestamp with time zone '2005-04-02 12:00:00-07' + interval '24 hours';
Result: 2005-04-03 13:00:00-06
```
 This happens because an hour was skipped due to a change in daylight saving time at `2005-04-03 02:00:00` in time zone `America/Denver`.


 Note there can be ambiguity in the `months` field returned by `age` because different months have different numbers of days. PostgreSQL's approach uses the month from the earlier of the two dates when calculating partial months. For example, `age('2004-06-01', '2004-04-30')` uses April to yield `1 mon 1 day`, while using May would yield `1 mon 2 days` because May has 31 days, while April has only 30.


 Subtraction of dates and timestamps can also be complex. One conceptually simple way to perform subtraction is to convert each value to a number of seconds using `EXTRACT(EPOCH FROM ...)`, then subtract the results; this produces the number of *seconds* between the two values. This will adjust for the number of days in each month, timezone changes, and daylight saving time adjustments. Subtraction of date or timestamp values with the “`-`” operator returns the number of days (24-hours) and hours/minutes/seconds between the values, making the same adjustments. The `age` function returns years, months, days, and hours/minutes/seconds, performing field-by-field subtraction and then adjusting for negative field values. The following queries illustrate the differences in these approaches. The sample results were produced with `timezone = 'US/Eastern'`; there is a daylight saving time change between the two dates used:


```

SELECT EXTRACT(EPOCH FROM timestamptz '2013-07-01 12:00:00') -
       EXTRACT(EPOCH FROM timestamptz '2013-03-01 12:00:00');
Result: 10537200.000000
SELECT (EXTRACT(EPOCH FROM timestamptz '2013-07-01 12:00:00') -
        EXTRACT(EPOCH FROM timestamptz '2013-03-01 12:00:00'))
        / 60 / 60 / 24;
Result: 121.9583333333333333
SELECT timestamptz '2013-07-01 12:00:00' - timestamptz '2013-03-01 12:00:00';
Result: 121 days 23:00:00
SELECT age(timestamptz '2013-07-01 12:00:00', timestamptz '2013-03-01 12:00:00');
Result: 4 mons
```
 <a id="functions-datetime-extract"></a>

### `EXTRACT`, `date_part`


```

EXTRACT(FIELD FROM SOURCE)
```


 The `extract` function retrieves subfields such as year or hour from date/time values. *source* must be a value expression of type `timestamp`, `date`, `time`, or `interval`. (Timestamps and times can be with or without time zone.) *field* is an identifier or string that selects what field to extract from the source value. Not all fields are valid for every input data type; for example, fields smaller than a day cannot be extracted from a `date`, while fields of a day or more cannot be extracted from a `time`. The `extract` function returns values of type `numeric`.


 The following are valid field names:

`century`
:   The century; for `interval` values, the year field divided by 100


    ```

    SELECT EXTRACT(CENTURY FROM TIMESTAMP '2000-12-16 12:21:13');
    Result: 20
    SELECT EXTRACT(CENTURY FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 21
    SELECT EXTRACT(CENTURY FROM DATE '0001-01-01 AD');
    Result: 1
    SELECT EXTRACT(CENTURY FROM DATE '0001-12-31 BC');
    Result: -1
    SELECT EXTRACT(CENTURY FROM INTERVAL '2001 years');
    Result: 20
    ```

`day`
:   The day of the month (1–31); for `interval` values, the number of days


    ```

    SELECT EXTRACT(DAY FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 16
    SELECT EXTRACT(DAY FROM INTERVAL '40 days 1 minute');
    Result: 40
    ```

`decade`
:   The year field divided by 10


    ```

    SELECT EXTRACT(DECADE FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 200
    ```

`dow`
:   The day of the week as Sunday (`0`) to Saturday (`6`)


    ```

    SELECT EXTRACT(DOW FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 5
    ```


     Note that `extract`'s day of the week numbering differs from that of the `to_char(..., 'D')` function.

`doy`
:   The day of the year (1–365/366)


    ```

    SELECT EXTRACT(DOY FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 47
    ```

`epoch`
:   For `timestamp with time zone` values, the number of seconds since 1970-01-01 00:00:00 UTC (negative for timestamps before that); for `date` and `timestamp` values, the nominal number of seconds since 1970-01-01 00:00:00, without regard to timezone or daylight-savings rules; for `interval` values, the total number of seconds in the interval


    ```

    SELECT EXTRACT(EPOCH FROM TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40.12-08');
    Result: 982384720.120000
    SELECT EXTRACT(EPOCH FROM TIMESTAMP '2001-02-16 20:38:40.12');
    Result: 982355920.120000
    SELECT EXTRACT(EPOCH FROM INTERVAL '5 days 3 hours');
    Result: 442800.000000
    ```


     You can convert an epoch value back to a `timestamp with time zone` with `to_timestamp`:


    ```

    SELECT to_timestamp(982384720.12);
    Result: 2001-02-17 04:38:40.12+00
    ```


     Beware that applying `to_timestamp` to an epoch extracted from a `date` or `timestamp` value could produce a misleading result: the result will effectively assume that the original value had been given in UTC, which might not be the case.

`hour`
:   The hour field (0–23 in timestamps, unrestricted in intervals)


    ```

    SELECT EXTRACT(HOUR FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 20
    ```

`isodow`
:   The day of the week as Monday (`1`) to Sunday (`7`)


    ```

    SELECT EXTRACT(ISODOW FROM TIMESTAMP '2001-02-18 20:38:40');
    Result: 7
    ```


     This is identical to `dow` except for Sunday. This matches the ISO 8601 day of the week numbering.

`isoyear`
:   The ISO 8601 week-numbering year that the date falls in


    ```

    SELECT EXTRACT(ISOYEAR FROM DATE '2006-01-01');
    Result: 2005
    SELECT EXTRACT(ISOYEAR FROM DATE '2006-01-02');
    Result: 2006
    ```


     Each ISO 8601 week-numbering year begins with the Monday of the week containing the 4th of January, so in early January or late December the ISO year may be different from the Gregorian year. See the `week` field for more information.

`julian`
:   The *Julian Date* corresponding to the date or timestamp. Timestamps that are not local midnight result in a fractional value. See [Julian Dates](../../appendixes/date-time-support/julian-dates.md#datetime-julian-dates) for more information.


    ```

    SELECT EXTRACT(JULIAN FROM DATE '2006-01-01');
    Result: 2453737
    SELECT EXTRACT(JULIAN FROM TIMESTAMP '2006-01-01 12:00');
    Result: 2453737.50000000000000000000
    ```

`microseconds`
:   The seconds field, including fractional parts, multiplied by 1 000 000; note that this includes full seconds


    ```

    SELECT EXTRACT(MICROSECONDS FROM TIME '17:12:28.5');
    Result: 28500000
    ```

`millennium`
:   The millennium; for `interval` values, the year field divided by 1000


    ```

    SELECT EXTRACT(MILLENNIUM FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 3
    SELECT EXTRACT(MILLENNIUM FROM INTERVAL '2001 years');
    Result: 2
    ```


     Years in the 1900s are in the second millennium. The third millennium started January 1, 2001.

`milliseconds`
:   The seconds field, including fractional parts, multiplied by 1000. Note that this includes full seconds.


    ```

    SELECT EXTRACT(MILLISECONDS FROM TIME '17:12:28.5');
    Result: 28500.000
    ```

`minute`
:   The minutes field (0–59)


    ```

    SELECT EXTRACT(MINUTE FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 38
    ```

`month`
:   The number of the month within the year (1–12); for `interval` values, the number of months modulo 12 (0–11)


    ```

    SELECT EXTRACT(MONTH FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 2
    SELECT EXTRACT(MONTH FROM INTERVAL '2 years 3 months');
    Result: 3
    SELECT EXTRACT(MONTH FROM INTERVAL '2 years 13 months');
    Result: 1
    ```

`quarter`
:   The quarter of the year (1–4) that the date is in; for `interval` values, the month field divided by 3 plus 1


    ```

    SELECT EXTRACT(QUARTER FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 1
    SELECT EXTRACT(QUARTER FROM INTERVAL '1 year 6 months');
    Result: 3
    ```

`second`
:   The seconds field, including any fractional seconds


    ```

    SELECT EXTRACT(SECOND FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 40.000000
    SELECT EXTRACT(SECOND FROM TIME '17:12:28.5');
    Result: 28.500000
    ```

`timezone`
:   The time zone offset from UTC, measured in seconds. Positive values correspond to time zones east of UTC, negative values to zones west of UTC. (Technically, PostgreSQL does not use UTC because leap seconds are not handled.)

`timezone_hour`
:   The hour component of the time zone offset

`timezone_minute`
:   The minute component of the time zone offset

`week`
:   The number of the ISO 8601 week-numbering week of the year. By definition, ISO weeks start on Mondays and the first week of a year contains January 4 of that year. In other words, the first Thursday of a year is in week 1 of that year.


     In the ISO week-numbering system, it is possible for early-January dates to be part of the 52nd or 53rd week of the previous year, and for late-December dates to be part of the first week of the next year. For example, `2005-01-01` is part of the 53rd week of year 2004, and `2006-01-01` is part of the 52nd week of year 2005, while `2012-12-31` is part of the first week of 2013. It's recommended to use the `isoyear` field together with `week` to get consistent results.


     For `interval` values, the week field is simply the number of integral days divided by 7.


    ```

    SELECT EXTRACT(WEEK FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 7
    SELECT EXTRACT(WEEK FROM INTERVAL '13 days 24 hours');
    Result: 1
    ```

`year`
:   The year field. Keep in mind there is no `0 AD`, so subtracting `BC` years from `AD` years should be done with care.


    ```

    SELECT EXTRACT(YEAR FROM TIMESTAMP '2001-02-16 20:38:40');
    Result: 2001
    ```


 When processing an `interval` value, the `extract` function produces field values that match the interpretation used by the interval output function. This can produce surprising results if one starts with a non-normalized interval representation, for example:

```

SELECT INTERVAL '80 minutes';
Result: 01:20:00
SELECT EXTRACT(MINUTES FROM INTERVAL '80 minutes');
Result: 20
```


!!! note

    When the input value is +/-Infinity, `extract` returns +/-Infinity for monotonically-increasing fields (`epoch`, `julian`, `year`, `isoyear`, `decade`, `century`, and `millennium` for `timestamp` inputs; `epoch`, `hour`, `day`, `year`, `decade`, `century`, and `millennium` for `interval` inputs). For other fields, NULL is returned. PostgreSQL versions before 9.6 returned zero for all cases of infinite input.


 The `extract` function is primarily intended for computational processing. For formatting date/time values for display, see [Data Type Formatting Functions](data-type-formatting-functions.md#functions-formatting).


 The `date_part` function is modeled on the traditional Ingres equivalent to the SQL-standard function `extract`:

```

date_part('FIELD', SOURCE)
```
 Note that here the *field* parameter needs to be a string value, not a name. The valid field names for `date_part` are the same as for `extract`. For historical reasons, the `date_part` function returns values of type `double precision`. This can result in a loss of precision in certain uses. Using `extract` is recommended instead.


```

SELECT date_part('day', TIMESTAMP '2001-02-16 20:38:40');
Result: 16
SELECT date_part('hour', INTERVAL '4 hours 3 minutes');
Result: 4
```
  <a id="functions-datetime-trunc"></a>

### `date_trunc`


 The function `date_trunc` is conceptually similar to the `trunc` function for numbers.


```

date_trunc(FIELD, SOURCE [, TIME_ZONE ])
```
 *source* is a value expression of type `timestamp`, `timestamp with time zone`, or `interval`. (Values of type `date` and `time` are cast automatically to `timestamp` or `interval`, respectively.) *field* selects to which precision to truncate the input value. The return value is likewise of type `timestamp`, `timestamp with time zone`, or `interval`, and it has all fields that are less significant than the selected one set to zero (or one, for day and month).


 Valid values for *field* are:

- `microseconds`
- `milliseconds`
- `second`
- `minute`
- `hour`
- `day`
- `week`
- `month`
- `quarter`
- `year`
- `decade`
- `century`
- `millennium`


 When the input value is of type `timestamp with time zone`, the truncation is performed with respect to a particular time zone; for example, truncation to `day` produces a value that is midnight in that zone. By default, truncation is done with respect to the current [TimeZone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) setting, but the optional *time_zone* argument can be provided to specify a different time zone. The time zone name can be specified in any of the ways described in [Time Zones](../data-types/date-time-types.md#datatype-timezones).


 A time zone cannot be specified when processing `timestamp without time zone` or `interval` inputs. These are always taken at face value.


 Examples (assuming the local time zone is `America/New_York`):

```

SELECT date_trunc('hour', TIMESTAMP '2001-02-16 20:38:40');
Result: 2001-02-16 20:00:00
SELECT date_trunc('year', TIMESTAMP '2001-02-16 20:38:40');
Result: 2001-01-01 00:00:00
SELECT date_trunc('day', TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40+00');
Result: 2001-02-16 00:00:00-05
SELECT date_trunc('day', TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40+00', 'Australia/Sydney');
Result: 2001-02-16 08:00:00-05
SELECT date_trunc('hour', INTERVAL '3 days 02:47:33');
Result: 3 days 02:00:00
```

  <a id="functions-datetime-bin"></a>

### `date_bin`


 The function `date_bin` “bins” the input timestamp into the specified interval (the *stride*) aligned with a specified origin.


```

date_bin(STRIDE, SOURCE, ORIGIN)
```
 *source* is a value expression of type `timestamp` or `timestamp with time zone`. (Values of type `date` are cast automatically to `timestamp`.) *stride* is a value expression of type `interval`. The return value is likewise of type `timestamp` or `timestamp with time zone`, and it marks the beginning of the bin into which the *source* is placed.


 Examples:

```

SELECT date_bin('15 minutes', TIMESTAMP '2020-02-11 15:44:17', TIMESTAMP '2001-01-01');
Result: 2020-02-11 15:30:00
SELECT date_bin('15 minutes', TIMESTAMP '2020-02-11 15:44:17', TIMESTAMP '2001-01-01 00:02:30');
Result: 2020-02-11 15:32:30
```


 In the case of full units (1 minute, 1 hour, etc.), it gives the same result as the analogous `date_trunc` call, but the difference is that `date_bin` can truncate to an arbitrary interval.


 The `stride` interval must be greater than zero and cannot contain units of month or larger.
  <a id="functions-datetime-zoneconvert"></a>

### `AT TIME ZONE` and `AT LOCAL`


 The `AT TIME ZONE` operator converts time stamp *without* time zone to/from time stamp *with* time zone, and `time with time zone` values to different time zones. [`AT TIME ZONE` and `AT LOCAL` Variants](#functions-datetime-zoneconvert-table) shows its variants.
 <a id="functions-datetime-zoneconvert-table"></a>

**Table: `AT TIME ZONE` and `AT LOCAL` Variants**

<table>
<thead>
<tr>
<th>Operator</th>
<th>Description</th>
<th>Example(s)</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>timestamp without time zone</code> <code>AT TIME ZONE</code> <em>zone</em> <code>timestamp with time zone</code></td>
<td>Converts given time stamp <em>without</em> time zone to time stamp <em>with</em> time zone, assuming the given value is in the named time zone.</td>
<td><code>timestamp '2001-02-16 20:38:40' AT TIME ZONE 'America/Denver'</code> <code>2001-02-17 03:38:40+00</code></td>
</tr>
<tr>
<td><code>timestamp without time zone</code> <code>AT LOCAL</code> <code>timestamp with time zone</code></td>
<td>Converts given time stamp <em>without</em> time zone to time stamp <em>with</em> the session's <code>TimeZone</code> value as time zone.</td>
<td><code>timestamp '2001-02-16 20:38:40' at local</code> <code>2001-02-17 03:38:40+00</code></td>
</tr>
<tr>
<td><code>timestamp with time zone</code> <code>AT TIME ZONE</code> <em>zone</em> <code>timestamp without time zone</code></td>
<td>Converts given time stamp <em>with</em> time zone to time stamp <em>without</em> time zone, as the time would appear in that zone.</td>
<td><code>timestamp with time zone '2001-02-16 20:38:40-05' AT TIME ZONE 'America/Denver'</code> <code>2001-02-16 18:38:40</code></td>
</tr>
<tr>
<td><code>timestamp with time zone</code> <code>AT LOCAL</code> <code>timestamp without time zone</code></td>
<td>Converts given time stamp <em>with</em> time zone to time stamp <em>without</em> time zone, as the time would appear with the session's <code>TimeZone</code> value as time zone.</td>
<td><code>timestamp with time zone '2001-02-16 20:38:40-05' at local</code> <code>2001-02-16 18:38:40</code></td>
</tr>
<tr>
<td><code>time with time zone</code> <code>AT TIME ZONE</code> <em>zone</em> <code>time with time zone</code></td>
<td>Converts given time <em>with</em> time zone to a new time zone. Since no date is supplied, this uses the currently active UTC offset for the named destination zone.</td>
<td><code>time with time zone '05:34:17-05' AT TIME ZONE 'UTC'</code> <code>10:34:17+00</code></td>
</tr>
<tr>
<td><code>time with time zone</code> <code>AT LOCAL</code> <code>time with time zone</code></td>
<td>Converts given time <em>with</em> time zone to a new time zone. Since no date is supplied, this uses the currently active UTC offset for the session's <code>TimeZone</code> value.</td>
<td>Assuming the session's <code>TimeZone</code> is set to <code>UTC</code>:<br><code>time with time zone '05:34:17-05' at local</code> <code>10:34:17+00</code></td>
</tr>
</tbody>
</table>


 In these expressions, the desired time zone *zone* can be specified either as a text value (e.g., `'America/Los_Angeles'`) or as an interval (e.g., `INTERVAL '-08:00'`). In the text case, a time zone name can be specified in any of the ways described in [Time Zones](../data-types/date-time-types.md#datatype-timezones). The interval case is only useful for zones that have fixed offsets from UTC, so it is not very common in practice.


 The syntax `AT LOCAL` may be used as shorthand for <code>AT TIME ZONE </code><em>local</em>, where *local* is the session's `TimeZone` value.


 Examples (assuming the current [TimeZone](../../server-administration/server-configuration/client-connection-defaults.md#guc-timezone) setting is `America/Los_Angeles`):

```

SELECT TIMESTAMP '2001-02-16 20:38:40' AT TIME ZONE 'America/Denver';
Result: 2001-02-16 19:38:40-08
SELECT TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40-05' AT TIME ZONE 'America/Denver';
Result: 2001-02-16 18:38:40
SELECT TIMESTAMP '2001-02-16 20:38:40' AT TIME ZONE 'Asia/Tokyo' AT TIME ZONE 'America/Chicago';
Result: 2001-02-16 05:38:40
SELECT TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40-05' AT LOCAL;
Result: 2001-02-16 17:38:40
SELECT TIMESTAMP WITH TIME ZONE '2001-02-16 20:38:40-05' AT TIME ZONE '+05';
Result: 2001-02-16 20:38:40
SELECT TIME WITH TIME ZONE '20:38:40-05' AT LOCAL;
Result: 17:38:40
```
 The first example adds a time zone to a value that lacks it, and displays the value using the current `TimeZone` setting. The second example shifts the time stamp with time zone value to the specified time zone, and returns the value without a time zone. This allows storage and display of values different from the current `TimeZone` setting. The third example converts Tokyo time to Chicago time. The fourth example shifts the time stamp with time zone value to the time zone currently specified by the `TimeZone` setting and returns the value without a time zone. The fifth example demonstrates that the sign in a POSIX-style time zone specification has the opposite meaning of the sign in an ISO-8601 datetime literal, as described in [Time Zones](../data-types/date-time-types.md#datatype-timezones) and [Date/Time Support](../../appendixes/date-time-support/index.md#datetime-appendix).


 The sixth example is a cautionary tale. Due to the fact that there is no date associated with the input value, the conversion is made using the current date of the session. Therefore, this static example may show a wrong result depending on the time of the year it is viewed because `'America/Los_Angeles'` observes Daylight Savings Time.


 The function <code>`timezone`(</code><em>zone</em><code>,
    </code><em>timestamp</em><code>)</code> is equivalent to the SQL-conforming construct <em>timestamp</em><code> AT TIME ZONE
    </code><em>zone</em>.


 The function <code>`timezone`(</code><em>zone</em><code>,
    </code><em>time</em><code>)</code> is equivalent to the SQL-conforming construct <em>time</em><code> AT TIME ZONE
    </code><em>zone</em>.


 The function <code>`timezone`(</code><em>timestamp</em><code>)</code> is equivalent to the SQL-conforming construct <em>timestamp</em><code>
    AT LOCAL</code>.


 The function <code>`timezone`(</code><em>time</em><code>)</code> is equivalent to the SQL-conforming construct <em>time</em><code>
    AT LOCAL</code>.
  <a id="functions-datetime-current"></a>

### Current Date/Time


 PostgreSQL provides a number of functions that return values related to the current date and time. These SQL-standard functions all return values based on the start time of the current transaction:

```

CURRENT_DATE
CURRENT_TIME
CURRENT_TIMESTAMP
CURRENT_TIME(PRECISION)
CURRENT_TIMESTAMP(PRECISION)
LOCALTIME
LOCALTIMESTAMP
LOCALTIME(PRECISION)
LOCALTIMESTAMP(PRECISION)
```


 `CURRENT_TIME` and `CURRENT_TIMESTAMP` deliver values with time zone; `LOCALTIME` and `LOCALTIMESTAMP` deliver values without time zone.


 `CURRENT_TIME`, `CURRENT_TIMESTAMP`, `LOCALTIME`, and `LOCALTIMESTAMP` can optionally take a precision parameter, which causes the result to be rounded to that many fractional digits in the seconds field. Without a precision parameter, the result is given to the full available precision.


 Some examples:

```

SELECT CURRENT_TIME;
Result: 14:39:53.662522-05
SELECT CURRENT_DATE;
Result: 2019-12-23
SELECT CURRENT_TIMESTAMP;
Result: 2019-12-23 14:39:53.662522-05
SELECT CURRENT_TIMESTAMP(2);
Result: 2019-12-23 14:39:53.66-05
SELECT LOCALTIMESTAMP;
Result: 2019-12-23 14:39:53.662522
```


 Since these functions return the start time of the current transaction, their values do not change during the transaction. This is considered a feature: the intent is to allow a single transaction to have a consistent notion of the “current” time, so that multiple modifications within the same transaction bear the same time stamp.


!!! note

    Other database systems might advance these values more frequently.


 PostgreSQL also provides functions that return the start time of the current statement, as well as the actual current time at the instant the function is called. The complete list of non-SQL-standard time functions is:

```

transaction_timestamp()
statement_timestamp()
clock_timestamp()
timeofday()
now()
```


 `transaction_timestamp()` is equivalent to `CURRENT_TIMESTAMP`, but is named to clearly reflect what it returns. `statement_timestamp()` returns the start time of the current statement (more specifically, the time of receipt of the latest command message from the client). `statement_timestamp()` and `transaction_timestamp()` return the same value during the first statement of a transaction, but might differ during subsequent statements. `clock_timestamp()` returns the actual current time, and therefore its value changes even within a single SQL statement. `timeofday()` is a historical PostgreSQL function. Like `clock_timestamp()`, it returns the actual current time, but as a formatted `text` string rather than a `timestamp with time zone` value. `now()` is a traditional PostgreSQL equivalent to `transaction_timestamp()`.


 All the date/time data types also accept the special literal value `now` to specify the current date and time (again, interpreted as the transaction start time). Thus, the following three all return the same result:

```sql

SELECT CURRENT_TIMESTAMP;
SELECT now();
SELECT TIMESTAMP 'now';  -- but see tip below
```


!!! tip

    Do not use the third form when specifying a value to be evaluated later, for example in a `DEFAULT` clause for a table column. The system will convert `now` to a `timestamp` as soon as the constant is parsed, so that when the default value is needed, the time of the table creation would be used! The first two forms will not be evaluated until the default value is used, because they are function calls. Thus they will give the desired behavior of defaulting to the time of row insertion. (See also [Special Values](../data-types/date-time-types.md#datatype-datetime-special-values).)
  <a id="functions-datetime-delay"></a>

### Delaying Execution


 The following functions are available to delay execution of the server process:

```

pg_sleep ( double precision )
pg_sleep_for ( interval )
pg_sleep_until ( timestamp with time zone )
```
 `pg_sleep` makes the current session's process sleep until the given number of seconds have elapsed. Fractional-second delays can be specified. `pg_sleep_for` is a convenience function to allow the sleep time to be specified as an `interval`. `pg_sleep_until` is a convenience function for when a specific wake-up time is desired. For example:

```sql

SELECT pg_sleep(1.5);
SELECT pg_sleep_for('5 minutes');
SELECT pg_sleep_until('tomorrow 03:00');
```


!!! note

    The effective resolution of the sleep interval is platform-specific; 0.01 seconds is a common value. The sleep delay will be at least as long as specified. It might be longer depending on factors such as server load. In particular, `pg_sleep_until` is not guaranteed to wake up exactly at the specified time, but it will not wake up any earlier.


!!! warning

    Make sure that your session does not hold more locks than necessary when calling `pg_sleep` or its variants. Otherwise other sessions might have to wait for your sleeping process, slowing down the entire system.
