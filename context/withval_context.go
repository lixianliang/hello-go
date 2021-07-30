package main

var Username = FooKey("user-name")
var UserId = FooKey("user-id")

func foo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), UserId, "1")
		ctx2 := context.WithValue(ctx, Username, "lixianliang")
		next(w, r.WithContext(ctx2))
	}
}

func GetUserName(context context.Context) string {
	if ret, ok := context.Value(Username).(string); ok {
		return ret
	}

	return ""
}

func GetUserId(context context.Context) string {
	if ret, ok := context.Value(UserId).(string); ok {
		return ret
	}

	return ""
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome: "))
	w.Write([]byte(GetUserId(r.Context())))
	w.Write([]byte(" "))
	w.Write([]byte(GetUserName(r.Context())))
}

func main() {
	http.Handle("/", foo(test))
	http.ListenAndServe(":8080", nil)
}
