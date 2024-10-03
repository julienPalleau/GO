// https://www.boot.dev/lessons/7f503d3f-7425-496a-b50a-70815384a00c

/*
Multiple Parameters

When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

Here are some examples:

func addToDatabase(hp, damage int) {
  // ...
}

func addToDatabase(hp, damage int, name string) {
  // ?
}

func addToDatabase(hp, damage int, name string, level int) {
  // ?
}

*/

/*
Which of the following is the most succinct (and valid) way to write a function signature?

func createUser(firstName, lastName string, age int)
*/