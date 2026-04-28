<a id="Address_Standardizer"></a>

## Address Standardizer


This is a fork of the [PAGC standardizer](http://www.pagcgeo.org/docs/html/pagc-11.html) (original code for this portion was [PAGC PostgreSQL Address Standardizer](http://sourceforge.net/p/pagc/code/360/tree/branches/sew-refactor/postgresql)).


The address standardizer is a single line address parser that takes an input address and normalizes it based on a set of rules stored in a table and helper lex and gaz tables.


The code is built into a single PostgreSQL extension library called <code>address_standardizer</code> which can be installed with <code>CREATE EXTENSION address_standardizer;</code>. In addition to the address_standardizer extension, a sample data extension called <code>address_standardizer_data_us</code> extensions is built, which contains gaz, lex, and rules tables for US data. This extensions can be installed via: <code>CREATE EXTENSION address_standardizer_data_us;</code>


The code for this extension can be found in the PostGIS `extensions/address_standardizer` and is currently self-contained.


For installation instructions refer to: [Installing and Using the address standardizer](../postgis-installation/installing-and-using-the-address-standardizer.md#installing_pagc_address_standardizer).
 <a id="Address_Standardizer_Basics"></a>

## How the Parser Works


The parser works from right to left looking first at the macro elements for postcode, state/province, city, and then looks micro elements to determine if we are dealing with a house number street or intersection or landmark. It currently does not look for a country code or name, but that could be introduced in the future.


Country code
:   Assumed to be US or CA based on: postcode as US or Canada state/province as US or Canada else US

Postcode/zipcode
:   These are recognized using Perl compatible regular expressions. These regexs are currently in the parseaddress-api.c and are relatively simple to make changes to if needed.

State/province
:   These are recognized using Perl compatible regular expressions. These regexs are currently in the parseaddress-api.c but could get moved into includes in the future for easier maintenance.
  <a id="Address_Standardizer_Types"></a>

## Address Standardizer Types
  <a id="stdaddr"></a>

# stdaddr

A composite type that consists of the elements of an address. This is the return type for `standardize_address` function.

## Description


A composite type that consists of elements of an address. This is the return type for [standardize_address](#standardize_address) function. Some descriptions for elements are borrowed from [PAGC Postal Attributes](http://www.pagcgeo.org/docs/html/pagc-12.html#ss12.1).


The token numbers denote the output reference number in the [rulestab](#rulestab).


`building`
:   is text (token number <code>0</code>): Refers to building number or name. Unparsed building identifiers and types. Generally blank for most addresses.

`house_num`
:   is a text (token number <code>1</code>): This is the street number on a street. Example *75* in <code>75 State Street</code>.

`predir`
:   is text (token number <code>2</code>): STREET NAME PRE-DIRECTIONAL such as North, South, East, West etc.

`qual`
:   is text (token number <code>3</code>): STREET NAME PRE-MODIFIER Example *OLD* in <code>3715 OLD HIGHWAY 99</code>.

`pretype`
:   is text (token number <code>4</code>): STREET PREFIX TYPE

`name`
:   is text (token number <code>5</code>): STREET NAME

`suftype`
:   is text (token number <code>6</code>): STREET POST TYPE e.g. St, Ave, Cir. A street type following the root street name. Example *STREET* in <code>75 State Street</code>.

`sufdir`
:   is text (token number <code>7</code>): STREET POST-DIRECTIONAL A directional modifier that follows the street name.. Example *WEST* in <code>3715 TENTH AVENUE WEST</code>.

`ruralroute`
:   is text (token number <code>8</code>): RURAL ROUTE . Example *7* in <code>RR 7</code>.

`extra`
:   is text: Extra information like Floor number.

`city`
:   is text (token number <code>10</code>): Example Boston.

`state`
:   is text (token number <code>11</code>): Example <code>MASSACHUSETTS</code>

`country`
:   is text (token number <code>12</code>): Example <code>USA</code>

`postcode`
:   is text POSTAL CODE (ZIP CODE) (token number <code>13</code>): Example <code>02109</code>

`box`
:   is text POSTAL BOX NUMBER (token number <code>14 and 15</code>): Example <code>02109</code>

`unit`
:   is text Apartment number or Suite Number (token number <code>17</code>): Example *3B* in <code>APT 3B</code>.
   <a id="Address_Standardizer_Tables"></a>

## Address Standardizer Tables
  <a id="rulestab"></a>

# rules table

The rules table contains a set of rules that maps address input sequence tokens to standardized output sequence. A rule is defined as a set of input tokens followed by -1 (terminator) followed by set of output tokens followed by -1 followed by number denoting kind of rule followed by ranking of rule.

## Description


A rules table must have at least the following columns, though you are allowed to add more for your own uses.


`id`
:   Primary key of table

`rule`
:   text field denoting the rule. Details at [PAGC Address Standardizer Rule records](http://www.pagcgeo.org/docs/html/pagc-12.html#--r-rec--).


    A rule consists of a set of non-negative integers representing input tokens, terminated by a -1, followed by an equal number of non-negative integers representing postal attributes, terminated by a -1, followed by an integer representing a rule type, followed by an integer representing the rank of the rule. The rules are ranked from 0 (lowest) to 17 (highest).


    So for example the rule <code>2 0 2 22 3 -1 5 5 6 7 3 -1 2 6</code> maps to sequence of output tokens *TYPE NUMBER TYPE DIRECT QUALIF* to the output sequence *STREET STREET SUFTYP SUFDIR QUALIF*. The rule is an ARC_C rule of rank 6.


    Numbers for corresponding output tokens are listed in [stdaddr](#stdaddr).
 <a id="rule_input_tokens"></a>

## Input Tokens


Each rule starts with a set of input tokens followed by a terminator <code>-1</code>. Valid input tokens excerpted from [PAGC Input Tokens](http://www.pagcgeo.org/docs/html/pagc-12.html#ss12.2) are as follows:


**Form-Based Input Tokens**


`AMPERS`
:   (13). The ampersand (&) is frequently used to abbreviate the word "and".

`DASH`
:   (9). A punctuation character.

`DOUBLE`
:   (21). A sequence of two letters. Often used as identifiers.

`FRACT`
:   (25). Fractions are sometimes used in civic numbers or unit numbers.

`MIXED`
:   (23). An alphanumeric string that contains both letters and digits. Used for identifiers.

`NUMBER`
:   (0). A string of digits.

`ORD`
:   (15). Representations such as First or 1st. Often used in street names.

`ORD`
:   (18). A single letter.

`WORD`
:   (1). A word is a string of letters of arbitrary length. A single letter can be both a SINGLE and a WORD.


**Function-based Input Tokens**


`BOXH`
:   (14). Words used to denote post office boxes. For example *Box* or *PO Box*.

`BUILDH`
:   (19). Words used to denote buildings or building complexes, usually as a prefix. For example: *Tower* in *Tower 7A*.

`BUILDT`
:   (24). Words and abbreviations used to denote buildings or building complexes, usually as a suffix. For example: *Shopping Centre*.

`DIRECT`
:   (22). Words used to denote directions, for example *North*.

`MILE`
:   (20). Words used to denote milepost addresses.

`ROAD`
:   (6). Words and abbreviations used to denote highways and roads. For example: the *Interstate* in *Interstate 5*

`RR`
:   (8). Words and abbreviations used to denote rural routes. *RR*.

`TYPE`
:   (2). Words and abbreviation used to denote street typess. For example: *ST* or *AVE*.

`UNITH`
:   (16). Words and abbreviation used to denote internal subaddresses. For example, *APT* or *UNIT*.


**Postal Type Input Tokens**


`QUINT`
:   (28). A 5 digit number. Identifies a Zip Code

`QUAD`
:   (29). A 4 digit number. Identifies ZIP4.

`PCH`
:   (27). A 3 character sequence of letter number letter. Identifies an FSA, the first 3 characters of a Canadian postal code.

`PCT`
:   (26). A 3 character sequence of number letter number. Identifies an LDU, the last 3 characters of a Canadian postal code.


**Stopwords**


STOPWORDS combine with WORDS. In rules a string of multiple WORDs and STOPWORDs will be represented by a single WORD token.


`STOPWORD`
:   (7). A word with low lexical significance, that can be omitted in parsing. For example: *THE*.


## Output Tokens


After the first -1 (terminator), follows the output tokens and their order, followed by a terminator <code>-1</code>. Numbers for corresponding output tokens are listed in [stdaddr](#stdaddr). What are allowed is dependent on kind of rule. Output tokens valid for each rule type are listed in [Rule Types and Rank](#rule_types_rank).
 <a id="rule_types_rank"></a>

## Rule Types and Rank


The final part of the rule is the rule type which is denoted by one of the following, followed by a rule rank. The rules are ranked from 0 (lowest) to 17 (highest).


**`MACRO_C`**


(token number = "**0**"). The class of rules for parsing MACRO clauses such as *PLACE STATE ZIP*


**`MACRO_C` output tokens** (excerpted from [http://www.pagcgeo.org/docs/html/pagc-12.html#--r-typ--](http://www.pagcgeo.org/docs/html/pagc-12.html#--r-typ--).


`CITY`
:   (token number "10"). Example "Albany"

`STATE`
:   (token number "11"). Example "NY"

`NATION`
:   (token number "12"). This attribute is not used in most reference files. Example "USA"

`POSTAL`
:   (token number "13"). (SADS elements "ZIP CODE" , "PLUS 4" ). This attribute is used for both the US Zip and the Canadian Postal Codes.


**`MICRO_C`**


(token number = "**1**"). The class of rules for parsing full MICRO clauses (such as House, street, sufdir, predir, pretyp, suftype, qualif) (ie ARC_C plus CIVIC_C). These rules are not used in the build phase.


**`MICRO_C` output tokens** (excerpted from [http://www.pagcgeo.org/docs/html/pagc-12.html#--r-typ--](http://www.pagcgeo.org/docs/html/pagc-12.html#--r-typ--).


`HOUSE`
:   is a text (token number <code>1</code>): This is the street number on a street. Example *75* in <code>75 State Street</code>.

`predir`
:   is text (token number <code>2</code>): STREET NAME PRE-DIRECTIONAL such as North, South, East, West etc.

`qual`
:   is text (token number <code>3</code>): STREET NAME PRE-MODIFIER Example *OLD* in <code>3715 OLD HIGHWAY 99</code>.

`pretype`
:   is text (token number <code>4</code>): STREET PREFIX TYPE

`street`
:   is text (token number <code>5</code>): STREET NAME

`suftype`
:   is text (token number <code>6</code>): STREET POST TYPE e.g. St, Ave, Cir. A street type following the root street name. Example *STREET* in <code>75 State Street</code>.

`sufdir`
:   is text (token number <code>7</code>): STREET POST-DIRECTIONAL A directional modifier that follows the street name.. Example *WEST* in <code>3715 TENTH AVENUE WEST</code>.


**`ARC_C`**


(token number = "**2**"). The class of rules for parsing MICRO clauses, excluding the HOUSE attribute. As such uses same set of output tokens as MICRO_C minus the HOUSE token.


**`CIVIC_C`**


(token number = "**3**"). The class of rules for parsing the HOUSE attribute.


**`EXTRA_C`**


(token number = "**4**"). The class of rules for parsing EXTRA attributes - attributes excluded from geocoding. These rules are not used in the build phase.


**`EXTRA_C` output tokens** (excerpted from [http://www.pagcgeo.org/docs/html/pagc-12.html#--r-typ--](http://www.pagcgeo.org/docs/html/pagc-12.html#--r-typ--).


`BLDNG`
:   (token number <code>0</code>): Unparsed building identifiers and types.

`BOXH`
:   (token number <code>14</code>): The **BOX** in <code>BOX 3B</code>

`BOXT`
:   (token number <code>15</code>): The **3B** in <code>BOX 3B</code>

`RR`
:   (token number <code>8</code>): The **RR** in <code>RR 7</code>

`UNITH`
:   (token number <code>16</code>): The **APT** in <code>APT 3B</code>

`UNITT`
:   (token number <code>17</code>): The **3B** in <code>APT 3B</code>

`UNKNWN`
:   (token number <code>9</code>): An otherwise unclassified output.
  <a id="lextab"></a>

# lex table

A lex table is used to classify alphanumeric input and associate that input with (a) input tokens ( See [Input Tokens](#rule_input_tokens)) and (b) standardized representations.

## Description


A lex (short for lexicon) table is used to classify alphanumeric input and associate that input with [Input Tokens](#rule_input_tokens) and (b) standardized representations. Things you will find in these tables are <code>ONE</code> mapped to stdword: <code>1</code>.


A lex has at least the following columns in the table. You may add


`id`
:   Primary key of table

`seq`
:   integer: definition number?

`word`
:   text: the input word

`stdword`
:   text: the standardized replacement word

`token`
:   integer: the kind of word it is. Only if it is used in this context will it be replaced. Refer to [PAGC Tokens](http://www.pagcgeo.org/docs/html/pagc-12.html#--i-tok--).
  <a id="gaztab"></a>

# gaz table

A gaz table is used to standardize place names and associate that input with (a) input tokens ( See [Input Tokens](#rule_input_tokens)) and (b) standardized representations.

## Description


A gaz (short for gazeteer) table is used to standardize place names and associate that input with [Input Tokens](#rule_input_tokens) and (b) standardized representations. For example if you are in US, you may load these with State Names and associated abbreviations.


A gaz table has at least the following columns in the table. You may add more columns if you wish for your own purposes.


`id`
:   Primary key of table

`seq`
:   integer: definition number? - identifier used for that instance of the word

`word`
:   text: the input word

`stdword`
:   text: the standardized replacement word

`token`
:   integer: the kind of word it is. Only if it is used in this context will it be replaced. Refer to [PAGC Tokens](http://www.pagcgeo.org/docs/html/pagc-12.html#--i-tok--).
   <a id="Address_Standardizer_Functions"></a>

## Address Standardizer Functions
 <a id="debug_standardize_address"></a>

# debug_standardize_address

Returns a json formatted text listing the parse tokens and standardizations

## Synopsis


```sql
text debug_standardize_address(text  lextab, text  gaztab, text  rultab, text  micro, text  macro=NULL)
```


## Description


This is a function for debugging address standardizer rules and lex/gaz mappings. It returns a json formatted text that includes the matching rules, mapping of tokens, and best standardized address [stdaddr](#stdaddr) form of an input address utilizing [lextab](#lextab) table name, [gaztab](#gaztab), and [rulestab](#rulestab) table names and an address.


For single line addresses use just `micro`


For two line address A `micro` consisting of standard first line of postal address e.g. <code>house_num street</code>, and a macro consisting of standard postal second line of an address e.g <code>city, state postal_code country</code>.


Elements returned in the json document are


`input_tokens`
:   For each word in the input address, returns the position of the word, token categorization of the word, and the standard word it is mapped to. Note that for some input words, you might get back multiple records because some inputs can be categorized as more than one thing.

`rules`
:   The set of rules matching the input and the corresponding score for each. The first rule (highest scoring) is what is used for standardization

`stdaddr`
:   The standardized address elements [stdaddr](#stdaddr) that would be returned when running [standardize_address](#standardize_address)


Availability: 3.4.0


## Examples


Using address_standardizer_data_us extension


```sql
CREATE EXTENSION address_standardizer_data_us; -- only needs to be done once
```


Variant 1: Single line address and returning the input tokens


```sql
SELECT it->>'pos' AS position, it->>'word' AS word, it->>'stdword' AS standardized_word,
            it->>'token' AS token, it->>'token-code' AS token_code
    FROM jsonb(
            debug_standardize_address('us_lex',
                'us_gaz', 'us_rules', 'One Devonshire Place, PH 301, Boston, MA 02109')
                 ) AS s, jsonb_array_elements(s->'input_tokens') AS it;
```


```
position |    word    | standardized_word | token  | token_code
----------+------------+-------------------+--------+------------
 0        | ONE        | 1                 | NUMBER | 0
 0        | ONE        | 1                 | WORD   | 1
 1        | DEVONSHIRE | DEVONSHIRE        | WORD   | 1
 2        | PLACE      | PLACE             | TYPE   | 2
 3        | PH         | PATH              | TYPE   | 2
 3        | PH         | PENTHOUSE         | UNITT  | 17
 4        | 301        | 301               | NUMBER | 0
(7 rows)
```


Variant 2: Multi line address and returning first rule input mappings and score


```sql
SELECT (s->'rules'->0->>'score')::numeric AS score, it->>'pos' AS position,
        it->>'input-word' AS word, it->>'input-token' AS input_token, it->>'mapped-word' AS standardized_word,
            it->>'output-token' AS output_token
    FROM jsonb(
            debug_standardize_address('us_lex',
                'us_gaz', 'us_rules', 'One Devonshire Place, PH 301', 'Boston, MA 02109')
                 ) AS s, jsonb_array_elements(s->'rules'->0->'rule_tokens') AS it;
```


```
 score   | position |    word    | input_token | standardized_word | output_token
----------+----------+------------+-------------+-------------------+--------------
 0.876250 | 0        | ONE        | NUMBER      | 1                 | HOUSE
 0.876250 | 1        | DEVONSHIRE | WORD        | DEVONSHIRE        | STREET
 0.876250 | 2        | PLACE      | TYPE        | PLACE             | SUFTYP
 0.876250 | 3        | PH         | UNITT       | PENTHOUSE         | UNITT
 0.876250 | 4        | 301        | NUMBER      | 301               | UNITT
(5 rows)
```


## See Also


[stdaddr](#stdaddr), [rulestab](#rulestab), [lextab](#lextab), [gaztab](#gaztab), [Pagc_Normalize_Address](tiger-geocoder.md#Pagc_Normalize_Address)
  <a id="parse_address"></a>

# parse_address

Takes a 1 line address and breaks into parts

## Synopsis


```sql
record parse_address(text  address)
```


## Description


Returns takes an address as input, and returns a record output consisting of fields *num*, *street*, *street2*, *address1*, *city*, *state*, *zip*, *zipplus*, *country*.


Availability: 2.2.0


## Examples


Single Address


```sql
SELECT num, street, city, zip, zipplus
	FROM parse_address('1 Devonshire Place, Boston, MA 02109-1234') AS a;
```


```

 num |      street      |  city  |  zip  | zipplus
-----+------------------+--------+-------+---------
 1   | Devonshire Place | Boston | 02109 | 1234
```


Table of addresses


```
-- basic table
CREATE TABLE places(addid serial PRIMARY KEY, address text);

INSERT INTO places(address)
VALUES ('529 Main Street, Boston MA, 02129'),
 ('77 Massachusetts Avenue, Cambridge, MA 02139'),
 ('25 Wizard of Oz, Walaford, KS 99912323'),
 ('26 Capen Street, Medford, MA'),
 ('124 Mount Auburn St, Cambridge, Massachusetts 02138'),
 ('950 Main Street, Worcester, MA 01610');

 -- parse the addresses
 -- if you want all fields you can use (a).*
SELECT addid, (a).num, (a).street, (a).city, (a).state, (a).zip, (a).zipplus
FROM (SELECT addid, parse_address(address) As a
 FROM places) AS p;
```


```
 addid | num |        street        |   city    | state |  zip  | zipplus
-------+-----+----------------------+-----------+-------+-------+---------
     1 | 529 | Main Street          | Boston    | MA    | 02129 |
     2 | 77  | Massachusetts Avenue | Cambridge | MA    | 02139 |
     3 | 25  | Wizard of Oz         | Walaford  | KS    | 99912 | 323
     4 | 26  | Capen Street         | Medford   | MA    |       |
     5 | 124 | Mount Auburn St      | Cambridge | MA    | 02138 |
     6 | 950 | Main Street          | Worcester | MA    | 01610 |
(6 rows)
```


## See Also


  <a id="standardize_address"></a>

# standardize_address

Returns an stdaddr form of an input address utilizing lex, gaz, and rule tables.

## Synopsis


```sql
stdaddr standardize_address(text  lextab, text  gaztab, text  rultab, text  address)
stdaddr standardize_address(text  lextab, text  gaztab, text  rultab, text  micro, text  macro)
```


## Description


Returns an [stdaddr](#stdaddr) form of an input address utilizing [lextab](#lextab) table name, [gaztab](#gaztab), and [rulestab](#rulestab) table names and an address.


Variant 1: Takes an address as a single line.


Variant 2: Takes an address as 2 parts. A `micro` consisting of standard first line of postal address e.g. <code>house_num street</code>, and a macro consisting of standard postal second line of an address e.g <code>city, state postal_code country</code>.


Availability: 2.2.0


## Examples


Using address_standardizer_data_us extension


```sql
CREATE EXTENSION address_standardizer_data_us; -- only needs to be done once
```


Variant 1: Single line address. This doesn't work well with non-US addresses


```sql
SELECT house_num, name, suftype, city, country, state, unit  FROM standardize_address('us_lex',
			   'us_gaz', 'us_rules', 'One Devonshire Place, PH 301, Boston, MA 02109');
```


```
house_num |    name    | suftype |  city  | country |     state     |      unit
----------+------------+---------+--------+---------+---------------+-----------------
1         | DEVONSHIRE | PLACE   | BOSTON | USA     | MASSACHUSETTS | # PENTHOUSE 301
```


Using tables packaged with tiger geocoder. This example only works if you installed `postgis_tiger_geocoder`.


```sql
SELECT *  FROM standardize_address('tiger.pagc_lex',
         'tiger.pagc_gaz', 'tiger.pagc_rules', 'One Devonshire Place, PH 301, Boston, MA 02109-1234');
```


Make easier to read we'll dump output using hstore extension CREATE EXTENSION hstore; you need to install


```sql
SELECT (each(hstore(p))).*
 FROM standardize_address('tiger.pagc_lex', 'tiger.pagc_gaz',
   'tiger.pagc_rules', 'One Devonshire Place, PH 301, Boston, MA 02109') As p;
```


```
    key     |      value
------------+-----------------
 box        |
 city       | BOSTON
 name       | DEVONSHIRE
 qual       |
 unit       | # PENTHOUSE 301
 extra      |
 state      | MA
 predir     |
 sufdir     |
 country    | USA
 pretype    |
 suftype    | PL
 building   |
 postcode   | 02109
 house_num  | 1
 ruralroute |
(16 rows)

```


Variant 2: As a two part Address


```sql
SELECT (each(hstore(p))).*
 FROM standardize_address('tiger.pagc_lex', 'tiger.pagc_gaz',
   'tiger.pagc_rules', 'One Devonshire Place, PH 301', 'Boston, MA 02109, US') As p;
```


```
    key     |      value
------------+-----------------
 box        |
 city       | BOSTON
 name       | DEVONSHIRE
 qual       |
 unit       | # PENTHOUSE 301
 extra      |
 state      | MA
 predir     |
 sufdir     |
 country    | USA
 pretype    |
 suftype    | PL
 building   |
 postcode   | 02109
 house_num  | 1
 ruralroute |
(16 rows)
```


## See Also


[stdaddr](#stdaddr), [rulestab](#rulestab), [lextab](#lextab), [gaztab](#gaztab), [Pagc_Normalize_Address](tiger-geocoder.md#Pagc_Normalize_Address)
