<a id="planner-stats-details"></a>

# How the Planner Uses Statistics

 This chapter builds on the material covered in [Using `EXPLAIN`](../../the-sql-language/performance-tips/using-explain.md#using-explain) and [Statistics Used by the Planner](../../the-sql-language/performance-tips/statistics-used-by-the-planner.md#planner-stats) to show some additional details about how the planner uses the system statistics to estimate the number of rows each part of a query might return. This is a significant part of the planning process, providing much of the raw material for cost calculation.

 The intent of this chapter is not to document the code in detail, but to present an overview of how it works. This will perhaps ease the learning curve for someone who subsequently wishes to read the code.

- [Row Estimation Examples](row-estimation-examples.md#row-estimation-examples)
- [Multivariate Statistics Examples](multivariate-statistics-examples.md#multivariate-statistics-examples)
- [Planner Statistics and Security](planner-statistics-and-security.md#planner-stats-security)
