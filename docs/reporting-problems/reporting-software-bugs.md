## Reporting Software Bugs


Reporting bugs effectively is a fundamental way to help PostGIS development. The most effective bug report is that enabling PostGIS developers to reproduce it, so it would ideally contain a script triggering it and every information regarding the environment in which it was detected. Good enough info can be extracted running <code>SELECT postgis_full_version()</code> [for PostGIS] and <code>SELECT version()</code> [for postgresql].


If you aren't using the latest release, it's worth taking a look at its [release changelog](https://git.osgeo.org/gitea/postgis/postgis/raw/branch/master/NEWS) first, to find out if your bug has already been fixed.


Using the [PostGIS bug tracker](http://trac.osgeo.org/postgis/) will ensure your reports are not discarded, and will keep you informed on its handling process. Before reporting a new bug please query the database to see if it is a known one, and if it is please add any new information you have about it.


You might want to read Simon Tatham's paper about [How to Report Bugs Effectively](http://www.chiark.greenend.org.uk/~sgtatham/bugs.html) before filing a new report.
