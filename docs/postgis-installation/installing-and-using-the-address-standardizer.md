<a id="installing_pagc_address_standardizer"></a>

## Installing and Using the address standardizer


The <code>address_standardizer</code> extension used to be a separate package that required separate download. From PostGIS 2.2 on, it is now bundled in. For more information about the address_standardize, what it does, and how to configure it for your needs, refer to [Address Standardizer](../postgis-extras/address-standardizer.md#Address_Standardizer).


This standardizer can be used in conjunction with the PostGIS packaged tiger geocoder extension as a replacement for the [Normalize_Address](../postgis-extras/tiger-geocoder.md#Normalize_Address) discussed. To use as replacement refer to [Using Address Standardizer Extension with Tiger geocoder](installing-upgrading-tiger-geocoder-and-loading-data.md#tiger_pagc_address_standardizing). You can also use it as a building block for your own geocoder or use it to standardize your addresses for easier compare of addresses.


The address standardizer relies on PCRE which is usually already installed on many Nix systems, but you can download the latest at: [http://www.pcre.org](http://www.pcre.org). If during [Build configuration](compiling-and-install-from-source.md#installation_configuration), PCRE is found, then the address standardizer extension will automatically be built. If you have a custom pcre install you want to use instead, pass to configure <code>--with-pcredir=/path/to/pcre</code> where `/path/to/pcre` is the root folder for your pcre include and lib directories.


For Windows users, the PostGIS 2.1+ bundle is packaged with the address_standardizer already so no need to compile and can move straight to <code>CREATE EXTENSION</code> step.


Once you have installed, you can connect to your database and run the SQL:


```sql
CREATE EXTENSION address_standardizer;
```


The following test requires no rules, gaz, or lex tables


```sql
SELECT num, street, city, state, zip
 FROM parse_address('1 Devonshire Place PH301, Boston, MA 02109');
```


Output should be


```
 num |         street         |  city  | state |  zip
-----+------------------------+--------+-------+-------
 1   | Devonshire Place PH301 | Boston | MA    | 02109
```
