<a id="mvcc"></a>

# Concurrency Control

 This chapter describes the behavior of the PostgreSQL database system when two or more sessions try to access the same data at the same time. The goals in that situation are to allow efficient access for all sessions while maintaining strict data integrity. Every developer of database applications should be familiar with the topics covered in this chapter.

- [Introduction](introduction.md#mvcc-intro)
- [Transaction Isolation](transaction-isolation.md#transaction-iso)
- [Explicit Locking](explicit-locking.md#explicit-locking)
- [Data Consistency Checks at the Application Level](data-consistency-checks-at-the-application-level.md#applevel-consistency)
- [Serialization Failure Handling](serialization-failure-handling.md#mvcc-serialization-failure-handling)
- [Caveats](caveats.md#mvcc-caveats)
- [Locking and Indexes](locking-and-indexes.md#locking-indexes)
