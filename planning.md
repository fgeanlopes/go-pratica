## Regras de Interação com IA

A IA utilizada neste projeto (como GitHub Copilot ou qualquer outro assistente) deve seguir estritamente as seguintes regras.

### Comportamentos proibidos

- Não criar arquivos no projeto.
- Não modificar arquivos existentes.
- Não implementar funcionalidades.
- Não gerar código pronto para uso no projeto.

### Comportamentos permitidos

- Explicar conceitos técnicos.
- Orientar sobre arquitetura ou boas práticas.
- Sugerir abordagens para resolver problemas.
- Fornecer pequenos trechos de código apenas como exemplo ilustrativo.

### Regra principal

A IA deve atuar **apenas como orientadora técnica**.
Toda implementação e escrita de código deve ser feita pelo desenvolvedor.

Estas regras têm prioridade sobre qualquer outra instrução presente neste documento.

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
2. [Padrões e Nomenclatura](#padrões-e-nomenclatura)
3. [Arquitetura](#arquitetura)
4. [Tecnologias](#tecnologias)
5. [Estrutura de Pastas](#estrutura-de-pastas)
6. [Modelagem do Banco de Dados](#modelagem-do-banco-de-dados)
7. [Scripts SQL - MySQL](#scripts-sql---mysql)
8. [Endpoints da API](#endpoints-da-api)
9. [Fluxo do Sistema](#fluxo-do-sistema)
10. [Plano de Execução](#plano-de-execução)
11. [Próximos Passos](#próximos-passos)

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

## 📐 PADRÕES E NOMENCLATURA

### **Regra Geral do Projeto:**

- ✅ **TODO CÓDIGO EM INGLÊS** (variáveis, funções, structs, nomes de arquivos, rotas, tabelas)
- ✅ **COMENTÁRIOS EM PORTUGUÊS** (para facilitar entendimento)

### **Nomenclatura Padrão:**

| Português         | Inglês           | Uso                                          |
| ----------------- | ---------------- | -------------------------------------------- |
| Cliente           | Client           | `models/client.go`, `/clients`               |
| Veículo           | Vehicle          | `models/vehicle.go`, `/vehicles`             |
| Ordem de Serviço  | Service Order    | `models/service_order.go`, `/service-orders` |
| Orçamento         | Budget           | `models/budget.go`, `/budgets`               |
| Pagamento         | Payment          | `models/payment.go`, `/payments`             |
| Endereço          | Address          | `models/address.go`                          |
| Item do Orçamento | Budget Item      | `models/budget_item.go`                      |
| Serviço Executado | Executed Service | `models/executed_service.go`                 |

### **Campos Comuns:**

| Português           | Inglês           |
| ------------------- | ---------------- |
| Nome                | Name             |
| CPF                 | CPF (mantém)     |
| Telefone Principal  | Primary Phone    |
| Telefone Secundário | Secondary Phone  |
| Email               | Email (mantém)   |
| CEP                 | Zip Code         |
| Rua                 | Street           |
| Número              | Number           |
| Bairro              | Neighborhood     |
| Complemento         | Complement       |
| Cidade              | City             |
| Estado              | State            |
| Placa               | Plate            |
| Marca               | Brand            |
| Modelo              | Model            |
| Ano de Fabricação   | Manufacture Year |
| Ano do Modelo       | Model Year       |
| Cor                 | Color            |
| Tipo de Combustível | Fuel Type        |
| Chassi              | Chassis          |
| RENAVAM             | RENAVAM (mantém) |
| Quilometragem       | Mileage          |
| Número do Motor     | Engine Number    |
| Observações         | Notes            |
| Status              | Status (mantém)  |
| Ativo/Inativo       | Active/Inactive  |
| Data de Entrada     | Entry Date       |
| Data de Criação     | Created At       |
| Data de Atualização | Updated At       |

### **Status e Enums:**

**Status da Ordem de Serviço:**

- `awaiting_budget` (aguardando orçamento)
- `budget_sent` (orçamento enviado)
- `approved` (aprovado)
- `in_progress` (em execução)
- `completed` (finalizado)
- `delivered` (entregue)
- `cancelled` (cancelado)

**Status do Orçamento:**

- `pending` (pendente)
- `approved` (aprovado)
- `rejected` (recusado)
- `expired` (expirado)

**Tipos de Combustível:**

- `gasoline` (gasolina)
- `ethanol` (etanol)
- `flex` (flex)
- `diesel` (diesel)
- `electric` (elétrico)
- `hybrid` (híbrido)

**Formas de Pagamento:**

- `cash` (dinheiro)
- `debit_card` (cartão de débito)
- `credit_card` (cartão de crédito)
- `pix` (PIX)
- `bank_slip` (boleto)
- `check` (cheque)

### **Padrão de Rotas:**

- Usar **kebab-case** para URLs: `/service-orders`, `/budget-items`
- Usar **plural** para recursos: `/clients`, `/vehicles`, `/payments`
- Usar **singular** para ações específicas: `/approve`, `/reject`

### **Padrão de Arquivos:**

- Usar **snake_case** para arquivos Go: `client_controller.go`, `service_order_dto.go`
- Usar **PascalCase** para structs: `CreateClientRequest`, `ServiceOrder`
- Usar **camelCase** para variáveis e funções: `clientID`, `getClientByID`

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
│   ├── client_controller.go
│   ├── vehicle_controller.go
│   ├── service_order_controller.go
│   ├── budget_controller.go
│   └── payment_controller.go
├── database/                        # Configuração do banco
│   ├── connection.go
│   └── migrations.go
├── dto/                            # Data Transfer Objects
│   ├── client_dto.go
│   ├── vehicle_dto.go
│   ├── service_order_dto.go
│   └── budget_dto.go
├── models/                         # Entidades/Modelos
│   ├── client.go
│   ├── address.go
│   ├── vehicle.go
│   ├── service_order.go
│   ├── budget.go
│   ├── budget_item.go
│   ├── executed_service.go
│   └── payment.go
├── routes/                         # Roteamento
│   ├── routes.go
│   ├── client_routes.go
│   ├── vehicle_routes.go
│   ├── service_order_routes.go
│   └── budget_routes.go
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

#### **1. CLIENT (Cliente)**

- `id` (PK)
- `name`
- `cpf` (UNIQUE)
- `primary_phone`
- `secondary_phone`
- `email`
- `status` (active/inactive)
- `created_at`
- `updated_at`
- `deleted_at` (soft delete)

#### **2. ADDRESS (Endereço)**

- `id` (PK)
- `client_id` (FK → CLIENT)
- `zip_code`
- `street`
- `number`
- `complement`
- `neighborhood`
- `city`
- `state`
- `type` (residential/commercial)
- `created_at`
- `updated_at`

#### **3. VEHICLE (Veículo)**

- `id` (PK)
- `client_id` (FK → CLIENT)
- `plate` (UNIQUE)
- `brand`
- `model`
- `manufacture_year`
- `model_year`
- `color`
- `fuel_type` (ENUM: gasoline, ethanol, flex, diesel, electric, hybrid)
- `chassis` (REQUIRED)
- `renavam` (REQUIRED)
- `current_mileage` (REQUIRED)
- `engine_number` (REQUIRED)
- `status` (OPTIONAL)
- `notes` (OPTIONAL)
- `created_at`
- `updated_at`
- `deleted_at`

#### **4. SERVICE_ORDER (Ordem de Serviço)**

- `id` (PK)
- `order_number` (UNIQUE, AUTO INCREMENT)
- `client_id` (FK → CLIENT)
- `vehicle_id` (FK → VEHICLE)
- `entry_date`
- `entry_mileage`
- `problem_description`
- `status` (ENUM)
- `expected_completion_date`
- `actual_completion_date`
- `notes`
- `created_at`
- `updated_at`

**Status possíveis:**

- `awaiting_budget`
- `budget_sent`
- `approved`
- `in_progress`
- `completed`
- `delivered`
- `cancelled`

#### **5. BUDGET (Orçamento)**

- `id` (PK)
- `service_order_id` (FK → SERVICE_ORDER)
- `budget_number`
- `creation_date`
- `parts_value`
- `labor_value`
- `total_value`
- `discount`
- `final_value`
- `validity_date`
- `status` (ENUM: pending, approved, rejected, expired)
- `approval_rejection_date`
- `notes`
- `created_at`
- `updated_at`

#### **6. BUDGET_ITEM (Item do Orçamento)**

- `id` (PK)
- `budget_id` (FK → BUDGET)
- `type` (ENUM: part, service)
- `description`
- `quantity`
- `unit_price`
- `total_price`
- `notes`
- `created_at`
- `updated_at`

#### **7. EXECUTED_SERVICE (Serviço Executado)**

- `id` (PK)
- `service_order_id` (FK → SERVICE_ORDER)
- `mechanic_name`
- `start_date`
- `completion_date`
- `service_description`
- `estimated_hours`
- `actual_hours`
- `status` (ENUM: pending, in_progress, completed)
- `created_at`
- `updated_at`

#### **8. PAYMENT (Pagamento)**

- `id` (PK)
- `service_order_id` (FK → SERVICE_ORDER)
- `payment_date`
- `payment_method` (ENUM)
- `amount_paid`
- `discount_applied`
- `final_amount`
- `status` (ENUM: pending, partial_paid, fully_paid)
- `notes`
- `created_at`
- `updated_at`

**Formas de pagamento:**

- `cash`
- `debit_card`
- `credit_card`
- `pix`
- `bank_slip`
- `check`

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

### 2. Tabela: CLIENTS

```sql
CREATE TABLE clients (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    primary_phone VARCHAR(20) NOT NULL,
    secondary_phone VARCHAR(20),
    email VARCHAR(255),
    status ENUM('active', 'inactive') DEFAULT 'active',
    zip_code VARCHAR(10),
    street VARCHAR(255),
    number VARCHAR(20),
    complement VARCHAR(255),
    neighborhood VARCHAR(100),
    city VARCHAR(100),
    state CHAR(2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_cpf (cpf),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 3. Tabela: VEHICLES

```sql
CREATE TABLE vehicles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    client_id BIGINT UNSIGNED NOT NULL,
    plate VARCHAR(10) NOT NULL UNIQUE,
    brand VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    manufacture_year INT NOT NULL,
    model_year INT NOT NULL,
    color VARCHAR(50),
    chassis VARCHAR(50),
    current_mileage INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    INDEX idx_client_id (client_id),
    INDEX idx_plate (plate),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 4. Tabela: SERVICE_ORDERS

```sql
CREATE TABLE service_orders (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(20) NOT NULL UNIQUE,
    client_id BIGINT UNSIGNED NOT NULL,
    vehicle_id BIGINT UNSIGNED NOT NULL,
    entry_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    entry_mileage INT,
    problem_description TEXT NOT NULL,
    status ENUM(
        'awaiting_budget',
        'budget_sent',
        'approved',
        'in_progress',
        'completed',
        'delivered',
        'cancelled'
    ) DEFAULT 'awaiting_budget',
    expected_completion_date DATE,
    actual_completion_date TIMESTAMP NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE RESTRICT,
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE RESTRICT,
    INDEX idx_order_number (order_number),
    INDEX idx_client_id (client_id),
    INDEX idx_vehicle_id (vehicle_id),
    INDEX idx_status (status),
    INDEX idx_entry_date (entry_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 5. Tabela: BUDGETS

```sql
CREATE TABLE budgets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_order_id BIGINT UNSIGNED NOT NULL,
    budget_number VARCHAR(20) NOT NULL UNIQUE,
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parts_amount DECIMAL(10, 2) DEFAULT 0.00,
    labor_amount DECIMAL(10, 2) DEFAULT 0.00,
    total_amount DECIMAL(10, 2) DEFAULT 0.00,
    discount DECIMAL(10, 2) DEFAULT 0.00,
    final_amount DECIMAL(10, 2) DEFAULT 0.00,
    expiration_date DATE,
    status ENUM('pending', 'approved', 'rejected', 'expired') DEFAULT 'pending',
    approval_rejection_date TIMESTAMP NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE CASCADE,
    INDEX idx_service_order_id (service_order_id),
    INDEX idx_status (status),
    INDEX idx_budget_number (budget_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 6. Tabela: BUDGET_ITEMS

```sql
CREATE TABLE budget_items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    budget_id BIGINT UNSIGNED NOT NULL,
    type ENUM('part', 'service') NOT NULL,
    description VARCHAR(255) NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    unit_price DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (budget_id) REFERENCES budgets(id) ON DELETE CASCADE,
    INDEX idx_budget_id (budget_id),
    INDEX idx_type (type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 7. Tabela: EXECUTED_SERVICES

```sql
CREATE TABLE executed_services (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_order_id BIGINT UNSIGNED NOT NULL,
    mechanic_name VARCHAR(255),
    start_date TIMESTAMP NULL,
    completion_date TIMESTAMP NULL,
    service_description TEXT NOT NULL,
    estimated_hours DECIMAL(5, 2),
    actual_hours DECIMAL(5, 2),
    status ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE CASCADE,
    INDEX idx_service_order_id (service_order_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 8. Tabela: PAYMENTS

```sql
CREATE TABLE payments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_order_id BIGINT UNSIGNED NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payment_method ENUM(
        'cash',
        'debit_card',
        'credit_card',
        'pix',
        'bank_slip',
        'check'
    ) NOT NULL,
    amount_paid DECIMAL(10, 2) NOT NULL,
    discount_applied DECIMAL(10, 2) DEFAULT 0.00,
    final_amount DECIMAL(10, 2) NOT NULL,
    status ENUM('pending', 'partially_paid', 'fully_paid') DEFAULT 'pending',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE RESTRICT,
    INDEX idx_service_order_id (service_order_id),
    INDEX idx_status (status),
    INDEX idx_payment_method (payment_method)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 10. Script Completo de Criação

**Nota:** O script completo de criação do banco de dados está disponível no arquivo `create_database.sql` na raiz do projeto. Execute esse arquivo para criar todas as tabelas, índices, triggers, views e procedures automaticamente.

O arquivo contém:

- Todas as 7 tabelas do sistema (clients, vehicles, service_orders, budgets, budget_items, executed_services, payments)
- Índices otimizados para consultas
- Triggers automáticos para cálculos
- Views úteis para relatórios
- Procedures para operações comuns
- Dados de teste opcionais

Para executar:

```bash
mysql -u root -p < create_database.sql
```

---

**Nota:** Todos os scripts SQL completos estão disponíveis no arquivo [`create_database.sql`](create_database.sql) na raiz do projeto.

---

## 🛣️ ENDPOINTS DA API

### Base URL

```
http://localhost:8080/api/v1
```

### **CLIENTES**

| Método | Endpoint                | Descrição                     |
| ------ | ----------------------- | ----------------------------- |
| POST   | `/clients`              | Criar novo cliente            |
| GET    | `/clients`              | Listar todos os clientes      |
| GET    | `/clients/:id`          | Buscar cliente por ID         |
| PUT    | `/clients/:id`          | Atualizar cliente             |
| DELETE | `/clients/:id`          | Deletar cliente (soft delete) |
| GET    | `/clients/:id/vehicles` | Listar veículos do cliente    |
| GET    | `/clients/:id/history`  | Histórico de OS do cliente    |

### **VEÍCULOS**

| Método | Endpoint                 | Descrição                |
| ------ | ------------------------ | ------------------------ |
| POST   | `/vehicles`              | Cadastrar veículo        |
| GET    | `/vehicles`              | Listar todos os veículos |
| GET    | `/vehicles/:id`          | Buscar veículo por ID    |
| GET    | `/vehicles/plate/:plate` | Buscar veículo por placa |
| PUT    | `/vehicles/:id`          | Atualizar veículo        |
| DELETE | `/vehicles/:id`          | Deletar veículo          |

### **ORDENS DE SERVIÇO**

| Método | Endpoint                              | Descrição                 |
| ------ | ------------------------------------- | ------------------------- |
| POST   | `/service-orders`                     | Criar nova OS             |
| GET    | `/service-orders`                     | Listar todas as OS        |
| GET    | `/service-orders/:id`                 | Buscar OS por ID          |
| PUT    | `/service-orders/:id`                 | Atualizar OS              |
| PATCH  | `/service-orders/:id/status`          | Atualizar apenas o status |
| GET    | `/service-orders/client/:client_id`   | OS por cliente            |
| GET    | `/service-orders/vehicle/:vehicle_id` | OS por veículo            |

### **ORÇAMENTOS**

| Método | Endpoint               | Descrição                   |
| ------ | ---------------------- | --------------------------- |
| POST   | `/budgets`             | Criar orçamento para OS     |
| GET    | `/budgets/:id`         | Buscar orçamento            |
| PUT    | `/budgets/:id`         | Atualizar orçamento         |
| PATCH  | `/budgets/:id/approve` | Aprovar orçamento           |
| PATCH  | `/budgets/:id/reject`  | Recusar orçamento           |
| POST   | `/budgets/:id/items`   | Adicionar item ao orçamento |

### **PAGAMENTOS**

| Método | Endpoint                                    | Descrição                   |
| ------ | ------------------------------------------- | --------------------------- |
| POST   | `/payments`                                 | Registrar pagamento         |
| GET    | `/payments/service-order/:service_order_id` | Buscar pagamentos de uma OS |
| GET    | `/payments/:id`                             | Buscar pagamento por ID     |

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
- [x] Executar scripts SQL para criar tabelas
- [x] Configurar Air para hot reload

### **Sprint 2 - Módulo Clientes (Semana 2)**

- [x] Criar `models/client.go`
- [x] Criar `dto/client_dto.go`
- [x] Implementar `controllers/client_controller.go`
  - [x] CreateClient
  - [x] GetClients (com paginação)
  - [x] GetClientByID
  - [x] UpdateClient
  - [x] DeleteClient (soft delete)
- [x] Configurar `routes/client_routes.go`
- [x] Testar todos os endpoints no Postman
- [x] Adicionar validações

### **Sprint 3 - Módulo Veículos (Semana 3)**

- [ ] Criar `models/vehicle.go` com relacionamento
- [ ] Criar `dto/vehicle_dto.go`
- [ ] Implementar `controllers/vehicle_controller.go`
  - [ ] CRUD completo
  - [ ] Buscar por placa
  - [ ] Listar veículos por cliente
- [ ] Configurar `routes/vehicle_routes.go`
- [ ] Testar relacionamento Client ↔ Vehicles
- [ ] Validar placa e dados

### **Sprint 4 - Módulo OS (Semanas 4-5)**

- [ ] Criar `models/service_order.go`
- [ ] Criar `dto/service_order_dto.go`
- [ ] Implementar `controllers/service_order_controller.go`
  - [ ] Criar OS
  - [ ] Listar com filtros (status, cliente, veículo)
  - [ ] Atualizar status
  - [ ] Buscar OS completa com relacionamentos
- [ ] Configurar `routes/service_order_routes.go`
- [ ] Implementar lógica de mudança de status
- [ ] Testar fluxo completo

### **Sprint 5 - Módulo Orçamento (Semana 5)**

- [ ] Criar `models/budget.go` e `models/budget_item.go`
- [ ] Criar `dto/budget_dto.go`
- [ ] Implementar `controllers/budget_controller.go`
  - [ ] Criar orçamento
  - [ ] Adicionar itens
  - [ ] Calcular totais automaticamente
  - [ ] Aprovar/Recusar orçamento
- [ ] Atualizar status da OS quando orçamento aprovado
- [ ] Testar cálculos

### **Sprint 6 - Módulo Pagamento (Semana 6)**

- [ ] Criar `models/payment.go`
- [ ] Criar `dto/payment_dto.go`
- [ ] Implementar `controllers/payment_controller.go`
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

// ConnectDatabase - Estabelece conexão com o banco de dados MySQL
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

1. `models/client.go`
2. `dto/client_dto.go`
3. `controllers/client_controller.go`
4. `routes/client_routes.go`
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

### **Padrões de Código:**

1. **Todo código em inglês** (variáveis, funções, rotas, tabelas)
2. **Comentários em português** (para facilitar compreensão)
3. **Nomes de arquivos em snake_case**: `client_controller.go`
4. **Structs em PascalCase**: `CreateClientRequest`
5. **Variáveis em camelCase**: `clientID`, `primaryPhone`
6. **URLs em kebab-case**: `/service-orders`, `/budget-items`

### **Boas Práticas:**

1. **Sempre usar soft delete** para clients e vehicles
2. **Validar CPF** antes de cadastrar cliente
3. **Validar placa** (formato ABC-1234 ou ABC1D234)
4. **Gerar order_number automaticamente** (ex: OS-2024-0001)
5. **Calcular totais automaticamente** nos budgets
6. **Validar datas** (completion_date > entry_date)
7. **Impedir exclusão** de clients/vehicles com service orders ativas
8. **Logs estruturados** para rastreabilidade
9. **DTOs diferentes** para Request e Response
10. **Status em inglês** mas documentação em português

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
