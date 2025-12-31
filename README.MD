# Sistema de Empacotamento (Go API)

## Descrição do Projeto

Este projeto implementa uma API em Go para automatizar o processo de empacotamento de pedidos.  
Dada uma lista de pedidos contendo produtos e suas dimensões (altura, largura, comprimento), a API retorna:

- Quais caixas de papelão devem ser usadas para cada pedido
- Quais produtos vão em cada caixa

A API também possui documentação Swagger e pode ser executada via Docker.
## 🛠️ 4. Arquitetura e Tecnologias

### Arquitetura

O projeto segue uma arquitetura simples e modular:

- **Controllers**: Recebem requisições HTTP, validam dados e chamam os serviços correspondentes.
- **Services**: Contêm a lógica principal da aplicação, como empacotamento de produtos em caixas.
- **Models**: Definem as estruturas de dados usadas pela aplicação (`Pedido`, `Produto`, `Caixa`).
- **Docs**: Contém a documentação Swagger gerada automaticamente (`swag init`).
- **Servidor**: Configura rotas, middlewares e inicializa o servidor HTTP com Gin.

### Tecnologias e Ferramentas

| Tecnologia / Ferramenta | Uso no projeto |
|-------------------------|----------------|
| **Go 1.21+**            | Linguagem principal para API e lógica de empacotamento |
| **Gin**                 | Framework HTTP para criação de endpoints RESTful |
| **Swaggo / Swagger UI** | Geração e visualização da documentação da API |
| **Docker**              | Containerização da aplicação para execução consistente |
| **Git**                 | Controle de versão e versionamento incremental do código |
| **PowerShell / Terminal** | Ambiente de execução para build e testes locais |


## 🚀 Como executar

### ✅ Pré-requisitos

Antes de começar, é necessário ter os seguintes softwares instalados na sua máquina:

- **Go 1.21+** (para rodar localmente, se não usar Docker)  
- **Docker Desktop** (para executar em container)  
- **Git** (para clonar o repositório)  
- **PowerShell ou terminal** (Windows, Linux ou macOS)

---

### ⚙️ Clonando o projeto

Clone o repositório do GitHub:

```bash
git clone https://github.com/LucasEpifanio1/sistema-empacotamento.git
`````

O projeto será salvo no diretório onde você rodou o comando git clone. Por exemplo:

```bash
C:\Users\Lucas\sistema-empacotamento
`````
Entre na pasta do projeto:

```bash
cd sistema-empacotamento
`````
### 💻 Executando localmente com Go
Abra o PowerShell ou terminal.

Inicialize o módulo do projeto (caso ainda não tenha feito):

```bash
go mod init sistema-empacotamento
`````
Instale as dependências do projeto:

```bash
go get github.com/gin-gonic/gin
go install github.com/swaggo/swag/cmd/swag@latest

`````
Gere a documentação Swagger:
```bash
swag init -g app/servidor/main.go
`````
Execute a aplicação:

```bash
go run app/servidor/main.go
`````
Acesse os endpoints:

- Health Check: http://localhost:8081/health

- Swagger UI: http://localhost:8081/swagger/index.html


###  🐳 Executando via Docker

Certifique-se de que o Docker Desktop está instalado e rodando.

Build da imagem Docker:

```bash
docker build -t sistema-empacotamento:latest .
`````
Execute o container:
```bash
docker run -p 8081:8081 sistema-empacotamento:latest
`````
Teste os endpoints:

- Health Check: http://localhost:8081/health
- Swagger UI: http://localhost:8081/swagger/index.html