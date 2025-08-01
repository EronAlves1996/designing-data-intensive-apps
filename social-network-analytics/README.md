# Social Network Analytics

In this exercises, are made some comparation of approaches for querying data.

## MapReduce implementation

One of key points on this repository is the map reduce implementation. The challenge here was to build in **Go** a data structure that will support concurrent concurrent transformation of data of one shape to another, generally associated with analytics.

The general schema is to receive a batch of data and distribute the data to concurrent processes to speed up the data processing and finally, in an unique process, to aggregate and reduce the data to the desired form.

## Queries

Queries are an abstracted way to fetch data than using imperative apis.
That way, we can use some form of DSL or descriptive language to describe what we want, and not manipulating it manually.
For querying data, we can use SQL, Cypher, mongo json.

The database vendor can optimize away the data fetching, make the data request decoupled from the data retrieving
