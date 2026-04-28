# Usage

The extension is mainly composed by three parts, tokenizer, bm25vector and bm25vector index. The tokenizer is used to convert the text into a bm25vector, and the bm25vector is similar to a sparse vector, which stores the vocabulary id and frequency. The bm25vector index is used to speed up the search and ranking process.

To tokenize a text, you can use the `tokenize` function. The `tokenize` function takes two arguments, the text to tokenize and the tokenizer name. 


!!! note
    Tokenizer part is completed by a separate extension [pg_tokenizer.rs](https://github.com/tensorchord/pg_tokenizer.rs), more details can be found [here](https://github.com/tensorchord/pg_tokenizer.rs/tree/main/docs).


```sql
-- create a tokenizer
SELECT create_tokenizer('bert', $$
model = "bert_base_uncased"  # using pre-trained model
$$);
-- tokenize text with bert tokenizer
SELECT tokenize('A quick brown fox jumps over the lazy dog.', 'bert')::bm25vector;
-- Output: {1012:1, 1037:1, 1996:1, 2058:1, 2829:1, 3899:1, 4248:1, 4419:1, 13971:1, 14523:1}
-- The output is a bm25vector, 1012:1 means the word with id 1012 appears once in the text.
```

One thing special about bm25 score is that it depends on a global document frequency, which means the score of a word in a document depends on the frequency of the word in all documents. To calculate the bm25 score between a bm25vector and a query, you need had a document set first and then use the `<&>` operator.

```sql
-- Setup the document table
CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    passage TEXT,
    embedding bm25vector
);

INSERT INTO documents (passage) VALUES
('PostgreSQL is a powerful, open-source object-relational database system. It has over 15 years of active development.'),
('Full-text search is a technique for searching in plain-text documents or textual database fields. PostgreSQL supports this with tsvector.'),
('BM25 is a ranking function used by search engines to estimate the relevance of documents to a given search query.'),
('PostgreSQL provides many advanced features like full-text search, window functions, and more.'),
('Search and ranking in databases are important in building effective information retrieval systems.'),
('The BM25 ranking algorithm is derived from the probabilistic retrieval framework.'),
('Full-text search indexes documents to allow fast text queries. PostgreSQL supports this through its GIN and GiST indexes.'),
('The PostgreSQL community is active and regularly improves the database system.'),
('Relational databases such as PostgreSQL can handle both structured and unstructured data.'),
('Effective search ranking algorithms, such as BM25, improve search results by understanding relevance.');
```

Then tokenize it 

```sql
UPDATE documents SET embedding = tokenize(passage, 'bert');
```

Create the index on the bm25vector column so that we can collect the global document frequency.

```sql
CREATE INDEX documents_embedding_bm25 ON documents USING bm25 (embedding bm25_ops);
```

Now we can calculate the BM25 score between the query and the vectors. Note that the BM25 score in VectorChord-BM25 is negative, which means the more negative the score, the more relevant the document is. We intentionally make it negative so that you can use the default order by to get the most relevant documents first.

```sql
-- bm25query(index_name, query, tokenizer_name)
-- <&> is the operator to compute the bm25 score
SELECT id, passage, embedding <&> bm25query('documents_embedding_bm25', tokenize('PostgreSQL', 'bert')) AS bm25_score FROM documents;
```

And you can use the order by to utilize the index to get the most relevant documents first and faster.
```sql
SELECT id, passage, embedding <&> bm25query('documents_embedding_bm25', tokenize('PostgreSQL', 'bert')) AS rank
FROM documents
ORDER BY rank
LIMIT 10;
```

