<a id="RasterBand_Stats"></a>

## Raster Band Statistics and Analytics
  <a id="RT_ST_Count"></a>

# ST_Count

Returns the number of pixels in a given band of a raster or raster coverage. If no band is specified defaults to band 1. If exclude_nodata_value is set to true, will only count pixels that are not equal to the nodata value.

## Synopsis


```sql
bigint ST_Count(raster  rast, integer  nband=1, boolean  exclude_nodata_value=true)
bigint ST_Count(raster  rast, boolean  exclude_nodata_value)
```


## Description


Returns the number of pixels in a given band of a raster or raster coverage. If no band is specified `nband` defaults to 1.


!!! note

    If `exclude_nodata_value` is set to true, will only count pixels with value not equal to the `nodata` value of the raster. Set `exclude_nodata_value` to false to get count all pixels


Changed: 3.1.0 - The ST_Count(rastertable, rastercolumn, ...) variants removed. Use [RT_ST_CountAgg](#RT_ST_CountAgg) instead.


Availability: 2.0.0


## Examples


```

--example will count all pixels not 249 and one will count all pixels.  --
SELECT rid, ST_Count(ST_SetBandNoDataValue(rast,249)) As exclude_nodata,
        ST_Count(ST_SetBandNoDataValue(rast,249),false) As include_nodata
    FROM dummy_rast WHERE rid=2;

rid | exclude_nodata | include_nodata
-----+----------------+----------------
   2 |             23 |             25

```


## See Also


 [RT_ST_CountAgg](#RT_ST_CountAgg), [RT_ST_SummaryStats](#RT_ST_SummaryStats), [RT_ST_SetBandNoDataValue](raster-band-editors.md#RT_ST_SetBandNoDataValue)
  <a id="RT_ST_CountAgg"></a>

# ST_CountAgg

Aggregate. Returns the number of pixels in a given band of a set of rasters. If no band is specified defaults to band 1. If exclude_nodata_value is set to true, will only count pixels that are not equal to the NODATA value.

## Synopsis


```sql
bigint ST_CountAgg(raster  rast, integer  nband, boolean  exclude_nodata_value, double precision  sample_percent)
bigint ST_CountAgg(raster  rast, integer  nband, boolean  exclude_nodata_value)
bigint ST_CountAgg(raster  rast, boolean  exclude_nodata_value)
```


## Description


Returns the number of pixels in a given band of a set of rasters. If no band is specified `nband` defaults to 1.


 If `exclude_nodata_value` is set to true, will only count pixels with value not equal to the `NODATA` value of the raster. Set `exclude_nodata_value` to false to get count all pixels


By default will sample all pixels. To get faster response, set `sample_percent` to value between zero (0) and one (1)


Availability: 2.2.0


## Examples


```sql

WITH foo AS (
    SELECT
        rast.rast
    FROM (
        SELECT ST_SetValue(
            ST_SetValue(
                ST_SetValue(
                    ST_AddBand(
                        ST_MakeEmptyRaster(10, 10, 10, 10, 2, 2, 0, 0,0)
                        , 1, '64BF', 0, 0
                    )
                    , 1, 1, 1, -10
                )
                , 1, 5, 4, 0
            )
            , 1, 5, 5, 3.14159
        ) AS rast
    ) AS rast
    FULL JOIN (
        SELECT generate_series(1, 10) AS id
    ) AS id
        ON 1 = 1
)
SELECT
    ST_CountAgg(rast, 1, TRUE)
FROM foo;

 st_countagg
-------------
          20
(1 row)

```


## See Also


 [RT_ST_Count](#RT_ST_Count), [RT_ST_SummaryStats](#RT_ST_SummaryStats), [RT_ST_SetBandNoDataValue](raster-band-editors.md#RT_ST_SetBandNoDataValue)
  <a id="RT_ST_Histogram"></a>

# ST_Histogram

Returns a set of record summarizing a raster or raster coverage data distribution separate bin ranges. Number of bins are autocomputed if not specified.

## Synopsis


```sql
SETOF record ST_Histogram(raster  rast, integer  nband=1, boolean  exclude_nodata_value=true, integer  bins=autocomputed, double precision[]  width=NULL, boolean  right=false)
SETOF record ST_Histogram(raster  rast, integer  nband, integer  bins, double precision[]  width=NULL, boolean  right=false)
SETOF record ST_Histogram(raster  rast, integer  nband, boolean  exclude_nodata_value, integer  bins, boolean  right)
SETOF record ST_Histogram(raster  rast, integer  nband, integer  bins, boolean  right)
```


## Description


Returns set of records consisting of min, max, count, percent for a given raster band for each bin. If no band is specified `nband` defaults to 1.


!!! note

    By default only considers pixel values not equal to the `nodata` value . Set `exclude_nodata_value` to false to get count all pixels.


`width`
:   width: an array indicating the width of each category/bin. If the number of bins is greater than the number of widths, the widths are repeated.


    Example: 9 bins, widths are [a, b, c] will have the output be [a, b, c, a, b, c, a, b, c]

`bins`
:   Number of breakouts -- this is the number of records you'll get back from the function if specified. If not specified then the number of breakouts is autocomputed.

`right`
:   compute the histogram from the right rather than from the left (default). This changes the criteria for evaluating a value x from [a, b) to (a, b]


Changed: 3.1.0 Removed ST_Histogram(table_name, column_name) variant.


Availability: 2.0.0


## Example: Single raster tile - compute histograms for bands 1, 2, 3 and autocompute bins


```sql
SELECT band, (stats).*
FROM (SELECT rid, band, ST_Histogram(rast, band) As stats
    FROM dummy_rast CROSS JOIN generate_series(1,3) As band
     WHERE rid=2) As foo;

 band |  min  |  max  | count | percent
------+-------+-------+-------+---------
    1 |   249 |   250 |     2 |    0.08
    1 |   250 |   251 |     2 |    0.08
    1 |   251 |   252 |     1 |    0.04
    1 |   252 |   253 |     2 |    0.08
    1 |   253 |   254 |    18 |    0.72
    2 |    78 | 113.2 |    11 |    0.44
    2 | 113.2 | 148.4 |     4 |    0.16
    2 | 148.4 | 183.6 |     4 |    0.16
    2 | 183.6 | 218.8 |     1 |    0.04
    2 | 218.8 |   254 |     5 |     0.2
    3 |    62 | 100.4 |    11 |    0.44
    3 | 100.4 | 138.8 |     5 |     0.2
    3 | 138.8 | 177.2 |     4 |    0.16
    3 | 177.2 | 215.6 |     1 |    0.04
    3 | 215.6 |   254 |     4 |    0.16
```


## Example: Just band 2 but for 6 bins


```sql
SELECT (stats).*
FROM (SELECT rid, ST_Histogram(rast, 2,6) As stats
    FROM dummy_rast
     WHERE rid=2) As foo;

    min     |    max     | count | percent
------------+------------+-------+---------
         78 | 107.333333 |     9 |    0.36
 107.333333 | 136.666667 |     6 |    0.24
 136.666667 |        166 |     0 |       0
        166 | 195.333333 |     4 |    0.16
 195.333333 | 224.666667 |     1 |    0.04
 224.666667 |        254 |     5 |     0.2
(6 rows)

-- Same as previous but we explicitly control the pixel value range of each bin.
SELECT (stats).*
FROM (SELECT rid, ST_Histogram(rast, 2,6,ARRAY[0.5,1,4,100,5]) As stats
    FROM dummy_rast
     WHERE rid=2) As foo;

  min  |  max  | count | percent
-------+-------+-------+----------
    78 |  78.5 |     1 |     0.08
  78.5 |  79.5 |     1 |     0.04
  79.5 |  83.5 |     0 |        0
  83.5 | 183.5 |    17 |   0.0068
 183.5 | 188.5 |     0 |        0
 188.5 |   254 |     6 | 0.003664
(6 rows)
```


## See Also


 [RT_ST_Count](#RT_ST_Count), [RT_ST_SummaryStats](#RT_ST_SummaryStats), [RT_ST_SummaryStatsAgg](#RT_ST_SummaryStatsAgg)
  <a id="RT_ST_Quantile"></a>

# ST_Quantile

Compute quantiles for a raster or raster table coverage in the context of the sample or population. Thus, a value could be examined to be at the raster's 25%, 50%, 75% percentile.

## Synopsis


```sql
SETOF record ST_Quantile(raster  rast, integer  nband=1, boolean  exclude_nodata_value=true, double precision[]  quantiles=NULL)
SETOF record ST_Quantile(raster  rast, double precision[]  quantiles)
SETOF record ST_Quantile(raster  rast, integer  nband, double precision[]  quantiles)
double precision ST_Quantile(raster  rast, double precision  quantile)
double precision ST_Quantile(raster  rast, boolean  exclude_nodata_value, double precision  quantile=NULL)
double precision ST_Quantile(raster  rast, integer  nband, double precision  quantile)
double precision ST_Quantile(raster  rast, integer  nband, boolean  exclude_nodata_value, double precision  quantile)
double precision ST_Quantile(raster  rast, integer  nband, double precision  quantile)
```


## Description


Compute quantiles for a raster or raster table coverage in the context of the sample or population. Thus, a value could be examined to be at the raster's 25%, 50%, 75% percentile.


!!! note

    If `exclude_nodata_value` is set to false, will also count pixels with no data.


Changed: 3.1.0 Removed ST_Quantile(table_name, column_name) variant.


Availability: 2.0.0


## Examples


```sql

UPDATE dummy_rast SET rast = ST_SetBandNoDataValue(rast,249) WHERE rid=2;
--Example will consider only pixels of band 1 that are not 249 and in named quantiles --

SELECT (pvq).*
FROM (SELECT ST_Quantile(rast, ARRAY[0.25,0.75]) As pvq
    FROM dummy_rast WHERE rid=2) As foo
    ORDER BY (pvq).quantile;

 quantile | value
----------+-------
     0.25 |   253
     0.75 |   254

SELECT ST_Quantile(rast, 0.75) As value
    FROM dummy_rast WHERE rid=2;

value
------
  254
```


```

--real live example.  Quantile of all pixels in band 2 intersecting a geometry
SELECT rid, (ST_Quantile(rast,2)).* As pvc
    FROM o_4_boston
        WHERE ST_Intersects(rast,
            ST_GeomFromText('POLYGON((224486 892151,224486 892200,224706 892200,224706 892151,224486 892151))',26986)
            )
ORDER BY value, quantile,rid
;


 rid | quantile | value
-----+----------+-------
   1 |        0 |     0
   2 |        0 |     0
  14 |        0 |     1
  15 |        0 |     2
  14 |     0.25 |    37
   1 |     0.25 |    42
  15 |     0.25 |    47
   2 |     0.25 |    50
  14 |      0.5 |    56
   1 |      0.5 |    64
  15 |      0.5 |    66
   2 |      0.5 |    77
  14 |     0.75 |    81
  15 |     0.75 |    87
   1 |     0.75 |    94
   2 |     0.75 |   106
  14 |        1 |   199
   1 |        1 |   244
   2 |        1 |   255
  15 |        1 |   255
```


## See Also


 [RT_ST_Count](#RT_ST_Count), [RT_ST_SummaryStats](#RT_ST_SummaryStats), [RT_ST_SummaryStatsAgg](#RT_ST_SummaryStatsAgg), [RT_ST_SetBandNoDataValue](raster-band-editors.md#RT_ST_SetBandNoDataValue)
  <a id="RT_ST_SummaryStats"></a>

# ST_SummaryStats

Returns summarystats consisting of count, sum, mean, stddev, min, max for a given raster band of a raster or raster coverage. Band 1 is assumed is no band is specified.

## Synopsis


```sql
summarystats ST_SummaryStats(raster  rast, boolean  exclude_nodata_value)
summarystats ST_SummaryStats(raster  rast, integer  nband, boolean  exclude_nodata_value)
```


## Description


Returns [summarystats](raster-support-data-types.md#summarystats) consisting of count, sum, mean, stddev, min, max for a given raster band of a raster or raster coverage. If no band is specified `nband` defaults to 1.


!!! note

    By default only considers pixel values not equal to the `nodata` value. Set `exclude_nodata_value` to false to get count of all pixels.


!!! note

    By default will sample all pixels. To get faster response, set `sample_percent` to lower than 1


Changed: 3.1.0 ST_SummaryStats(rastertable, rastercolumn, ...) variants are removed. Use [RT_ST_SummaryStatsAgg](#RT_ST_SummaryStatsAgg) instead.


Availability: 2.0.0


## Example: Single raster tile


```sql

SELECT rid, band, (stats).*
FROM (SELECT rid, band, ST_SummaryStats(rast, band) As stats
    FROM dummy_rast CROSS JOIN generate_series(1,3) As band
     WHERE rid=2) As foo;

 rid | band | count | sum  |    mean    |  stddev   | min | max
-----+------+-------+------+------------+-----------+-----+-----
   2 |    1 |    23 | 5821 | 253.086957 |  1.248061 | 250 | 254
   2 |    2 |    25 | 3682 |     147.28 | 59.862188 |  78 | 254
   2 |    3 |    25 | 3290 |      131.6 | 61.647384 |  62 | 254

```


## Example: Summarize pixels that intersect buildings of interest


This example took 574ms on PostGIS windows 64-bit with all of Boston Buildings and aerial Tiles (tiles each 150x150 pixels ~ 134,000 tiles), ~102,000 building records


```sql

WITH
-- our features of interest
   feat AS (SELECT gid As building_id, geom_26986 As geom FROM buildings AS b
    WHERE gid IN(100, 103,150)
   ),
-- clip band 2 of raster tiles to boundaries of builds
-- then get stats for these clipped regions
   b_stats AS
    (SELECT  building_id, (stats).*
FROM (SELECT building_id, ST_SummaryStats(ST_Clip(rast,2,geom)) As stats
    FROM aerials.boston
        INNER JOIN feat
    ON ST_Intersects(feat.geom,rast)
 ) As foo
 )
-- finally summarize stats
SELECT building_id, SUM(count) As num_pixels
  , MIN(min) As min_pval
  ,  MAX(max) As max_pval
  , SUM(mean*count)/SUM(count) As avg_pval
    FROM b_stats
 WHERE count > 0
    GROUP BY building_id
    ORDER BY building_id;
 building_id | num_pixels | min_pval | max_pval |     avg_pval
-------------+------------+----------+----------+------------------
         100 |       1090 |        1 |      255 | 61.0697247706422
         103 |        655 |        7 |      182 | 70.5038167938931
         150 |        895 |        2 |      252 | 185.642458100559
```


## Example: Raster coverage


```

-- stats for each band --
SELECT band, (stats).*
FROM (SELECT band, ST_SummaryStats('o_4_boston','rast', band) As stats
    FROM generate_series(1,3) As band) As foo;

 band |  count  |  sum   |       mean       |      stddev      | min | max
------+---------+--------+------------------+------------------+-----+-----
    1 | 8450000 | 725799 | 82.7064349112426 | 45.6800222638537 |   0 | 255
    2 | 8450000 | 700487 | 81.4197705325444 | 44.2161184161765 |   0 | 255
    3 | 8450000 | 575943 |  74.682739408284 | 44.2143885481407 |   0 | 255

-- For a table -- will get better speed if set sampling to less than 100%
-- Here we set to 25% and get a much faster answer
SELECT band, (stats).*
FROM (SELECT band, ST_SummaryStats('o_4_boston','rast', band,true,0.25) As stats
    FROM generate_series(1,3) As band) As foo;

 band |  count  |  sum   |       mean       |      stddev      | min | max
------+---------+--------+------------------+------------------+-----+-----
    1 | 2112500 | 180686 | 82.6890480473373 | 45.6961043857248 |   0 | 255
    2 | 2112500 | 174571 |  81.448503668639 | 44.2252623171821 |   0 | 255
    3 | 2112500 | 144364 | 74.6765884023669 | 44.2014869384578 |   0 | 255

```


## See Also


 [summarystats](raster-support-data-types.md#summarystats), [RT_ST_SummaryStatsAgg](#RT_ST_SummaryStatsAgg), [RT_ST_Count](#RT_ST_Count), [RT_ST_Clip](raster-processing-map-algebra.md#RT_ST_Clip)
  <a id="RT_ST_SummaryStatsAgg"></a>

# ST_SummaryStatsAgg

Aggregate. Returns summarystats consisting of count, sum, mean, stddev, min, max for a given raster band of a set of raster. Band 1 is assumed is no band is specified.

## Synopsis


```sql
summarystats ST_SummaryStatsAgg(raster  rast, integer  nband, boolean  exclude_nodata_value, double precision  sample_percent)
summarystats ST_SummaryStatsAgg(raster  rast, boolean  exclude_nodata_value, double precision  sample_percent)
summarystats ST_SummaryStatsAgg(raster  rast, integer  nband, boolean  exclude_nodata_value)
```


## Description


Returns [summarystats](raster-support-data-types.md#summarystats) consisting of count, sum, mean, stddev, min, max for a given raster band of a raster or raster coverage. If no band is specified `nband` defaults to 1.


!!! note

    By default only considers pixel values not equal to the `NODATA` value. Set `exclude_nodata_value` to False to get count of all pixels.


!!! note

    By default will sample all pixels. To get faster response, set `sample_percent` to value between 0 and 1


Availability: 2.2.0


## Examples


```sql

WITH foo AS (
    SELECT
        rast.rast
    FROM (
        SELECT ST_SetValue(
            ST_SetValue(
                ST_SetValue(
                    ST_AddBand(
                        ST_MakeEmptyRaster(10, 10, 10, 10, 2, 2, 0, 0,0)
                        , 1, '64BF', 0, 0
                    )
                    , 1, 1, 1, -10
                )
                , 1, 5, 4, 0
            )
            , 1, 5, 5, 3.14159
        ) AS rast
    ) AS rast
    FULL JOIN (
        SELECT generate_series(1, 10) AS id
    ) AS id
        ON 1 = 1
)
SELECT
    (stats).count,
    round((stats).sum::numeric, 3),
    round((stats).mean::numeric, 3),
    round((stats).stddev::numeric, 3),
    round((stats).min::numeric, 3),
    round((stats).max::numeric, 3)
FROM (
    SELECT
        ST_SummaryStatsAgg(rast, 1, TRUE, 1) AS stats
    FROM foo
) bar;

 count |  round  | round  | round |  round  | round
-------+---------+--------+-------+---------+-------
    20 | -68.584 | -3.429 | 6.571 | -10.000 | 3.142
(1 row)

```


## See Also


 [summarystats](raster-support-data-types.md#summarystats), [RT_ST_SummaryStats](#RT_ST_SummaryStats), [RT_ST_Count](#RT_ST_Count), [RT_ST_Clip](raster-processing-map-algebra.md#RT_ST_Clip)
  <a id="RT_ST_ValueCount"></a>

# ST_ValueCount

Returns a set of records containing a pixel band value and count of the number of pixels in a given band of a raster (or a raster coverage) that have a given set of values. If no band is specified defaults to band 1. By default nodata value pixels are not counted. and all other values in the pixel are output and pixel band values are rounded to the nearest integer.

## Synopsis


```sql
SETOF record ST_ValueCount(raster  rast, integer  nband=1, boolean  exclude_nodata_value=true, double precision[]  searchvalues=NULL, double precision  roundto=0, double precision  OUT value, integer  OUT count)
SETOF record ST_ValueCount(raster  rast, integer  nband, double precision[]  searchvalues, double precision  roundto=0, double precision  OUT value, integer  OUT count)
SETOF record ST_ValueCount(raster  rast, double precision[]  searchvalues, double precision  roundto=0, double precision  OUT value, integer  OUT count)
bigint ST_ValueCount(raster  rast, double precision  searchvalue, double precision  roundto=0)
bigint ST_ValueCount(raster  rast, integer  nband, boolean  exclude_nodata_value, double precision  searchvalue, double precision  roundto=0)
bigint ST_ValueCount(raster  rast, integer  nband, double precision  searchvalue, double precision  roundto=0)
SETOF record ST_ValueCount(text  rastertable, text  rastercolumn, integer  nband=1, boolean  exclude_nodata_value=true, double precision[]  searchvalues=NULL, double precision  roundto=0, double precision  OUT value, integer  OUT count)
SETOF record ST_ValueCount(text  rastertable, text  rastercolumn, double precision[]  searchvalues, double precision  roundto=0, double precision  OUT value, integer  OUT count)
SETOF record ST_ValueCount(text  rastertable, text  rastercolumn, integer  nband, double precision[]  searchvalues, double precision  roundto=0, double precision  OUT value, integer  OUT count)
bigintST_ValueCount(text  rastertable, text  rastercolumn, integer  nband, boolean  exclude_nodata_value, double precision  searchvalue, double precision  roundto=0)
bigint ST_ValueCount(text  rastertable, text  rastercolumn, double precision  searchvalue, double precision  roundto=0)
bigint ST_ValueCount(text  rastertable, text  rastercolumn, integer  nband, double precision  searchvalue, double precision  roundto=0)
```


## Description


Returns a set of records with columns `value` `count` which contain the pixel band value and count of pixels in the raster tile or raster coverage of selected band.


If no band is specified `nband` defaults to 1. If no `searchvalues` are specified, will return all pixel values found in the raster or raster coverage. If one searchvalue is given, will return an integer instead of records denoting the count of pixels having that pixel band value


!!! note

    If `exclude_nodata_value` is set to false, will also count pixels with no data.


Availability: 2.0.0


## Examples


```sql

UPDATE dummy_rast SET rast = ST_SetBandNoDataValue(rast,249) WHERE rid=2;
--Example will count only pixels of band 1 that are not 249. --

SELECT (pvc).*
FROM (SELECT ST_ValueCount(rast) As pvc
    FROM dummy_rast WHERE rid=2) As foo
    ORDER BY (pvc).value;

 value | count
-------+-------
   250 |     2
   251 |     1
   252 |     2
   253 |     6
   254 |    12

-- Example will coount all pixels of band 1 including 249 --
SELECT (pvc).*
FROM (SELECT ST_ValueCount(rast,1,false) As pvc
    FROM dummy_rast WHERE rid=2) As foo
    ORDER BY (pvc).value;

 value | count
-------+-------
   249 |     2
   250 |     2
   251 |     1
   252 |     2
   253 |     6
   254 |    12

-- Example will count only non-nodata value pixels of band 2
SELECT (pvc).*
FROM (SELECT ST_ValueCount(rast,2) As pvc
    FROM dummy_rast WHERE rid=2) As foo
    ORDER BY (pvc).value;
 value | count
-------+-------
    78 |     1
    79 |     1
    88 |     1
    89 |     1
    96 |     1
    97 |     1
    98 |     1
    99 |     2
   112 |     2
:


```


```

--real live example.  Count all the pixels in an aerial raster tile band 2 intersecting a geometry
-- and return only the pixel band values that have a count > 500
SELECT (pvc).value, SUM((pvc).count) As total
FROM (SELECT ST_ValueCount(rast,2) As pvc
    FROM o_4_boston
        WHERE ST_Intersects(rast,
            ST_GeomFromText('POLYGON((224486 892151,224486 892200,224706 892200,224706 892151,224486 892151))',26986)
             )
        ) As foo
    GROUP BY (pvc).value
    HAVING SUM((pvc).count) > 500
    ORDER BY (pvc).value;

 value | total
-------+-----
    51 | 502
    54 | 521
```


```

-- Just return count of pixels in each raster tile that have value of 100 of tiles that intersect  a specific geometry --
SELECT rid, ST_ValueCount(rast,2,100) As count
    FROM o_4_boston
        WHERE ST_Intersects(rast,
            ST_GeomFromText('POLYGON((224486 892151,224486 892200,224706 892200,224706 892151,224486 892151))',26986)
             ) ;

 rid | count
-----+-------
   1 |    56
   2 |    95
  14 |    37
  15 |    64
```


## See Also


[RT_ST_Count](#RT_ST_Count), [RT_ST_SetBandNoDataValue](raster-band-editors.md#RT_ST_SetBandNoDataValue)
