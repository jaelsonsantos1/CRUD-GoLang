# Sistema de Livros

## Como funciona o projeto?
O projeto foi criado com intuito de testar os conhecimentos com Go lang e PostgresSQL. Este projeto conta com função para Cadastrar Livro, Ver Livros cadastrados, Atualizar livros e Deletar Livro.

## Como usar?
### Clonando reposiório
`git clone https://github.com/jaelsonsantos1/CRUD-GoLang.git`

### Abra a pasta do projeto
`cd .\CRUD-GoLang`

### Atualizar os pacotes do Go
`go mod tidy`

### Configurar conexão com banco de dados
No arquivo main.go na raiz do projeto, vá para variável *cfg* e configure as informações para seu banco de dados, como mostra o código a seguir.
```
var cfg db.ConfigDb = db.ConfigDb{
	Host: "localhost",
	Port: 5432,
	User: "postgres",
	Password: "sua-senha",
	NameDataBase: "nome-banco",
}
```

### Rodar o projeto
`go run .`

<hr>
