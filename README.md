# Simple

## A Programming Language - Made with **GO**

### Introduction

*Simple* is a  programming language based on *GO lang* as a base language.
The objective to create this language is not only to learn GO lang but also to understand the journey of *idea* from *source code* to a *finished product*.

---

### **Lexical Elements**

#### **Variables**

A variable is a storage location for holding a value.

```js
// Bind values to names with let-statements

let version = 1;
let name = "Simple programming language";
let myArray = [1, 2, 3, 4, 5];
let coolBooleanLiteral = true;
let add = fn(){}
```

#### **Use expressions to produce values**

```js
let awesomeValue = (10 / 2) * 5 + 30;
let arrayWithValues = [1 + 1, 2 * 2, 3];
```

#### `;` - semicolon

The formal grammar uses semicolons `;` as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:

- When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
- an identifier
- an integer, floating-point or string literal
- one of the operators and punctuation `++`, `--`, `)`, `]`, or `}`
- To allow complex statements to occupy a single line, a semicolon may
be omitted before a closing `)` or `}`.
- indentifier - name program entities such as variables and types `identifier = letter { letter | unicode_digit }`

#### **Operators**

- **Precedence**

```js
Precedence    Operator
5             *  /  %
4             +  -
3             ==  !=  <  <=  >  >=
2             &&
1             ||
```

`x / y * z` is same as `(x / y) * z`

- **Arthematic**

```js
+    sum                    integers, floats, strings
-    difference             integers, floats
*    product                integers, floats
/    quotient               integers, floats
%    remainder              integers
```

- **Comparison**

Comparison operators compare two operands and yield an boolean value.

```js

==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal

```

- **Logical**

Logical operators apply to boolean values and yield a result of the same type as the operands. The right operand is evaluated conditionally.

```js
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

#### **Keywords**

`fn` - is a built in keyword to declare a function in the program.

```js
fn(){}
fn(x int) { return int}
fn(a, b ,z) { return bool }
fn(n){ return task(p) }
```

`let` - is used of bind a name with the value or expression

```js
let a = 5;
let a = a + 6
let add = task ()
```

- `true`
- `false`
- `if`
- `else`
- `elseif`
- `return`

#### **Types**

```js

- `int` - the set of all signed 32-bit integers (-2147483648 to 2147483647)
- `float` - the set of all IEEE-754 32-bit floating-point numbers
- `char`
- `string`
```

```js

// Here is an array containing two hashes, that use strings as keys and integers
// and strings as values
let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

// Getting elements out of the data types is also supported.
// Here is how we can access array elements by using index expressions:
fibonacci(myArray[4]);
// => 5

// We can also access hash elements with index expressions:
let getName = fn(person) { person["name"]; };

// And here we access array elements and call a function with the element as
// argument:
getName(people[0]); // => "Anna"
getName(people[1]); // => "Bob"

```

#### Define a fibonacci function

```js
let fibonacci = fn(x) {
  if (x == 0) {
    0                // Simple supports implicit returning of values
  } else {
    if (x == 1) {
      return 1;      // ... and explicit return statements
    } else {
      fibonacci(x - 1) + fibonacci(x - 2); // Recursion! Yay!
    }
  }
};
```

#### functions are first-class citizens, higher-order functions and pass functions around as values

```js

// Define the higher-order function `map`, that calls the given function `f`
// on each element in `arr` and returns an array of the produced values.
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };

  iter(arr, []);
};

// Now let's take the `people` array and the `getName` function from above and
// use them with `map`.
map(people, getName); // => ["Anna", "Bob"]
```

#### closures

```js

// newGreeter returns a new function, that greets a `name` with the given
// `greeting`.
let newGreeter = fn(greeting) {

  return fn(name) { return (greeting + " " + name); }
};

// `hello` is a greeter function that says "Hello"
let hello = newGreeter("Hello");

// Calling it outputs the greeting:
hello("dear, simple programmer!"); // => dear, simple programmer!

```
