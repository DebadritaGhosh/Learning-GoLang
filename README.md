
# Learning Golang

- Download Go : 
 https://golang.org/dl/
---
- Some Important Command for golang :  
  1. Run go code : go run <filename>.go (Example : go run main.go)
  2. Generate go mod file : go mod init <name> (Example : go mod init variable)
---
- Some Important Package :  
  1. fmt (Format package) : This package allows to format basic strings, values, or, anything and print them or collect user input from the console, or write into a file using a writer or even print customized fancy error messages. This package is all about formatting input and output. 
  2. bufio : 
  3. os : 
---
# I/O in Golang :  
  1. Output :

    Printf - "Print Formatter" this function allows you to format numbers, variables, and strings into the first string parameter you give it.  
    Example : fmt.Printf("Debadrita Ghosh")

    Print - "Print" This cannot format anything, it simply takes a string and prints it.  
    Example : fmt.Print("Debadrita Ghosh")  

    Println - "Print Line" same thing as Printf() however it will append a newline character \n . 
    Example : fmt.Println("Debadrita Ghosh")  
      
    
  2. Take Input from User

    reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Comma Ok Syntax || Error Ok Syntax

---
# Data Types in Go :  
  - ***Basic type:*** 
    - **Numbers**  

      - Integers
      1.     int8	8-bit signed integer
      2.     int16	16-bit signed integer
      3.     int32	32-bit signed integer
      4.     int64	64-bit signed integer
      5.     uint8	8-bit unsigned integer
      6.     uint16	16-bit unsigned integer
      7.     uint32	32-bit unsigned integer
      8.     uint64	64-bit unsigned integer
      9.     int	Both in and uint contain the same size, either 32 or 64 bit.
      10.     uint	Both in and uint contain the  same size, either 32 or 64 bit.
      11.     rune	It is a synonym of int32 and also represents Unicode code points.
      12.     byte	It is a synonym of int8.
      13.     uintptr	It is an unsigned integer type. Its width is not defined, but it can hold all the bits of a pointer value.   

      
      - Float
      1.     float32	32-bit IEEE 754 floating-point number
      2.     float64	64-bit IEEE 754 floating-point number

      - Complex Number
      1.     complex64	Complex numbers which contain float32 as a real and imaginary component.
      2.     complex128	Complex numbers which contain float64 as a real and imaginary component.
    - **Strings** 
    - **booleans**
  - ***Aggregate type:*** 
      - **Array** 
      - **Structs**
  - ***Reference type:*** 
      - **Pointers**
      - **Slices**
      - **Maps**
      - **Functions**
      - **Channels** 
  - ***Interface type***
  ---
  # Declare Variables in Golang :  

  - Var : var keyword in Golang is used to create the variables
  - Const : Const keyword in Golang is used to create the constant variables
  ---

  There are three ways.
---
  **Explicit Type**

    1. var website string = "debadritaghosh.tech"  
       const website string = "debadritaghosh.tech"
---
  **Implicit Type**

    2. var website = "debadritaghosh.tech"

---
  **No var Style Or Walrus Operator (:=)**

    3. website := "debadritaghosh.tech" // Inside the metod we can use := operator
---
# Pointer    
   - Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable.
   - Pointers in Golang is also termed as the special variables. 
   - The variables are used to store some data at a particular memory address in the system. The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

__* Operator__ :    

    * also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.

**& Operator**:    

    Termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.
---
# Arrays
 - Arrays in Golang or Go programming language is much similar to other programming languages. In the program, sometimes we need to store a collection of data of the same type
 - In Go language, arrays are mutable.
---
# Slices
 - In Go language slice is more powerful, flexible, convenient than an array, and is a lightweight data structure. 
 - Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store different type of elements in the same slice. 
 - It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array.
 - Internally, slice and an array are connected with each other, a slice is a reference to an underlying array. 
 - It is allowed to store duplicate elements in the slice. 
 - The first index position in a slice is always 0 and the last one will be (length of slice – 1)

