# Reference

## Data Types

- `bm25vector`: A specialized vector type for storing BM25 tokenized text. Structured as a sparse vector, it stores token IDs and their corresponding frequencies. For example, `{1:2, 2:1}` indicates that token ID 1 appears twice and token ID 2 appears once in the document.
- `bm25query`: A query type for BM25 ranking.

## Functions

- `bm25query(regclass, bm25vector) RETURNS bm25query`: Convert the input text into a BM25 query.

## Operators

- `bm25vector = bm25vector RETURNS boolean`: Check if two BM25 vectors are equal.
- `bm25vector <> bm25vector RETURNS boolean`: Check if two BM25 vectors are not equal.
- `bm25vector <&> bm25query RETURNS float4`: Calculate the **negative** BM25 score between the BM25 vector and query. The lower the score, the more relevant the document is. (This is intentionally designed to be negative for easier sorting.)

## Casts

- `int[]::bm25vector (implicit)`: Cast an integer array to a BM25 vector. The integer array represents token IDs, and the cast aggregates duplicates into frequencies, ignoring token order. For example, `{1, 2, 1}` will be cast to `{1:2, 2:1}` (token ID 1 appears twice, token ID 2 appears once).

## GUCs

- `bm25_catalog.bm25_limit (integer)`: The maximum number of documents to return in a search. Default is 100, minimum is -1, and maximum is 65535. When set to -1, it will perform brute force search and return all documents with scores greater than 0.
- `bm25_catalog.enable_index (boolean)`: Whether to enable the bm25 index. Default is true.
- `bm25_catalog.segment_growing_max_page_size (integer)`: The maximum page count of the growing segment. When the size of the growing segment exceeds this value, the segment will be sealed into a read-only segment. Default is 4,096, minimum is 1, and maximum is 1,000,000.

