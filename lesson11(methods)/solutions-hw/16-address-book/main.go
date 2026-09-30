package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}

type AddressBook struct {
	contacts map[string]Contact
}

func (b *AddressBook) Add(key string, c Contact) {
	if b.contacts == nil {
		b.contacts = make(map[string]Contact)
	}
	b.contacts[key] = c
}

func (b AddressBook) Find(key string) (Contact, bool) {
	c, ok := b.contacts[key]
	return c, ok
}

func (b *AddressBook) Delete(key string) {
	delete(b.contacts, key)
}

func main() {
	book := &AddressBook{}

	book.Add("mama", Contact{Name: "Мама", Phone: "+7 700 111 22 33"})
	book.Add("boss", Contact{Name: "Начальник", Phone: "+7 700 444 55 66"})

	if c, ok := book.Find("mama"); ok {
		fmt.Println("Найден:", c.Name, c.Phone)
	}

	book.Delete("boss")

	if _, ok := book.Find("boss"); !ok {
		fmt.Println("Контакт 'boss' не найден — уже удалён")
	}

	if _, ok := book.Find("friend"); !ok {
		fmt.Println("Контакт 'friend' не найден")
	}
}
