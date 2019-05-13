# Level 1: Exercise 5

Building on the code from the previous example
1. at the package level scope, using the `var` keyword, create a VARIABLE with the IDENTIFIER `y`. The variable should be of the UNDERLYING TYPE of your custom TYPE `x`
eg:
    ```go
    type hotdog int
    var x hotdog
    var y int
    ```
2. in func main
    - this should already be done
        1. print out the value of the variable “x”
        2. print out the type of the variable “x”
        3. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
        4. print out the value of the variable “x”
    - now do this
        - now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
            1. then use the “=” operator to ASSIGN that value to “y”
            2. print out the value stored in “y”
            3. print out the type of “y”
