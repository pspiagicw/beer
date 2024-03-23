# `beer`

`beer` is a programming language.

## `inspiration`

Most languages aim to provide strong features like 

- `concurrency`.
- `blazingly fast performance`.
- `confusing memory safety on steroids`.
- `A new package/framework/library/runtime every week`.

Can't you write a project on a lazy Sunday afternoon, without worrying about the future ?

`beer` is a lighly brewed language, that aims to provide a simple and fun way to write code.

## `features`

### beautiful codebase
- The codebase is designed for beginner coders to understand and contribute.
- Aimed to be extensible and easy to hack on.
- Smartly documented with beautiful write-up and self-explanatory code.

### expressive syntax
- The syntax is not ambiguous and is easy to read.
- Looks beautiful and is fun to write.
- Errors don't scold you.

### batteries included
- Comes with a suite of tools to write, test and deploy code.
- Write, Test, Refactor quickly.

### fun
- You fall in love with coding.

## `syntax`

### `variables`

`a: 10`
`b: 20`

### `control flow`

`if a > b { printf("a is greater than b") }`

### `functions`

`fn add(a, b) { ret a + b }`

Call using

`add(a:10, b:20)`

`fn something() {
    printf("something")
}`
}`

### `comments`

`;; this is a comment`
`;;; this is a docstring`
`;;; before a function it becomes a docstring`

### `data types`

- `number`, combination of `int` and `float`.
- `string`, a sequence of characters.
- `boolean`, `true` or `false`.
- `list`, a sequence of elements.
- `table`.

### `operators`

- `+`, `-`, `*`, `/`, `%`.
- `==`, `!=`, `>`, `<`, `>=`, `<=`.
- `&&`, `||`, `!`.
- `++`, `--`.

### `tests`

- Write a test anywhere.

- Just use `test` instead of `fn`.

`
test add {
    a: 10
    b: 20
    assert a + b == 30, "addition failed"
}`

- `assert` can test anything!

### `cask`

- A `cask` is a something like a `package` in other languages.

`
import math
`

- Export functions using `cask` keyword.

`cask {
    ;; mention the functions you want to export
    ;; you can only export functions
    add,
}`

- Should be at the end of the file.
- A directory with a set of files becomes a `cask`, that can be imported with the directory name.
- A directory when run using `beer run` will find the `main` function and run it. It shouldn't be exported.

## `language`

- Open a file with `.b` as extension and write code.
- Run code with `beer run <filename>`
- Generate a new project with `beer create`
- Test code with `beer test`.


