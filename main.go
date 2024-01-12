package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	id        int
	title     string
	completed bool
}

type Toggleable interface {
	toggle()
}

func (t *Todo) toggle() {
	t.completed = !t.completed
}

type TodoApi struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var name string = "John"
	var name2 *string = &name // => pointer to name
	*name2 = "Jean"           // => change value of name in memory address
	// or name := "John"
	fmt.Printf("Hello %#v : %#v\n", name, *name2)

	// Array and slice
	a := [3]int{1, 2, 3}
	b := a
	b[1] = 4
	c := a[1:] // Display with first element
	fmt.Printf("Array: %#v : %#v\n", a, b)
	fmt.Printf("Array first element: %#v : %#v\n", a, c)

	d := []int{1, 2, 3}                    // slice
	fmt.Printf("Array len: %#v\n", len(d)) // length of slice
	fmt.Printf("Array cap: %#v\n", cap(d)) // capacity of the slice

	e := append(d, 4) // push new element in the slice
	fmt.Printf("Array push: %#v : %#v\n", d, e)

	// Objects and maps
	maps := map[string]string{
		"title":       "Faire le tutoriel",
		"description": "Parler le golang",
	}
	fmt.Printf("Map: %#v : %#v\n", maps, maps["title"])
	maps["title"] = "Paul"
	fmt.Printf("Map: %#v\n", maps)

	// Struct
	todo := Todo{1, "Faire la cuisine", true}
	// Method
	todo.toggle()
	// Interface
	toggleTodo(&todo)
	fmt.Printf("Struct: %#v\n", todo)

	arrayTodo := []Toggleable{
		&Todo{1, "Faire la cuisine", true},
		&Todo{1, "Faire la cuisine", true},
		&Todo{1, "Faire la cuisine", true},
	}
	fmt.Printf("Interface: %#v\n", arrayTodo)

	// Goroutine
	ch := make(chan string)   // chanel content the responce
	errCh := make(chan error) // chanel content the error
	go func() {
		title, err := fetchTodoTitle()
		if err != nil {
			errCh <- err
		} else {
			ch <- title
		}
	}()
	fmt.Println("Hello")
	defer close(ch)
	defer close(errCh)
	select {
	case err := <-errCh:
		panic(err)
	case title := <-ch:
		fmt.Printf("Titre de la todo : %s", title)
	}

	// Web server
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	title, err := fetchTodoTitle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(title))
}

func toggleTodo(t Toggleable) {
	t.toggle()
}

func fetchTodoTitle() (string, error) { // () => argument, (string, error) => types of retrun
	// Call Api ans json parser
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos?_limit=3")
	if err != nil { // try/catch but GO force the developpeur a make implicite error it's better for find the error
		return "", fmt.Errorf("impossible de contacter le server %w", err) // retrun string and error
	}

	defer r.Body.Close() // close the fonction after execution of principal function (main())
	var todos []TodoApi
	err = json.NewDecoder(r.Body).Decode(&todos)
	if err != nil { // try/catch but GO force the developpeur a make implicite error it's better for find the error
		return "", fmt.Errorf("impossible de parser la reponse du server %s", err)
	}
	if len(todos) > 0 {
		return todos[0].Title, nil
	}

	return "", fmt.Errorf("aucune todo trouvé par le server")
}
