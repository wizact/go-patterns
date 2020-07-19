package main

import (
	"fmt"
	pb "github.com/wizact/go-patterns/encoding/protobuf/addressbookpb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

const FILENAME = "addressbook.dat"

func main() {
	writeAddress()
	readAddress()
}

func writeAddress() {
	phone := &pb.Person_PhoneNumber{Type: pb.Person_MOBILE, Number: "1234321"}
	person := &pb.Person{Name: "Amir", Email: "amir@example.com", Phones: []*pb.Person_PhoneNumber{phone}}
	book := &pb.AddressBook{People: []*pb.Person{person}}

	out, err := proto.Marshal(book)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile(FILENAME, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func readAddress() {
	in, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", FILENAME)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	book := &pb.AddressBook{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(book)
}
