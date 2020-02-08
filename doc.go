/*
Package nexus focuses on simplifying error checking in go software.

Definition of the word 'nexus', according to WordNet

      n 1: the means of connection between things linked in series
           [syn: {link}, {nexus}]
      2: a connected series or group

Examples

When writing multiple values sequentially and you want to stop
at the first error.

  p, err := nexus.NewPrinter(os.Stdout)
  p.Print("My long")
  p.Println("text")
  p.Printf("Hello, %s", "World")
  return *err

One go idiom is to return the result of a function together with an
error if one occured. It's a powerful idiom and works in most cases.
When you are working with sequential flows however, the idiom is less
suitable. You end up disrupting the flow with multiple error checks.

Redesign your funcs to only return an error and supply the result holder
as an argument, ie.

  func double(result *int, v int) error {...}
  func toString(result *string, v int) error {...}
  func concat(result *string, a,b string) error {...}

and then use

  var (
     a = 1
     b = 3
     da, db int
     as, bs string
     result string
  )
  err := nexus.First(
     double(&da, a),
     double(&db, b),
     toString(&as, da),
     toString(&bs, db),
     concat(&result, as, bs),
  )

The downside being that each step is actually executed. No silver bullet
here either.

*/
package nexus
