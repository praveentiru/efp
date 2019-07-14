# efp

There are lot of people who use excel in their day to day operations. They are well versed with excel formulas. Having support for excel formulas in a software system would be of great value of end users. Some target groups:

* Project Managers
* Process Planners
* BOM Managers
* Finance Managers

## Pre-requisites

Go version - 1.12 and above

## Roadmap

1. Drop 1:
    * Support for all Logical functions
    * ~~Support for all text functions~~
2. Drop 2:
    * Support for building a range
    * Support for Math & Trig with exception of array functions
    * Support for Date & Time functions

## Excel functions supported

Text Functions

* CONCAT
* CONCATENATE
* EXACT
* FIND
* FIXED - *Todo*
* LEFT
* LEN
* LOWER
* MID
* PROPER
* REPLACE
* REPT
* RIGHT
* SEARCH
* SUBSTITUTE
* TRIM
* UPPER

## Approach

Use [gval](https://github.com/PaesslerAG/gval) to implement Excel formula language

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
