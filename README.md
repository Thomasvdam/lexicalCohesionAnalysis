# Visualising Lexical Cohesion over Token Time

An attempt at developing an application that is able to parse texts, identify cohesive ties within it, and then present the data collected in a visual manner.

## Why

For my bachelor thesis I want to develop this piece of software to help me better investigate the differences in structure between learner texts and expert texts. I hypothesise that learner texts will show a smaller degree of semantic cohesion when compared to expert texts. In order to test this I make use of WordNet to gather information about the words in the text which I then transform into a score of sorts which tells something about the amount of semantic overlap between lexical tokens.

## Rough Overview

There are two main components to this application: the parsing of a text, and the displaying of the results.

#### Display
Since I intend to display the results in a visual manner it is probably best if I use the javascript D3 library for this, since this contains excellent methods in order to display and manipulate graphs of varying sorts and sizes. This does mean that I should present the results from the parsing process in a JSON.

#### Parsing
The parsing of the texts is of course the most difficult step and involves several stages:  
* Identifying lexical tokens and the relation between them within a particular frame of observation.  

## Notes

The entirety of the WordNet interface is terrible since I spawn shell process for every query. If I ever want to improve this program I'll have to write a wrapper for the native C code, which should speed up the process significantly.

* Currently I strip all special characters from the tokens, which will be problematic with possessives and and/or constructions.
* I decided to drop all wordTypes other than nouns, since nouns are the easiest to classify and determine 'similarity' between. This change goes hand in hand with the decision to use the ontological tree to help determine how related 2 word senses are. It boils down to how far removed from each other the two senses are in the tree.  
  * This decision is partly due to the fact that it significantly increases performance.
* It might be worthwhile to drop words that have a low polysemy count as nouns altogether. At the moment I have the threshold set to 3, but this can be easily changed.
* The ontological tree is now constructed based on the first path that WordNet displays, mainly because this saves a lot of time and from what I have seen it should not make too big of a difference.

## To Do

1. [ ] Create a functional core program that can parse a text and produces some kind of result.  
  - [ ] Make an accessible interface to WordNet, albeit not very efficient.  
    - [x] Create a function that tokenises words.
      - [x] Places all the senses in an ontology tree which is constructed simultaneously.
    - [ ] Create a function that compares two senses based on their ontological paths.
    - [ ] From there decide which sense is most likely. ???
  - [ ] Write the main loop and source text cleaning.
  - [ ] Determine how to score the different types of relations lexical tokens can have.  
2. [ ] Package the results in a nice JSON.
3. [ ] Write a small JS script that displays the results.
