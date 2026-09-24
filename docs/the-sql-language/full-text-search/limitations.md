<a id="textsearch-limitations"></a>

## Limitations


 The current limitations of PostgreSQL's text search features are:

- The length of each lexeme must be less than 2 kilobytes
- No more than 256 positions per lexeme
- Position values in `tsvector` must be greater than 0 and no more than 16,383
- The length of a `tsvector`'s data (lexemes + positions) must be less than 1 megabyte
- The length of a `tsquery`'s data (lexemes only) must be less than 1 megabyte
- The match distance in a <code><</code><em>N</em><code>></code> (FOLLOWED BY) `tsquery` operator cannot be more than 16,384


 For comparison, the PostgreSQL 8.1 documentation contained 10,441 unique words, a total of 335,420 words, and the most frequent word “postgresql” was mentioned 6,127 times in 655 documents.


 Another example — the PostgreSQL mailing list archives contained 910,989 unique words with 57,491,343 lexemes in 461,020 messages.
