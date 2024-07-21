# A fast and convinient Go library for solving quadratic equations
## ⚠️ Attention
**Equation** v1.x.x does not support complex numbers, only floating point numbers.
## ⚙️ Installation

First, initialize your project with Go modules by executing the following command in your terminal:

```bash
go mod init github.com/your/repo
```

To learn more about Go modules and how they work, you can check out the [Using Go Modules](https://go.dev/blog/using-go-modules) blog post.

After setting up your project, you can install **Equation** with the `go get` command:

```bash
go get -u github.com/Gregory-coder/quadratic-equation-solver
```

This command fetches the package and adds it to your project's dependencies, allowing you to start building your web applications with **Equation**.

## ⚡️ Quickstart

Once you have installed the package, follow this short documentation to start using **Equation** easily.

```go
package main

import (
    "fmt"
    "github.com/Gregory-coder/quadratic-equation-solver"
)

func main() {
    // parse an equation from a string
    e, _ := equation.Parse("x^2=4")
    // solve it
    res, _ := e.Solve()

    fmt.Println(res) // [-2 2]
}
```

### Parse equations
**Equation** allows you to parse an equation from the string via **equation.Parse(string)** method 

The unknown value must be determined as *x*, *x* in the second degree is defined as *x^2*.
In other words, **Equation** can solve equations looking like:
*ax^2 + bx + c = 0*, 
where *a* and *x* are floating point numbers.

```go
e, err := equation.Parse("x^2=4")
```
