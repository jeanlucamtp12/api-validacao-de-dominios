//Author: Jeanluca Martins de Abreu
//Documentação da API: https://documenter.getpostman.com/view/20002232/UVsJw7Eo

package main

//importe da função net/http
import (
	"net/http"
)


func main() {
    http.HandleFunc("/", Handler)              //especifica a rota do servidor, e chama a função Handler para trabalhar com as rotas 
    http.ListenAndServe(":8080", nil)          //inicia o server, informando a porta do servidor local e nil para definir o servidor padrao 
}

