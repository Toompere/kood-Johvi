# ASCII-Art-Web
Ascii-Art-Web creates a http server on localhost which hosts GUI where you can create ASCII art. 

It uses GET / to send html response(index) and POST /ascii-art sends data(text, banner) to Go server to process ASCII art.
## Usage
```golang
$ go run .
Open http://localhost:8000
```
You can click the link in the terminal or open your browser at localhost:8000

In the textfield enter the text you want to use to generate ASCII art, select banner(standard, shadow or thinkertoy) and press submit to generate ASCII art.

Press Ctrl+c in terminal to stop the server.
## Testing
You can use [ascii-art-web/audit](https://github.com/01-edu/public/tree/master/subjects/ascii-art-web/audit) to test if the program is working correctly.
## Author
[Margus Toompere](https://01.kood.tech/git/MargusT)