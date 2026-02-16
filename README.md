# 🛍️ API de Gerenciamento de Produtos

Uma API RESTful simples e eficiente para gerenciar produtos, desenvolvida em Go com foco em boas práticas e padrões de mercado.

## 📋 Sobre o Projeto

Este projeto é uma API completa de CRUD (Create, Read, Update, Delete) para gerenciamento de produtos. Foi desenvolvido seguindo padrões profissionais da indústria, com arquitetura em camadas, validações robustas e código limpo.

## 🚀 Tecnologias Utilizadas

- **Go 1.25+** - Linguagem de programação
- **Gin** - Framework web de alta performance
- **GORM** - ORM para manipulação do banco de dados
- **MySQL** - Banco de dados relacional
- **govalidator** - Validação de dados
- **Air** - Hot reload para desenvolvimento

## 📁 Estrutura do Projeto

```
go-pratica/
├── controllers/       # Lógica de negócio e handlers
├── models/           # Estruturas de dados do banco
├── dto/              # Data Transfer Objects
├── routes/           # Definição de rotas
├── database/         # Configuração do banco de dados
├── tmp/              # Arquivos temporários (Air)
├── .air.toml         # Configuração do hot reload
├── go.mod            # Dependências do projeto
└── main.go           # Ponto de entrada da aplicação
```

## 🔧 Como Rodar o Projeto

### Pré-requisitos

- Go 1.25 ou superior instalado
- MySQL rodando localmente
- Arquivo `.env` configurado (veja exemplo abaixo)

### Configuração do Banco de Dados

Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco
```

### Instalação

1. Clone o repositório
2. Instale as dependências:

```bash
go mod download
```

3. Execute a aplicação:

```bash
go run main.go
```

Ou use o Air para hot reload durante o desenvolvimento:

```bash
air
```

A API estará disponível em `http://localhost:3000`

## 📍 Endpoints Disponíveis

| Método | Endpoint       | Descrição                   |
| ------ | -------------- | --------------------------- |
| POST   | `/clients`     | Criar um novo produto       |
| GET    | `/clients`     | Listar todos os produtos    |
| GET    | `/clients/:id` | Buscar produto por ID       |
| PUT    | `/clients/:id` | Atualizar produto existente |
| DELETE | `/clients/:id` | Deletar produto             |

### Exemplos de Uso

**Criar produto:**

```json
POST /clients
{
  "name": "Notebook",
  "price": 2999.99
}
```

**Atualizar produto:**

```json
PUT /clients/1
{
  "name": "Notebook Gamer",
  "price": 3499.99
}
```

## ✨ Melhorias Implementadas

### Arquitetura e Organização

- ✅ Separação de responsabilidades com DTOs
- ✅ Validação de dados em múltiplas camadas
- ✅ Soft delete para produtos (dados não são perdidos)
- ✅ Timestamps automáticos (criação e atualização)

### Qualidade do Código

- ✅ Validações robustas com govalidator
- ✅ Tratamento adequado de erros
- ✅ Status HTTP corretos para cada cenário
- ✅ Código limpo e bem organizado

### Funcionalidades

- ✅ CRUD completo e funcional
- ✅ Update parcial (campos opcionais)
- ✅ Validação de IDs nas rotas
- ✅ Mensagens de erro claras e descritivas

## 🔄 Próximas Melhorias Planejadas

- [x] Middleware de validação nas rotas
- [ ] Sistema de logging estruturado
- [ ] Paginação na listagem de produtos
- [ ] Documentação com Swagger
- [ ] Testes unitários e de integração
- [ ] Suporte a variáveis de ambiente

## 📝 Licença

Este projeto foi desenvolvido para fins educacionais e de aprendizado.

---

Desenvolvido com ❤️ em Go
