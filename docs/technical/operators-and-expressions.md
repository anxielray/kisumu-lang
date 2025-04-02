# Kisumu Programming Language Specification

## 5. Operators and Expressions

Operators and expressions in Kisumu Lang allow developers to perform various computations and logical evaluations. This section outlines the supported operators, their behavior, and the precedence rules governing their execution.

### 5.1 Overview of Operators

Kisumu supports the following categories of operators:

1. **Arithmetic Operators**  
   Perform mathematical operations on numerical values.
   
   | Operator | Description       | Example      | Result |
   |----------|-------------------|--------------|--------|
   | `+`      | Addition          | `5 + 3`      | `8`    |
   | `-`      | Subtraction       | `5 - 3`      | `2`    |
   | `*`      | Multiplication    | `5 * 3`      | `15`   |
   | `/`      | Division          | `6 / 2`      | `3`    |
   | `%`      | Modulus (remainder)| `5 % 2`     | `1`    |

2. **Relational Operators**  
   Compare two values and return a Boolean result.
   
   | Operator | Description         | Example     | Result |
   |----------|---------------------|-------------|--------|
   | `==`     | Equal to            | `5 == 5`    | `true` |
   | `!=`     | Not equal to        | `5 != 3`    | `true` |
   | `<`      | Less than           | `3 < 5`     | `true` |
   | `>`      | Greater than        | `5 > 3`     | `true` |
   | `<=`     | Less than or equal  | `3 <= 5`    | `true` |
   | `>=`     | Greater than or equal | `5 >= 5` | `true` |

3. **Logical Operators**  
   Perform Boolean logic operations.
   
   | Operator | Description           | Example         | Result |
   |----------|-----------------------|-----------------|--------|
   | `&&`     | Logical AND           | `true && false` | `false`|
   | `||`     | Logical OR            | `true || false` | `true` |
   | `!`      | Logical NOT           | `!true`         | `false`|

4. **Assignment Operators**  
   Assign values to variables.
   
   | Operator | Description         | Example      | Result |
   |----------|---------------------|--------------|--------|
   | `=`      | Assign              | `a = 5`      | `a = 5`|
   | `+=`     | Add and assign       | `a += 3`     | `a = 8`|
   | `-=`     | Subtract and assign  | `a -= 2`     | `a = 6`|
   | `*=`     | Multiply and assign  | `a *= 2`     | `a = 10`|
   | `/=`     | Divide and assign    | `a /= 2`     | `a = 5`|

5. **Bitwise Operators**  
   Perform operations at the binary level.
   
   | Operator | Description   | Example      | Result |
   |----------|---------------|--------------|--------|
   | `&`      | Bitwise AND   | `5 & 3`      | `1`    |
   | `|`      | Bitwise OR    | `5 | 3`      | `7`    |
   | `^`      | Bitwise XOR   | `5 ^ 3`      | `6`    |
   | `~`      | Bitwise NOT   | `~5`         | `-6`   |
   | `<<`     | Left shift    | `5 << 1`     | `10`   |
   | `>>`     | Right shift   | `5 >> 1`     | `2`    |

6. **Other Operators**
   Additional operators for specialized purposes.
   
   | Operator | Description       | Example       |
   |----------|-------------------|---------------|
   | `.`      | Member access     | `obj.prop`    |
   | `[]`     | Indexing          | `arr[0]`      |
   | `()`     | Function call     | `func(5)`     |

### 5.2 Operator Precedence and Associativity

Operators in Kisumu Lang have a defined precedence, which determines the order in which they are evaluated. Associativity specifies the direction of evaluation when operators have the same precedence.

| Precedence Level | Operators           | Associativity |
|-------------------|---------------------|---------------|
| 1 (Highest)       | `()` `[]` `.`       | Left-to-right |
| 2                 | `!` `~`             | Right-to-left |
| 3                 | `*` `/` `%`         | Left-to-right |
| 4                 | `+` `-`             | Left-to-right |
| 5                 | `<<` `>>`           | Left-to-right |
| 6                 | `<` `<=` `>` `>=`   | Left-to-right |
| 7                 | `==` `!=`           | Left-to-right |
| 8                 | `&`                 | Left-to-right |
| 9                 | `^`                 | Left-to-right |
| 10                | `|`                 | Left-to-right |
| 11                | `&&`                | Left-to-right |
| 12                | `||`                | Left-to-right |
| 13 (Lowest)       | `=` `+=` `-=` `*=` `/=` | Right-to-left |

### 5.3 Expressions

Expressions combine variables, literals, and operators to produce a value. Examples:

```k
// Arithmetic expression
int sum = 5 + 3 * 2

// Relational expression
bool isAdult = age >= 18

// Logical expression
bool canVote = isAdult && hasID
```
