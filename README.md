# gomux

Package includes:
1. A simple wrapper for `http.HandleFunc`

## Usage
```go
type RequestUser struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type ResponseData struct {
    FullName string `json:"full_name"`
}

func main() {
    postUserHandler := gomux.SimpleHandler(func(u *RequestUser, w http.ResponseWriter, r *http.Request) any {
        rd := ResponseData{
            FullName: fmt.Sprintf("%s, %d", u.Name, u.Age),
        }

        // you can return nil if the body is empty
        return rd
    })

    router := chi.NewRouter()
    router.Post("/api/users", postUserHandler)

    http.ListenAndServe(":8181", router)
}
```