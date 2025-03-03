# API students

<h2> Introdução </h2>
O objetivo da aplicação é criar um sistema para controlar os estudantes cadastrados e ativos. Com base nas requisições de GET, POST, PUT e DELETE para cadastrar, listar, atualizar e deletar os perfis. Usando o banco de dados SQLite.

## Tecnologias utilizadas
* VS Code
* Golang (Go)
* Insomnia ou Postman 
* Echo v4
* SQLite
* GORM
* Zerolog
* Swagger


## Ferramentas
Foi instalado a extensão Go no VS Code com a versão 0.41.2

Foi instalado a extensão SQLite no VS Code com a versão 0.14.1

Foi instalado a extensão SQLite Viewer no VS Code com a versão 0.5.8

Foi instalado a extensão C/C++ no VS Code com a versão 1.20.5

Foi instalado a extensão Makefile Tools no VS Code com a versão 0.9.10

Foi usado  aplicação <a href="https://insomnia.rest/download" target="_blank" > Insomnia </a>  e/ou <a href="https://www.postman.com/downloads" target="_blank" >  postman  </a> para testar as requisições das rotas simulando o Front-end.

Foi usado o framework web <a href="https://github.com/labstack/echo" target="_blank"> ECHO </a> na versão 4.

Foi usado a ferramenta de mapeamento relacional de objetos (ORM) <a href="https://gorm.io/docs/connecting_to_the_database.html" target="_blank"> GORM </a> com conexão ao banco de dados SQLite.

Foi usado o pacote <a href="https://github.com/rs/zerolog" target="_blank"> Zerolog </a> .

Foi usado o <a href="https://github.com/swaggo/echo-swagger" target="_blank"> Swagger </a> com o Echo.

## Rodando o projeto
- Comando para rodar o servidor:
```
go run main.go
```

- Parar de rodar o servidor: no terminal clicar nas teclas de "Ctrl" e "C".

- Comando para inicializar o gerenciador de pacotes (_go.mod_):
```
go mod init
```

- Comando para verificar, atualizar e baixar os módulos no gerenciador de pacotes (go.mod):
```
go mod tidy
```

- Depois do Makefile configurado, o novo comando para rodar o servidor:
```
make run
```

## Rotas

| Método | URL             | Descrição                                        |
| ------ | --------------  | -------------------------------------------------|
| GET    | /estudantes     | Listar todos os alunos                           |
| POST   | /estudantes     | Cadastrar um estudante                           |
| GET    | /estudantes/:id | Pegar a informação de um estudante específico    |
| PUT    | /estudantes/:id | Atualizar informações de um estudante específico |
| DELETE | /estudantes/:id | Deletar um estudante específico                  |


## Estrutura do estudante
- Nome
- CPF
- Email
- Idade
- Ativo

## Swagger
A partir da <a href="https://github.com/swaggo/echo-swagger" target="_blank"> documentação </a>, siga com os seguintes passos:

- Comandos para o download do pacote:
```
go get -d github.com/swaggo/swag/cmd/swag
```
```
go install github.com/swaggo/swag/cmd/swag@latest
```

- Comando para inicializar:
```
swag init
```

- Comando para o download da lib echo-swagger:
```
go get -u github.com/swaggo/echo-swagger
```

- Importar no arquivo da _api.go_ seguindo meu usuário neste repositório:
```
echoSwagger "github.com/swaggo/echo-swagger"

_ "github.com/themarcosramos/api-students/docs"
```

- Comando para fazer atualização no gerenciador de pacotes (_go.mod_):
```
go mod tidy
```

- Comando para rodar o servidor:
```            
go run main.go
```

- Para testar o Swagger, escreva no navegador a seguinte URL:  
```
http://localhost:8080/swagger/index.html
```

## Status do projeto
:construction: Aplicação em andamento.
