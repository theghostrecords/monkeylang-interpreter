# Monkey Programming Language

Monkey is a simple, C-like programming language built for learning and exploration. It was introduced in the book [_Writing an Interpreter in Go_](https://interpreterbook.com) by Thorsten Ball. Monkey supports variable bindings, functions, conditionals, arrays, hashes, and first-class functions.

---

## Language Features

### Variable Bindings
You can bind values to names using the `let` keyword.

```monkey
let x = 5;
let y = 10;
let foobar = 838383;
```

---

### Data Types
Monkey supports integers, booleans, strings, arrays, and hashes.

```monkey
let age = 28;           // Integer
let isCool = true;      // Boolean
let name = "John Doe";  // String
```

---

### Arithmetic Expressions
Monkey supports standard arithmetic operations.

```monkey
let x = 5;
let y = 10;
let result = (x + y) * 2 / 10 - 3;  // result is -1
```

---

### Conditional Expressions
Conditionals use `if`/`else`. In Monkey, `if` is an expression, meaning it returns a value.

```monkey
let x = 10;

if (x > 5) {
  return true;
} else {
  return false;
}

let value = if (x > 0) { 100 } else { -100 };  // value is 100
```

---

### Functions
Functions are first-class values. You can assign them to variables and pass them around.

```monkey
let add = fn(a, b) {
  return a + b;
};

add(5, 10);  // returns 15
```

---

### Closures
Functions can capture variables from their surrounding environment.

```monkey
let newAdder = fn(x) {
  return fn(y) { x + y; };
};

let addTwo = newAdder(2);
addTwo(3);  // returns 5
```

---

### Strings
Strings are wrapped in double quotes and support concatenation with the `+` operator.

```monkey
let greeting = "Hello";
let subject = "World";
let message = greeting + " " + subject + "!";  // "Hello World!"
```

---

### Arrays
Arrays are ordered collections of elements, which can be of any type.

```monkey
let myArray = [1, "two", fn(x) { x * x }];
let firstNumber = myArray[0];       // 1
let len = len(myArray);             // 3
let lastElement = last(myArray);    // fn(x) { x * x }
```

---

### Hashes
Hashes (similar to maps or dictionaries) store key-value pairs. Keys can be integers, booleans, or strings.

```monkey
let myHash = {
  "name": "Gandalf",
  "age": 2019,
  true: "Wizard"
};

myHash["name"];   // "Gandalf"
myHash[true];     // "Wizard"
```

---

## Reference
This language is based on the book [_Writing an Interpreter in Go_](https://interpreterbook.com) by Thorsten Ball. It’s a great resource if you want to learn how interpreters and programming languages work from the ground up.
