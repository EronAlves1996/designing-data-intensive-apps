# Library Inventory System

This is a code kata comparing the query data model between Postgresql and Mongodb

## The proposal

Mongodb offers a non-relational schema based on documents, where I found that modelling some relationships can be easier, and simple than relational models, by nesting some data and taking some advantage of denormalization and embedding data. You can compare the complexity of the two tasks by inspecting the files `dune_book.sql` and `create_book.js`.

By creating the same entity on the two databases, you find that the document model is more simple.

## Querying

Querying can be more challenging in document model than in relational model. Maybe I experienced this because I am more used to SQL than mongodb syntax, but in fact, for only one document, querying in relational model is faster than querying in document model.

In mongo, to emulate joins or perform some aggregations, we need to use aggregate pipelines. In sql, we only use the good old sql syntax for perform all these things.

## Usages

SQL will be present in every system you use.
Mongo will be present in a small portion, where the data needs are more stratosferic.
