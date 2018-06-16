# Readme
## FizzBuzz in Golang
Written by Charlene Bastians and Sam Niechial

### What is FizzBuzz
Numbers from 1 to 100 are printed. If a number is divisible by 15, "FizzBuzz" replaces the number. If the number is divisible by 3, "Fizz" replaces the number. If a number is divisible by 5, "Buzz" replaces the number.


### How and why we made this
We wanted to familiarise ourselves with Golang and it's accompanying testing frameworks, Ginkgo and Gomega. We wanted to build something small so we could get to grips with the basic syntax and structure of the language. After doing some research, we played around in gore (a golang repl) and got a very basic version of fizzbuzz working. After that, we did some research into setting up a go programme and set up a basic project. We installed Golang, Ginkgo and Gomega and followed the associated docs we founds online to fully test our small programme and concentrated on further encapsulating the version we spiked on gore.

### Technologies used
Golang, Ginkgo, Gomega and Gore for practice.

### How to run
Assuming you have golang installed and go added to your path environment variables, from the command line run ``` go run fizzBuzz.go``` to see the full print out of numbers. To run the tests, run ``` go test ``` or ```ginkgo```.
