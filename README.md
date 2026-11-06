# Conditional Statements in Go (Golang)

This project demonstrates various conditional statements, branching logic, and switch statement variations in Go (Golang).

---

## 📌 Topics Covered

1. **If-Else Ladder (`if`, `else if`, `else`)**
2. **Nested If-Else**
3. **Logical Operators (`&&`, `||`, `!`)**
4. **Switch with Exact Value**
5. **Switch with Multiple Values per Case**
6. **Switch Without Expression (Tagless Switch)**
7. **Float and Boolean Conditions**
8. **String Switch**
9. **Type Switch (`value.(type)`)**

---

## 📖 Detailed Explanations & Examples

### 1. `if`, `else if`, `else`
Go's `if` statements do not require parentheses `()` around the condition, but curly braces `{}` are mandatory.

```go
if marks >= 80 {
    fmt.Println("Grade: A+")
} else if marks >= 70 {
    fmt.Println("Grade: A")
} else {
    fmt.Println("Grade: Fail")
}
```

### 2. Nested `if-else`
You can nest `if` blocks inside other `if` blocks for multi-level validation:

```go
if age >= 18 {
    fmt.Println("Adult")
    if isStudent {
        fmt.Println("You are a Student")
    }
} else {
    fmt.Println("Minor")
}
```

### 3. Logical Operators
- `&&` (Logical AND): Evaluates to true if both conditions are true.
- `||` (Logical OR): Evaluates to true if at least one condition is true.
- `!` (Logical NOT): Inverts the boolean value.

```go
if age >= 18 && marks >= 80 {
    fmt.Println("Eligible")
}

if !isActive {
    fmt.Println("Account is Inactive")
}
```

### 4. Switch Statement with Exact Values
In Go, switch cases break automatically; you do not need an explicit `break` statement at the end of each case.

```go
switch age {
case 18:
    fmt.Println("Age is 18")
case 25:
    fmt.Println("Age is 25")
default:
    fmt.Println("Age Not Found")
}
```

### 5. Switch with Multiple Case Values
You can combine multiple match criteria in a single case by separating them with commas `,`.

```go
switch day {
case 1, 2, 3:
    fmt.Println("First Group")
case 4, 5, 6:
    fmt.Println("Second Group")
case 7:
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid Day")
}
```

### 6. Switch Without Expression (Tagless Switch)
A switch without a target expression acts as a clean, idiomatic alternative to a long `if - else if` chain. Each case evaluates a boolean condition:

```go
switch {
case age < 18:
    fmt.Println("Minor")
case age >= 18 && age < 30:
    fmt.Println("Young Adult")
default:
    fmt.Println("Senior")
}
```

### 7. String Switch
Switch statements work seamlessly on strings:

```go
switch role {
case "admin":
    fmt.Println("Admin Dashboard")
case "developer":
    fmt.Println("Developer Dashboard")
default:
    fmt.Println("Unknown Role")
}
```

### 8. Type Switch
A type switch compares dynamic types using the syntax `value.(type)` on an interface (`any` or `interface{}`):

```go
func checkType(value any) {
    switch v := value.(type) {
    case string:
        fmt.Println(v, "is String")
    case int:
        fmt.Println(v, "is Integer")
    case float64:
        fmt.Println(v, "is Float64")
    case bool:
        fmt.Println(v, "is Boolean")
    default:
        fmt.Println("Unknown Type")
    }
}
```

---

## 🚀 How to Run

Make sure you have Go installed on your machine.

1. Open your terminal in this project directory:
   ```bash
   cd "4. Conditional Statements in GO"
   ```

2. Run the program directly:
   ```bash
   go run main.go
   ```

---

## 📋 Expected Output

```text
===== GO CONDITION PRACTICE =====

--- IF ELSE ---
Grade: A+

--- NESTED IF ---
Adult
You are a Student

--- LOGICAL OPERATORS ---
Eligible
Account is Inactive

--- SWITCH VALUE ---
Age is 25

--- MULTIPLE CASE ---
Second Group

--- SWITCH CONDITION ---
Young Adult

--- FLOAT CONDITION ---
Excellent CGPA

--- STRING SWITCH ---
Developer Dashboard

--- BOOLEAN CONDITION ---
Inactive Student

--- TYPE SWITCH ---
Aftab is String
25 is Integer
3.5 is Float64
true is Boolean

===== END OF PROGRAM =====
```

---

## 💡 Key Takeaways for Go Conditionals

- **No Parentheses Required**: Conditions do not require parentheses: `if x > 0` instead of `if (x > 0)`.
- **Braces are Mandatory**: The opening brace `{` must always be on the same line as the `if` or `switch` statement.
- **Automatic Break in Switch**: Unlike C/C++/Java, Go executes only the matching case and automatically breaks out. Use the `fallthrough` keyword if you explicitly want to fall through to the next case.
