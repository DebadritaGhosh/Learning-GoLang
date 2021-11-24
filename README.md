
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


#
