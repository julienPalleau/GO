// https://www.boot.dev/lessons/4c0f0f05-38dc-45c3-af26-802434325407
/*
Compiled vs Interpreted

You can run a compiled program without the original source code. You don't need the compiler anymore after it's done its job. That's how most videogames are distributed! Players 
don't need to install the correct version of python to run a PC game: they just download the executable game and run it.

compiler vs interpreter

With interpreted languages like Python and Ruby, the code is interpreted at runtime by a separate program known as the "interpreter". Distributing code for users to run can be a pain 
because they need to have an interpreter installed, and they need access to the source code.
Examples of compiled languages

    Go
    C
    C++
    Rust

Examples of interpreted languages

    JavaScript (sometimes JIT-compiled, but a similar concept)
    Python
    Ruby

Why build Textio in a compiled language?

One of the most convenient things about using a compiled language like Go for Textio is that when we deploy our server we don't need to include any runtime language dependencies like 
Node or a Python interpreter. We just add the pre-compiled binary to the server and start it up!
*/

/*
Which language is interpreted?

Python
*/