package handlers 

import (
	"net/http"
)
func HomeHandler(w http.ResponseWriter, r *http.Request) { // Handler function for the "/" route
	w.Write([]byte(`Welcome to my web page!
My name is Asael Tobar.
I will be working on the Banking Account & Transaction System project for the semester.
The reason I have chosen this project is because I want to challenge myself for this semester, and I believe this project is the best way to do so.
I am excited to become a domain expert to build a database for this system and learn more about web development along the way.
`))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) { // Handler function for the "/about" route
	w.Write([]byte(`Hello! My name is Asael Tobar.
I am 20 years old and currently studying at the University of Belize pursuing a degree in Information Technology.
I aspire to become a software developer and web developer, as I have an interest for creation and problem-solving. I enjoy exploring complex ideas and figuring out how everyday things work through observation and analysis. 
In my free time, I enjoy playing video games with my friends, listening to music, singing, exploring various artforms, and reading books about pseudoscientific theories.
I like to be a balanced person, developing my intellectual, emotional, and spiritual sides while pursuing my career goals and personal interests.
`))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) { // Handler function for the "/contact" route
	w.Write([]byte(`Contact Me!
Via email at: asael2bar@gmail.com, inthespector@gmail.com
Via phone at: +501 606-8871
Via Github at: https://github.com/spector-asael
`))
}

func QuoteHandler(w http.ResponseWriter, r *http.Request) { // Handler function for the "/quote" route
	w.Write([]byte("“One does not become enlightened by imagining figures of light, but by making the darkness conscious.” \n– Carl Jung\n"))
}
