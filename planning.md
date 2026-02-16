# 🚗 SISTEMA DE GERENCIAMENTO DE MECÂNICA

**Projeto:** Sistema de Gerenciamento de Pátio de Mecânica  
**Versão:** 1.0  
**Backend:** Go (Golang) + Gin + GORM  
**Frontend:** Next.js (futuro)  
**Banco de Dados:** MySQL  
**Data:** Fevereiro 2026

---

## 📋 ÍNDICE

1. [Objetivo do Projeto](#objetivo-do-projeto)
2. [Arquitetura](#arquitetura)
3. [Tecnologias](#tecnologias)
4. [Estrutura de Pastas](#estrutura-de-pastas)
5. [Modelagem do Banco de Dados](#modelagem-do-banco-de-dados)
6. [Scripts SQL - MySQL](#scripts-sql---mysql)
7. [Endpoints da API](#endpoints-da-api)
8. [Fluxo do Sistema](#fluxo-do-sistema)
9. [Plano de Execução](#plano-de-execução)
10. [Próximos Passos](#próximos-passos)

---

## 🎯 OBJETIVO DO PROJETO

Sistema completo para gerenciar o pátio de uma mecânica, controlando:

- ✅ Cadastro de clientes e seus veículos
- ✅ Criação de Ordens de Serviço (OS)
- ✅ Geração de orçamentos detalhados
- ✅ Aprovação/recusa de orçamentos
- ✅ Execução e acompanhamento de serviços
- ✅ Registro de pagamentos
- ✅ Histórico completo de serviços por cliente/veículo

---

## 🏗️ ARQUITETURA

**Padrão:** MVC + DTO Pattern

```
┌─────────────────────────────────────────────┐
│          CLIENTE (Frontend/Postman)         │
└──────────────────┬──────────────────────────┘
                   │ HTTP Request
                   ↓
┌─────────────────────────────────────────────┐
│         ROUTES (Router Groups)              │
│  /api/v1/clientes, /api/v1/veiculos...     │
└──────────────────┬──────────────────────────┘
                   │
                   ↓
┌─────────────────────────────────────────────┐
│         CONTROLLERS (Lógica de Controle)    │
│  - Valida dados (DTO)                       │
│  - Processa requisições                     │
│  - Chama Models                             │
└──────────────────┬──────────────────────────┘
                   │
                   ↓
┌─────────────────────────────────────────────┐
│         MODELS (Entidades GORM)             │
│  - Define estrutura das tabelas             │
│  - Relacionamentos                          │
└──────────────────┬──────────────────────────┘
                   │
                   ↓
┌─────────────────────────────────────────────┐
│         DATABASE (MySQL)                    │
└─────────────────────────────────────────────┘
```

---

## 🛠️ TECNOLOGIAS

### Backend

- **Linguagem:** Go 1.21+
- **Framework Web:** Gin (github.com/gin-gonic/gin)
- **ORM:** GORM (gorm.io/gorm)
- **Driver MySQL:** gorm.io/driver/mysql
- **Validação:** go-playground/validator
- **Env Variables:** godotenv
- **Hot Reload:** Air

### Banco de Dados

- **SGBD:** MySQL 8.0+
- **Charset:** utf8mb4
- **Collation:** utf8mb4_unicode_ci

### Futuro (Frontend)

- **Framework:** Next.js 15+
- **Linguagem:** TypeScript

---

## 📁 ESTRUTURA DE PASTAS

```
go-pratica/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point alternativo
├── controllers/                     # Controladores (lógica)
│   ├── cliente_controller.go
│   ├── veiculo_controller.go
│   ├── os_controller.go
│   ├── orcamento_controller.go
│   └── pagamento_controller.go
├── database/                        # Configuração do banco
│   ├── connection.go
│   └── migrations.go
├── dto/                            # Data Transfer Objects
│   ├── cliente_dto.go
│   ├── veiculo_dto.go
│   ├── os_dto.go
│   └── orcamento_dto.go
├── models/                         # Entidades/Modelos
│   ├── cliente.go
│   ├── endereco.go
│   ├── veiculo.go
│   ├── ordem_servico.go
│   ├── orcamento.go
│   ├── item_orcamento.go
│   ├── servico_executado.go
│   └── pagamento.go
├── routes/                         # Roteamento
│   ├── routes.go
│   ├── cliente_routes.go
│   ├── veiculo_routes.go
│   ├── os_routes.go
│   └── orcamento_routes.go
├── middlewares/                    # Middlewares (futuro)
│   ├── auth.go
│   ├── cors.go
│   └── logger.go
├── pkg/                           # Pacotes reutilizáveis
│   ├── validator/
│   ├── response/
│   └── utils/
├── tmp/                           # Arquivos temporários (Air)
├── .air.toml                      # Config do Air
├── .env                           # Variáveis de ambiente
├── .env.example                   # Exemplo de .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go                        # Entry point principal
└── README.md
```

---

## 📊 MODELAGEM DO BANCO DE DADOS

### Diagrama de Relacionamentos

```
┌─────────────┐
│   CLIENTE   │
└──────┬──────┘
       │ 1
       │
       │ N
┌──────┴──────┐         ┌──────────────┐
│   ENDEREÇO  │         │   VEÍCULO    │
└─────────────┘         └──────┬───────┘
                               │ 1
                               │
                               │ N
                        ┌──────┴───────┐
                        │ ORDEM_SERVICO│
                        └──────┬───────┘
                               │ 1
                ┌──────────────┼──────────────┐
                │              │              │
                │ N            │ N            │ N
        ┌───────┴──────┐ ┌────┴─────┐ ┌──────┴─────────┐
        │  ORÇAMENTO   │ │ SERVIÇO  │ │   PAGAMENTO    │
        └───────┬──────┘ │EXECUTADO │ └────────────────┘
                │        └──────────┘
                │ 1
                │
                │ N
        ┌───────┴───────┐
        │ITEM_ORÇAMENTO │
        └───────────────┘
```

### Entidades e Atributos

#### **1. CLIENTE**

- `id` (PK)
- `nome`
- `cpf` (UNIQUE)
- `telefone_principal`
- `telefone_secundario`
- `email`
- `status` (ativo/inativo)
- `created_at`
- `updated_at`
- `deleted_at` (soft delete)

#### **2. ENDERECO**

- `id` (PK)
- `cliente_id` (FK → CLIENTE)
- `cep`
- `rua`
- `numero`
- `complemento`
- `bairro`
- `cidade`
- `estado`
- `tipo` (residencial/comercial)
- `created_at`
- `updated_at`

#### **3. VEICULO**

- `id` (PK)
- `cliente_id` (FK → CLIENTE)
- `placa` (UNIQUE)
- `marca`
- `modelo`
- `ano_fabricacao`
- `ano_modelo`
- `cor`
- `chassi`
- `quilometragem_atual`
- `created_at`
- `updated_at`
- `deleted_at`

#### **4. ORDEM_SERVICO**

- `id` (PK)
- `numero_os` (UNIQUE, AUTO INCREMENT)
- `cliente_id` (FK → CLIENTE)
- `veiculo_id` (FK → VEICULO)
- `data_entrada`
- `quilometragem_entrada`
- `descricao_problema`
- `status` (ENUM)
- `data_prevista_conclusao`
- `data_conclusao_real`
- `observacoes`
- `created_at`
- `updated_at`

**Status possíveis:**

- `aguardando_orcamento`
- `orcamento_enviado`
- `aprovado`
- `em_execucao`
- `finalizado`
- `entregue`
- `cancelado`

#### **5. ORCAMENTO**

- `id` (PK)
- `os_id` (FK → ORDEM_SERVICO)
- `numero_orcamento`
- `data_criacao`
- `valor_pecas`
- `valor_mao_obra`
- `valor_total`
- `desconto`
- `valor_final`
- `data_validade`
- `status` (ENUM: pendente, aprovado, recusado, expirado)
- `data_aprovacao_recusa`
- `observacoes`
- `created_at`
- `updated_at`

#### **6. ITEM_ORCAMENTO**

- `id` (PK)
- `orcamento_id` (FK → ORCAMENTO)
- `tipo` (ENUM: peca, servico)
- `descricao`
- `quantidade`
- `valor_unitario`
- `valor_total`
- `observacao`
- `created_at`
- `updated_at`

#### **7. SERVICO_EXECUTADO**

- `id` (PK)
- `os_id` (FK → ORDEM_SERVICO)
- `mecanico_responsavel`
- `data_inicio`
- `data_conclusao`
- `descricao_servico`
- `tempo_estimado_horas`
- `tempo_real_horas`
- `status` (ENUM: pendente, em_andamento, concluido)
- `created_at`
- `updated_at`

#### **8. PAGAMENTO**

- `id` (PK)
- `os_id` (FK → ORDEM_SERVICO)
- `data_pagamento`
- `forma_pagamento` (ENUM)
- `valor_pago`
- `desconto_aplicado`
- `valor_final`
- `status` (ENUM: pendente, pago_parcial, pago_total)
- `observacoes`
- `created_at`
- `updated_at`

**Formas de pagamento:**

- `dinheiro`
- `cartao_debito`
- `cartao_credito`
- `pix`
- `boleto`
- `cheque`

---

## 💾 SCRIPTS SQL - MYSQL

### 1. Criar Banco de Dados

```sql
-- Criar banco de dados
CREATE DATABASE IF NOT EXISTS mecanica_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- Usar o banco
USE mecanica_db;
```

### 2. Tabela: CLIENTE

```sql
CREATE TABLE clientes (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    telefone_principal VARCHAR(20) NOT NULL,
    telefone_secundario VARCHAR(20),
    email VARCHAR(255),
    status ENUM('ativo', 'inativo') DEFAULT 'ativo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_cpf (cpf),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 3. Tabela: ENDERECO

```sql
CREATE TABLE enderecos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cliente_id BIGINT UNSIGNED NOT NULL,
    cep VARCHAR(10) NOT NULL,
    rua VARCHAR(255) NOT NULL,
    numero VARCHAR(20) NOT NULL,
    complemento VARCHAR(255),
    bairro VARCHAR(100) NOT NULL,
    cidade VARCHAR(100) NOT NULL,
    estado CHAR(2) NOT NULL,
    tipo ENUM('residencial', 'comercial') DEFAULT 'residencial',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE,
    INDEX idx_cliente_id (cliente_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 4. Tabela: VEICULO

```sql
CREATE TABLE veiculos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cliente_id BIGINT UNSIGNED NOT NULL,
    placa VARCHAR(10) NOT NULL UNIQUE,
    marca VARCHAR(100) NOT NULL,
    modelo VARCHAR(100) NOT NULL,
    ano_fabricacao INT NOT NULL,
    ano_modelo INT NOT NULL,
    cor VARCHAR(50),
    chassi VARCHAR(50),
    quilometragem_atual INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE,
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_placa (placa),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 5. Tabela: ORDEM_SERVICO

```sql
CREATE TABLE ordens_servico (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    numero_os VARCHAR(20) NOT NULL UNIQUE,
    cliente_id BIGINT UNSIGNED NOT NULL,
    veiculo_id BIGINT UNSIGNED NOT NULL,
    data_entrada TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    quilometragem_entrada INT,
    descricao_problema TEXT NOT NULL,
    status ENUM(
        'aguardando_orcamento',
        'orcamento_enviado',
        'aprovado',
        'em_execucao',
        'finalizado',
        'entregue',
        'cancelado'
    ) DEFAULT 'aguardando_orcamento',
    data_prevista_conclusao DATE,
    data_conclusao_real TIMESTAMP NULL,
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE RESTRICT,
    FOREIGN KEY (veiculo_id) REFERENCES veiculos(id) ON DELETE RESTRICT,
    INDEX idx_numero_os (numero_os),
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_veiculo_id (veiculo_id),
    INDEX idx_status (status),
    INDEX idx_data_entrada (data_entrada)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 6. Tabela: ORCAMENTO

```sql
CREATE TABLE orcamentos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    numero_orcamento VARCHAR(20) NOT NULL UNIQUE,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valor_pecas DECIMAL(10, 2) DEFAULT 0.00,
    valor_mao_obra DECIMAL(10, 2) DEFAULT 0.00,
    valor_total DECIMAL(10, 2) DEFAULT 0.00,
    desconto DECIMAL(10, 2) DEFAULT 0.00,
    valor_final DECIMAL(10, 2) DEFAULT 0.00,
    data_validade DATE,
    status ENUM('pendente', 'aprovado', 'recusado', 'expirado') DEFAULT 'pendente',
    data_aprovacao_recusa TIMESTAMP NULL,
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_numero_orcamento (numero_orcamento)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 7. Tabela: ITEM_ORCAMENTO

```sql
CREATE TABLE itens_orcamento (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    orcamento_id BIGINT UNSIGNED NOT NULL,
    tipo ENUM('peca', 'servico') NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    quantidade INT NOT NULL DEFAULT 1,
    valor_unitario DECIMAL(10, 2) NOT NULL,
    valor_total DECIMAL(10, 2) NOT NULL,
    observacao TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (orcamento_id) REFERENCES orcamentos(id) ON DELETE CASCADE,
    INDEX idx_orcamento_id (orcamento_id),
    INDEX idx_tipo (tipo)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 8. Tabela: SERVICO_EXECUTADO

```sql
CREATE TABLE servicos_executados (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    mecanico_responsavel VARCHAR(255),
    data_inicio TIMESTAMP NULL,
    data_conclusao TIMESTAMP NULL,
    descricao_servico TEXT NOT NULL,
    tempo_estimado_horas DECIMAL(5, 2),
    tempo_real_horas DECIMAL(5, 2),
    status ENUM('pendente', 'em_andamento', 'concluido') DEFAULT 'pendente',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 9. Tabela: PAGAMENTO

```sql
CREATE TABLE pagamentos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    data_pagamento TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    forma_pagamento ENUM(
        'dinheiro',
        'cartao_debito',
        'cartao_credito',
        'pix',
        'boleto',
        'cheque'
    ) NOT NULL,
    valor_pago DECIMAL(10, 2) NOT NULL,
    desconto_aplicado DECIMAL(10, 2) DEFAULT 0.00,
    valor_final DECIMAL(10, 2) NOT NULL,
    status ENUM('pendente', 'pago_parcial', 'pago_total') DEFAULT 'pendente',
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE RESTRICT,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_forma_pagamento (forma_pagamento)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 10. Script Completo de Criação

```sql
-- =====================================================
-- SCRIPT COMPLETO DE CRIAÇÃO DO BANCO DE DADOS
-- Sistema de Gerenciamento de Mecânica
-- =====================================================

-- Criar e usar o banco
CREATE DATABASE IF NOT EXISTS mecanica_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE mecanica_db;

-- Tabela: CLIENTES
CREATE TABLE clientes (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    telefone_principal VARCHAR(20) NOT NULL,
    telefone_secundario VARCHAR(20),
    email VARCHAR(255),
    status ENUM('ativo', 'inativo') DEFAULT 'ativo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_cpf (cpf),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: ENDERECOS
CREATE TABLE enderecos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cliente_id BIGINT UNSIGNED NOT NULL,
    cep VARCHAR(10) NOT NULL,
    rua VARCHAR(255) NOT NULL,
    numero VARCHAR(20) NOT NULL,
    complemento VARCHAR(255),
    bairro VARCHAR(100) NOT NULL,
    cidade VARCHAR(100) NOT NULL,
    estado CHAR(2) NOT NULL,
    tipo ENUM('residencial', 'comercial') DEFAULT 'residencial',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE,
    INDEX idx_cliente_id (cliente_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: VEICULOS
CREATE TABLE veiculos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cliente_id BIGINT UNSIGNED NOT NULL,
    placa VARCHAR(10) NOT NULL UNIQUE,
    marca VARCHAR(100) NOT NULL,
    modelo VARCHAR(100) NOT NULL,
    ano_fabricacao INT NOT NULL,
    ano_modelo INT NOT NULL,
    cor VARCHAR(50),
    chassi VARCHAR(50),
    quilometragem_atual INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE,
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_placa (placa),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: ORDENS_SERVICO
CREATE TABLE ordens_servico (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    numero_os VARCHAR(20) NOT NULL UNIQUE,
    cliente_id BIGINT UNSIGNED NOT NULL,
    veiculo_id BIGINT UNSIGNED NOT NULL,
    data_entrada TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    quilometragem_entrada INT,
    descricao_problema TEXT NOT NULL,
    status ENUM(
        'aguardando_orcamento',
        'orcamento_enviado',
        'aprovado',
        'em_execucao',
        'finalizado',
        'entregue',
        'cancelado'
    ) DEFAULT 'aguardando_orcamento',
    data_prevista_conclusao DATE,
    data_conclusao_real TIMESTAMP NULL,
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE RESTRICT,
    FOREIGN KEY (veiculo_id) REFERENCES veiculos(id) ON DELETE RESTRICT,
    INDEX idx_numero_os (numero_os),
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_veiculo_id (veiculo_id),
    INDEX idx_status (status),
    INDEX idx_data_entrada (data_entrada)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: ORCAMENTOS
CREATE TABLE orcamentos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    numero_orcamento VARCHAR(20) NOT NULL UNIQUE,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valor_pecas DECIMAL(10, 2) DEFAULT 0.00,
    valor_mao_obra DECIMAL(10, 2) DEFAULT 0.00,
    valor_total DECIMAL(10, 2) DEFAULT 0.00,
    desconto DECIMAL(10, 2) DEFAULT 0.00,
    valor_final DECIMAL(10, 2) DEFAULT 0.00,
    data_validade DATE,
    status ENUM('pendente', 'aprovado', 'recusado', 'expirado') DEFAULT 'pendente',
    data_aprovacao_recusa TIMESTAMP NULL,
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_numero_orcamento (numero_orcamento)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: ITENS_ORCAMENTO
CREATE TABLE itens_orcamento (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    orcamento_id BIGINT UNSIGNED NOT NULL,
    tipo ENUM('peca', 'servico') NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    quantidade INT NOT NULL DEFAULT 1,
    valor_unitario DECIMAL(10, 2) NOT NULL,
    valor_total DECIMAL(10, 2) NOT NULL,
    observacao TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (orcamento_id) REFERENCES orcamentos(id) ON DELETE CASCADE,
    INDEX idx_orcamento_id (orcamento_id),
    INDEX idx_tipo (tipo)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: SERVICOS_EXECUTADOS
CREATE TABLE servicos_executados (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    mecanico_responsavel VARCHAR(255),
    data_inicio TIMESTAMP NULL,
    data_conclusao TIMESTAMP NULL,
    descricao_servico TEXT NOT NULL,
    tempo_estimado_horas DECIMAL(5, 2),
    tempo_real_horas DECIMAL(5, 2),
    status ENUM('pendente', 'em_andamento', 'concluido') DEFAULT 'pendente',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Tabela: PAGAMENTOS
CREATE TABLE pagamentos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    data_pagamento TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    forma_pagamento ENUM(
        'dinheiro',
        'cartao_debito',
        'cartao_credito',
        'pix',
        'boleto',
        'cheque'
    ) NOT NULL,
    valor_pago DECIMAL(10, 2) NOT NULL,
    desconto_aplicado DECIMAL(10, 2) DEFAULT 0.00,
    valor_final DECIMAL(10, 2) NOT NULL,
    status ENUM('pendente', 'pago_parcial', 'pago_total') DEFAULT 'pendente',
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE RESTRICT,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_forma_pagamento (forma_pagamento)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Mensagem de conclusão
SELECT 'Banco de dados criado com sucesso!' AS status;
```

---

## 🛣️ ENDPOINTS DA API

### Base URL

```
http://localhost:8080/api/v1
```

### **CLIENTES**

| Método | Endpoint                  | Descrição                     |
| ------ | ------------------------- | ----------------------------- |
| POST   | `/clientes`               | Criar novo cliente            |
| GET    | `/clientes`               | Listar todos os clientes      |
| GET    | `/clientes/:id`           | Buscar cliente por ID         |
| PUT    | `/clientes/:id`           | Atualizar cliente             |
| DELETE | `/clientes/:id`           | Deletar cliente (soft delete) |
| GET    | `/clientes/:id/veiculos`  | Listar veículos do cliente    |
| GET    | `/clientes/:id/historico` | Histórico de OS do cliente    |

### **VEÍCULOS**

| Método | Endpoint                 | Descrição                |
| ------ | ------------------------ | ------------------------ |
| POST   | `/veiculos`              | Cadastrar veículo        |
| GET    | `/veiculos`              | Listar todos os veículos |
| GET    | `/veiculos/:id`          | Buscar veículo por ID    |
| GET    | `/veiculos/placa/:placa` | Buscar veículo por placa |
| PUT    | `/veiculos/:id`          | Atualizar veículo        |
| DELETE | `/veiculos/:id`          | Deletar veículo          |

### **ORDENS DE SERVIÇO**

| Método | Endpoint                              | Descrição                 |
| ------ | ------------------------------------- | ------------------------- |
| POST   | `/ordens-servico`                     | Criar nova OS             |
| GET    | `/ordens-servico`                     | Listar todas as OS        |
| GET    | `/ordens-servico/:id`                 | Buscar OS por ID          |
| PUT    | `/ordens-servico/:id`                 | Atualizar OS              |
| PATCH  | `/ordens-servico/:id/status`          | Atualizar apenas o status |
| GET    | `/ordens-servico/cliente/:cliente_id` | OS por cliente            |
| GET    | `/ordens-servico/veiculo/:veiculo_id` | OS por veículo            |

### **ORÇAMENTOS**

| Método | Endpoint                  | Descrição                   |
| ------ | ------------------------- | --------------------------- |
| POST   | `/orcamentos`             | Criar orçamento para OS     |
| GET    | `/orcamentos/:id`         | Buscar orçamento            |
| PUT    | `/orcamentos/:id`         | Atualizar orçamento         |
| PATCH  | `/orcamentos/:id/aprovar` | Aprovar orçamento           |
| PATCH  | `/orcamentos/:id/recusar` | Recusar orçamento           |
| POST   | `/orcamentos/:id/itens`   | Adicionar item ao orçamento |

### **PAGAMENTOS**

| Método | Endpoint                | Descrição                   |
| ------ | ----------------------- | --------------------------- |
| POST   | `/pagamentos`           | Registrar pagamento         |
| GET    | `/pagamentos/os/:os_id` | Buscar pagamentos de uma OS |
| GET    | `/pagamentos/:id`       | Buscar pagamento por ID     |

---

## 🔄 FLUXO DO SISTEMA

```
┌─────────────────────────────────────────────────────────────┐
│                    FLUXO COMPLETO                           │
└─────────────────────────────────────────────────────────────┘

1️⃣ Cliente chega com problema no veículo
   └─> Funcionário cria ORDEM DE SERVIÇO
       ├─> Registra: cliente, veículo, problema
       └─> Status: "aguardando_orcamento"

2️⃣ Mecânico avalia e cria ORÇAMENTO
   └─> Lista peças e serviços necessários
       ├─> Calcula valores (peças + mão de obra)
       └─> Status da OS: "orcamento_enviado"

3️⃣ Cliente recebe orçamento
   └─> Decisão:
       ├─> ✅ Aprova → Status: "aprovado"
       └─> ❌ Recusa → Status: "cancelado"

4️⃣ Serviços são executados
   └─> Status da OS: "em_execucao"
       ├─> Mecânico registra serviços
       └─> Atualiza progresso

5️⃣ Serviços concluídos
   └─> Status da OS: "finalizado"
       └─> Sistema notifica cliente

6️⃣ Cliente vem buscar o veículo
   └─> Registra PAGAMENTO
       ├─> Forma de pagamento
       ├─> Valor pago
       └─> Status da OS: "entregue"

7️⃣ Histórico salvo
   └─> Fica disponível para consultas futuras
```

---

## 📅 PLANO DE EXECUÇÃO

### **Sprint 1 - Setup Inicial (Semana 1)**

- [x] Configurar ambiente Go
- [x] Instalar dependências (Gin, GORM, MySQL Driver)
- [x] Criar estrutura de pastas
- [x] Configurar conexão com MySQL
- [x] Criar arquivo `.env`
- [ ] Executar scripts SQL para criar tabelas
- [x] Configurar Air para hot reload

### **Sprint 2 - Módulo Clientes (Semana 2)**

- [ ] Criar `models/cliente.go`
- [ ] Criar `dto/cliente_dto.go`
- [ ] Implementar `controllers/cliente_controller.go`
  - [ ] CreateCliente
  - [ ] GetClientes (com paginação)
  - [ ] GetClienteByID
  - [ ] UpdateCliente
  - [ ] DeleteCliente (soft delete)
- [ ] Configurar `routes/cliente_routes.go`
- [ ] Testar todos os endpoints no Postman
- [ ] Adicionar validações

### **Sprint 3 - Módulo Veículos (Semana 3)**

- [ ] Criar `models/veiculo.go` com relacionamento
- [ ] Criar `dto/veiculo_dto.go`
- [ ] Implementar `controllers/veiculo_controller.go`
  - [ ] CRUD completo
  - [ ] Buscar por placa
  - [ ] Listar veículos por cliente
- [ ] Configurar `routes/veiculo_routes.go`
- [ ] Testar relacionamento Cliente ↔ Veículos
- [ ] Validar placa e dados

### **Sprint 4 - Módulo OS (Semanas 4-5)**

- [ ] Criar `models/ordem_servico.go`
- [ ] Criar `dto/os_dto.go`
- [ ] Implementar `controllers/os_controller.go`
  - [ ] Criar OS
  - [ ] Listar com filtros (status, cliente, veículo)
  - [ ] Atualizar status
  - [ ] Buscar OS completa com relacionamentos
- [ ] Configurar `routes/os_routes.go`
- [ ] Implementar lógica de mudança de status
- [ ] Testar fluxo completo

### **Sprint 5 - Módulo Orçamento (Semana 5)**

- [ ] Criar `models/orcamento.go` e `models/item_orcamento.go`
- [ ] Criar `dto/orcamento_dto.go`
- [ ] Implementar `controllers/orcamento_controller.go`
  - [ ] Criar orçamento
  - [ ] Adicionar itens
  - [ ] Calcular totais automaticamente
  - [ ] Aprovar/Recusar orçamento
- [ ] Atualizar status da OS quando orçamento aprovado
- [ ] Testar cálculos

### **Sprint 6 - Módulo Pagamento (Semana 6)**

- [ ] Criar `models/pagamento.go`
- [ ] Criar `dto/pagamento_dto.go`
- [ ] Implementar `controllers/pagamento_controller.go`
  - [ ] Registrar pagamento
  - [ ] Validar valor com orçamento
  - [ ] Atualizar status da OS
- [ ] Finalizar fluxo completo

### **Sprint 7 - Melhorias e Testes (Semana 7)**

- [ ] Implementar endpoint de histórico
- [ ] Adicionar filtros avançados
- [ ] Tratamento de erros robusto
- [ ] Logs estruturados
- [ ] Documentação com Swagger (opcional)
- [ ] Testes completos

---

## ✅ PRÓXIMOS PASSOS IMEDIATOS

### **1. Configuração Inicial**

**Instalar dependências:**

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/joho/godotenv
go get -u github.com/go-playground/validator/v10
```

**Criar arquivo `.env`:**

```env
# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=sua_senha
DB_NAME=mecanica_db

# Server
PORT=8080
GIN_MODE=debug

# JWT (futuro)
JWT_SECRET=sua_chave_secreta_aqui
```

### **2. Criar Banco de Dados**

Executar o script SQL completo no MySQL:

```bash
mysql -u root -p < create_database.sql
```

### **3. Estruturar Projeto**

Criar todas as pastas:

```bash
mkdir -p controllers database dto models routes middlewares pkg/validator pkg/response
```

### **4. Implementar Conexão com Banco**

Arquivo `database/connection.go`:

```go
package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Erro ao conectar no banco de dados:", err)
    }

    log.Println("✅ Conexão com banco de dados estabelecida!")
}
```

### **5. Começar pelo Módulo de Clientes**

Seguir a ordem:

1. `models/cliente.go`
2. `dto/cliente_dto.go`
3. `controllers/cliente_controller.go`
4. `routes/cliente_routes.go`
5. Testar no Postman

---

## 🚀 COMANDOS ÚTEIS

### **Iniciar projeto:**

```bash
# Com Air (hot reload)
air

# Sem Air
go run main.go
```

### **Gerenciar dependências:**

```bash
# Baixar dependências
go mod download

# Limpar módulos não utilizados
go mod tidy
```

### **Testar banco:**

```bash
# Conectar no MySQL
mysql -u root -p

# Ver databases
SHOW DATABASES;

# Usar o banco
USE mecanica_db;

# Listar tabelas
SHOW TABLES;

# Ver estrutura de uma tabela
DESCRIBE clientes;
```

---

## 📚 RECURSOS DE ESTUDO

### **Documentação Oficial:**

- [Go Documentation](https://go.dev/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [GORM](https://gorm.io/docs/)

### **Conceitos importantes:**

- REST API
- CRUD Operations
- SQL Relationships (1:1, 1:N, N:N)
- HTTP Status Codes
- JSON Marshaling/Unmarshaling

---

## 🔮 FASE 2 - FEATURES FUTURAS

- [ ] Autenticação JWT
- [ ] Permissões de usuário (admin, mecânico, atendente)
- [ ] Controle de estoque de peças
- [ ] Gestão de fornecedores
- [ ] Relatórios financeiros
- [ ] Dashboard com métricas
- [ ] Notificações por email/SMS
- [ ] Upload de fotos do veículo
- [ ] Assinatura digital do cliente
- [ ] Histórico de manutenções preventivas

---

## 📝 NOTAS IMPORTANTES

1. **Sempre usar soft delete** para clientes e veículos
2. **Validar CPF** antes de cadastrar cliente
3. **Validar placa** (formato ABC-1234 ou ABC1D234)
4. **Gerar número de OS automaticamente** (ex: OS-2024-0001)
5. **Calcular totais automaticamente** nos orçamentos
6. **Validar datas** (data de conclusão > data de entrada)
7. **Impedir exclusão** de clientes/veículos com OS ativas
8. **Logs estruturados** para rastreabilidade

---

## 👥 CONTRIBUIÇÃO

Este é um projeto de aprendizado. Para contribuir:

1. Fork o projeto
2. Crie uma branch (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

---

## 📄 LICENÇA

Este projeto é privado e destinado a fins educacionais.

---

## ✍️ AUTOR

Desenvolvido como projeto de aprendizado - Backend em Go

**Data de criação:** Fevereiro 2026  
**Versão:** 1.0  
**Status:** Em desenvolvimento

---

**🎯 Objetivo:** Aprender desenvolvimento backend com Go, GORM e arquitetura MVC!
