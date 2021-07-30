package main

func main() {
	url := "https://httpbin.org/delay/3"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("request failed:%s\n", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
	req = req.WithContext(ctx)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("DefaultClient failed: %s\n", err)
	}
	defer rsp.Body.Close()
}
