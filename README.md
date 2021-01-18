# esfilter

esfilter is the json generater for Elasticsearch.

Elasticsearch's filter condition has ...

- must ... The clause (query) must appear in matching documents. However unlike must the score of the query will be ignored. Filter clauses are executed in filter context, meaning that scoring is ignored and clauses are considered for caching.
- should ... The clause (query) should appear in the matching document. If the bool query is in a query context and has a must or filter clause then a document will match the bool query even if none of the should queries match. In this case these clauses are only used to influence the score. If the bool query is a filter context or has neither must or filter then at least one of the should queries must match a document for it to match the bool query. This behavior may be explicitly controlled by settings the minimum_should_match parameter.
- must_not ... The clause (query) must not appear in the matching documents. Clauses are executed in filter context meaning that scoring is ignored and clauses are considered for caching. Because scoring is ignored, a score of 0 for all documents is returned.

There filter condition corresponds to `AND` (must), `OR` (should), `NOT`(must_not).

So, we have to use filter other logical fomula (xor, nor, nand or more...) need to combine those operators to create other operators.

# operator

input two value (A, B)

- NAND ... `(not A) OR (not B)`
- NOR ... `(not A) AND (not B)`
- XOR ... `(A and (not B)) or ((not A) and B)`
- XNOR ... `(A and B) or ((not A) and (not B))`


