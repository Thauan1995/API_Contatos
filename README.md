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

- Em seguida suba os containers rodando o comando abaixo na pasta raiz do projeto

```docker-compose up -d```

- Entre na pasta '/site' e digite o seguinte comando

```go run router.go```

*Prontinho, sua aplicação está no ar!*

## Como testar a aplicação?

- Sua aplicação irá rodar na porta 5000 ```http://localhost:5000```

- A aplicação contém um prefixo rota '/api' ```http://localhost:5000/api```

- Primeiro será necessário criar uma tabela de usuario no seu banco MYSQL 

    CREATE table contacts (
	id serial PRIMARY KEY,
	nome VARCHAR ( 200 ) NOT NULL,
	senha VARCHAR ( 200 ) NOT NULL
);
