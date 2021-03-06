package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/utils"
)

/*func init() {
	HashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(HashKey)

	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockkey)

}*/

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	//	fmt.Printf()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
