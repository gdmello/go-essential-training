*Notes*

* Golang functions by default _pass by value_, slices _pass by reference_
* _package main_ indicates that the go program will be an executable
* _go.mod_ indicates a go module, and all import paths will begin with the module path. 
* Struct methods - A method belonging to a type must use a "Pointer Receiver" to receive the reference to the object of that type, instead of the value of that type
```go
type Point struct {
  x int,
  y int
}

func (p Point) Move(dx, dy) {
  p.x += dx 
  p.y += dy
}

func (p *Point) MovePoint(dx, dy) {
  p.x += dx 
  p.y += dy
}

func main() {
  p := Point{1,2}
  p.Move(2,3)
  fmt.Printf("%+v",p) // Prints {x:1, y:2} since p is now passed by value
  p.MovePoint(2,3)
  fmt.Printf("%+v",p) // Prints {x:3, y:5} since p is now passed by reference
}
```
* Points - 
  * _&_ generates a pointer to the operand (the object on which the operation is done)
  ```go
    i := 7
    p := &i // p is given a pointer to i
  ```
  * _*_ dereferences a pointer - i.e. used when referencing the object behind the pointer 
  ```go
    fmt.Println(*p) // dereferences p and prints 7 (the value of i)
  ```  
