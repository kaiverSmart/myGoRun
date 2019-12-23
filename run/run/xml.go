package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

var t, token xml.Token
var err error

/*这个结构会保存解析后的返回数据。
他们会形成有层级的XML，可以忽略一些无用的数据*/
type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Note  Status
}



func main() {
	fmt.Println("xml start")
	respon,err := http.Get("https://www.runoob.com/try/xml/note_error.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	user := User{xml.Name{"", "user"}, Status{""}}
	but := new(bytes.Buffer)
	 but.ReadFrom(respon.Body)
	xml.Unmarshal(but.Bytes(),&user)
	fmt.Printf("status: %s", user.Note.Text)
	fmt.Println("\n-----buff",string(but.Bytes()))
}


func xmlstr()  {
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	//decode
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	for t, err := p.Token();err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
	}
}