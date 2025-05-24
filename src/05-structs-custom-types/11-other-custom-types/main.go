package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "John Doe"

	name.log()
}

/*
   Explanation:

   ✅ In Go, you can define **custom types** based on built-in types (e.g., `string`, `int`, etc.).
      Here, `str` is a custom type based on `string`.

   ✅ You can attach **methods** to these custom types.

   ❌ However, you CANNOT attach methods directly to built-in types.
      For example, this is NOT allowed:

          func (s string) log() { ... } // ❌ Compile-time error!

      This is because built-in types like `string` are defined in the standard library
      and cannot be modified or extended with methods directly.

   ✅ To work around this, define a new type (like `type str string`) and then attach methods to it.

   ➕ This also improves code safety and clarity, by giving semantic meaning to values.
      For example, `type UserID int` is clearer and safer than just `int`.
*/
