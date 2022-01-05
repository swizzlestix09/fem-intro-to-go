# Exercise 1A: Find Stuff

## Goals

- Familiarize yourself with the golang.org website and resources

## Setup

- Visit `golang.org`

## Directions

Answer the following questions

1. Read about `for loops` in the _Effective Go_ document

- What kind of loop doesn’t exist in Go?
- Go's for loop unifies for & while, and there is no do while
```
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

2. Read about the `fmt` _package_

- What does `fmt.Println()` return?

- Println is the equivilant of console.log.. prints what is in between it's parenthesis.
Will space variables seperated by commas.  It returns the number of bytes written and any write error encountered

3. Find a _blog post_ about the recent release of Go 1.18

- What are some of the new features?

Generics have been added
Generics - functions/types that are written to work with any of a set of types provided by calling code.
Fuzz testing is added - particularly valuble for finding security exploits or vunerabilities. Automated testing that continuously manipulated inputs that existing unit codes may miss.