# � Sistema de Gerenciamento de Mecânica

Uma API RESTful completa para gerenciar um pátio de mecânica, desenvolvida em Go com foco em boas práticas e padrões de mercado.

## 📋 Sobre o Projeto

Este projeto é um sistema completo para gerenciar o pátio de uma mecânica, controlando cadastro de clientes e seus veículos, criação de Ordens de Serviço, geração de orçamentos, aprovação/recusa de orçamentos, execução e acompanhamento de serviços, registro de pagamentos e histórico completo de serviços por cliente/veículo.

Foi desenvolvido seguindo padrões profissionais da indústria, com arquitetura em camadas, validações robustas e código limpo.

### ✅ Status Atual

**Implementado:**

- ✅ CRUD completo de clientes (Client)
- ✅ Validação de CPF, CEP e telefone
- ✅ Soft delete
- ✅ Estrutura modular com controllers, models, DTOs e routes

**Em Planejamento/Desenvolvimento:**

- 🔄 Cadastro de veículos
- 🔄 Ordens de Serviço
- 🔄 Orçamentos
- 🔄 Pagamentos
- 🔄 Histórico de serviços

## 🚀 Tecnologias Utilizadas

- **Go 1.25.6** - Linguagem de programação
- **Gin** - Framework web de alta performance
- **GORM** - ORM para manipulação do banco de dados
- **MySQL** - Banco de dados relacional
- **govalidator** - Validação de dados
- **godotenv** - Gerenciamento de variáveis de ambiente

## 📁 Estrutura do Projeto

```
go-pratica/
├── controllers/       # Lógica de negócio e handlers
│   └── client_controller.go
├── models/           # Estruturas de dados do banco
│   └── client.go
├── dto/              # Data Transfer Objects
│   └── client_dto.go
├── routes/           # Definição de rotas
│   ├── routes.go
│   └── client_routes.go
├── database/         # Configuração do banco de dados
│   └── database.go
├── utils/            # Funções utilitárias
│   └── validators.go
├── tmp/              # Arquivos temporários
├── create_database.sql   # Script de criação do banco
├── planning.md       # Planejamento detalhado do projeto
├── go.mod            # Dependências do projeto
└── main.go           # Ponto de entrada da aplicação
```

## 🔧 Como Rodar o Projeto

### Pré-requisitos

- Go 1.25 ou superior instalado
- MySQL rodando localmente
- Arquivo `.env` configurado (veja exemplo abaixo)

### Configuração do Banco de Dados

1. Execute o script SQL para criar o banco de dados:

```bash
mysql -u seu_usuario -p < create_database.sql
```

2. Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=garage_management
```

### Instalação e Execução

1. Clone o repositório:

```bash
git clone <url-do-repositorio>
cd go-pratica
```

2. Instale as dependências:

```bash
go mod download
```

3. Configure o arquivo `.env` com as credenciais do banco de dados

4. Execute a aplicação:

```bash
go run main.go
```

A API estará disponível em `http://localhost:3000`

## 📍 Endpoints da API

### Clientes (Clients)

| Método | Endpoint       | Descrição                     |
| ------ | -------------- | ----------------------------- |
| POST   | `/clients`     | Criar um novo cliente         |
| GET    | `/clients`     | Listar todos os clientes      |
| GET    | `/clients/:id` | Buscar cliente por ID         |
| PUT    | `/clients/:id` | Atualizar cliente existente   |
| DELETE | `/clients/:id` | Deletar cliente (soft delete) |

### Exemplos de Requisições

**Criar cliente:**

```json
POST /clients
{
  "name": "João Silva",
  "cpf": "123.456.789-00",
  "primary_phone": "(11) 98765-4321",
  "secondary_phone": "(11) 3456-7890",
  "email": "joao@email.com",
  "zip_code": "01310-100",
  "street": "Avenida Paulista",
  "number": "1578",
  "complement": "Apto 101",
  "neighborhood": "Bela Vista",
  "city": "São Paulo",
  "state": "SP"
}
```

**Atualizar cliente:**

```json
PUT /clients/1
{
  "name": "João Silva Santos",
  "email": "joao.santos@email.com",
  "primary_phone": "(11) 99999-8888"
}
```

**Listar clientes:**

```
GET /clients
```

**Buscar cliente específico:**

```
GET /clients/1
```

**Deletar cliente:**

```
DELETE /clients/1
```

## ✨ Features Implementadas

### Arquitetura e Organização

- ✅ Arquitetura em camadas (Controllers, Models, DTOs, Routes)
- ✅ Separação de responsabilidades com DTOs
- ✅ Validação de dados em múltiplas camadas
- ✅ Soft delete (dados não são perdidos permanentemente)
- ✅ Timestamps automáticos (created_at e updated_at)

### Validações

- ✅ Validação de CPF (formato e dígitos verificadores)
- ✅ Validação de CEP
- ✅ Validação de telefone brasileiro
- ✅ Validação de campos obrigatórios
- ✅ Sanitização de dados (remoção de caracteres especiais)

### Qualidade do Código

- ✅ Validações robustas com govalidator
- ✅ Tratamento adequado de erros
- ✅ Status HTTP corretos para cada cenário
- ✅ Código limpo e bem organizado
- ✅ Padrão de nomenclatura consistente (inglês)
- ✅ Comentários em português para facilitar manutenção

### Funcionalidades

- ✅ CRUD completo de clientes
- ✅ Update parcial (campos opcionais)
- ✅ Validação de IDs nas rotas
- ✅ Mensagens de erro claras e descritivas
- ✅ Suporte a variáveis de ambiente com godotenv
- ✅ Conexão otimizada com MySQL

## 📐 Padrões do Projeto

### Nomenclatura

- **Código em inglês**: Todas as variáveis, funções, structs e rotas
- **Comentários em português**: Para facilitar a manutenção
- **Tabelas e campos**: Snake_case no banco, CamelCase no Go

### Campos Padrão

| Português           | Inglês          |
| ------------------- | --------------- |
| Nome                | Name            |
| CPF                 | CPF             |
| Telefone Principal  | Primary Phone   |
| Telefone Secundário | Secondary Phone |
| CEP                 | Zip Code        |
| Rua                 | Street          |
| Bairro              | Neighborhood    |

Veja o [planning.md](planning.md) para mais detalhes sobre padrões e planejamento completo.

## 🔄 Próximas Implementações

- [ ] CRUD de Veículos (Vehicles)
- [ ] CRUD de Ordens de Serviço (Service Orders)
- [ ] Sistema de Orçamentos (Budgets)
- [ ] Módulo de Pagamentos (Payments)
- [ ] Histórico de serviços por cliente
- [ ] Sistema de logging estruturado
- [ ] Paginação na listagem
- [ ] Documentação com Swagger
- [ ] Testes unitários e de integração
- [ ] Autenticação e autorização

## 👨‍💻 Autor

Desenvolvido por **fgeanlopes** (f.geanlopes@gmail.com)

## 📝 Licença

Este projeto foi desenvolvido para fins de aprendizado e portfólio.

---

Desenvolvido com ❤️ em Go
