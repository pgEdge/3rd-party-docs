<a id="Raster_Inputs"></a>

## Raster Inputs
  <a id="RT_ST_RastFromWKB"></a>

# ST_RastFromWKB

Return a raster value from a Well-Known Binary (WKB) raster.

## Synopsis


```sql
raster ST_RastFromWKB(bytea  wkb)
```


## Description


 Given a Well-Known Binary (WKB) raster, return a raster.


Availability: 2.5.0


## Examples


```sql

SELECT (ST_Metadata(
    ST_RastFromWKB(
        '\001\000\000\000\000\000\000\000\000\000\000\000@\000\000\000\000\000\000\010@\000\000\000\000\000\000\340?\000\000\000\000\000\000\340?\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\000\012\000\000\000\012\000\024\000'::bytea
    )
)).* AS metadata;

 upperleftx | upperlefty | width | height | scalex | scaley | skewx | skewy | srid | numbands
------------+------------+-------+--------+--------+--------+-------+-------+------+----------
        0.5 |        0.5 |    10 |     20 |      2 |      3 |     0 |     0 |   10 |        0

```


## See Also


 [RT_ST_MetaData](raster-accessors.md#RT_ST_MetaData), [RT_ST_RastFromHexWKB](#RT_ST_RastFromHexWKB), [RT_ST_AsBinary](raster-outputs.md#RT_ST_AsBinary), [RT_ST_AsHexWKB](raster-outputs.md#RT_ST_AsHexWKB)
  <a id="RT_ST_RastFromHexWKB"></a>

# ST_RastFromHexWKB

Return a raster value from a Hex representation of Well-Known Binary (WKB) raster.

## Synopsis


```sql
raster ST_RastFromHexWKB(text  wkb)
```


## Description


 Given a Well-Known Binary (WKB) raster in Hex representation, return a raster.


Availability: 2.5.0


## Examples


```sql

SELECT (ST_Metadata(
    ST_RastFromHexWKB(
        '010000000000000000000000400000000000000840000000000000E03F000000000000E03F000000000000000000000000000000000A0000000A001400'
    )
)).* AS metadata;

 upperleftx | upperlefty | width | height | scalex | scaley | skewx | skewy | srid | numbands
------------+------------+-------+--------+--------+--------+-------+-------+------+----------
        0.5 |        0.5 |    10 |     20 |      2 |      3 |     0 |     0 |   10 |        0

```


## See Also


 [RT_ST_MetaData](raster-accessors.md#RT_ST_MetaData), [RT_ST_RastFromWKB](#RT_ST_RastFromWKB), [RT_ST_AsBinary](raster-outputs.md#RT_ST_AsBinary), [RT_ST_AsHexWKB](raster-outputs.md#RT_ST_AsHexWKB)
