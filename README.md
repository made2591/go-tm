# go-tm - Make a Turing Machine easily

[![Build Status](https://travis-ci.org/made2591/go-tm.svg?branch=master)](https://travis-ci.org/made2591/go-tm)
[![codebeat badge](https://codebeat.co/badges/53c8e4e9-5bed-485f-9a18-570bce089e1b)](https://codebeat.co/projects/github-com-made2591-go-tm-master)
[![Coverage Status](https://coveralls.io/repos/github/made2591/go-tm/badge.svg?branch=master)](https://coveralls.io/github/made2591/go-tm?branch=master)
[![License](https://img.shields.io/github/license/made2591/go-tm.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/made2591/go-tm?status.svg)](https://godoc.org/github.com/made2591/go-tm)

![header](header.gif)

**go-tm** is a library to implement both deterministic and non deterministic Turing machines easily. A related [blog-post](https://made2591.github.io/posts/golang-turing-machine) is here.

The Turing machine is a mathematical model of computation that defines an abstract machine which manipulates symbols on a strip of tape according to a list of rules. Formally, it is defined as an infinite tape of 0 with a pointer (identified by square bracket) to one specific zero,

```sh
... 0 0 0 0 0 [0] 0 0 0 0 0  ...
```

a set of states identified by letters (or numbers),

```sh
{A, B, C}
```

and a list of transactions, like the one

```sh
(A, 0, B, 1, R);
```

where a single transaction like *(A, 0, B, 1, R)* has to been read as

> *Given a Turing machine in state A with the pointer over a 0, write 1, evolve to state B and move the pointer by one position in right direction over the tape*.

The pointer can only be moved by one position at time, left, right, or stay where it is.

## Table of Contents

- [go-tm - Make a Turing Machine easily](#go-tm---make-a-turing-machine-easily)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Support](#support)
  - [Contributing](#contributing)

## Installation

Just get the library by:

```sh
go get github.com/made2591/go-tm
```

## Usage

To create a new ```Symbol```, use the ```NewSymbol()``` factory method

```sh
// will create a new Symbol with value 0
sym := NewSymbol(uint8(0))
```

To create a new ```State```, use the ```NewState()``` or ```NewStateInitial()``` or ```NewStateFinal()``` factory methods

```sh
// will create a three new State, one with Identifier "A", one INITIAL and one FINAL
sts := NewState("A")
ins := NewStateInitial()
fis := NewStateFinal()
```

To create a new ```Transaction```, use the ```NewTransaction()``` factory method

```sh
// will move from INITIAL State, reading 0, writing 0, to state "B"
tr0 := transaction.NewTransaction(state.NewInitialState(), symbol.NewSymbol(uint8(0)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "N")
```

To init a new ```TuringMachine``` to play the [BB-2 Game](https://en.wikipedia.org/wiki/Busy_beaver)

```sh
// create states set
iss := set.NewSet()
fss := set.NewSet()
trs := set.NewSet()

// create transaction
tr0 := transaction.NewTransaction(state.NewInitialState(), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(0)), "N")
tr1 := transaction.NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(0)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "R")
tr2 := transaction.NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(1)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "L")
tr3 := transaction.NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(1)), "L")
tr4 := transaction.NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(1)), state.NewFinalState(), symbol.NewSymbol(uint8(1)), "R")
// add transaction to set
trs.Add(tr0)
trs.Add(tr1)
trs.Add(tr2)
trs.Add(tr3)
trs.Add(tr4)

// init the Turing Machine
tm := turingMachine.NewTuringMachine(iss, fss, trs, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))

// run the Turing Machine
tm.Run()

```

or, from inside the ```root``` folder of the project

```sh
go run main.go
```

## Support

Please [open an issue](https://github.com/made2591/go-tm/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/made2591/go-tm/compare/).
