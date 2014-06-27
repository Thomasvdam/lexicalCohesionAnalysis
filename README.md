# Visualising Lexical Cohesion over Token Time

For my bachelor thesis I developed an application that is able to parse texts, identify cohesive ties within it, and then present the data collected visually.

## Disclaimer

I realise this code is far from perfect and needs significant refactoring. It is also flawed in the sense that it does not interact with the database directly, but rather through the CLI, which is of course highly detrimental to the performance of the program. However, for the purpose of my thesis the current functionality and runtime was sufficient, which is why I did not bother with all the defects.
Should I ever revisit this project I will try to address the most glaring issues, but for now I'm open sourcing this terrible mess because open source is the way to go. :D

## Notes

* Currently I strip all special characters from the tokens, which will be problematic with possessives and 'and/or' constructions.
* I decided to drop all wordTypes other than nouns, since nouns are the easiest to classify and determine 'similarity' between. This can be addressed by creating a global root node between the different POS trees in WordNet.
  * This decision is partly due to the fact that it significantly increases performance.
* The ontological tree is now constructed based on the first path that WordNet displays, mainly because this saves a lot of time and from what I have seen it should not make too big of a difference.
