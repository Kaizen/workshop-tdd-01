# Word Wrap Kata
From: https://codingdojo.org/kata/WordWrap/

# Problem Description

You write a class called `Wrapper`, that has a *single static function* named `wrap` that takes two arguments: a `string`, and a `column number`. 

The function returns the `string`, but with line breaks inserted at just the right places to make sure that no line is longer than the column number. You try to break lines at word boundaries.

Like a word processor, break the line by replacing the last space in a line with a newline.

## Example
Given: `column number=10` and `input=Lorem ipsum dolor sit amet consectetur adipiscing elit.`, result should be:

```
Lorem\n 
ipsum\n
dolor sit\n
amet\n 
consectetu\n
r\n
adipiscing\n 
elit.
```
