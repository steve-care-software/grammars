# Grammar
The grammar is a schema that makes it easy and visual to generate data (using compose elements) or combine arbitrary data to the grammar to create an AST and therefore validate the composition of the data and order it in a logical way.

The following section explains:
* The composition of each part of the grammar structure
* How to assign a grammar instance to a variable
* How to compose an AST by combining arbitrary data with a grammar instance
* How to save a grammar instance into the database
* How to retrieve a grammar instance from the database
* How to execute the token tests of a grammar instance and receive its results and coverage reports

## Variable
The first letter of a variable must be a lower-case letter, then can be followed by any letter, lower-case or upper-case.  A variable can contain only one (1) letter.
```
// valid one-letter variable:
a

// valid variable where the second letter is upper-case:
aVariable

// valid variable where the second letter is lower-case:
variable

```
## Value
The "value" is represented by an unsigned integer and therefore must be a number between 0 and 255 (range: [0,255]).  It must always be assigned to a variable in order to be used in the composition of tokens.

There is no test suites for values.
```
// valid assignment of the value '122' to a variable:
myVariable: 122;

```

## Compose
The "compose" is a series of "value" and other "compose" elements exclusively.  A "compose" always contains a specific series of bytes.  Each element of its composition is repeated only once if there is no occurence following the element.  If there is an "occurence" value, the element is repeated by this amount of time.  The "occurence" is an unsigned integer.

Please note each element of a "compose" is created using a variable name, a pipe (|) and an occurence.  The pipe (|) only acts as a syntaxic separator.  If the element is only repeated once, the occurence one (1) and its pipe (|) is optional.

There is no test suites for "compose".
```
// compose the world "actually"
wordActually: letterA letterC letterT letterU letterA letterL|2 letterY|1;

// values, where each byte is a ASCII representation of a letter:
letterA: 97;
letterC: 99;
letterT: 116;
letterU: 117;
letterA: 97;
letterL: 108;
letterY: 108;

```

## Test suites
## Cardinality
## Token
## Everything
## External Grammar
## Channels
## Root
