# go-server
- O *build* desta imagem está publicado em [repositório do DockerHub](https://hub.docker.com/r/eduardoleal1981/go-server).
- O código-fonte desta imagem está armazenado em [repositório do GitHub](https://github.com/eduardoleal1981/go-server).

- As estruturas indicadas acima devem ser usadas na customização implementada em SEU REPOSITÓRIO GITLAB (fluxo de ***build*** final).
- Depois do *BUILD* do aplicativo *WEB*, deve-se configurar a publicação final com solução ***DEVOPS*** de ***Cloud***.

#### Servidor 
- Servidor ***GOLANG*** apenas executa comandos pré-definidos e genéricos deste projeto *e-commerce*.
  - Para alterar comportamento do servidor, deve-se alterar o código-fonte que se encontra em [repositório do GitHub](https://github.com/eduardoleal1981/go-server). 
- Há modelo de aplicativo WEB que deve ser customizado por você, em **SEU REPOSITÓRIO**.
  - Código-fonte do aplicativo WEB segue especificações publicadas por [W3C](https://www.w3.org/2019/06/htmlwg-charter.html).
  - Para a customização, deve-se acrescentar imagens gráficas, arquivos textuais e vídeos em **SEU REPOSITÓRIO**.
  - Como exemplo de customização e uso, veja:
    - [gitlab.com/eduardoleal1981/no-luar](https://gitlab.com/eduardoleal1981/no-luar)
    - [gitlab.com/eduardoleal1981/no-frete](https://gitlab.com/eduardoleal1981/no-frete)
    - [gitlab.com/eduardoleal1981/na-porteira](https://gitlab.com/eduardoleal1981/na-porteira)
    - [gitlab.com/eduardoleal1981/tcarvi](https://gitlab.com/eduardoleal1981/tcarvi)

# Instalação/execução sem container *docker*:
- `sudo apt install golang`
- `emacs ~/.bashrc`
  - add line: `export GOBIN=/yourGoBinDirectory`
  - add line: `export PATH=$PATH:$GOBIN`
  - add line: `export GOPATH=/yourGoPathDirectory`
  - salve o arquivo e o feche. (`ctrl x ctrl z`)
- `go install github.com/eduardoleal1981/go-server`
- `go-server`

# Instalação com container *docker*:
- `docker run -d -p 8080:8080 eduardoleal1981/go-server:latest`
  - -p [host port]:[container port]
  - A porta do container é a mesma publicada/escutada pelo código do servidor

# Visualização local da template servida por go-server:
  - Tanto para execução em container como fora do container
    - No *browser*, acesse localhost:8080

#### Estrutura HTML:
- Seção com toda a tela (fundo)
  - id="map"
  - Google Maps no centro
  - 3 menus incorporados (apresentação na superfície superior) com tratamento de cor e de sombreado
- Menu superior:
  - Logomarca
  - "Botão de acesso ao e-commerce"
  - "Acompanhamento do servico no Google Maps"
- Menu Lateral esquerdo:
  - Lista de tópicos e sub-tópicos
- Menu lateral direito:
  - Apresentação de Vídeos
- Menu inferior:
  - Contato: Inclusive Nomes e WhatsApp

#### Referências relevantes para aplicativos WEB

```  
The W3C DOM standard is separated into 3 different parts:
Core DOM - standard model for all document types
XML DOM - standard model for XML documents
HTML DOM - standard model for HTML documents 
```

```   
The HTML DOM is a standard object model and programming interface for HTML.
It defines:
The HTML elements as objects
The properties of all HTML elements
The methods to access all HTML elements
The events for all HTML elements
In other words: The HTML DOM is a standard for how to get, change, add, or delete HTML elements.
DOM APIs are designed such that no script can ever detect the simultaneous execution of other scripts.
The security model of the web is based on the concept of "origins". Many of potential attacks on the web involve cross-origin actions.
HTML namespace: http://www.w3.org/1999/xhtml
```  

```  
DOM Interface:
The Document Object Model (DOM) is a representation - a model - of a document and its content. 
Dom trees have nodes: DocumentType node, Element nodes, Text nodes, Comment nodes, and ProcessingInstruction nodes. 
html tag    ->  interface HTMLHtmlElement : HTMLElement {};
Javascript  ->  document.getElementsByTagName(name) 
```  

```  
  Especificações também relevantes:
  - URL
  - HTTP
  - COOKIES
  - FETCH
  - WEBIDL
  - JAVASCRIPT / ECMA-262
  - UIEVENTS
  - DOM
  - TOUCH
  - FILEAPI
  - XHR
  - MQ
  - CSS
  - GEOMETRY
  - WEBGL
  - WSP (WebSocket protocol)
  - MATHML
  - SVG
```  

```  
IDL 
    [Exposed=Window]
    interface Example {
      // this is an IDl definition
    };
```  

```  
document.getElementById(id)
document.getElementsByTagName(name)
document.getElementsByClassName(name)
```  

```  
element.innerHTML =  new html content
element.setAttribute(attribute, value)
element.attribute = new value
```  

```  
document.createElement(element)	  Create an HTML element
document.removeChild(element)	  Remove an HTML element
document.appendChild(element)	  Add an HTML element
document.replaceChild(new, old)	  Replace an HTML element 
```  

```   
document.getElementById(id).onclick = function(){code}	Adding event handler code to an onclick event 
```  

```   
Animation:
function myMove() {
    var elem = document.getElementById("animate"); 
    var pos = 0;
    var id = setInterval(frame, 5);
    function frame() {
      if (pos == 350) {
        clearInterval(id);
      } else {
        pos++; 
        elem.style.top = pos + 'px'; 
        elem.style.left = pos + 'px'; 
      }
    }
  }
```  

```  
Tag events:
<h1 onclick="this.innerHTML = 'Ooops!'">Click on this text!</h1>
```  

```   
Event listeners
element.addEventListener(event, function, useCapture);
The first parameter is the type of the event (like "click" or "mousedown").
The second parameter is the function we want to call when the event occurs.
The third parameter is a boolean value specifying whether to use event bubbling or event capturing. This parameter is optional.
Ex.:
document.getElementById("myBtn").addEventListener("click", displayDate);
element.addEventListener("click", function(){ alert("Hello World!"); });
element.addEventListener("click", myFunction);
function myFunction() {
  alert ("Hello World!");
}
```  

```  
Navigation:
parentNode
childNodes[nodenumber]
firstChild
lastChild
nextSibling
previousSibling
Ex.: 
document.getElementById("id02").innerHTML = document.getElementById("id01").firstChild; 
```  

```  
Two big elements:
document.body - The body of the document
document.documentElement - The full document 
```  

```  
Get the name of a tag
document.getElementById("id02").innerHTML = document.getElementById("id01").nodeName;
```  

```  
NodeValue
nodeValue for element nodes is null
nodeValue for text nodes is the text itself
nodeValue for attribute nodes is the attribute value
```  

```  
TextNode
var node = document.createTextNode("This is a new paragraph.");
```  

```  
Remove child tag
<div id="div1">
    <p id="p1">This is a paragraph.</p>
    <p id="p2">This is another paragraph.</p>
  </div>
  
  <script>
  var parent = document.getElementById("div1");
  var child = document.getElementById("p1");
  parent.removeChild(child);
  </script>
```  

```  
Replace a child tag
<div id="div1">
    <p id="p1">This is a paragraph.</p>
    <p id="p2">This is another paragraph.</p>
  </div>
  
  <script>
  var para = document.createElement("p");
  var node = document.createTextNode("This is new.");
  para.appendChild(node);
  
  var parent = document.getElementById("div1");
  var child = document.getElementById("p1");
  parent.replaceChild(para, child);
  </script>
```  
  
```  
APIs
HTML Geolocation
HTML Drag and Drop
HTML Web Workers
HTML SSE
```  

```  
Creating new tags
<head>
...
<script>document.createElement("myTag")</script>
</head>
<body>
<myTag>My personal tag</myTag>
</body>
```  

```  
WASM
To source toolchains: source ./emsdk_env.sh --build=Release
Compiling:
We have to pass the linker flag -s WASM=1 to emcc (otherwise by default emcc will emit asm.js).
If we want Emscripten to generate an HTML page that runs our program, in addition to the Wasm binary and JavaScript wrapper, we have to specify an output filename with a .html extension.
Finally, to actually run the program, we cannot simply open the HTML file in a web browser because cross-origin requests are not supported for the file protocol scheme. We have to actually serve the output files over HTTP. 
```  