# Simple
###### A Programming Language - Made with **GO**

- **Introduction**
    *Simple* is a  programming language based on *GO lang* as a base language.
    The objective to create this language is not only to learn GO lang but also to understand the journey of *idea* from *source code* to a *finished product*.

---

- **Modern Features**

    - **Spread Operator**
    The spread operator allows an iterable to spread or expand individually inside a receiver. Iterables are anything that can be looped over such as strings, arrays, and sets.
    ```
    let arr1 = [1,2,3,4,5]
    let arr2 = [2,4,1,...arr1,5,6,7]
    // arr2  =  2,4,1, 1,2,3,4,5, 5,6,7
    ```

    - **Rest Operator**
    Syntactically, it looks exactly the same as spread operator. But it’s function is the exact opposite of the spread Operator. If Spread operator expands individual items, then rest operator collects a bunch of items and puts them into arrays and objects.
    ```
    const numbers = [1, 2, 3];
    const [ firstNumber, ...restOfTheNumbers ] = numbers;
    console.log(firstNumber, restOfTheNumbers);
    // 1 [ 2, 3 ]
    ```

    - **Destructuring**
    Destructuring means taking out ( technically, making a copy of ) individual items from an object or an array and assigning them to a variable.

    ```
    let address = [221, 'Baker Street', 'London'];
    let [ houseNo, , city ] = address;
    // houseNo = 221
    // city = 'London'
    ```

    - **String interpolation**
    Template String or Template literals are enclosed with ` (back-tick) . You can embed variables in Template string just by surrounding variable by ${}, Multi-line strings.

    ```
    var name = "Bob", time = "today";
    var multiLine = `This
    Line
    Spans Multiple
    Lines`

    print(`Hello ${name},how are you ${time}?`) //Hello Bob,how are you today?
    print(multiLine) //This Line Spans Multiple Lines
    ```

    - **Nil Coalescing Operators**
    Nil coalescing is an operator used to guarantee that an optional contains a value.
    If the optional has a value, the nil coalescing operator will unwrap it. If the optional’s value is nil, it will assign it a default value instead.
    ```
    let myOptional =  newVar?
    let myExample = myOptional ?? "this string is my default value"

    ```
    `a != nul ? a! : b` With nil coalescing, this becomes `a ?? b`

---

- **Lexical Elements**
    - `\\` - single line comments
    ```
    let a = 32
    \\ this is a single line comment
    let b = "test"
    ```
    - `//` - multiline comments
    ```
    //
        this is a multiline comment
    //
    ```
    - `;` - semicolon
        The formal grammar uses semicolons `;` as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:
        - When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
            - an identifier
            - an integer, floating-point or string literal
            - one of the keywords `break`, `continue`, or `return`
            - one of the operators and punctuation `++`, `--`, `)`, `]`, or `}`
        - To allow complex statements to occupy a single line, a semicolon may
        be omitted before a closing `)` or `}`.
    - indentifier - name program entities such as variables and types.
    `identifier = letter { letter | unicode_digit }`

---

- **Variables**
    A variable is a storage location for holding a value.
    ```
    let name = ""
    let age = 21
    const PI = 3.14
    ```

---

- **Types**
    - `int` - the set of all signed 32-bit integers (-2147483648 to 2147483647)
    - `float` - the set of all IEEE-754 32-bit floating-point numbers
    - `char`
    - `string`

---

- **Blocks**

---

- **Scope**

---

- **Statements**

---

- **Built-in Functions**

.

- **Errors**

---

- **Operators**
    - **Precedence**
    ```
    Precedence    Operator
    5             *  /  %
    4             +  -
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
    ```
    `x / y * z` is same as ` (x / y) * z`

    - **Arthematic**
    ```
    +    sum                    integers, floats, strings
    -    difference             integers, floats
    *    product                integers, floats
    **   square                 integers
    /    quotient               integers, floats
    //   division               integers, floats
    %    remainder              integers
    ```

    - **Comparison**
    Comparison operators compare two operands and yield an boolean value.
    ```

    ==    equal
    !=    not equal
    <     less
    <=    less or equal
    >     greater
    >=    greater or equal
    ```
    - **Logical**
    Logical operators apply to boolean values and yield a result of the same type as the operands. The right operand is evaluated conditionally.
    ```
    &&    conditional AND    p && q  is  "if p then q else false"
    ||    conditional OR     p || q  is  "if p then true else q"
    !     NOT                !p      is  "not p"
    ```

---

- **Keywords**
    - `task` - is a built in keyword to declare a function in the program.

        ```
        task(){}
        task(x int) { return int}
        task(a, b ,z) { return bool }
        task(n){ return task(p) }
        ```
    - `let` - is used of bind a name with the value or expression
        ```
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
