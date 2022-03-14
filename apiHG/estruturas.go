package main

//estrutura utilizada para armazenar os dados retornados pelo endpoint

type DadosJson struct {
	Status            int      `json:"status"`
	Fqdn              string   `json:"fqdn"`
	//Fqdnace           string   `json:"fqdnace"`
	//Exempt            bool     `json:"exempt"`
	//Hosts             []string `json:"hosts"`
	//PublicationStatus string   `json:"publication-status"`
	//ExpiresAt         string   `json:"expires-at"`
	//Suggestions       []string `json:"suggestions"`
	Reasons           []string `json:"reasons"`
	
}

type ConteudoImpressao struct {
	Available         string     `json:"Available"`
	Price             float32     `json:"Price:"`
	Reasons           string     `json:"Reasons"`
}

//construtor da estrutura DadosJson
func (s DadosJson) Data(fdqn string, price float32, reasons string) bool{
	return true
}


//construtor da estrutura ConteudoImpressao
func (s ConteudoImpressao) Data(available string, price float32, reasons string) bool{
	return true
}


