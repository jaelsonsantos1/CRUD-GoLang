package main

type Book struct {
	Id 			int `db:"id"`
	Name 		string `db:"name"`
	Author 		string `db:"author"`
	Genre		string `db:"genre"`
	Price 		float64 `db:"price"`
}


// Setters
func (b *Book) SetName(name string) {
	b.Name = name
}
func (b *Book) SetAuthor(author string) {
	b.Author = author
}
func (b *Book) SetGenre(genre string) {
	b.Genre = genre
}
func (b *Book) SetPrice(price float64) {
	b.Price = price
}

// Getters
func (b Book) GetName() string {
	return b.Name
}
func (b Book) getAuthor() string {
	return b.Author
}
func (b Book) getGenre() string {
	return b.Genre
}
func (b Book) getPrice() float64 {
	return b.Price
}

