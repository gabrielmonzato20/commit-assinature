type Car struct{
Name string

}

func (c Car) run(){
fmt.Println(c,Name, "run")
}

func main(){
  car := Car{
    Name:"Tesla model 3",
  }
  car.run()

}
