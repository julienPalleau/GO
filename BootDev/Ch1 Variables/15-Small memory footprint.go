// https://www.boot.dev/lessons/a011d7b7-209e-43ff-8d6a-aefa26e9772f

/*
Small memory footprint

Go programs are fairly lightweight. Each program includes a small amount of extra code that's included in the executable binary called the Go Runtime. One of the purposes of the Go runtime is to clean up unused memory at runtime. It includes a garbage collector that automatically frees up memory that's no longer in use.
Comparison

As a general rule, Java programs use more memory than comparable Go programs. There are several reasons for this, but one of them is that Java uses a virtual machine to interpret bytecode at runtime and typically allocates more on the heap.

On the other hand, Rust and C programs use slightly less memory than Go programs because more control is given to the developer to optimize the memory usage of the program. The Go runtime just handles it for us automatically.
Idle memory usage

idle memory

In the chart above, Dexter Darwich compares the memory usage of three very simple programs written in Java, Go, and Rust. As you can see, Go and Rust use very little memory when compared to Java.
*/

/*
Generally speaking, which language uses the most memory?

Java
*/