<a id="RasterBand_Accessors"></a>

## Raster Band Accessors
  <a id="RT_ST_BandMetaData"></a>

# ST_BandMetaData

Returns basic meta data for a specific raster band. band num 1 is assumed if none-specified.

## Synopsis


```sql
(1) record ST_BandMetaData(raster rast, integer band=1)
(2) record ST_BandMetaData(raster rast, integer[] band)
```


## Description


Returns basic meta data about a raster band. Columns returned: pixeltype, nodatavalue, isoutdb, path, outdbbandnum, filesize, filetimestamp.


!!! note

    If raster contains no bands then an error is thrown.


!!! note

    If band has no NODATA value, nodatavalue are NULL.


!!! note

    If isoutdb is False, path, outdbbandnum, filesize and filetimestamp are NULL. If outdb access is disabled, filesize and filetimestamp will also be NULL.


Enhanced: 2.5.0 to include *outdbbandnum*, *filesize* and *filetimestamp* for outdb rasters.


## Examples: Variant 1


```sql

SELECT
    rid,
    (foo.md).*
FROM (
    SELECT
        rid,
        ST_BandMetaData(rast, 1) AS md
    FROM dummy_rast
    WHERE rid=2
) As foo;

 rid | pixeltype | nodatavalue | isoutdb | path | outdbbandnum
-----+-----------+---- --------+---------+------+--------------
   2 | 8BUI      |           0 | f       |      |

```


## Examples: Variant 2


```sql

WITH foo AS (
    SELECT
        ST_AddBand(NULL::raster, '/home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif', NULL::int[]) AS rast
)
SELECT
    *
FROM ST_BandMetadata(
    (SELECT rast FROM foo),
    ARRAY[1,3,2]::int[]
);

 bandnum | pixeltype | nodatavalue | isoutdb |                                      path                                      | outdbbandnum  | filesize | filetimestamp |
---------+-----------+-------------+---------+--------------------------------------------------------------------------------+---------------+----------+---------------+-
       1 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif      |            1  |    12345 |    1521807257 |
       3 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif      |            3  |    12345 |    1521807257 |
       2 | 8BUI      |             | t       | /home/pele/devel/geo/postgis-git/raster/test/regress/loader/Projected.tif      |            2  |    12345 |    1521807257 |

```


## See Also


