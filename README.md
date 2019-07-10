# efp

There are lot of people who use excel in their day to day operations. They are well versed with excel formulas. Having support for excel formulas in a software system would be of great value of end users. Some target groups:

* Project Managers
* Process Planners
* BOM Managers
* Finance Managers

## Roadmap

1. Drop 1:
    * Lexer for excel formulas
    * Parser for orchestrating lexer
    * Support for all Logical functions
    * Support for all text functions
2. Drop 2:
    * Support for building a range
    * Support for Math & Trig with exception of array functions
    * Support for Date & Time functions

## Approach

The intent is to leverage golang approach in parsing text templates. You can find a great talk by Rob Pike here -> [YouTube](https://youtu.be/HxaD_trXwRE)

## How to install

TODO

## Examples

TODO - Draft only

### Use-case 1

Initialize the excel formula parse

```golang
package main

include (
    praveentiru/efp
)

func main() int {
    parser := efp.CreateParser()
}
```
