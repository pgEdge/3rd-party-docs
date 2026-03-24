# Limitation
- The index will return up to `bm25_catalog.bm25_limit` results to PostgreSQL. Users need to adjust the `bm25_catalog.bm25_limit` for more results when using larger limit values or stricter filter conditions.
- We currently have only tested against English. Other language can be supported with bpe tokenizer with larger vocab like tiktoken out of the box. Feel free to talk to us or raise issue if you need more language support.

