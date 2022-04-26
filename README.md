# Klever API Challenge

## Seja bem-vindo(a)!
Olá, este repositório foi feito baseado no desafio proposto pela [Klever](https://www.linkedin.com/company/klever-br/?trk=similar-pages&originalSubdomain=mt).

## Descrição e requisitos do projeto:
The Technical Challenge consists of creating an API with Golang using gRPC with stream pipes that exposes an upvote service endpoints.

Technical requirements:

- [X] Keep the code in Github

API:

- [X] The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
- [X] The structs used with your mongo model should support Marshal/Unmarshal with bson, json and struct
- [X] The API should contain unit test of methods it uses

Extra:

- [ ] Deliver the whole solution running in some free cloud service

# Utilizando a aplicação

## Clone o Repositório
- Faça o clone do repositório na sua máquina por meio da chave SSH, basta rodar o comando:

``` 
git clone git@github.com:LuizFJP/KLEVER-CHALLENGE-currency-coin.git
```

# Instalação do Insomnia

- **O que é?** 
  Uma ferramenta cliente de API REST, como o Postman, mas tem alguns recursos adicionais, como suporte a GraphQL, gRPC, entre outros.

- **Por quê?** 
  Para ter acesso ao server da aplicação, faz-se necessário a utilização de um client como meio de envio das requisições.

  **[Site oficial para a instalação](https://docs.insomnia.rest/insomnia/install/)**

---------
## Setup e server

Para **instalar as bibliotecas** necessárias da aplicação, rode na raiz do projeto o comando: ```make mod``` 

Para **inicializar o servidor**, rode na raiz do projeto o comando: ```make server```

### Pronto!!! Seu servidor já está pronto. É esperada a seguinte mensagem:
```
Currency Coin Service Started
2022/04/26 05:20:11 Listening on 0.0.0.0:50051
Starting Server...
```
---
## Testes
Para checar se a aplicação está funcionando devidamente, rode na raiz do projeto o comando: ```make test```

---
# Setup do Insomnia

Abra o Insomnia e clique no botão **Create** e logo em seguida no botão **Request Collection**

------
![step-one](/assets/images/1.png)

Defina o nome de preferência ou deixe como padrão. Logo em seguida, clique **Create**.

------
![step-one](/assets/images/2.png)

Aperte no botão **+** e depois em **New Request**

----
![step-one](/assets/images/3.png)

Defina o nome de preferência ou deixe como padrão. Logo em seguida, clique em **GET** e selecione **gRPC**.

----
![step-one](/assets/images/4.png)

Clique no botão **Add Proto File** e selecione o arquivo **service.proto**, localizado em ```./proto``` referente a raiz onde o projeto foi clonado.

----
![step-one](/assets/images/5.png)

No local indicado na seta 1, insira o valor *localhost:50051*, fazendo isso o client do Insomnia conseguirá se conectar com o nosso server.
Por fim, selecione um dos métodos da aplicação gRPC e estará pronto para uso.

----

![step-one](/assets/images/6.png)

----
# Sobre os métodos

- **/CurrencyCoinService/createCoin**

Responsável pela criação de uma moeda. Envia uma requisição com nome e preço. O retorno deve ser o nome, preço e a quantidade de votos (não informado pelo usuário, moedas criadas recebem o valor 0 aos votos automaticamente).

Exemplo de requisição:
```
{
	"name": "BITCOIN",
	"price": 5.467
}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 5.467,
  "vote": 0
}
```

- **/CurrencyCoinService/ListCoins**

Responsável pelo retorno de todas as moedas. Envia uma requisição vazia e será retornado todas as moedas.

Exemplo de requisição:
```
{}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 5.467,
  "vote": 0
}
```

Exemplo de response:
```
{
	"name": "Ethereum",
	"price": 942.78,
  "vote": 91
}
```

Exemplo de response:
```
{
	"name": "Shiba",
	"price": 00000.1,
  "vote": 4984
}
```

- **/CurrencyCoinService/UpvoteCoin**

Responsável por aumentar o número de votos em 1. O corpo da requisição deve ser o nome da moeda e o retorno será o nome, preço e voto atualizado.
**ATENÇÃO: A API diferencia letras maiúsculas e minúsculas, por isso é importante inserir o nome da mesma forma como está cadastrado**

Exemplo de requisição:
```
{"name":"BITCOIN"}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 5.467,
  "vote": "1"
}
```

- **/CurrencyCoinService/DownvoteCoin**

Responsável por decrementar o número de votos em 1. O corpo da requisição deve ser o nome da moeda e o retorno será o nome, preço e voto atualizado.
**ATENÇÃO: A API diferencia letras maiúsculas e minúsculas, por isso é importante inserir o nome da mesma forma como está cadastrado**

Exemplo de requisição:
```
{"name":"BITCOIN"}
```

Exemplo de response:
```
{
	"name": "BITCOIN",
	"price": 5.467,
  "vote": 0
}
```

- **/CurrencyCoinService/Delete**

Responsável por deletar uma moeda. O corpo da requisição deve ser o nome da moeda e o retorno será uma mensagem confirmando a deleção.
**ATENÇÃO: A API diferencia letras maiúsculas e minúsculas, por isso é importante inserir o nome da mesma forma como está cadastrado**

Exemplo de requisição:
```
{"name":"BITCOIN"}
```

Exemplo de response:
```
{
	"message": "BITCOIN was deleted successful!
}
```

### Fontes de estudo:
[Documentação do mongoDB](https://www.mongodb.com/docs/) \
[Documentação do gRPC](https://grpc.io/docs/) \
[Documentação do GO](https://go.dev/doc/) \
[Documentação do mongodb driver](https://www.mongodb.com/docs/drivers/go/current/) \
Indispensável mencionar o uso de diversos tutoriais no Youtube, StackOverFLow, GitHub e o Médium. \

---
## Sugestões e feedbacks serão sempre bem-vindos :)

