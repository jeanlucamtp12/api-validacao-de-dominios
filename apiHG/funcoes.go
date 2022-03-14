package main

import (                  //importa os pacotes necessarios 
    "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"strings"
	"regexp"
)

func validaTamanhoString (tam int) bool{           //recebe um numero e valida se ele é > ou < que 26, retorna true ou false
	if (tam > 26){                   
		return false
	}
	return true
}

func checaString (dominio string) bool {                 //recebe um nome de dominio 

	err := regexp.MustCompile(`\S*[^.com.br]`)           //regex que realiza a remoção do tld: .com.br
	verificador := err.FindAllString(dominio, -1)        //funçao findAllString retorna uma slice para verificador

	sldDominio := strings.Join(verificador, "")          //realiza a conversão da slice em uma string, se tornando assim um sld

	for _, i := range sldDominio {                       //for para verificar se o sld se encaixa nos requisitos, conter apenas a-z e 0-9 sem caracteres especiais 
		if (i < 'a' || i > 'z'){
			if (i < '0' || i > '9'){
			    return false                             // caso contenha caracteres especiais, retorna false 
			}
		}
	}
    //apos verificar a inexistencia de caracteres especiais, chama a funçao validaTamanhoString, para validar o tamanho do sld
	var tamanhoString = validaTamanhoString(len([]rune(sldDominio)))  // fução len retorna o tamanho da string 
	return tamanhoString                                              // retorna false para string maior que 26 e true para menor 
}

func verificaTld (w http.ResponseWriter,dominioCompleto string) bool{   //funçao para validar se o domínio informado contem TLD
	if !(strings.Contains(string(dominioCompleto), ".com")) {           // função contains verifica se: .com está incluso no dominio
		formataJson("invalid domain", 0 , "erro", w)                         //se não estiver incluso, printa que é um dominio invalido e retorna false
		return false
	}
	return true                                          //retorna true para um dominio com tld valido
}  

func pesquisaEndpoint (dominioCompleto string, w http.ResponseWriter) {              //função responsavel por realizar a pesquisa no endpoint

	endereco := "https://registro.br/v2/ajax/avail/raw/"                             
	busca := endereco + dominioCompleto                                              //junta a url do endpoint + o dominio

	response, err := http.Get(busca)                                                 //captura as informações retornadas pelo endpoint 

	if err != nil {                                                                  //verifica a possibilidade de erro durante a execução 
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)                             // responseData realiza a leitura do corpo da resposta
	if err != nil {
		log.Fatal(err)
	}

	var responseObject = DadosJson{}                                               //criação de uma variavel do tipo DadosJson

    json.Unmarshal(responseData, &responseObject)                                  //realiza a decodificação do Json e atribui os dados a struct response Object
	exibiResultados(responseObject, dominioCompleto, w)                            // chama a função exibir resultados

}

func formataJson(available string, price float32, reason string, w http.ResponseWriter){     //função responsavel por criar a Json para impressao

	conteudoExibicao := ConteudoImpressao {                                                  //cria a struct que sera usada para impressao
		Available : available,
		Price: price,                                                                        //faz a atribução dos dados recebidos
		Reasons: reason,
	} 

	r, err := json.MarshalIndent(conteudoExibicao,""," ")                                    //usa a função Marshal para tranformar a struct em Json
	if err != nil{
		fmt.Println(err)
	} 

	fmt.Fprint(w,string(r))                                                                 //printa a Json para o usuario

}

func exibiResultados (responseObject DadosJson, dominioCompleto string, w http.ResponseWriter){       //função responsavel por exibir os dados 

	var result = checaString(dominioCompleto)                                                         //chama a função que checa se o dominio é valido 

	if (result  == false){                                                                             //caso o retorno seja falso, o domínio é printado como invalido 
		formataJson("invalid domain", 0 , "erro", w)                                                      //passa os dados para a funcao formataJson realizar a impressao
	}else {
		
		if(strings.Contains(responseObject.Fqdn, "hostgator") || strings.Contains(responseObject.Fqdn, "endurance")  ){   //verifica se o domínio contem as palavras reservadas, hostgator e endurance
			
			formataJson("false", 0 , "Registered trademark", w)                                                          //chama a função reponsavel por imprimir os dados
			return
		}
	
    	if(responseObject.Status != 2){                                        // padrao encontrado/ caso o status do domínio retornado pelo endpoint seja diferente de 2, este é um domínio não registrado
			if(strings.Contains(responseObject.Fqdn, "loja") || strings.Contains(responseObject.Fqdn, "premium")  ){    //verifica se o dominio contem palavras reservadas 
				formataJson("yes", 99.90 , "within the norms", w)                                                       //printa as informações atraves da funcao 
				                         
    		}else{
	    		formataJson("yes", 26.90 , "within the norms", w)
    		}
        }else {                                                              //caso o dominio já esteja registrado, printa a informação 
			formataJson("no", 0 , "domain registered", w)
        }


	}


}

func Handler(w http.ResponseWriter, r *http.Request) {          //função Handler, possui dois argumentos ResponseWriter: Resposta do pedido e Resquest: Requisição ao pedido 
	
    keys, ok := r.URL.Query()["domain"]                         //usa URL.Query para capturar o que contem o campo "domain", digitado no browser pelo usuario
    
    if !ok || len(keys[0]) < 1 {                                //verifica se o usuario não digitou um dominio vazio, caso tenha digitado, é retornado que este é invalido 
		formataJson("invalid domain", 0 , "erro", w)
        return
    }

    key := keys[0]
	var dominioCompleto = string(key)                          // transforma key em uma string 

	var retorno = verificaTld(w, dominioCompleto)               //chama a função verifica Tld e armazena o retorno na variavel 
	if retorno == false {                                       //verifica se o retorno é false ou true/ false: dominio sem tld, true: dominio com tld
		return 
	}else {                                                     //caso o dominio contenha o tld, as tratativas continuam
		pesquisaEndpoint (dominioCompleto, w)                   //chama a função endpoint e passa o dominio 
	} 
}