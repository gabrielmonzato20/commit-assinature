type Car struct{
Name string

}

func (c Car) run(){
fmt.Println(c,Name, "run")
}

func main(){
  car := Car{
    Name:"1Tesla",
  }
  car.run()

}
