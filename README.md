# go-server
- O *build* desta imagem está publicado em [repositório do DockerHub](https://hub.docker.com/r/tcarvi/go-server).
- O código-fonte desta imagem está armazenado em [repositório do GitHub](https://github.com/tcarvi/go-server).

#### Servidor 
- Servidor ***GOLANG*** apenas executa comandos pré-definidos e genéricos deste projeto *e-commerce*.
  - Para alterar comportamento do servidor, deve-se alterar o código-fonte que se encontra em [repositório do GitHub](https://github.com/tcarvi/go-server). 
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
- `go install github.com/tcarvi/go-server`
- `go-server`

# Instalação com container *docker*:
- `docker run -d -p 8080:8080 tcarvi/go-server:latest`
  - -p [host port]:[container port]
  - A porta do container é a mesma publicada/escutada pelo código do servidor

# Visualização local da template servida por go-server:
  - Tanto para execução em container como fora do container
    - No *browser*, acesse localhost:8080