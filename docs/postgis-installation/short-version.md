<a id="install_short_version"></a>

## Short Version


To compile assuming you have all the dependencies in your search path:


```
tar -xvzf postgis-3.5.6.tar.gz
cd postgis-3.5.6
./configure
make
make install
```


 Once PostGIS is installed, it needs to be enabled ([Creating spatial databases](../postgis-administration/creating-spatial-databases.md#create_spatial_db)) or upgraded ([Upgrading spatial databases](../postgis-administration/upgrading-spatial-databases.md#upgrading)) in each individual database you want to use it in.
