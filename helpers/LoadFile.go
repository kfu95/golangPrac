package helpers

import(
	"io/ioutil"
	"fmt"
	"log"
)

func LoadFile(fileName string)(string,error){
	
	contents, err :=ioutil.ReadFile(fileName);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", contents)
	return string(contents), nil
}