package main

func main(n int) bool {
   
if n == 0 {
   
  return true

 } 
   if n == 0 {

	return false

}

return main(n - 2)

}