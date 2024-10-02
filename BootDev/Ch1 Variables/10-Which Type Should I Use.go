// https://www.boot.dev/lessons/fc3fe54e-6106-44a6-abe0-7cf2961c240c

/*
Which Type Should I Use?

With so many types for what is essentially just a number, developers coming from languages that only have one kind of Number type (like JavaScript) may find the choices daunting.
Prefer "default" types

A problem arises when we have a uint16, and the function we are trying to pass it into takes an int. We're forced to write code riddled with type conversions like:

var myAge uint16 = 25
myAgeInt := int(myAge)

This style of code can be slow and annoying to read. When Go developers stray from the “default” type for any given type family, the code can get messy quickly. Unless you have a good performance related reason, you'll typically just want to use the "default" types:

    bool
    string
    int
    uint
    byte
    rune
    float64
    complex128

When should I use a more specific type?

When you're super concerned about performance and memory usage.

That’s about it. The only reason to deviate from the defaults is to squeeze out every last bit of performance when you are writing an application that is resource-constrained. (Or, in the special case of uint64, you need an absurd range of unsigned integers).

You can read more on this subject here if you'd like, but it's not required.
*/

/*
What does the size of a type indicate?

Bits
*/