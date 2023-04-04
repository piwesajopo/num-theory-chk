# num-theory-chk
A number properties checker based on easy to understand number theory algorithms

## Introduction

I recently became insterested in number theory, mainly because what I first saw as a funny and aesthetic way to mathematically represent numbers turned out to have practical applications. I soon realized this was a part of arithmetic that is not covered deeply on schools as it has to do with arithmetic functions.

Hence came the idea of doing this repository, it serves several purposes:

* Provides the code for a fun program that tells you the properties of a given number.
* Teachs you about the concepts behind these properties (more documentation comming later).
* Have easy to understand short functions that can be used to teach children and young people about programming logic and math.
* Has some interesting extra programming easter eggs for beginners that were not really needed like:
  - How to properly handle named parameters before and after non-named ones.
  - Passing functions to be used inside other functions.
  - Optimized Integer versions of some math operations like IntPower and *Interger Powers* and *Integer Square Roots* based on bitwise operations.

## Checks Performed

- Happy Number
- Self Number
- Odious Number
- Perfect Number
- Abundant Number
- Deficient Number
- Perfect Square Number
- Prime Number
- Fibonacci Number
- Narcissistic Number
- Münchhausen Number
- Factorial Number
<br>

## Pending Checks

- Semiperfect Number
<br>

## Some Definitions

#### Number Theory Numbners

- **Fibonacci Number:** A quick method to check if a number if Fibonacci number or not, is as below: N is a Fibonacci number if and only if `( 5*N^2 + 4 )` or `( 5*N^2 – 4 )` is a perfect square! For Example: 3 is a Fibonacci number since `(5*3*3 + 4)` is `49` which is `7*7`.

- **Münchhausen Number:** A Münchhausen number (sometimes spelled Münchausen number, with a single 'h') is a number equal to the sum of its digits raised to each digit's power. If 0s are disallowed (since 0^0 is not well-defined), the only Münchhausen numbers are 1 and `3435 = 3³ + 4⁴ + 3³ + 5⁵`. If the definition `0^0=0` is adopted, then there are exactly four Münchhausen numbers: `0`, `1`, `3435`, and `438579088`

- **Narcissistic or Armstrong number:** An n-digit number that is the sum of the  nth powers of its digits is called an n-narcissistic number.

- **Semiperfect Number:** In number theory, a semiperfect number or pseudoperfect number is a natural number n that is equal to the sum of all or some of its proper divisors. A semiperfect number that is equal to the sum of all its proper divisors is a perfect number.

- **Odious Number:** An odious number is a nonnegative number that has an odd number of 1s in its binary expansion.


#### Combinatorial Mathermatics Numbers

- Bell Number
- Catalan Number

## Some examples

- Narcissistic:
```
153 = 1³ + 5³ + 3³

1634 = 1⁴ + 6⁴ + 3⁴ + 4⁴

54748 = 5⁵ + 4⁵ + 7⁵ + 4⁵ + 8⁵

548834 = 5⁶ + 4⁶ + 8⁶ + 8⁶ + 3⁶ + 4⁶ (notable)

1741725 = 1⁷ + 7⁷ + 4⁷ + 1⁷ + 7⁷ + 2⁷ + 5⁷
```

- Münchhausen:
```
3435 = 3³ + 4⁴ + 3³ + 5⁵
```

