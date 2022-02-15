# String Calculator Kata
Based on the original: https://osherove.com/tdd-kata-1

## Before you start: 
- Try not to read ahead.
- Do one task at a time. The trick is to learn to work incrementally.
- Make sure you only test for correct inputs. There is no need to test for invalid inputs for this kata

# Kata:

## 1. Create a simple String calculator with a method signature:

In Java/C#/Kotlin like languages
```
int Add(string numbers)
```

In Go:
```
fun Add(numbers string) : result, err
```

The method can take up to two numbers, separated by commas, and will return their sum. 
for example `“”` or `“1”` or `“1,2”` as inputs.

(for an empty string it will return 0) 

### Hints

 - Start with the simplest test case of an empty string and move to one and two numbers
 - Remember to solve things as simply as possible so that you force yourself to write tests you did not think about
 - Remember to refactor after each passing test

## 2. Allow the Add method to handle an unknown amount of numbers

Allow the Add method to handle new lines between numbers (instead of commas).

the following input is ok: `“1\n2,3”` (will equal 6)

the following input is NOT ok: `“1,\n”` (not need to prove it - just clarifying)


## 3. Support different delimiters
To change a delimiter, the beginning of the string will contain a separate line that looks like this: `“//[delimiter]\n[numbers…]”` for example `“//;\n1;2”` should return three where the default delimiter is `‘;’` .

The first line is optional. All existing scenarios should still be supported.


## 4. Calling Add with a negative number will throw an exception “negatives not allowed” - and the negative that was passed. 

If there are multiple negatives, show *all* of them in the exception message.

## 5. Numbers bigger than 1000 should be ignored

Example: adding `2 + 1001 = 2`

## 6. Delimiters can be of any length with the following format: `“//[delimiter]\n”` 

Example: `“//[***]\n1***2***3”` should return 6


## 7. Allow multiple delimiters like this: `“//[delim1][delim2]\n”`

Example `“//[*][%]\n1*2%3”` should return `6`.

## 8. make sure you can also handle multiple delimiters with length longer than one char