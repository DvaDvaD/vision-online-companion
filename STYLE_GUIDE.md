# Coding Style

### Go Conventions

1. File and Folder Names

    * Use snake_case for file and folder names.
    * Avoid uppercase letters and use underscores to separate words

```
user_repository.go
config_loader.go
models/
handlers/
```

2. Variable Names
    * Use camelCase for local variables.
    * Keep variable names descriptive but concise.
    * Avoid single-letter names, except for simple iterators (i, j, etc.)
    * Abbreviation should be UPPERCASE (e.g., ID, URL).

```go
var userName string
var userID int
count := 10
for i := 0; i < count; i++ {
    // ...
}
```

3. Constant Names

    * Use camelCase for unexported constants and PascalCase for exported constants.
    * Define constants at the top of a file if they’re file-specific or in a separate package if shared.
    * Use iota for enumerated constants.

```go
const maxRetryCount = 5
const DefaultTimeout = 30

// Integer enum
type Color int

const (
    Red Color = iota
    Green
    Blue
)

// String enum
type Color string

const (
    Red   Color = "Red"
    Green Color = "Green"
    Blue  Color = "Blue"
)

func (c Color) String() string {
    return string(c)
}
```

4. Function, Type and Interface Names

    * Use PascalCase for both exported and unexported type names.
    * Avoid abbreviations unless they are well-known (e.g., ID, URL).

```go
func NewUser(name string) *User {
    // ...
}

func calculateTotal(amount, tax float64) float64 {
    // ...
}

type User struct {
    ID   int
    Name string
}

type Config struct {
    Timeout int
    URL     string
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Stringer interface {
    String() string
}
```

5. Understand when to use a technique vs another (TODO: Make your own explanation)

    * [Generic Type](https://www.youtube.com/watch?v=WpTKqnfp5dY)
    * [Pointer vs Value as function parameters](https://www.youtube.com/watch?v=3WsEDZRif6U)
    * [Interface as a contract](https://www.youtube.com/watch?v=rH0bpx7I2Dk)
    * Go [constructor factory pattern](https://refactoring.guru/design-patterns/factory-method/go/example) and [functional options](https://www.youtube.com/watch?v=MDy7JQN5MN4)

## General

* Write readable and maintainable code with meaningful variable names.
* Avoid large functions; break them down if necessary.
* Comment complex logic but avoid over-commenting obvious code.

| **Convention**        | **Go**                                            | **JavaScript / React**                                             |
| --------------------- | ------------------------------------------------- | ------------------------------------------------------------------ |
| **File Names**        | `snake_case`                                      | `kebab-case` for non-components;<br /> `PascalCase` for components |
| **Folder Names**      | `snake_case`                                      | `kebab-case`                                                       |
| **Variable Names**    | `camelCase`                                       | `camelCase`                                                        |
| **Constant Names**    | `camelCase` (unexported), `PascalCase` (exported) | `UPPER_SNAKE_CASE`                                                 |
| **Function Names**    | `camelCase` (unexported), `PascalCase` (exported) | `camelCase`                                                        |
| **Component Names**   | N/A                                               | `PascalCase`                                                       |
| **Type Names**        | `PascalCase`                                      | `PascalCase`                                                       |
| **Interface Names**   | `PascalCase`                                      | `PascalCase`                                                       |
| **Custom Hook Names** | N/A                                               | `camelCase`, starts with `use`                                     |
| **Imports**           | N/A                                               | Named imports preferred, organized with index files                |
| **Exports**           | N/A                                               | Named exports preferred (no default)                               |
| **Package Naming**    | `lowercase`                                       | N/A                                                                |
