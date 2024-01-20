package main

import "fmt"

type contactInfo struct { // its not variable declaration it's type declaration using this type we can create variable
	email   string
	zipCode int
}

type Person struct {
	first_name string
	last_name  string
}

type contactInfoParent struct {
	street      string
	city        string
	state       string
	contactInfo contactInfo // you can simply give contactInfo if key and value are same you can omit value but i mentioned value
	mobile      mobileNumberDetails
}
type mobileNumberDetails struct {
	mobileNumber string
	email        string
	zipcode      string
}

func main() {

	//first way to define variable

	// person := Person{"Dharmadurai", "Muthusamy"}
	// fmt.Println(person)

	//second way to define variable

	// person := Person{first_name: "Dharmadurai", last_name: "Muthusamy"}
	// fmt.Printf("person struck %+v", person)

	//third way to create

	var person Person

	person.first_name = "Dharmadurai"
	person.last_name = "Muthusamy"
	// fmt.Printf("%+v", person)

	contactDetails := contactInfoParent{
		street: "west street",
		city:   "Karur",
		state:  "TN",
		contactInfo: contactInfo{
			email:   "dharmadurai@gmail.com",
			zipCode: 621313,
		},
		mobile: mobileNumberDetails{
			mobileNumber: "98654789658",
			email:        "dharma@gmail.com",
			zipcode:      "6321248",
		},
	}

	fmt.Printf("%+v", contactDetails)

	//first way to update the details
	// contactRef := &contactDetails //get contactDetails pointer
	// contactRef.updateCity()

	//second way to update the details
	customStruct()
	contactDetails.updateCity()

}

//pointer

/*
  &variablename = > give me the memory address of the value this variable is pointing at
  *pointer = > give me the value this memory address pointing at.
*/

//update city using pointer

func (contactInfo *contactInfoParent) updateCity() {
	(*contactInfo).city = "Kulithalai"
}

//custom struct
type ContactInfo struct {
	mobileNumber string
	email        string
}

type basicAddressInfo struct {
	street      string
	city        string
	district    string
	state       string
	pincode     string
	contactInfo ContactInfo
}
type priyaAddress struct {
	addressOne basicAddressInfo
	addressTwo basicAddressInfo
}

func customStruct() {
	priyaAddr := priyaAddress{
		addressOne: basicAddressInfo{
			street:   "West street",
			city:     "Kulithalai",
			district: "karur",
			state:    "Tamilnadu",
			pincode:  "621313",
			contactInfo: ContactInfo{
				mobileNumber: "9874563210",
				email:        "priya@zohomail.com",
			},
		},
		addressTwo: basicAddressInfo{
			street:   "West street",
			city:     "Manaparai",
			district: "Trichy",
			state:    "Tamilnadu",
			pincode:  "6000010",
			contactInfo: ContactInfo{
				mobileNumber: "9630125874",
				email:        "priya.dharma@zohomail.com",
			},
		},
	}

	priyaAddr.printEmail()
}

//attach a receiver to the struct

func (addr priyaAddress) printEmail() {

	fmt.Println(addr.addressOne.contactInfo.email)
	fmt.Println(addr.addressTwo.contactInfo.email)
}
