# Kisumu Programming Language Specification

## 6. Control Structures

Control structures allow developers to dictate the flow of execution in a program based on conditions, iterations, or other mechanisms. Kisumu Lang offers a range of flow-control constructs for efficient and readable code.

### 6.1 Conditional Statements

Conditional statements enable decision-making in programs. Kisumu provides `if`, `else if`, and `else` constructs for evaluating conditions.

#### Syntax

```go
if (condition) {
    // Code to execute if condition is true
} else if (anotherCondition) {
    // Code to execute if anotherCondition is true
} else {
    // Code to execute if none of the above conditions are true
}
```

#### Example

```go
int age = 20;

if (age >= 18) {
    print("You are an adult.");
} else {
    print("You are a minor.");
}
```

### 6.2 Loops

Loops are used to execute a block of code repeatedly based on a condition or iteration.

#### 6.2.1 `for` Loop

The `for` loop is a versatile construct for iteration.

##### Syntax

```go
for (initialization; condition; update) {
    // Code to execute repeatedly
}
```

##### Example

```go
for (int i = 0; i < 5; i++) {
    print(i);
}
```

#### 6.2.2 `while` Loop

The `while` loop executes a block of code as long as the condition evaluates to true.

##### Syntax

```go
while (condition) {
    // Code to execute repeatedly
}
```

##### Example

```go
int count = 0;

while (count < 5) {
    print(count);
    count++;
}
```

#### 6.2.3 `do-while` Loop

The `do-while` loop ensures the block of code executes at least once before evaluating the condition.

##### Syntax

```go
do {
    // Code to execute at least once
} while (condition);
```

##### Example

```go
int count = 0;

do {
    print(count);
    count++;
} while (count < 5);
```

### 6.3 Switch Statements

The `switch` statement evaluates a variable or expression and executes the matching case.

#### Syntax

```go
switch (expression) {
    case value1:
        // Code for case value1
        break;
    case value2:
        // Code for case value2
        break;
    default:
        // Code if no case matches
}
```

#### Example

```go
int day = 3;

switch (day) {
    case 1:
        print("Monday");
        break;
    case 2:
        print("Tuesday");
        break;
    case 3:
        print("Wednesday");
        break;
    default:
        print("Invalid day");
}
```

### 6.4 Control Keywords

Kisumu provides keywords for altering the flow of execution within loops and conditionals:

- **`break`**: Exits the nearest enclosing loop or `switch`.
- **`continue`**: Skips the remaining code in the current loop iteration and proceeds to the next iteration.
- **`return`**: Exits from the current function and optionally returns a value.

#### Examples

```go
// Using break
for (int i = 0; i < 10; i++) {
    if (i == 5) {
        break; // Exit loop when i is 5
    }
    print(i);
}

// Using continue
for (int i = 0; i < 10; i++) {
    if (i % 2 == 0) {
        continue; // Skip even numbers
    }
    print(i);
}

// Using return
int square(int num) {
    return num * num;
}
```

### 6.5 Nested Control Structures

Kisumu allows nesting of loops, conditionals, and other control structures for more complex logic.

#### Example

```go
for (int i = 1; i <= 3; i++) {
    print("Outer loop iteration:", i);

    for (int j = 1; j <= 2; j++) {
        print("  Inner loop iteration:", j);
    }
}
```

### 6.6 Error Handling in Control Structures

Kisumu Lang is designed to handle runtime errors gracefully, especially within control structures. The language will introduce structured exception handling in a future update.
