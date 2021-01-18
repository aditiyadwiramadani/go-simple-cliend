package main

import (
	"encoding/json"
	"fmt"
	"log"
    "strings"
    "Cliend/cliend"
    
)

type Data struct { 
    NameProduct string `json:"name_product"`
    Price int  `json:"price"` 
}



func main() { 
     url := "http://localhost:8000"
     
     data := cliend.GetData(url)
     dec :=json.NewDecoder(strings.NewReader(data))
     dec.Token()

     for dec.More() {
         d := Data{}
         err := dec.Decode(&d)
         if err != nil { 
             log.Fatal(err)
         }
         fmt.Println(d.NameProduct, d.Price)
     }
    
     
}








