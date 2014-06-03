# Visualising Structure

An attempt at developing an application that is able to parse texts, identify cohesive ties within it, and then present the data collected in an undirected graph.

## Why

For my bachelor thesis I want to develop this piece of software to help me better investigate the differences in structure between learner texts and expert texts. I hypothesise that, when presented as un undirected force graph, expert texts will generally have a higher clustering coefficient, meaning they will look more dense and interconnected.

## Rough Overview

There are two main components to this application: the parsing of a text, and the displaying of the results.

#### Display
Since I intend to display the results in a graph it is probably best if I use the javascript D3 library for this, since this contains excellent methods in order to display and manipulate graphs of varying sorts and sizes. This does mean that I should present the results from the parsing process in a particular manner.

#### Parsing
The parsing of the texts is of course the most difficult step and involves several stages:  
* Identifying relevant words and keeping track of their count, and the relation between it and other relevant words.  
* Identifying overlap in relevant words: two different words can refer to the same entity. An easy example are synonyms, but there are also pronouns to take into account.

## Notes

* Currently I strip all special characters from the tokens, which will be problematic with possessives and and/or constructions.  
* I filter out common words ('a', 'the', etc) at the end. I'll want to change this when I add distance tracking.  
* I have WordNet 3.0 working on my system. I feel that this tool can end up being really useful so I am going to dedicate some time to writing a wrapper for it in go.  
  * WordNet distinguishes between nouns, verbs, adjectives, and adverbs.  
  * Within these categories it will distinguish between the different 'senses' a word can have.  
  * Finally, with the above two pieces of information WN can produce all sorts of information such as a list of synonyms, derivates, etc.
