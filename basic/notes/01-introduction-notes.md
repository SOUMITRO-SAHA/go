# Introduction

- `Golang` is a procedural and statically typed programming language with syntax similar to C programming language.
- Sometimes it is termed as Go Programming Language.
- It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language and mainly used in Google’s production systems.
- Go is a statically typed, concurrent, and garbage-collected programming language created at Google in 2009. It is designed to be simple, efficient, and easy to learn, making it a popular choice for building scalable network services, web applications, and command-line tools.
- Golang is one of the most trending programming languages among developers.
- Programs are assembled by using packages, for efficient management of dependencies. This language also supports environment adopting patterns alike to dynamic languages
- Go is known for its support for concurrency, which is the ability to run multiple tasks simultaneously. Concurrency is achieved in Go through the use of Goroutines and Channels, which allow you to write code that can run multiple operations at the same time. This makes Go an ideal choice for building high-performance and scalable network services, as well as for solving complex computational problems.
- Another important feature of Go is its garbage collection, which automatically manages memory for you. This eliminates the need for manual memory management, reducing the likelihood of memory leaks and other bugs that can arise from manual memory management.

#### Key Features

- Simplicity: Go is designed to be easy to learn and use. Its syntax is simple and straightforward, making it a good choice for beginners and experienced programmers alike.

- Concurrency: Go has built-in support for concurrency, allowing developers to write efficient and scalable code for multi-core and distributed systems.

- Garbage collection: Go has automatic memory management, which frees developers from having to worry about memory allocation and de-allocation.

- Fast compile times: Go has a fast compiler, which makes it easy to iterate quickly during development.

- Cross-platform support: Go can be compiled to run on many different platforms, including Windows, Linux, and macOS.

- Strong typing: Go is a statically typed language, which helps catch errors at compile time rather than at runtime.

- Go has a large and growing community of developers and is used by many well-known companies, including Google, Uber, and Dropbox.

#### Important Points

- Go is a statically typed language, which means that the type of a variable must be declared before it can be used.

- Go has a built-in garbage collector that automatically frees up memory when it is no longer needed.
  Go has strong support for concurrency, allowing developers to write efficient and scalable code for multi-core and distributed systems.

- Go has a minimalist syntax that is easy to learn and read.

- Go has a fast compiler that generates code that is optimized for modern hardware architectures.

- Go has a standard library that provides support for a wide range of functionality, including networking, encryption, and file handling.

- Go has a growing community of developers and a vibrant ecosystem of third-party packages and tools.

- Go is used by many well-known companies for building large-scale distributed systems and high-performance applications.

- Overall, Go is a powerful and efficient programming language that is well-suited for building modern applications and distributed systems. Its strong support for concurrency and minimalist syntax make it an attractive choice for developers who want to build scalable and efficient applications.

#### `Hello World` program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

/*
Output:
Hello, World!
*/
```

#### Explanation of the syntax of Go program:

- Line 1: It contains the package main of the program, which have overall content of the program.It is the initial point to run the program, So it is compulsory to write.
- Line 2: It contains import `fmt`, it is a preprocessor command which tells the compiler to include the files lying in the package.
- Line 3: main function, it is beginning of execution of program.
- Line 4: `fmt.Println()` is a standard library function to print something as a output on screen.In this, `fmt` package has transmitted `Println` method which is used to display the output.

### `Comment`:

Comments are used for explaining code and are used in similar manner as in Java or C or C++. Compilers ignore the comment entries and does not execute them. Comments can be of single line or multiple lines.

#### Single Line Comment

Syntax:

```go
// single line comment
```

#### Multi-line Comment

Syntax:

```go
/* multiline comment */
```

### Why "Go language"?

Because Go language is an effort to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aims to be modern, with support for networked and multi-core computing.

### Some popular Applications developed in Go Language

- Docker: a set of tools for deploying linux containers
- Openshift: a cloud computing platform as a service by Red Hat.
- Kubernetes: The future of seamlessly automated deployment processes
- Dropbox: migrated some of their critical components from Python to Go.
- Netflix: for two part of their server architecture.
- InfluxDB: is an open-source time series database developed by InfluxData.
- Golang: The language itself was written in Go.

[]
