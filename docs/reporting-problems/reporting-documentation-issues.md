## Reporting Documentation Issues


The documentation should accurately reflect the features and behavior of the software. If it doesn't, it could be because of a software bug or because the documentation is in error or deficient.


Documentation issues can also be reported to the [PostGIS bug tracker](http://trac.osgeo.org/postgis).


If your revision is trivial, just describe it in a new bug tracker issue, being specific about its location in the documentation.


If your changes are more extensive, a patch is definitely preferred. This is a four step process on Unix (assuming you already have [git](https://git-scm.com/) installed):


1. Clone the PostGIS' git repository. On Unix, type:

`git clone https://git.osgeo.org/gitea/postgis/postgis.git`

   This will be stored in the directory postgis
2. Make your changes to the documentation with your favorite text editor. On Unix, type (for example):

`vim doc/postgis.xml`

   Note that the documentation is written in DocBook XML rather than HTML, so if you are not familiar with it please follow the example of the rest of the documentation.
3. Make a patch file containing the differences from the master copy of the documentation. On Unix, type:

`git diff doc/postgis.xml > doc.patch`
4. Attach the patch to a new issue in bug tracker.
