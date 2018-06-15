# Short introduction to CLI creation in Go

![Go CLI tutorial](/img/gocli1.png)

## Introduction

There already exists many third-party packages that you can use to create command line interfaces in Go, but this tutorial is going to focus on making CLIs using only the standard library.

In other words, "build your own fucking birdfeeder".

[All files presented can be found on my GitHub](https://github.com/vagnes/go-cli-tutorial).

## Getting started

We are going to use three main standard library packages for this tutorial:

1. [fmt](https://golang.org/pkg/fmt/) - Formatting simple I/O
2. [flag](https://golang.org/pkg/flag/) - CLI flag parsing
3. [os](https://golang.org/pkg/os/) - Platform independent OS functionality

This tutorial assumes basic knowledge of Go, but should be pretty easy to understand regardless.

## Simple CLI

Let's first make a CLI without flags, which would be the simplest CLI we could make:

ref. 1-1

```go
package main

import (
	"fmt"
	"os"
)

func main() {
    programName := os.Args[0]
    fmt.Printf("Program name: %s\n", programName)
}
```

[os.Args](https://golang.org/pkg/os/#pkg-variables) holds the CLI arguments, which starts with the program name as its first string in the slice, and then continues with the supplied arguments.

We can demonstrate this with the following snippet of code.

ref 1-2

```go
package main

import (
	"fmt"
	"os"
)

func main() {
    for i := 0; i<5; i++ {
		programName := os.Args[i]
		fmt.Printf("Argument No. : %d (%s)\n", i, programName)
	}
    
}
```

It will have the following output:

```
go run 1-2.go foo bar

Argument No. : 0 (/tmp/go-build501384166/command-line-arguments/_obj/exe/1-2)
Argument No. : 1 (foo)
Argument No. : 2 (bar)
panic: runtime error: index out of range

goroutine 1 [running]:
panic(0x4db1e0, 0xc82000e070)
        /usr/lib/go-1.6/src/runtime/panic.go:481 +0x3e6
main.main()
        /mnt/d/oneDrive/code/go/testing/cli/1/1-2.go:10 +0x20f
exit status 2
```

The error is expected as we are looping over a slice of strings, which are the arguments stored in os.Args, and there are only three arguments; the first argument being the name of the program itself, and the "foo" and "bar" supplied by us in the terminal. This can be easily fixed by finding the number of strings in os.Args, by doing the following.

ref. 1-3

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	numberOfArgs := len(os.Args[1:])
    for i := 0; i <= numberOfArgs; i++ {
		programName := os.Args[i]
		fmt.Printf("Argument No. : %d (%s)\n", i, programName)
	}
    
}
```

This will give us the expected output, as in 1-2, but without the "index out of range" error. Alternatively, you could use "range" to enumerate the arguments instead.

ref. 1-4

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	
	for i, arg := range os.Args[1:] { // 1, not 0 because we do not want the name of the program
        fmt.Printf("Argument No.%d: %s\n", i+1, arg)
    }
}
```

```text
go run 1-4.go foo bar

Argument No.1: foo
Argument No.2: bar
```

Why you would want to use "range" over "len" would of course be completely dependent on the context, but there is no wonder which way is more concise.

> flag.Arg(i) returns the i'th CLI argument, while flag.Args() returns the **non-flag** CLI arguments.

## Using the flag package

Using the flag package, we can take flags entered at the time of running the program and parse them, and hence we can do some more action with the flags/arguments that we pass. The **flag.StringVar** function takes the following input:

```go
StringVar func(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage ring. The > argument p points to a string variable in which to store the value of the flag.

StringVar is not the only thing that could be used; you could also use IntVar, Float64 and so on.

This means you could make programs like the following:

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	
	var password string
	flag.StringVar(&password, "p", "", "password for access")
	flag.Parse()

	if password == "password" {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

}
```

Please note that you should never do this, as it is really simple to disassemble the program and extract the plain text password even though the disassembly of Go is arguably more messy than of a C program, as demonstrated below.

```
mov     rcx, [rax]
mov     [rsp+98h+usage.str], rcx
mov     rcx, [rax+8]
mov     [rsp+98h+usage.len], rcx
lea     rbp, aPassword  ; "password"
mov     [rsp+98h+var_88], rbp
mov     [rsp+98h+var_80], 8
call    runtime_eqstring
movzx   ebx, byte ptr [rsp+98h+var_78]
cmp     bl, 0
jz      loc_4011B5
```

## More advanced CLI

The first more advanced CLI is using an operator, specified using the "-o" flag, on our arguments pairwise, for example 1 + 2, 3 + 4 and so forth.

ref 2-1

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	flag.Usage = func() {
		fmt.Println("\nMy Mediocre Program - By Me\n")
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

	var operator string
	flag.StringVar(&operator, "o", "+", "operator to use (+, -, * or /)")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for i := 0; i < len(flag.Args()); i = i + 2 {
		firstInt, _ := strconv.ParseFloat(flag.Arg(i), 64)
		secondInt, _ := strconv.ParseFloat(flag.Arg(i+1), 64)
		if operator == "+" {
			fmt.Printf("%.2f + %.2f = %.2f\n", firstInt, secondInt, firstInt+secondInt)
		} else if operator == "-" {
			fmt.Printf("%.2f - %.2f = %.2f\n", firstInt, secondInt, firstInt-secondInt)
		} else if operator == "*" {
			fmt.Printf("%.2f * %.2f = %.2f\n", firstInt, secondInt, firstInt*secondInt)
		} else if operator == "/" {
			fmt.Printf("%.2f / %.2f = %.2f\n", firstInt, secondInt, firstInt/secondInt)
		} else {
			println("Operator not recognised.")
		}
	}

}
```

```
go run 2-1.go -o "/" 1 2 3 4 5 6 7 8 9 10

1.00 / 2.00 = 0.50
3.00 / 4.00 = 0.75
5.00 / 6.00 = 0.83
7.00 / 8.00 = 0.88
9.00 / 10.00 = 0.90
```

## Conclusion

Compared to other implementations in other languages, like argparse in Python, the flag package might not be a very advanced or intuitive one. For simple and small programs, it might get the job done. There are some obvious shortcomings, for example the flag package does not support mandatory or required flags (meaning the flag must be specified explicitly).

Hence, you might want to check out more extensive implementations like [cli](https://github.com/urfave/cli) and [Cobra](https://github.com/spf13/cobra). Cobra is used in many projects written in Go like [Hugo](https://gohugo.io/) and [Kubernetes](https://kubernetes.io/).
