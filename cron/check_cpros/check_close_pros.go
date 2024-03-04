/*
Скрипт для запуска по крону, проверка на испавленные данные,
подробнее https://tc.jusanmobile.kz:9000/swagger/index.html#/Problem_Router_Onyma/check_close_problem_router_onyma_speed
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type token_val struct {
	Token string `json:"token"`
}

//Бот для выполнения скрипта, данные для подключения
var jsonStr = []byte(`{"email":"apicron@noreply.kz", "password": "*DGfv2g87S5ZS'6h"}`)

func main() {

	client := &http.Client{}

	/*caCert, err := ioutil.ReadFile("/etc/ssl/certs/ca-certificates.crt")
	if err != nil {
		fmt.Errorf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}*/

	url := "https://tc.jusanmobile.kz:9000/auth/sign-in/"
	fmt.Println("URL:>", url)

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Errorf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Failed reading response body: %s", err)
	}
	//fmt.Printf("Got response %d: %s %s", resp.StatusCode, resp.Proto, string(body))
	t := new(token_val)
	err = json.Unmarshal(body, &t)

	if err != nil {
		fmt.Errorf("Ошибка при получении токена: %s", err)
	}

	fmt.Println("token получен")
	fmt.Println("Запускаю проверку на наличие новых исправленых строк, ставлю статус 'Закрыто'")
	url = "https://tc.jusanmobile.kz:9000/api/router_onyma/problem/check_close_problem_router_onyma_speed"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+t.Token)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	//log.Println(resp)
	defer resp.Body.Close()
	fmt.Println("Скрипт успешно отработал")

}
