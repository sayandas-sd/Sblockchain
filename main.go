package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)

	data := string(b.Position) + b.TimeStamp + string(bytes) + b.PreviousHash

	hash := sha256.New()

	hash.Write([]byte(data))

	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func CreateBlock(prevblock *Block, data BookCheckout) *Block {

	block := &Block{
		Position:     prevblock.Position,
		Data:         data,
		TimeStamp:    time.Now().Format(time.RFC3339),
		PreviousHash: prevblock.Hash,
	}
	block.generateHash()

	return block
}

func (b *Blockchain) AddBlock(data BookCheckout) {

	prevBlock := b.blocks[len(b.blocks)-1]

	block := CreateBlock(prevBlock, data)

	if validBlock(block, prevBlock) {
		b.blocks = append(b.blocks, block)
	} else {
		log.Println("Invalid block: not added to the blockchain")
	}

}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PreviousHash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Position != block.Position {
		return false
	}

	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()

	if b.Hash != hash {
		return false
	}
	return true
}

func newBlock(w http.ResponseWriter, r *http.Request) {

	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error to create: %v", err)
		w.Write([]byte("error to create a new block"))
		return
	}

	new := md5.New()
	io.WriteString(new, book.Isbn+book.PublishDate)
	book.Id = fmt.Sprintf("%x", new.Sum(nil))

	res, err := json.MarshalIndent(book, "", "")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("couldnot parse the json %v", err)
		w.Write([]byte("could not save data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func writeBlock(w http.ResponseWriter, r *http.Request) {
	var bookCheckout BookCheckout

	if err := json.NewDecoder(r.Body).Decode(&bookCheckout); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error to create %v", err)
		w.Write([]byte("couldnot write a block"))
		return
	}

	blockchain.AddBlock(bookCheckout)

	// Get the last added block
	newBlock := blockchain.blocks[len(blockchain.blocks)-1]

	// Respond with the new block as JSON
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(newBlock)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)

		err := []byte(`{"error": "Failed to encode response"}`)
		w.Write(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBLock(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(blockchain.blocks, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	io.WriteString(w, string(bytes))
}

//create first genesis block

func GenesisBlock() *Block {
	genesisData := BookCheckout{
		Genesis: true,
	}
	genesisBlock := &Block{
		Position:     "0",
		Data:         genesisData,
		TimeStamp:    time.Now().Format(time.RFC3339),
		PreviousHash: "0",
	}
	genesisBlock.generateHash()
	return genesisBlock
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {

	blockchain = NewBlockChain()

	r := mux.NewRouter()

	r.HandleFunc("/", GetBLock).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBlock).Methods("POST")

	go func() {
		for _, block := range blockchain.blocks {
			fmt.Printf("Prev.hash: %x\n", block.PreviousHash)
			bytes, _ := json.MarshalIndent(block.Data, "", "")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Server is running on port: 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
