# pg_test_timing { #pgtesttiming }

measure timing overhead

## Synopsis


```
pg_test_timing [OPTION...]
```


## Description


 pg_test_timing is a tool to measure the timing overhead on your system and confirm that the system time never moves backwards. It simply reads the system clock over and over again as fast as it can for a specified length of time, and then prints statistics about the observed differences in successive clock readings.


 Smaller (but not zero) differences are better, since they imply both more-precise clock hardware and less overhead to collect a clock reading. Systems that are slow to collect timing data can give less accurate `EXPLAIN ANALYZE` results.


 This tool is also helpful to determine if the `track_io_timing` configuration parameter is likely to produce useful results.


## Options


 pg_test_timing accepts the following command-line options:

<code>-d </code><em>duration</em>, <code>--duration=</code><em>duration</em>
:   Specifies the test duration, in seconds. Longer durations give slightly better accuracy, and are more likely to discover problems with the system clock moving backwards. The default test duration is 3 seconds.

<code>-c </code><em>cutoff</em>, <code>--cutoff=</code><em>cutoff</em>
:   Specifies the cutoff percentage for the list of exact observed timing durations (that is, the changes in the system clock value from one reading to the next). The list will end once the running percentage total reaches or exceeds this value, except that the largest observed duration will always be printed. The default cutoff is 99.99.

`-V`, `--version`
:   Print the pg_test_timing version and exit.

`-?`, `--help`
:   Show help about pg_test_timing command line arguments, and exit.


## Usage


### Interpreting Results


 The first block of output has four columns, with rows showing a shifted-by-one log2(ns) histogram of timing durations (that is, the differences between successive clock readings). This is not the classic log2(n+1) histogram as it counts zeros separately and then switches to log2(ns) starting from value 1.


 The columns are:

-

  nanosecond value that is >= the durations in this bucket
-

  percentage of durations in this bucket
-

  running-sum percentage of durations in this and previous buckets
-

  count of durations in this bucket


 The second block of output goes into more detail, showing the exact timing differences observed. For brevity this list is cut off when the running-sum percentage exceeds the user-selectable cutoff value. However, the largest observed difference is always shown.


 The example results below show that 99.99% of timing loops took between 8 and 31 nanoseconds, with the worst case somewhere between 32768 and 65535 nanoseconds. In the second block, we can see that typical loop time is 16 nanoseconds, and the readings appear to have full nanosecond precision.


```

Testing timing overhead for 3 seconds.
Average loop time including overhead: 16.40 ns
Histogram of timing durations:
   <= ns   % of total  running %      count
       0       0.0000     0.0000          0
       1       0.0000     0.0000          0
       3       0.0000     0.0000          0
       7       0.0000     0.0000          0
      15       4.5452     4.5452    8313178
      31      95.4527    99.9979  174581501
      63       0.0001    99.9981        253
     127       0.0001    99.9982        165
     255       0.0000    99.9982         35
     511       0.0000    99.9982          1
    1023       0.0013    99.9994       2300
    2047       0.0004    99.9998        690
    4095       0.0000    99.9998          9
    8191       0.0000    99.9998          8
   16383       0.0002   100.0000        337
   32767       0.0000   100.0000          2
   65535       0.0000   100.0000          1

Observed timing durations up to 99.9900%:
      ns   % of total  running %      count
      15       4.5452     4.5452    8313178
      16      58.3785    62.9237  106773354
      17      33.6840    96.6078   61607584
      18       3.1151    99.7229    5697480
      19       0.2638    99.9867     482570
      20       0.0093    99.9960      17054
...
   38051       0.0000   100.0000          1
```


## See Also
  [sql-explain](../sql-commands/explain.md#sql-explain), [Wiki discussion about timing](https://wiki.postgresql.org/wiki/Pg_test_timing)
