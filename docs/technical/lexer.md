# Kisumu Programming Language Specification

## 2. Lexical Structure

The lexical structure defines the set of basic building blocks of the Kisumu programming language. These elements include characters, tokens, and rules for constructing identifiers, keywords, and other syntactic components.

### 2.1 Character Set

Kisumu uses the Unicode character set, supporting a wide range of characters from various languages and scripts. The standard encoding is UTF-8 to ensure compatibility and flexibility.

### 2.2 Tokens

Tokens are the smallest units of the language's syntax. Kisumu supports the following token types:

1. **Identifiers**
2. **Keywords**
3. **Literals**
4. **Operators**
5. **Punctuation**

#### 2.2.1 Identifiers

Identifiers are names used to represent variables, functions, and other user-defined elements. They must begin with a letter or an underscore (`_`) and can be followed by letters, digits, or underscores. Identifiers are case-sensitive.

**Examples:**

```
validName
_variable1
myFunction
```

#### 2.2.2 Keywords

Keywords are reserved words with specific meanings in the language. They cannot be used as identifiers. The initial set of keywords includes:

```
int, float, string, bool, func, return, if, else, for, while, break, continue, true, false
```

#### 2.2.3 Literals

Literals represent fixed values in the code, including:

- **Integer Literals** (e.g., `42`, `0`, `-7`)
- **Floating-Point Literals** (e.g., `3.14`, `-0.001`)
- **String Literals** (e.g., `"Hello, Kisumu!"`)
- **Boolean Literals** (e.g., `true`, `false`)

#### 2.2.4 Operators

Operators are symbols used to perform operations on data. Examples include:

- **Arithmetic Operators**: `+`, `-`, `*`, `/`, `%`
- **Relational Operators**: `==`, `!=`, `<`, `>`, `<=`, `>=`
- **Logical Operators**: `&&`, `||`, `!`

#### 2.2.5 Punctuation

Punctuation tokens are used for syntax structuring. Examples include:

- Comma `,`
- Parentheses `(` and `)`
- Braces `{` and `}`
- Brackets `[` and `]`

### 2.3 Comments

Comments are ignored by the compiler and are used to annotate code. Kisumu supports:

- **Single-Line Comments**: Begin with `//` and extend to the end of the line.
- **Multi-Line Comments**: Enclosed between `/*` and `*/`.

**Examples:**

```k
// This is a single-line comment
/* This is a
multi-line comment */
```

### 2.4 Whitespace

Whitespace characters (spaces, tabs, and newlines) are generally ignored by the compiler but are used to separate tokens.
