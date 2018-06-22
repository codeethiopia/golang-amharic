
#### Exercise 1

1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”

```
42
“Dr. Abiy Ahmed”
true
```

Now print the values stored in those variables using :

* a single print statement   

* multiple print statements

2. Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
```
identifier “x” type int
identifier “y” type string
identifier “z” type bool
```
3. in func main : 

* print out the values for each identifier

* The compiler assigned values to the variables. What are these values called?

Using the code from above

4. at the package level scope, assign the following values to the three variables

    a. for x assign 42
   
    b. for y assign “Dr. Abiy Ahmed”
    
    c. for z assign true

5. in func main
   
    a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
   
    b. print out the value stored by variable “s”

6. Create your own type. Have the underlying type be an int.

7. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword

8. in func main
   
    a. print out the value of the variable “x”
    
    b. print out the type of the variable “x”
   
    c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
   
    d. print out the value of the variable “x”

Building on the code from above

9. At the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

10. in func main
   
    a. this should already be done
        
        - print out the value of the variable “x”
       
        - print out the type of the variable “x”
       
        - assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
       
        - print out the value of the variable “x”
   
    b. now do this
       
        - now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
            
            * then use the “=” operator to ASSIGN that value to “y”
           
            * print out the value stored in “y”
            
            * print out the type of “y”