***Ways to declare slices***

    ` s := []int{3,4,5}  // Initializing with values
      s := []int{} //Initialize with values,momory is allocated    
      var s []int //Declared as slice but no memory allocated
      s := make([]int,5,10) //initialized with no values,Allocated memory 
    
---
# Maps
- In Go language, a map is a powerful, ingenious, and versatile data structure. 
- Golang Maps is a collection of unordered pairs of key-value. 
- It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.
- It is a reference to a hash table.
- Due to its reference type it is inexpensive to pass, for example, for a 64-bit machine it takes 8 bytes and for a 32-bit machine, it takes 4 bytes.
- In the maps, a key must be unique and always in the type which is comparable using == operator or the type which support != operator. So, most of the built-in type can be used as a key like an int, float64, rune, string, comparable array and structure, pointer, etc. The data types like slice and noncomparable arrays and structs or the custom data types which are not comparable don’t use as a map key.
- In maps, the values are not unique like keys and can be of any type like int, float64, rune, string, pointer, reference type, map type, etc.
- The type of keys and type of values must be of the same type, different types of keys and values in the same maps are not allowed. But the type of key and the type values can differ.
- The map is also known as a hash map, hash table, unordered map, dictionary, or associative array.
- In maps, you can only add value when the map is initialized if you try to add value in the uninitialized map, then the compiler will throw an error.
---
# Structs
- A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. 
- Any real-world entity which has some set of properties/fields can be represented as a struct. 
- This concept is generally compared with the classes in object-oriented programming. 
- It can be termed as a lightweight class that does not support inheritance but supports composition.
---
# Functions
- Functions are generally the block of codes or statements in a program that gives the user the ability to reuse the same code which ultimately saves the excessive use of memory, acts as a time saver and more importantly, provides better readability of the code. 
- So basically, a function is a collection of statements that perform some specific task and return the result to the caller. 
- A function can also perform some specific task without returning anything.
---
# Methods
- Go language support methods. 
- Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it. With the help of the receiver argument, the method can access the properties of the receiver. 
- Here, the receiver can be of struct type or non-struct type. 
- When you create a method in your code the receiver and receiver type must be present in the same package. 
- And you are not allowed to create a method in which the receiver type is already defined in another package including inbuilt type like int, string, etc. 
- If you try to do so, then the compiler will give an error.
---
# Defer
- In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns. 
---
# Panic
- In Go language, panic is just like an exception, it also arises at runtime. Or in other words, panic means an unexpected condition arises in your Go program due to which the execution of your program is terminated. 
- The panic function is an inbuilt function which is defined under the builtin package of the Go language.
---
# Goroutines
- Go language provides a special feature known as a Goroutines. 
- A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program.
---
# Channels
- In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. 
- In other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel 
---
# Exercise (Programs)
	1.  Add two numbers.  
	2.  swap two numbers(with extra variable & without extra variable).
    3.  Check if a number is palindrome or not
    4.  Find all Palindrome numbers in a given range
    5.  Check if a number is prime or not
    6.  Prime numbers in a given range
    7.  Check if a number is armstrong number of not
    8.  Check if a number is perfect number
    9.  Even or Odd
    10. Check weather a given number is positive or negative
    11. Sum of first N natural numbers
    12. Find Sum of AP Series
    13. Program to find sum of GP Series
    14. Greatest of two numbers
    15. Greatest of three numbers
    16. Leap Year or not
    17. Reverse digits of a number
    18. Maximum and Minimum digit in a number
    19. Print Fibonacci upto Nth Term
    20. Factorial of a number

---
# References

***Hitesh Choudhary***  
    https://www.youtube.com/watch?v=JoJ8Sw5Yb4c&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa

***GeeksForGeeks***  
    https://www.geeksforgeeks.org/golang

***freeCodeCamp.org***  
    https://www.youtube.com/watch?v=YS4e4q9oBaU&t=13552s
