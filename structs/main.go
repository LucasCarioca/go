package main

func main() {
	karen := Person{"Karen", "Ardila", Contact{"karen@gmail.com", 123456}}
	lucas := Person{
		firstname: "Lucas",
		lastname:  "Desouza",
		contact: Contact{
			email:   "lucas@gmail.com",
			zipcode: 123456,
		},
	}
	var toto Person
	toto.firstname = "Toto"
	toto.lastname = "Desouza"
	toto.contact = Contact{
		email:   "toto@gmail.com",
		zipcode: 12345,
	}
	karen.print()
	toto.print()
	lucas.print()

	lucas.updateName("Lucas Franklin")
	lucas.print()
}
