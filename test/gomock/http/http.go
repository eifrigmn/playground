package http

import (
	"testing"
)

func TestMockHttpRequest(t *testing.T) {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/me", func(w http.ResponseWriter, req *http.Request) {
	//	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//	w.WriteHeader(200)
	//	fmt.Fprintf(w, "Request OK")
	//})
	//
	//transport := gomock.NewTransport()
	//transport.Stub("graph.facebook.com", mux)
	//
	//client := &http.Client{
	//	Transport: transport,
	//}
	//
	//resp, err := client.Get("https://graph.facebook.com/me")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(body))
}
