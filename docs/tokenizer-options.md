# Tokenizer Options

Tokenizers can be configured in two primary ways:

- Pre-Trained Models: Suitable for most standard use cases, these models are efficient and require minimal setup. They are ideal for general-purpose applications where the text aligns with the model's training data.
- Custom Models: Offer flexibility and superior accuracy for specialized texts. These models are trained on specific corpora, making them suitable for domains with unique terminology, such as technical fields or industry-specific jargon.

Usage Details can be found in [pg_tokenizer doc](https://github.com/tensorchord/pg_tokenizer.rs/blob/main/docs/04-usage.md)

## Key Considerations

1. Language and Script:
- **Space-Separated Languages** (e.g., English, Spanish, German): Simple tokenizers such as `bert` (for English) or `unicode` tokenizers are effective here.
- **Non-Space-Separated Languages** (e.g., Chinese, Japanese): These require specialized algorithms (pre-tokenizer) that understand language structure beyond simple spaces. You can refer [Chinese](more-examples.md#using-jieba-for-chinese-text) and [Japanese](more-examples.md#using-lindera-for-japanese-text) example.
- **Multilingual Data**: Handling multiple languages within a single index requires tokenizers designed for multilingual support, such as `gemma2b` or `llmlingua2`, which efficiently manage diverse scripts and languages.

2. Vocabulary Complexity:
- **Standard Language**: For texts with common vocabulary, pre-trained models are sufficient. They handle everyday language efficiently without requiring extensive customization.
- **Specialized Texts**: Technical terms, abbreviations (e.g., "k8s" for Kubernetes), or compound nouns may need custom models. Custom models can be trained to recognize domain-specific terms, ensuring accurate tokenization. Custom synonyms may also be necessary for precise results. See [custom model](more-examples.md#using-custom-model) example.

## Preload (for performance)

For each connection, Postgresql will load the model at the first time you use it. This may cause a delay for the first query. You can use the `add_preload_model` function to preload the model at the server startup.

```sh
psql -c "SELECT add_preload_model('model1')"
# restart the PostgreSQL to take effects
sudo docker restart container_name         # for pg_tokenizer running with docker
sudo systemctl restart postgresql.service  # for pg_tokenizer running with systemd
```

The default preload model is `llmlingua2`. You can change it by using `add_preload_model`, `remove_preload_model` functions.

> Note: The pre-trained model may take a lot of memory (100MB for gemma2b, 200MB for llmlingua2). If you have a lot of models, you should consider the memory usage when you preload the model.

<!-- ## Performance Benchmark

We used datasets are from [xhluca/bm25-benchmarks](https://github.com/xhluca/bm25-benchmarks) and compare the results with ElasticSearch and Lucene. The QPS reflects the query efficiency with the index structure. And the NDCG@10 reflects the ranking quality of the search engine, which is totally based on the tokenizer. This means we can achieve the same ranking quality as ElasticSearch and Lucene if using the exact same tokenizer. 

## QPS Result

| Dataset          | VectorChord-BM25 | ElasticSearch |
| ---------------- | ---------------- | ------------- |
| trec-covid       | 28.38            | 27.31         |
| webis-touche2020 | 38.57            | 32.05         |

## NDCG@10 Result

| Dataset          | VectorChord-BM25 | ElasticSearch | Lucene |
| ---------------- | ---------------- | ------------- | ------ |
| trec-covid       | 67.67            | 68.80         | 61.0   |
| webis-touche2020 | 31.0             | 34.70         | 33.2   |

