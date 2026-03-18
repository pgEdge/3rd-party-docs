<a id="Troubleshooting_Functions"></a>

## Troubleshooting Functions
  <a id="PostGIS_AddBBox"></a>

# PostGIS_AddBBox

Add bounding box to the geometry.

## Synopsis


```sql
geometry PostGIS_AddBBox(geometry  geomA)
```


## Description


Add bounding box to the geometry. This would make bounding box based queries faster, but will increase the size of the geometry.


!!! note

    Bounding boxes are automatically added to geometries so in general this is not needed unless the generated bounding box somehow becomes corrupted or you have an old install that is lacking bounding boxes. Then you need to drop the old and re-add.


## Examples


```sql
UPDATE sometable
 SET geom =  PostGIS_AddBBox(geom)
 WHERE PostGIS_HasBBox(geom) = false;
```


## See Also


[PostGIS_DropBBox](#PostGIS_DropBBox), [PostGIS_HasBBox](#PostGIS_HasBBox)
  <a id="PostGIS_DropBBox"></a>

# PostGIS_DropBBox

Drop the bounding box cache from the geometry.

## Synopsis


```sql
geometry PostGIS_DropBBox(geometry  geomA)
```


## Description


Drop the bounding box cache from the geometry. This reduces geometry size, but makes bounding-box based queries slower. It is also used to drop a corrupt bounding box. A tale-tell sign of a corrupt cached bounding box is when your ST_Intersects and other relation queries leave out geometries that rightfully should return true.


!!! note

    Bounding boxes are automatically added to geometries and improve speed of queries so in general this is not needed unless the generated bounding box somehow becomes corrupted or you have an old install that is lacking bounding boxes. Then you need to drop the old and re-add. This kind of corruption has been observed in 8.3-8.3.6 series whereby cached bboxes were not always recalculated when a geometry changed and upgrading to a newer version without a dump reload will not correct already corrupted boxes. So one can manually correct using below and re-add the bbox or do a dump reload.


## Examples


```
--This example drops bounding boxes where the cached box is not correct
			--The force to ST_AsBinary before applying Box2D forces a recalculation of the box, and Box2D applied to the table geometry always
			-- returns the cached bounding box.
			UPDATE sometable
 SET geom =  PostGIS_DropBBox(geom)
 WHERE Not (Box2D(ST_AsBinary(geom)) = Box2D(geom));

	UPDATE sometable
 SET geom =  PostGIS_AddBBox(geom)
 WHERE Not PostGIS_HasBBOX(geom);


```


## See Also


[PostGIS_AddBBox](#PostGIS_AddBBox), [PostGIS_HasBBox](#PostGIS_HasBBox), [Box2D](bounding-box-functions.md#Box2D)
  <a id="PostGIS_HasBBox"></a>

# PostGIS_HasBBox

Returns TRUE if the bbox of this geometry is cached, FALSE otherwise.

## Synopsis


```sql
boolean PostGIS_HasBBox(geometry  geomA)
```


## Description


Returns TRUE if the bbox of this geometry is cached, FALSE otherwise. Use [PostGIS_AddBBox](#PostGIS_AddBBox) and [PostGIS_DropBBox](#PostGIS_DropBBox) to control caching.


## Examples


```sql
SELECT geom
FROM sometable WHERE PostGIS_HasBBox(geom) = false;
```


## See Also


[PostGIS_AddBBox](#PostGIS_AddBBox), [PostGIS_DropBBox](#PostGIS_DropBBox)
