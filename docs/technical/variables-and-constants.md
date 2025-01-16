# Kisumu Programming Language Specification

## 4. Variables and Constants

Variables and constants are essential components of Kisumu, providing mechanisms to store and manage data throughout the program's lifecycle. This section defines the rules for declaring, assigning, and using variables and constants in Kisumu.

### 4.1 Variables

Variables in Kisumu are used to store data that can change during program execution. All variables must be explicitly declared with a type before use.

#### 4.1.1 Variable Declaration

**Syntax:**

```ksm
type variableName
```

**Examples:**

```ksm
int age
float temperature
string name
```

#### 4.1.2 Variable Initialization

Variables can be assigned a value at the time of declaration or after declaration.

**Syntax:**

```ksm
type variableName = value
variableName = newValue
```

**Examples:**

```ksm
int age = 25
age = 30
string greeting = "Hello, Kisumu!"
```

#### 4.1.3 Variable Scope

Kisumu supports block-level scoping for variables, meaning variables are only accessible within the block in which they are declared.

**Example:**

```ksm
if (true) {
    int localVariable = 10
    print(localVariable) // Accessible here
}
// print(localVariable) // Error: Not accessible outside the block
```

### 4.2 Constants

Constants store immutable data, ensuring that their values cannot be changed after initialization.

#### 4.2.1 Constant Declaration

Constants must be declared with an initial value and cannot be reassigned.

**Syntax:**

```ksm
const type constantName = value
```

**Examples:**

```ksm
const int MAX_USERS = 100
const string GREETING = "Welcome to Kisumu!"
```

#### 4.2.2 Benefits of Constants

1. **Immutability**: Prevents accidental changes to critical values.
2. **Clarity**: Indicates that the value is fixed and non-modifiable.
3. **Optimization**: Allows the compiler to optimize constant usage for better performance.

### 4.3 Type Inference (Optional)

While Kisumu enforces explicit type declaration, it optionally supports type inference for variables where the type can be deduced from the assigned value.

**Syntax:**

```ksm
var variableName = value
```

**Example:**

```ksm
var count = 10 // Inferred as int
var message = "Hello" // Inferred as string
```

### 4.4 Best Practices

1. **Use Constants for Fixed Values**: Replace magic numbers or strings with constants to enhance readability.
2. **Declare Variables Close to Usage**: Minimize variable scope to improve code clarity.
3. **Name Variables Clearly**: Use descriptive names that convey the purpose of the variable or constant.
