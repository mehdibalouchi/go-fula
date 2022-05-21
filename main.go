package main

import (
	"fmt"
	fula "github.com/farhoud/go-fula/mobile"
	"log"
	"os"
	"os/signal"
	"runtime"
	// "time"
)



func main() {

	fula,_ := fula.NewFula("/home/farhoud")
	fula.AddBox("")
	fmt.Println("We are know connected")
	cid,err := fula.Send("/home/farhoud/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("cid", cid)
	meta,err := fula.Receive(cid)
	if err != nil {
		panic(err)
	}
	fmt.Println(meta)
	query := `
	mutation addTodo($values:JSON){
	  create(input:{
		collection:"todo",
		values: $values
	  }){
		id
		text
		isComplete
	  }
	}
  `
// query := `
// query {
//   read(input:{
// 	collection:"todo",
// 	filter:{text: {eq: "todo2"}}
//   }){
// 	id
// 	text
// 	isComplete
//   }
// } 
// `

// [
// 			{id: "1", text: "todo1", isComplete: false},
// 			  {id: "2", text: "todo2", isComplete: false}
// 		  ]
	//   var values map[string]interface{}
	//   values = append(values, map[string]interface{}{"id": "1", "text": "todo1", "isComplete": false})
	  
	// values := map[string]interface{}{
	// 	"values": []interface{}{
	// 		map[string]interface{}{"id": "1", "text": "todo1", "isComplete": false},
	// 		map[string]interface{}{"id": "2", "text": "todo2", "isComplete": true}}}

	values := `{"values": [{"id": "1", "text": "todo1", "isComplete": false}, {"id": "2", "text": "todo2", "isComplete": true}]}`
			
	res, err := fula.GraphQL(query, values)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	runtime.Goexit()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case <-c:
			log.Printf("Close gracefully")
			signal.Stop(c)
			os.Exit(0)
		}
	}()
	fmt.Println("Exit")
	fmt.Println("R u running")

}
