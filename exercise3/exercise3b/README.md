# Count the Number of Different Integers in a String

This Go package provides a function `NumDifferentIntegers` to count the number of different integers in a given string.

## Problem Statement

Given a string `word` containing alphanumeric characters, the task is to find the number of different integers present in the string. The string may contain leading zeros before an integer.

Example:
Consider the following string:

```go
word := "a1b01c001"
```

In this string, there is only one identical integer: 1. So, the `NumDifferentIntegers` function will return 1.

## Step-by-Step Explanation of `NumDifferentIntegers` Function

The `NumDifferentIntegers` function works as follows:

1. It takes a string `word` as input.

2. It initializes an empty `uniqueIntegers` map to store the unique integers encountered in the string.

3. The function iterates through each character in the string using a range loop.

4. For each character, it checks if it is a digit using the `unicode.IsDigit` function. If the character is not a digit, it replaces it with a space character. This step is done to separate the digits and non-digits.

5. The function then splits the modified string using spaces as separators using the `splitBySpace` helper function. This results in a slice of strings, where each element is a numeric token.

6. It then iterates through the tokens and removes any leading zeros from each token using the `removeLeadingZeros` helper function.

7. If the resulting token is not empty (after removing leading zeros), it is considered a valid numeric string. The function adds this numeric string as a key in the `uniqueIntegers` map to count the unique integers.

8. Finally, the function returns the length of the `uniqueIntegers` map, which represents the count of different integers in the original string.

9. The `splitBySpace` function is a helper function that splits the given string `s` using spaces as separators and returns a slice of strings containing the tokens.

10. The `removeLeadingZeros` function is another helper function that takes a string `s` and removes any leading zeros from it. If the resulting string becomes empty, it returns "0" to ensure a valid numeric string is always returned.

---

**Note:** This is a sample implementation for educational purposes. For production use, it is recommended to handle edge cases and error conditions appropriately.

Please feel free to modify and use this [Readme](README.md) file or navigate back to this [Readme](../../README.md) file to see the overview Golang Learning repository as needed.
