package main

import (
	"bytes"
    "encoding/json"
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/Thanghand/blockchain"
	"encoding/base64"
	"fmt"
)


func main() {
	// Init blockchain
	blockchain.InitBlockchain()

	// Init Router
	fmt.Printf("Server started with port 8069")
	router := mux.NewRouter()
	router.HandleFunc("/blockchain", AddNewBlockAPI).Methods("POST")
	router.HandleFunc("/blockchain/lastBlock", GetLastBlockAPI).Methods("GET")
	log.Fatal(http.ListenAndServe(":8069", router))
	
}

// BlockRequest is request block 
type BlockRequest struct {
	Text string `json:"text,omitempty"`
	OldHash string `json:"oldHash,omitempty"`
}

// ResponseBlockChain is response message
type ResponseBlockChain struct {
	Message string `json:"message,omitempty"`
	HashBlock []byte `json:"hashBlock,omitempty"`
}
// AddNewBlockAPI router 
func AddNewBlockAPI (w http.ResponseWriter, r *http.Request){

	var blockRequest BlockRequest
	var response ResponseBlockChain
	_ = json.NewDecoder(r.Body).Decode(&blockRequest)

	data, err := base64.StdEncoding.DecodeString(blockRequest.OldHash)

	lastBlock := blockchain.GetLastBlock()
	if bytes.Compare(lastBlock.Hash, data) == 0{
		newBlock := blockchain.NewBlock(lastBlock, blockRequest.Text)
		errAddBlock := blockchain.AddBlock(newBlock)
		
		
		if errAddBlock == nil {
			response.Message = "Add new block successfully with text : " + blockRequest.Text
			response.HashBlock = newBlock.Hash
			json.NewEncoder(w).Encode(response)
		} else {
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
		}
	} else {
		response.Message = "Add new block error"
		json.NewEncoder(w).Encode(response)
	}
	
}

// GetLastBlockAPI is a function to get last block in BlockChain
func GetLastBlockAPI (w http.ResponseWriter, r *http.Request){
	lastBlock := blockchain.GetLastBlock()
	var response ResponseBlockChain
	response.Message = "Get Last Block successfully"
	response.HashBlock = lastBlock.Hash
	json.NewEncoder(w).Encode(response)	
}