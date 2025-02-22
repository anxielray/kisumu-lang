# Kisumu Programming Language Specification

## 3. Data Structures

The data structures in Kisumu Lang provide robust and efficient ways to organize and manipulate data. They include built-in methods and support advanced structures for greater flexibility and functionality.

### 3.1 Primitive Data Structures

#### 3.1.1 Arrays

An array is an ordered collection of elements of the same type. Arrays in Kisumu are fixed-size and zero-indexed.

**Syntax:**

```ksm
type[] arrayName = [element1, element2, ...]
```

**Example:**

```ksm
int[] numbers = [1, 2, 3, 4, 5]
string[] names = ["Alice", "Bob", "Charlie"]
```

**Built-in Methods:**

- **`length()`**: Returns the number of elements in the array.
- **`append(element)`**: Returns a new array with the element added to the end (does not modify the original array).
- **`slice(start, end)`**: Returns a sub-array from `start` to `end` (exclusive).

**Usage:**

```ksm
int count = numbers.length()
int[] extendedNumbers = numbers.append(6)
int[] subNumbers = numbers.slice(1, 3) // [2, 3]
```

#### 3.1.2 Maps

A map is a collection of key-value pairs, where each key is unique.

**Syntax:**

```ksm
map<keyType, valueType> mapName = {key1: value1, key2: value2, ...}
```

**Example:**

```ksm
map<string, int> ages = {"Alice": 25, "Bob": 30}
int aliceAge = ages["Alice"]
ages["Charlie"] = 35
```

**Built-in Methods:**

- **`keys()`**: Returns an array of all keys.
- **`values()`**: Returns an array of all values.
- **`hasKey(key)`**: Returns a boolean indicating if the key exists.
- **`remove(key)`**: Deletes the key-value pair and returns the updated map.

**Usage:**

```k
array<string> allKeys = ages.keys()
bool exists = ages.hasKey("Alice")
ages = ages.remove("Bob")
```

### 3.2 Composite Data Structures

#### 3.2.1 Structs

Structs group related data into a single composite type.

**Syntax:**

```ksm
struct StructName {
    fieldType1 fieldName1
    fieldType2 fieldName2
    ...
}
```

**Example:**

```ksm
struct Person {
    string name
    int age
    bool isEmployed
}

Person alice = Person(name="Alice", age=25, isEmployed=true)
print(alice.name) // Output: Alice
```

**Built-in Methods:**

- **`toString()`**: Converts the struct into a readable string representation.
- **`equals(otherStruct)`**: Compares two structs for equality.

#### 3.2.2 Enums

Enums define a set of named constants.

**Syntax:**

```ksm
enum EnumName { Value1, Value2, ... }
```

**Example:**

```ksm
enum Day { Monday, Tuesday, Wednesday }
Day today = Day.Monday
if (today == Day.Monday) {
    print("Start of the week!")
}
```

### 3.3 Advanced Data Structures

#### 3.3.1 Tuples

Tuples are immutable collections of fixed-size, heterogeneous elements.

**Syntax:**

```ksm
(type1, type2, ...) tupleName = (value1, value2, ...)
```

**Example:**

```ksm
(string, int, bool) personData = ("Alice", 25, true)
print(personData.0) // Accesses the first element: "Alice"
```

#### 3.3.2 Unions

Unions allow a variable to hold one of several predefined types.

**Syntax:**

```ksm
union UnionName { type1, type2, ... }
```

**Example:**

```ksm
union Number { int, float }
Number value = 42
value = 3.14 // Allowed because `float` is part of the union
```

### 3.4 Memory Management

Kisumu provides automatic garbage collection for all data structures, ensuring efficient resource utilization without manual intervention.

### 3.5 Type Safety

All data structures are statically typed, enforcing strict type rules to prevent runtime errors.
