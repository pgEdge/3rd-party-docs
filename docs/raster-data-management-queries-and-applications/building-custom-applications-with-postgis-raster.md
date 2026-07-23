<a id="RT_Raster_Applications"></a>

## Building Custom Applications with PostGIS Raster


The fact that PostGIS raster provides you with SQL functions to render rasters in known image formats gives you a lot of options for rendering them. For example you can use OpenOffice / LibreOffice for rendering as demonstrated in [Rendering PostGIS Raster graphics with LibreOffice Base Reports](http://www.postgresonline.com/journal/archives/244-Rendering-PostGIS-Raster-graphics-with-LibreOffice-Base-Reports.html). In addition you can use a wide variety of languages as demonstrated in this section.
 <a id="RT_PHP_Output"></a>

## PHP Example Outputting using ST_AsPNG in concert with other raster functions


In this section, we'll demonstrate how to use the PHP PostgreSQL driver and the [RT_ST_AsGDALRaster](../raster-reference/raster-outputs.md#RT_ST_AsGDALRaster) family of functions to output band 1,2,3 of a raster to a PHP request stream that can then be embedded in an img src html tag.


The sample query demonstrates how to combine a whole bunch of raster functions together to grab all tiles that intersect a particular wgs 84 bounding box and then unions with [RT_ST_Union](../raster-reference/raster-processing-map-algebra.md#RT_ST_Union) the intersecting tiles together returning all bands, transforms to user specified projection using [RT_ST_Transform](../raster-reference/raster-editors.md#RT_ST_Transform), and then outputs the results as a png using [RT_ST_AsPNG](../raster-reference/raster-outputs.md#RT_ST_AsPNG).


You would call the below using

```
http://mywebserver/test_raster.php?srid=2249
```
 to get the raster image in Massachusetts state plane feet.


```


<?php
/** contents of test_raster.php **/
$conn_str ='dbname=mydb host=localhost port=5432 user=myuser password=mypwd';
$dbconn = pg_connect($conn_str);
header('Content-Type: image/png');
/**If a particular projection was requested use it otherwise use mass state plane meters **/
if (!empty( $_REQUEST['srid'] ) && is_numeric( $_REQUEST['srid']) ){
		$input_srid = intval($_REQUEST['srid']);
}
else { $input_srid = 26986; }
/** The set bytea_output may be needed for PostgreSQL 9.0+, but not for 8.4 **/
$sql = "set bytea_output='escape';
SELECT ST_AsPNG(ST_Transform(
			ST_AddBand(ST_Union(rast,1), ARRAY[ST_Union(rast,2),ST_Union(rast,3)])
				,$input_srid) ) As new_rast
 FROM aerials.boston
	WHERE
	 ST_Intersects(rast, ST_Transform(ST_MakeEnvelope(-71.1217, 42.227, -71.1210, 42.218,4326),26986) )";
$result = pg_query($sql);
$row = pg_fetch_row($result);
pg_free_result($result);
if ($row === false) return;
echo pg_unescape_bytea($row[0]);
?>
```
  <a id="RT_Net_Output_CS"></a>

## ASP.NET C# Example Outputting using ST_AsPNG in concert with other raster functions


In this section, we'll demonstrate how to use Npgsql PostgreSQL .NET driver and the [RT_ST_AsGDALRaster](../raster-reference/raster-outputs.md#RT_ST_AsGDALRaster) family of functions to output band 1,2,3 of a raster to a PHP request stream that can then be embedded in an img src html tag.


You will need the npgsql .NET PostgreSQL driver for this exercise which you can get the latest of from [http://npgsql.projects.postgresql.org/](http://npgsql.projects.postgresql.org/). Just download the latest and drop into your ASP.NET bin folder and you'll be good to go.


The sample query demonstrates how to combine a whole bunch of raster functions together to grab all tiles that intersect a particular wgs 84 bounding box and then unions with [RT_ST_Union](../raster-reference/raster-processing-map-algebra.md#RT_ST_Union) the intersecting tiles together returning all bands, transforms to user specified projection using [RT_ST_Transform](../raster-reference/raster-editors.md#RT_ST_Transform), and then outputs the results as a png using [RT_ST_AsPNG](../raster-reference/raster-outputs.md#RT_ST_AsPNG).


This is same example as [PHP Example Outputting using ST_AsPNG in concert with other raster functions](#RT_PHP_Output) except implemented in C#.


You would call the below using

```
http://mywebserver/TestRaster.ashx?srid=2249
```
 to get the raster image in Massachusetts state plane feet.


```

 -- web.config connection string section --
<connectionStrings>
    <add name="DSN"
        connectionString="server=localhost;database=mydb;Port=5432;User Id=myuser;password=mypwd"/>
</connectionStrings>
```


```

// Code for TestRaster.ashx
<%@ WebHandler Language="C#" Class="TestRaster" %>
using System;
using System.Data;
using System.Web;
using Npgsql;

public class TestRaster : IHttpHandler
{
	public void ProcessRequest(HttpContext context)
	{

		context.Response.ContentType = "image/png";
		context.Response.BinaryWrite(GetResults(context));

	}

	public bool IsReusable {
		get { return false; }
	}

	public byte[] GetResults(HttpContext context)
	{
		byte[] result = null;
		NpgsqlCommand command;
		string sql = null;
		int input_srid = 26986;
        try {
		    using (NpgsqlConnection conn = new NpgsqlConnection(System.Configuration.ConfigurationManager.ConnectionStrings["DSN"].ConnectionString)) {
			    conn.Open();

                if (context.Request["srid"] != null)
                {
                    input_srid = Convert.ToInt32(context.Request["srid"]);
                }
                sql = @"SELECT ST_AsPNG(
                            ST_Transform(
			                ST_AddBand(
                                ST_Union(rast,1), ARRAY[ST_Union(rast,2),ST_Union(rast,3)])
				                    ,:input_srid) ) As new_rast
                        FROM aerials.boston
	                        WHERE
	                            ST_Intersects(rast,
                                    ST_Transform(ST_MakeEnvelope(-71.1217, 42.227, -71.1210, 42.218,4326),26986) )";
			    command = new NpgsqlCommand(sql, conn);
                command.Parameters.Add(new NpgsqlParameter("input_srid", input_srid));


			    result = (byte[]) command.ExecuteScalar();
                conn.Close();
			}

		}
        catch (Exception ex)
        {
            result = null;
            context.Response.Write(ex.Message.Trim());
        }
		return result;
	}
}
```
  <a id="RT_Java_Console_App"></a>

## Java console app that outputs raster query as Image file


This is a simple java console app that takes a query that returns one image and outputs to specified file.


You can download the latest PostgreSQL JDBC drivers from [http://jdbc.postgresql.org/download.html](http://jdbc.postgresql.org/download.html)


You can compile the following code using a command something like:


```
set env CLASSPATH .:..\postgresql-9.0-801.jdbc4.jar
javac SaveQueryImage.java
jar cfm SaveQueryImage.jar Manifest.txt *.class
```


And call it from the command-line with something like


```
java -jar SaveQueryImage.jar "SELECT ST_AsPNG(ST_AsRaster(ST_Buffer(ST_Point(1,5),10, 'quad_segs=2'),150, 150, '8BUI',100));" "test.png"
```


```
 -- Manifest.txt --
Class-Path: postgresql-9.0-801.jdbc4.jar
Main-Class: SaveQueryImage
```


```
// Code for SaveQueryImage.java
import java.sql.Connection;
import java.sql.SQLException;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.io.*;

public class SaveQueryImage {
  public static void main(String[] argv) {
      System.out.println("Checking if Driver is registered with DriverManager.");

      try {
        //java.sql.DriverManager.registerDriver (new org.postgresql.Driver());
        Class.forName("org.postgresql.Driver");
      }
      catch (ClassNotFoundException cnfe) {
        System.out.println("Couldn't find the driver!");
        cnfe.printStackTrace();
        System.exit(1);
      }

      Connection conn = null;

      try {
        conn = DriverManager.getConnection("jdbc:postgresql://localhost:5432/mydb","myuser", "mypwd");
        conn.setAutoCommit(false);

        PreparedStatement sGetImg = conn.prepareStatement(argv[0]);

        ResultSet rs = sGetImg.executeQuery();

		FileOutputStream fout;
		try
		{
			rs.next();
			/** Output to file name requested by user **/
			fout = new FileOutputStream(new File(argv[1]) );
			fout.write(rs.getBytes(1));
			fout.close();
		}
		catch(Exception e)
		{
			System.out.println("Can't create file");
			e.printStackTrace();
		}

        rs.close();
		sGetImg.close();
        conn.close();
      }
      catch (SQLException se) {
        System.out.println("Couldn't connect: print out a stack trace and exit.");
        se.printStackTrace();
        System.exit(1);
      }
  }
}
```
  <a id="RT_PLPython"></a>

## Use PLPython to dump out images via SQL


This is a plpython stored function that creates a file in the server directory for each record. Requires you have plpython installed. Should work fine with both plpythonu and plpython3u.


```sql
CREATE OR REPLACE FUNCTION write_file (param_bytes bytea, param_filepath text)
RETURNS text
AS $$
f = open(param_filepath, 'wb+')
f.write(param_bytes)
return param_filepath
$$ LANGUAGE plpythonu;
```


```
--write out 5 images to the PostgreSQL server in varying sizes
-- note the postgresql daemon account needs to have write access to folder
-- this echos back the file names created;
 SELECT write_file(ST_AsPNG(
	ST_AsRaster(ST_Buffer(ST_Point(1,5),j*5, 'quad_segs=2'),150*j, 150*j, '8BUI',100)),
	 'C:/temp/slices'|| j || '.png')
	 FROM generate_series(1,5) As j;

     write_file
---------------------
 C:/temp/slices1.png
 C:/temp/slices2.png
 C:/temp/slices3.png
 C:/temp/slices4.png
 C:/temp/slices5.png
```
  <a id="RasterOutput_PSQL"></a>

## Outputting Rasters with PSQL


Sadly PSQL doesn't have easy to use built-in functionality for outputting binaries. This is a bit of a hack that piggy backs on PostgreSQL somewhat legacy large object support. To use first launch your psql commandline connected to your database.


Unlike the python approach, this approach creates the file on your local computer.


```
SELECT oid, lowrite(lo_open(oid, 131072), png) As num_bytes
 FROM
 ( VALUES (lo_create(0),
   ST_AsPNG( (SELECT rast FROM aerials.boston WHERE rid=1) )
  ) ) As v(oid,png);
-- you'll get an output something like --
   oid   | num_bytes
---------+-----------
 2630819 |     74860

-- next note the oid and do this replacing the c:/test.png to file path location
-- on your local computer
 \lo_export 2630819 'C:/temp/aerial_samp.png'

-- this deletes the file from large object storage on db
SELECT lo_unlink(2630819);

```
