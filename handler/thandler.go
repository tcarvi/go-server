// Package

package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type THandler int

func (tHandler THandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Identificação de padrões nas requisições
	// Padrão errado não gerará dado: apenas retornará para fluxo de main.go com erro a ser apresentado.

	// Segurança:
	// tratar de r.Host e de r.URL.Path para que estes dados apenas possam ser "strings"
	// urlPath = verifySecurityIssues(r.URL.Path)
	// urlHost = verifySecurityIssues(r.Host)
	// Implementação a ser definida por perito em tratamento de dados.
	// Solução escolhida não será mostrada neste repositório público.
	// No caso de identificação de falha neste quesito de segurança, retornar dados do solicitante, incluindo:
	//    IP da requisição,
	//    Porta de Comunicação do computador da requisição e
	//    Horário da Requisição;
	//    com mensagem de alerta sobre uso criminoso da rede mundial de computadores.
	// Fluxo seguinte não precisa de tratamento de segurança pois os dados a serem processados já foram validados.

	logs(r)

	if strings.HasPrefix(r.URL.Path, "/api") {
		tServeAPI(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".js") {
		tServeJs(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/layout") {
		tServeLayout(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".css") {
		tServeCss(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".svg") {
		tServeSvg(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".html") {
		tServeHtml(w, r)
	} else if strings.HasSuffix(r.URL.Path, "/") {
		tServeHtml(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".ico") {
		tServeIco(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".jpeg") {
		tServeJpeg(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".png") {
		tServePng(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".pdf") {
		tServePdf(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".webp") {
		tServeWebp(w, r)
	} else if strings.HasSuffix(r.URL.Path, ".pptx") {
		tServePptx(w, r)
	} else if strings.HasSuffix(r.URL.Path, "/v") {
		fmt.Fprint(w, "Server 1.0")
	}
}

func tServeAPI(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public"+r.URL.Path)
}

func tServeJs(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/js"+r.URL.Path)
}

func tServeLayout(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public"+r.URL.Path)
}

func tServeCss(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/css"+r.URL.Path)
}

func tServeSvg(w http.ResponseWriter, r *http.Request) {
	// Connection
	w.Header().Set("Transfer-Encoding", "gzip")
	w.Header().Set("Cache-Control", "no-cache")
	// Message
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Content-Language", "pt-br")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/img/svg"+r.URL.Path)
}

func tServeHtml(w http.ResponseWriter, r *http.Request) {
	// Protocol
	w.Header().Set("HTTP-Version", "HTTP/2.0") //Chrome OK
	w.Header().Set("X-Firefox-Spdy", "h2")     //Chrome OK
	w.Header().Set("Protocol", "HTTP/2.0")     //Chrome OK
	// Connection
	w.Header().Set("Cache-Control", "no-cache") //Chrome OK
	w.Header().Set("Date", time.Now().Format(time.UnixDate))
	// Message
	w.Header().Set("Content-Type", "text/html; charset=utf-8") //Chrome OK
	w.Header().Set("Content-Language", "pt-br")                //Chrome OK
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve File
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/index.html")
}

func tServeIco(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/x-icon")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/img/ico"+r.URL.Path)
}

func tServeJpeg(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/jpeg")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/img/jpeg"+r.URL.Path)
}

func tServePng(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/png")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/img/png"+r.URL.Path)
}

func tServeWebp(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "image/webp")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/img/webp"+r.URL.Path)
}

func tServePdf(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "/pdf")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/docs/pdf"+r.URL.Path)
}

func tServePptx(w http.ResponseWriter, r *http.Request) {
	// Message
	w.Header().Set("Content-Type", "application/vnd.mspowerpoint")
	// To be handled by webapp
	w.Header().Set("key-Code", "00000000001")
	//Serve Files
	http.ServeFile(w, r, "public/"+tTrimUrlHost(r.Host)+"/docs/ppt"+r.URL.Path)
}

func tTrimUrlHost(host string) string {
	var name = strings.TrimSuffix(host, ":8080")
	name = strings.TrimSuffix(name, ".com.br")
	name = strings.TrimSuffix(name, ".com")
	return name
}

func logs(r *http.Request) {
	fmt.Println("*** *** *** ***")
	fmt.Println("Requisição com RemoteAddr = " + r.RemoteAddr)
	fmt.Println("Requisição com RequestURI = " + r.RequestURI)
	fmt.Println("r.Host = " + r.Host)
	fmt.Println("tTrimUrlHost(r.Host) = " + tTrimUrlHost(r.Host))
	fmt.Println("r.URL.Path = " + r.URL.Path)
	fmt.Println("Horário da request:" + time.Now().Format(time.UnixDate))
	fmt.Println("*** *** *** ***")
}

// TODO: func database() string
// 	db, err := sql.Open("postgres", "user=postgres dbname=np sslmode=disable")

// TODO: func answer(q Question) Answer
// a, err := ai.Answers(q)
