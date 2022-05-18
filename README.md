# API_Contatos

Tem como objetivo receber um ou mais contatos através de uma API Rest e adicioná-los ao banco de dados especifico para cada cliente

## Requisitos

- Docker
- Docker-compose 1.29 ou compátivel com versão 3.8 do docker-compose.yml
- SGDB (Sistema Gerenciados de Banco de Dados) compátivel para Postgres e MYSQL
indico o Beekeeper, fácil de usar
- Postman para testes de endpoints

## Como executar o projeto?

- Antes de mais nada abra seu terminal e clone este repositorio em sua maquina

- Em seguida suba os containers rodando o comando abaixo na pasta raiz do projeto:

```docker-compose up -d```

- Entre na pasta '/site' e crie um arquivo .env para passar as credenciais dos serviços (Postgres/MYSQL) e da porta. Obrigatório seguir o seguinte padrão:

```
API_PORT=

POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

MYSQL_USER=
MYSQL_ROOT_PASSWORD=
MYSQL_DATABASE=

SECRET_KEY=
```
- Você deverá criar uma Secret Key para assinar o token

- Em seguida ainda na pasta site rode o seguinte comando no terminal:

```go run router.go```

*Prontinho, sua aplicação está no ar!*

## Como testar a aplicação?

- Primeiro vamos configurar nossos bancos

- Crie uma tabela de usuario no seu banco MYSQL 

 ```  
    CREATE table contacts (
	id serial PRIMARY KEY,
	nome VARCHAR ( 200 ) NOT NULL,
	senha VARCHAR ( 200 ) NOT NULL
);
```
- Depois popule essa tabela com dois usuarios sendo eles Macapá e Varejao

Aproveitando que esta no MYSQL ja crie também a tabela do Cliente Macapá

```
    CREATE table contacts (
	id serial PRIMARY KEY,
	nome VARCHAR ( 200 ) NOT NULL,
	celular VARCHAR ( 20 ) NOT NULL
);  
```

- Em seguida entre no banco Postgres e crie a tabela do Cliente Varejao
 
 ```
    CREATE table contacts (
	id serial PRIMARY KEY,
	nome VARCHAR ( 100 ) NOT NULL,
	celular VARCHAR ( 13 ) NOT NULL
);
 ```

- Por padrão sua aplicação irá rodar na porta 5000 ```http://localhost:5000```

- A aplicação contém um prefixo rota '/api' ```http://localhost:5000/api```

*Tudo pronto, vamos testar!*

- Primeiro efetue o login do usuario ultilizando o methodo POST```http://localhost:5000/api/login``` adicionando um json no body da requisição, exemplo:

```
{
    "nome":"varejao",
    "senha":"654321"
}
```
- Você receberá um token autenticado como resposta, ultilize este token para ter acesso ao  próximo endpoint cuja o methodo também é POST```http://localhost:5000/api/recebe-dados``` é só inserir ele no 'Barer Token' na aba 'Authorization' do Postman
e enviar um body de acordo com a estrutura:

```
{
    "contacts": [
        {
            "name": "Thauan Mendes",
            "cellphone": "5511960327601"
        },
     ]
}
```
