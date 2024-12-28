package main

type Book struct {
	Id          string `json:"id"`
	Title       string `json:"title`
	Author      string `json:"author"`
	PublishDate string `json:"publishdate`
	Isbn        string `json:"isbn"`
}

type BookCheckout struct {
	BookId       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkoutdate"`
	Genesis      bool   `json:"genesis"`
}

type Block struct {
	Position     string
	Data         BookCheckout
	TimeStamp    string
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	blocks []*Block
}

var blockchain *Blockchain
