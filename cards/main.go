package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
RewriteEngine On

# Rewrite request to index.php
RewriteCond %{REQUEST_FILENAME} -f
RewriteRule !index\.php$ index.php [L]

# Serve a 404 otherwise...
RewriteCond %{ENV:REDIRECT_STATUS} ^$
RewriteRule ^ - [R=404]