package modelos

// Senha representa formato da requisicao de alteração de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
