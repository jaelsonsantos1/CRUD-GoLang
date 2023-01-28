package main

import (
	"db"
	"log"
	"fmt"
)

var cfg db.ConfigDb = db.ConfigDb{
	Host: "localhost",
	Port: 5432,
	User: "postgres",
	Password: "sua-senha",
	NameDataBase: "nome-banco",
}

func main() {
	// Exibir menu de opções
	showMenu()
}

// 
// Função que mostra o menu de opções
// 
func showMenu() {
	var choiceOption int

	for {
		fmt.Print("\n[1] Ver dados\n[2] Cadastrar\n[3] Atualizar\n[4] Deletar\n[5] Saindo\n--> ")
		fmt.Scan(&choiceOption)
	
		switch choiceOption {
		case 1:
			fmt.Println("Opção de ver livros cadastrados!")
			
			// Ver livros cadastrados
			showBooks()

		case 2:
			fmt.Println("Opção de cadastrar um livro!")
			
			// Cadastrar livro
			insertItem()

		case 3:
			var id int
			fmt.Println("Opção de atualizar um livro!")
			
			// Atualizar livro
			fmt.Print("Informe um id de um livro que deseja atualizar: ")
			fmt.Scan(&id)
			updateItem(id)
			
		case 4:
			var id int
			fmt.Println("Opção de deletar um livro!")

			// Deletar livro
			fmt.Print("Informe um id de um livro que deseja deletar: ")
			fmt.Scan(&id)
			deleteItem(id)

		case 5:
			// Opção para sair do programa
			fmt.Println("Saindo")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}

}

// 
// Função que mostra de forma formatada todos os dados da lista de livros
// 
func showBooks() {
	books := getAllItems()

	for _, book := range books {
		fmt.Println("\n==================")
		fmt.Println("ID: ", book.Id)
		fmt.Println("Nome: ", book.Name)
		fmt.Println("Autor: ", book.Author)
		fmt.Println("Gênero: ", book.Genre)
		fmt.Println("Preço: ", book.Price)
	}
}

// ==========================================
// Funções de gerenciamento do banco de dados
// ==========================================
// 
// Função para buscar todos os cadastros do banco
// 
func getAllItems() []Book {
	db, err := db.ConnectDb(cfg)
	
	if (err != nil) {
		log.Fatal(err)
	}

	defer db.Close()

	// Sql para buscar dados do banco
	rows, err := db.Query("select * from books")
	if (err != nil) {
		log.Fatal(err)
	}

	var books []Book
	
	for rows.Next() {
		b := Book{}

		err := rows.Scan(
			&b.Id, &b.Name, &b.Author,
			&b.Genre,& b.Price)

		if (err != nil) {
			log.Fatal(err)
		}

		books = append(books, b)
	}

	return books
}

// 
// Função para incluir um novo registro no banco
// 
func insertItem() {
	book := Book{}

	// Solicitar o nome do livro
	fmt.Print("Qual nome do livro: ")
	fmt.Scan(&book.Name)

	if (book.Name == "") {
		fmt.Println("O nome informado é inválido!")
		return
	}
	
	// Solicitar o nome do autor do livro
	fmt.Print("Qual o autor do livro: ")
	fmt.Scan(&book.Author)
	
	if (book.Author == "") {
		fmt.Println("O nome do autor informado é inválido!")
		return
	}
	
	// Solicitar o gênero do livro
	fmt.Print("Qual o gênero literário do livro: ")
	fmt.Scan(&book.Genre)
	
	if (book.Genre == "") {
		fmt.Println("O gênero literário informado é inválido!")
		return
	}
	
	// Solicitar o valor do livro
	fmt.Print("Qual o preço do livro: ")
	fmt.Scan(&book.Price)
	
	if (book.Price <= 0) {
		fmt.Println("O valor informado é inválido!")
		return
	}
	
	// Criando conexão com banco de dados
	db, err := db.ConnectDb(cfg)

	if (err != nil) {
		log.Fatal(err)
	}

	defer db.Close()

	// Código slq responsável por inserir dados ao banco
	sqlInsert := `insert into books(nome, autor, genero, preco) values ($1, $2, $3, $4)`

	_, err = db.Exec(sqlInsert, book.Name, book.Author, book.Genre, book.Price)

	if (err != nil) {
		log.Fatal(err)
	} else {
		fmt.Println("Livro inserido ao banco com sucesso!")
	}
}

// 
// Função para incluir um novo registro no banco
// 
func updateItem(id int) {
	listBooks := getAllItems()
	book := Book{}
	var idBook int
	var bookItem Book

	for _, book := range listBooks {
		if (id == book.Id) {
			idBook = book.Id
			bookItem = book
			break
		}
	}

	if idBook == 0 {
		fmt.Println("Não foi encontrado nenhum registro com este ID!")
	} else {
		fmt.Println("\n==================")
		fmt.Println("ID: ", bookItem.Id)
		fmt.Println("Nome: ", bookItem.Name)
		fmt.Println("Autor: ", bookItem.Author)
		fmt.Println("Gênero: ", bookItem.Genre)
		fmt.Println("Preço: ", bookItem.Price)
		fmt.Println("====================")

		// Solicitando as novas informações
		fmt.Println("Insira as novas informações")
		// Solicitar o nome do livro
		fmt.Print("Qual nome do livro: ")
		fmt.Scan(&book.Name)

		if (book.Name == "") {
			fmt.Println("O nome informado é inválido!")
			return
		}
		
		// Solicitar o nome do autor do livro
		fmt.Print("Qual o autor do livro: ")
		fmt.Scan(&book.Author)
		
		if (book.Author == "") {
			fmt.Println("O nome do autor informado é inválido!")
			return
		}
		
		// Solicitar o gênero do livro
		fmt.Print("Qual o gênero literário do livro: ")
		fmt.Scan(&book.Genre)
		
		if (book.Genre == "") {
			fmt.Println("O gênero literário informado é inválido!")
			return
		}
		
		// Solicitar o valor do livro
		fmt.Print("Qual o preço do livro: ")
		fmt.Scan(&book.Price)
		
		if (book.Price <= 0) {
			fmt.Println("O valor informado é inválido!")
			return
		}
		
		db, err := db.ConnectDb(cfg)

		if (err != nil) {
			log.Fatal(err)
		}

		defer db.Close()

		sqlUpdate := `UPDATE books
					  SET (nome, autor, genero, preco) = ($1, $2, $3, $4)
					  WHERE ID=$5;`
		_, err = db.Exec(sqlUpdate, book.Name, book.Author, book.Genre, book.Price, idBook)

		if (err != nil) {
			log.Fatal(err)
		}

		fmt.Println("Dados atualizados com sucesso!")
	}
}

// 
// Função para incluir um novo registro no banco
// 
func deleteItem(id int)  {
	listBooks := getAllItems()
	var idBook int

	for _, book := range listBooks {
		if (id == book.Id) {
			idBook = book.Id
			break
		}
	}

	if idBook == 0 {
		fmt.Println("Não foi encontrado nenhum registro com este ID!")
	} else {
		db, err := db.ConnectDb(cfg)

		if (err != nil) {
			log.Fatal(err)
		}

		defer db.Close()

		sqlUpdate := `delete from books where id=$1;`
		_, err = db.Exec(sqlUpdate, idBook)

		if (err != nil) {
			log.Fatal(err)
		}

		fmt.Println("Informação deletada com sucesso!")
	}
}