[RT_ST_MetaData](raster-accessors.md#RT_ST_MetaData), [RT_ST_BandPixelType](#RT_ST_BandPixelType)
  <a id="RT_ST_BandNoDataValue"></a>

# ST_BandNoDataValue

Returns the value in a given band that represents no data. If no band num 1 is assumed.

## Synopsis


```sql
double precision ST_BandNoDataValue(raster  rast, integer  bandnum=1)
```


## Description


Returns the value that represents no data for the band


## Examples


```sql
SELECT ST_BandNoDataValue(rast,1) As bnval1,
    ST_BandNoDataValue(rast,2) As bnval2, ST_BandNoDataValue(rast,3) As bnval3
FROM dummy_rast
WHERE rid = 2;

 bnval1 | bnval2 | bnval3
--------+--------+--------
      0 |      0 |      0

```


## See Also


[RT_ST_NumBands](raster-accessors.md#RT_ST_NumBands)
  <a id="RT_ST_BandIsNoData"></a>

# ST_BandIsNoData

Returns true if the band is filled with only nodata values.

## Synopsis


```sql
boolean ST_BandIsNoData(raster  rast, integer  band, boolean  forceChecking=true)
boolean ST_BandIsNoData(raster  rast, boolean  forceChecking=true)
```


## Description


Returns true if the band is filled with only nodata values. Band 1 is assumed if not specified. If the last argument is TRUE, the entire band is checked pixel by pixel. Otherwise, the function simply returns the value of the isnodata flag for the band. The default value for this parameter is FALSE, if not specified.


Availability: 2.0.0


!!! note

    If the flag is dirty (this is, the result is different using TRUE as last parameter and not using it) you should update the raster to set this flag to true, by using ST_SetBandIsNodata(), or ST_SetBandNodataValue() with TRUE as last argument. See [RT_ST_SetBandIsNoData](raster-band-editors.md#RT_ST_SetBandIsNoData).


## Examples


```

-- Create dummy table with one raster column
create table dummy_rast (rid integer, rast raster);

-- Add raster with two bands, one pixel/band. In the first band, nodatavalue = pixel value = 3.
-- In the second band, nodatavalue = 13, pixel value = 4
insert into dummy_rast values(1,
(
'01' -- little endian (uint8 ndr)
||
'0000' -- version (uint16 0)
||
'0200' -- nBands (uint16 0)
||
'17263529ED684A3F' -- scaleX (float64 0.000805965234044584)
||
'F9253529ED684ABF' -- scaleY (float64 -0.00080596523404458)
||
'1C9F33CE69E352C0' -- ipX (float64 -75.5533328537098)
||
'718F0E9A27A44840' -- ipY (float64 49.2824585505576)
||
'ED50EB853EC32B3F' -- skewX (float64 0.000211812383858707)
||
'7550EB853EC32B3F' -- skewY (float64 0.000211812383858704)
||
'E6100000' -- SRID (int32 4326)
||
'0100' -- width (uint16 1)
||
'0100' -- height (uint16 1)
||
'6' -- hasnodatavalue and isnodata value set to true.
||
'2' -- first band type (4BUI)
||
'03' -- novalue==3
||
'03' -- pixel(0,0)==3 (same that nodata)
||
'0' -- hasnodatavalue set to false
||
'5' -- second band type (16BSI)
||
'0D00' -- novalue==13
||
'0400' -- pixel(0,0)==4
)::raster
);

select st_bandisnodata(rast, 1) from dummy_rast where rid = 1; -- Expected true
select st_bandisnodata(rast, 2) from dummy_rast where rid = 1; -- Expected false

```


## See Also


[RT_ST_BandNoDataValue](#RT_ST_BandNoDataValue), [RT_ST_NumBands](raster-accessors.md#RT_ST_NumBands), [RT_ST_SetBandNoDataValue](raster-band-editors.md#RT_ST_SetBandNoDataValue), [RT_ST_SetBandIsNoData](raster-band-editors.md#RT_ST_SetBandIsNoData)
  <a id="RT_ST_BandPath"></a>

# ST_BandPath

Returns system file path to a band stored in file system. If no bandnum specified, 1 is assumed.

## Synopsis


```sql
text ST_BandPath(raster  rast, integer  bandnum=1)
```


## Description


Returns system file path to a band. Throws an error if called with an in db band.


## Examples


```


```


## See Also


  <a id="RT_ST_BandFileSize"></a>

# ST_BandFileSize

Returns the file size of a band stored in file system. If no bandnum specified, 1 is assumed.

## Synopsis


```sql
bigint ST_BandFileSize(raster  rast, integer  bandnum=1)
```


## Description


Returns the file size of a band stored in file system. Throws an error if called with an in db band, or if outdb access is not enabled.


This function is typically used in conjunction with ST_BandPath() and ST_BandFileTimestamp() so a client can determine if the filename of a outdb raster as seen by it is the same as the one seen by the server.


Availability: 2.5.0


## Examples


```sql
SELECT ST_BandFileSize(rast,1) FROM dummy_rast WHERE rid = 1;

 st_bandfilesize
-----------------
          240574

```
  <a id="RT_ST_BandFileTimestamp"></a>

# ST_BandFileTimestamp

Returns the file timestamp of a band stored in file system. If no bandnum specified, 1 is assumed.

## Synopsis


```sql
bigint ST_BandFileTimestamp(raster  rast, integer  bandnum=1)
```


## Description


Returns the file timestamp (number of seconds since Jan 1st 1970 00:00:00 UTC) of a band stored in file system. Throws an error if called with an in db band, or if outdb access is not enabled.


This function is typically used in conjunction with ST_BandPath() and ST_BandFileSize() so a client can determine if the filename of a outdb raster as seen by it is the same as the one seen by the server.


Availability: 2.5.0


## Examples


```sql
SELECT ST_BandFileTimestamp(rast,1) FROM dummy_rast WHERE rid = 1;

 st_bandfiletimestamp
----------------------
           1521807257

```
  <a id="RT_ST_BandPixelType"></a>

# ST_BandPixelType

Returns the type of pixel for given band. If no bandnum specified, 1 is assumed.

## Synopsis


```sql
text ST_BandPixelType(raster  rast, integer  bandnum=1)
```


## Description


Returns name describing data type and size of values stored in each cell of given band.


There are 11 pixel types. Pixel Types supported are as follows:

- 1BB - 1-bit boolean
- 2BUI - 2-bit unsigned integer
- 4BUI - 4-bit unsigned integer
- 8BSI - 8-bit signed integer
- 8BUI - 8-bit unsigned integer
- 16BSI - 16-bit signed integer
- 16BUI - 16-bit unsigned integer
- 32BSI - 32-bit signed integer
- 32BUI - 32-bit unsigned integer
- 32BF - 32-bit float
- 64BF - 64-bit float


## Examples


```sql
SELECT ST_BandPixelType(rast,1) As btype1,
    ST_BandPixelType(rast,2) As btype2, ST_BandPixelType(rast,3) As btype3
FROM dummy_rast
WHERE rid = 2;

 btype1 | btype2 | btype3
--------+--------+--------
 8BUI   | 8BUI   | 8BUI

```


## See Also


[RT_ST_NumBands](raster-accessors.md#RT_ST_NumBands)
  <a id="ST_MinPossibleValue"></a>

# ST_MinPossibleValue

Returns the minimum value this pixeltype can store.

## Synopsis


```sql
integer ST_MinPossibleValue(text  pixeltype)
```


## Description


Returns the minimum value this pixeltype can store.


## Examples


```sql
SELECT ST_MinPossibleValue('16BSI');

 st_minpossiblevalue
---------------------
              -32768


SELECT ST_MinPossibleValue('8BUI');

 st_minpossiblevalue
---------------------
                   0

```


## See Also


[RT_ST_BandPixelType](#RT_ST_BandPixelType)
  <a id="RT_ST_HasNoBand"></a>

# ST_HasNoBand

Returns true if there is no band with given band number. If no band number is specified, then band number 1 is assumed.

## Synopsis


```sql
boolean ST_HasNoBand(raster  rast, integer  bandnum=1)
```


## Description


Returns true if there is no band with given band number. If no band number is specified, then band number 1 is assumed.


Availability: 2.0.0


## Examples


```sql
SELECT rid, ST_HasNoBand(rast) As hb1, ST_HasNoBand(rast,2) as hb2,
ST_HasNoBand(rast,4) as hb4, ST_NumBands(rast) As numbands
FROM dummy_rast;

rid | hb1 | hb2 | hb4 | numbands
-----+-----+-----+-----+----------
1 | t   | t   | t   |        0
2 | f   | f   | t   |        3

```


## See Also


[RT_ST_NumBands](raster-accessors.md#RT_ST_NumBands)
