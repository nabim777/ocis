---
title: "20. File Search Query Language"
date: 2022-03-18T09:00:00+01:00
weight: 20
geekdocRepo: https://github.com/owncloud/ocis
geekdocEditPath: edit/master/docs/ocis/adr
geekdocFilePath: 0020-file-search-query-language.md
---

* Status: proposed
* Deciders: @butonic, @micbar, @dragotin, @C0rby
* Date: 2022-03-18

## Context and Problem Statement

From the users perspective, the interface to search is just a single form field where the user enters one or more search terms. The minimum expectation is that the search returns file names and links to files that:

* have a file name that contains at least one of the search terms
* contain at least one of the search terms in the file contents
* have metadata that is equal or contains one of the search terms

## Decision Drivers

* The standard user should not be bothered by a query syntax
* The power user should also be able to narrow his search with an efficient and flexible syntax
* We need to consider different backend technologies which we need to access through an abstraction layer
* Using different indexing systems should lead to a slightly different feature set without changing the syntax completely

## Considered Options

* [Simple Query](#simplified-query)
* [Lucene Query Language](#lucene-query-language)
* [Solr Query Language](#solr-query-language)
* [Elasticsearch Query Language](#elasticsearch-query-language)

## Decision Outcome

Chosen option: [Simple Query](#simplified-query), because it is a suitable MVP.

### Positive Consequences

* We can do a quick implementation

### Negative consequences

* We cannot specify terms or other search criteria
* We will need to find a good point in time to use a query language before we start working around it

## Pros and Cons of the Options

### Simplified Query

Implement a very simple search approach: Return all files which contain at least one of the keywords in their name, path, alias or selected metadata.

* Good, because that covers 80% of the users needs
* Good, because it is very straightforward
* Good, because it is a suitable solution for GA
* Bad, because it is below the industry standard
* Bad, because it only provides one search query

### Lucene Query Language

The Lucene Query Parser syntax supports advanced queries like term, phrase, wildcard, fuzzy search, proximity search, regular expressions, boosting, boolean operators and grouping. It is a well known query syntax used by the Apache Lucene Project. Popular Platforms like Wikipedia are using Lucene or Solr, which is the successor of Lucene

* Good, because it is a well documented and powerful syntax
* Good, because it is very close to the Elasticsearch and the Solr syntax which enhances compatibility
* Bad, because there is no powerful and well tested query parser for golang available
* Bad, because it adds complexity and fulfilling all the different query use-cases can be an "uphill battle"

### Solr Query Language

Solr is highly reliable, scalable and fault-tolerant, providing distributed indexing, replication and load-balanced querying, automated failover and recovery, centralized configuration and more. Solr powers the search and navigation features of many of the world's largest internet sites.

* Good, because it is a well documented and powerful syntax
* Good, because it is very close to the Elasticsearch and the Lucene syntax which enhances compatibility
* Good, because it has a strong community with large resources and knowledge
* Bad, because it adds complexity and fulfilling all the different query use-cases can be an "uphill battle"

### Elasticsearch Query Language

Elasticsearch provides a full Query DSL (Domain Specific Language) based on JSON to define queries. Think of the Query DSL as an AST (Abstract Syntax Tree) of queries, consisting of two types of clauses. It is able to combine multiple query types into compound queries. It is also a successor of Solr.

* Good, because it is a well documented and powerful syntax
* Good, because it is very close to the Elasticsearch and the Solr syntax which enhances compatibility
* Good, because there is a stable and well tested go client which brings a query builder
* Good, because it could be used as the query language which supports different search backends by just implementing what is needed for our use-case
* Bad, because it adds complexity and fulfilling all the different query use-cases can be an "uphill battle"

## Links

* [Search API](0018-file-search-api.md)
* [Search Indexing](0019-file-search-index.md)
* [Apache Lucene](https://lucene.apache.org/)
* [Apache Solr](https://solr.apache.org/)
* [Elastic Search](https://solr.apache.org/)
* [Elastic Search for go](https://github.com/elastic/go-elasticsearch)
