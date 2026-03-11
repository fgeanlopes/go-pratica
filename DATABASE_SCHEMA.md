# 🗄️ Database Schema Map - Auto Shop Management System

## 📋 Overview

This document describes the complete structure of the `mecanica_db` database with all tables and relationships required for the system.

**Database:** `mecanica_db`  
**Charset:** `utf8mb4`  
**Collate:** `utf8mb4_unicode_ci`  
**Engine:** `InnoDB`

---

## 🔗 Relationship Diagram

```
┌──────────────────┐
│     CLIENTS      │
│    (clients)     │
└──────┬───────────┘
       │
       ├────────────────┐
       │                │
       ↓                ↓
┌──────────────┐  ┌──────────────┐
│   VEHICLES   │  │              │
│  (vehicles)  │  │              │
└──────┬───────┘  │              │
       │          │              │
       ├──────────┘              │
       │                         │
       ↓                         ↓
┌────────────────────────────────────┐
│       SERVICE ORDERS               │
│      (service_orders)              │
└───────┬───────────┬────────┬───────┘
        │           │        │
        ↓           ↓        ↓
  ┌──────────┐ ┌─────────┐ ┌──────────┐
  │ BUDGETS  │ │EXECUTED │ │ PAYMENTS │
  │(budgets) │ │SERVICES │ │(payments)│
  └────┬─────┘ │(executed│ └──────────┘
       │       │_services│
       ↓       └─────────┘
  ┌──────────────┐
  │BUDGET ITEMS  │
  │(budget_items)│
  └──────────────┘
```

---

## 📊 Detailed Tables

### 1️⃣ CLIENTS (`clients`)

**Description:** Client registration

| Field             | Type            | Constraints                 | Description                      |
| ----------------- | --------------- | --------------------------- | -------------------------------- |
| `id`              | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier                |
| `name`            | VARCHAR(255)    | NOT NULL                    | Client full name                 |
| `cpf`             | VARCHAR(14)     | NOT NULL, UNIQUE            | CPF (format: 000.000.000-00)     |
| `primary_phone`   | VARCHAR(20)     | NOT NULL                    | Primary phone                    |
| `secondary_phone` | VARCHAR(20)     | NULL                        | Secondary phone                  |
| `email`           | VARCHAR(255)    | NULL                        | Email address                    |
| `status`          | ENUM            | DEFAULT 'active'            | Status (active, inactive)        |
| `zip_code`        | VARCHAR(10)     | NULL                        | ZIP code (format: 00000-000)     |
| `street`          | VARCHAR(255)    | NULL                        | Street name                      |
| `number`          | VARCHAR(20)     | NULL                        | Address number                   |
| `complement`      | VARCHAR(255)    | NULL                        | Address complement               |
| `neighborhood`    | VARCHAR(100)    | NULL                        | Neighborhood                     |
| `city`            | VARCHAR(100)    | NULL                        | City                             |
| `state`           | CHAR(2)         | NULL                        | State abbreviation (SP, RJ, etc) |
| `created_at`      | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date                    |
| `updated_at`      | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date                 |
| `deleted_at`      | TIMESTAMP       | NULL                        | Soft delete                      |

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_cpf` (cpf)
- INDEX: `idx_status` (status)
- INDEX: `idx_deleted_at` (deleted_at)

**Relationships:**

- ➡️ `1:N` with `vehicles` (one client can have multiple vehicles)
- ➡️ `1:N` with `service_orders` (one client can have multiple service orders)

---

### 2️⃣ VEHICLES (`vehicles`)

**Description:** Registered vehicles (one client can have multiple vehicles)

| Field              | Type            | Constraints                 | Description                    |
| ------------------ | --------------- | --------------------------- | ------------------------------ |
| `id`               | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier              |
| `client_id`        | BIGINT UNSIGNED | NOT NULL, FK                | Client reference               |
| `plate`            | VARCHAR(10)     | NOT NULL, UNIQUE            | Plate (ABC-1234 or ABC1D23)    |
| `brand`            | VARCHAR(100)    | NOT NULL                    | Brand (Fiat, Volkswagen, Ford) |
| `model`            | VARCHAR(100)    | NOT NULL                    | Model (Uno, Gol, Fiesta)       |
| `manufacture_year` | INT             | NOT NULL                    | Manufacture year               |
| `model_year`       | INT             | NOT NULL                    | Model year                     |
| `color`            | VARCHAR(50)     | NULL                        | Vehicle color                  |
| `chassis`          | VARCHAR(50)     | NULL                        | Chassis number                 |
| `current_mileage`  | INT             | DEFAULT 0                   | Current mileage                |
| `created_at`       | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date                  |
| `updated_at`       | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date               |
| `deleted_at`       | TIMESTAMP       | NULL                        | Soft delete                    |

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_client_id` (client_id)
- INDEX: `idx_plate` (plate)
- INDEX: `idx_brand_model` (brand, model)
- INDEX: `idx_deleted_at` (deleted_at)

**Relationships:**

- ⬅️ `N:1` with `clients` (CASCADE ON DELETE)
- ➡️ `1:N` with `service_orders` (one vehicle can have multiple service orders)

---

### 3️⃣ SERVICE ORDERS (`service_orders`)

**Description:** Service orders - vehicle entry control

| Field                      | Type            | Constraints                 | Description                     |
| -------------------------- | --------------- | --------------------------- | ------------------------------- |
| `id`                       | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier               |
| `order_number`             | VARCHAR(20)     | NOT NULL, UNIQUE            | Order number (Ex: SO-2024-0001) |
| `client_id`                | BIGINT UNSIGNED | NOT NULL, FK                | Client reference                |
| `vehicle_id`               | BIGINT UNSIGNED | NOT NULL, FK                | Vehicle reference               |
| `entry_date`               | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Vehicle entry date              |
| `entry_mileage`            | INT             | NULL                        | Vehicle mileage at entry        |
| `problem_description`      | TEXT            | NOT NULL                    | Problem reported by client      |
| `status`                   | ENUM            | DEFAULT 'awaiting_budget'   | Service order status            |
| `expected_completion_date` | DATE            | NULL                        | Expected completion date        |
| `actual_completion_date`   | TIMESTAMP       | NULL                        | Actual completion date          |
| `notes`                    | TEXT            | NULL                        | General notes                   |
| `created_at`               | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date                   |
| `updated_at`               | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date                |

**Possible Status Values:**

- `awaiting_budget`
- `budget_sent`
- `approved`
- `in_progress`
- `completed`
- `delivered`
- `cancelled`

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_order_number` (order_number)
- INDEX: `idx_client_id` (client_id)
- INDEX: `idx_vehicle_id` (vehicle_id)
- INDEX: `idx_status` (status)
- INDEX: `idx_entry_date` (entry_date)

**Relationships:**

- ⬅️ `N:1` with `clients` (RESTRICT ON DELETE)
- ⬅️ `N:1` with `vehicles` (RESTRICT ON DELETE)
- ➡️ `1:N` with `budgets` (one service order can have multiple budgets)
- ➡️ `1:N` with `executed_services` (one service order can have multiple services)
- ➡️ `1:N` with `payments` (one service order can have multiple payments)

---

### 4️⃣ BUDGETS (`budgets`)

**Description:** Budgets generated for each service order

| Field                     | Type            | Constraints                 | Description                       |
| ------------------------- | --------------- | --------------------------- | --------------------------------- |
| `id`                      | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier                 |
| `service_order_id`        | BIGINT UNSIGNED | NOT NULL, FK                | Service order reference           |
| `budget_number`           | VARCHAR(20)     | NOT NULL, UNIQUE            | Budget number (Ex: BUD-2024-0001) |
| `creation_date`           | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date                     |
| `parts_amount`            | DECIMAL(10,2)   | DEFAULT 0.00                | Sum of parts values               |
| `labor_amount`            | DECIMAL(10,2)   | DEFAULT 0.00                | Sum of service values             |
| `total_amount`            | DECIMAL(10,2)   | DEFAULT 0.00                | parts_amount + labor_amount       |
| `discount`                | DECIMAL(10,2)   | DEFAULT 0.00                | Applied discount                  |
| `final_amount`            | DECIMAL(10,2)   | DEFAULT 0.00                | total_amount - discount           |
| `expiration_date`         | DATE            | NULL                        | Budget expiration date            |
| `status`                  | ENUM            | DEFAULT 'pending'           | Budget status                     |
| `approval_rejection_date` | TIMESTAMP       | NULL                        | Approval/rejection date           |
| `notes`                   | TEXT            | NULL                        | Notes                             |
| `created_at`              | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date                     |
| `updated_at`              | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date                  |

**Possible Status Values:**

- `pending`
- `approved`
- `rejected`
- `expired`

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_service_order_id` (service_order_id)
- INDEX: `idx_status` (status)
- INDEX: `idx_budget_number` (budget_number)
- INDEX: `idx_expiration_date` (expiration_date)

**Relationships:**

- ⬅️ `N:1` with `service_orders` (CASCADE ON DELETE)
- ➡️ `1:N` with `budget_items` (one budget has multiple items)

---

### 5️⃣ BUDGET ITEMS (`budget_items`)

**Description:** Detailed items of each budget (parts and services)

| Field         | Type            | Constraints                 | Description                 |
| ------------- | --------------- | --------------------------- | --------------------------- |
| `id`          | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier           |
| `budget_id`   | BIGINT UNSIGNED | NOT NULL, FK                | Budget reference            |
| `type`        | ENUM            | NOT NULL                    | Item type (part, service)   |
| `description` | VARCHAR(255)    | NOT NULL                    | Part or service description |
| `quantity`    | INT             | NOT NULL, DEFAULT 1         | Quantity                    |
| `unit_price`  | DECIMAL(10,2)   | NOT NULL                    | Unit price                  |
| `total_price` | DECIMAL(10,2)   | NOT NULL                    | quantity \* unit_price      |
| `notes`       | TEXT            | NULL                        | Notes                       |
| `created_at`  | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date               |
| `updated_at`  | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date            |

**Possible Type Values:**

- `part` - Parts/components
- `service` - Labor/services

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_budget_id` (budget_id)
- INDEX: `idx_type` (type)

**Relationships:**

- ⬅️ `N:1` with `budgets` (CASCADE ON DELETE)

---

### 6️⃣ EXECUTED SERVICES (`executed_services`)

**Description:** Record of services executed for each service order

| Field                 | Type            | Constraints                 | Description             |
| --------------------- | --------------- | --------------------------- | ----------------------- |
| `id`                  | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier       |
| `service_order_id`    | BIGINT UNSIGNED | NOT NULL, FK                | Service order reference |
| `mechanic_name`       | VARCHAR(255)    | NULL                        | Mechanic name           |
| `start_date`          | TIMESTAMP       | NULL                        | Service start date      |
| `completion_date`     | TIMESTAMP       | NULL                        | Service completion date |
| `service_description` | TEXT            | NOT NULL                    | Service details         |
| `estimated_hours`     | DECIMAL(5,2)    | NULL                        | Estimated time in hours |
| `actual_hours`        | DECIMAL(5,2)    | NULL                        | Actual time spent       |
| `status`              | ENUM            | DEFAULT 'pending'           | Service status          |
| `created_at`          | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date           |
| `updated_at`          | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date        |

**Possible Status Values:**

- `pending`
- `in_progress`
- `completed`

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_service_order_id` (service_order_id)
- INDEX: `idx_status` (status)
- INDEX: `idx_mechanic` (mechanic_name)

**Relationships:**

- ⬅️ `N:1` with `service_orders` (CASCADE ON DELETE)

---

### 7️⃣ PAYMENTS (`payments`)

**Description:** Service order payment records

| Field              | Type            | Constraints                 | Description                    |
| ------------------ | --------------- | --------------------------- | ------------------------------ |
| `id`               | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Unique identifier              |
| `service_order_id` | BIGINT UNSIGNED | NOT NULL, FK                | Service order reference        |
| `payment_date`     | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Payment date                   |
| `payment_method`   | ENUM            | NOT NULL                    | Payment method                 |
| `amount_paid`      | DECIMAL(10,2)   | NOT NULL                    | Amount paid                    |
| `discount_applied` | DECIMAL(10,2)   | DEFAULT 0.00                | Applied discount               |
| `final_amount`     | DECIMAL(10,2)   | NOT NULL                    | amount_paid - discount_applied |
| `status`           | ENUM            | DEFAULT 'pending'           | Payment status                 |
| `notes`            | TEXT            | NULL                        | Notes                          |
| `created_at`       | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Creation date                  |
| `updated_at`       | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Last update date               |

**Payment Methods:**

- `cash`
- `debit_card`
- `credit_card`
- `pix`
- `bank_slip`
- `check`

**Possible Status Values:**

- `pending`
- `partially_paid`
- `fully_paid`

**Indexes:**

- PRIMARY KEY: `id`
- INDEX: `idx_service_order_id` (service_order_id)
- INDEX: `idx_status` (status)
- INDEX: `idx_payment_method` (payment_method)
- INDEX: `idx_payment_date` (payment_date)

**Relationships:**

- ⬅️ `N:1` with `service_orders` (RESTRICT ON DELETE)

---

## 🔐 Foreign Keys Summary

### Table `vehicles`

```sql
FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
```

### Table `service_orders`

```sql
FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE RESTRICT
FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE RESTRICT
```

### Table `budgets`

```sql
FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE CASCADE
```

### Table `budget_items`

```sql
FOREIGN KEY (budget_id) REFERENCES budgets(id) ON DELETE CASCADE
```

### Table `executed_services`

```sql
FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE CASCADE
```

### Table `payments`

```sql
FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE RESTRICT
```

---

## ⚙️ Automatic Triggers

### 1. Automatic calculation of item total price

- **`trg_calculate_item_price_insert`**: Calculates `total_price = quantity * unit_price` on insert
- **`trg_calculate_item_price_update`**: Calculates `total_price = quantity * unit_price` on update

### 2. Automatic update of budget totals

- **`trg_update_budget_after_item_insert`**: After inserting an item, updates:
  - `parts_amount` (sum of items with type 'part')
  - `labor_amount` (sum of items with type 'service')
  - `total_amount` (parts_amount + labor_amount)
  - `final_amount` (total_amount - discount)

---

## 📊 Useful Views

### 1. `vw_complete_service_orders`

Displays service orders with complete client and vehicle data.

```sql
SELECT * FROM vw_complete_service_orders;
```

**Returned Fields:**

- id, order_number, entry_date, status
- client_name, client_cpf, client_phone
- vehicle_plate, vehicle_brand, vehicle_model
- problem_description, expected_completion_date

### 2. `vw_financial_summary`

Financial summary per service order.

```sql
SELECT * FROM vw_financial_summary;
```

**Returned Fields:**

- service_order_id, order_number
- budget_amount, amount_paid, pending_balance
- payment_status (paid/partial/pending)

---

## 🔧 Useful Procedures

### 1. `sp_client_history`

Retrieves complete client history.

```sql
CALL sp_client_history(1); -- Client ID
```

### 2. `sp_pending_budget_orders`

Lists service orders awaiting budget.

```sql
CALL sp_pending_budget_orders();
```

---

## 📝 Table Creation Order

To create the database manually, follow this order (respect dependencies):

1. ✅ `clients` (no dependencies)
2. ✅ `vehicles` (depends on `clients`)
3. ✅ `service_orders` (depends on `clients` and `vehicles`)
4. ✅ `budgets` (depends on `service_orders`)
5. ✅ `budget_items` (depends on `budgets`)
6. ✅ `executed_services` (depends on `service_orders`)
7. ✅ `payments` (depends on `service_orders`)

---

## 🎯 Important Business Rules

1. **Soft Delete**: Tables `clients` and `vehicles` use soft delete (`deleted_at` field)
2. **Cascade Delete**:
   - Delete client → deletes vehicles
   - Delete service order → deletes budgets, executed services
3. **Restrict Delete**:
   - Cannot delete client or vehicle with active service orders
   - Cannot delete service order with registered payments
4. **Automatic Calculations**:
   - Budget values are automatically calculated via triggers
5. **Validations**:
   - CPF must be unique
   - Vehicle plate must be unique
   - Service order number must be unique
   - Budget number must be unique

---

## 🚀 How to Use This Document

1. **To create the database**: Execute the `create_database.sql` file
2. **To understand the structure**: Use this document as reference
3. **To make queries**: Use the views and procedures already created
4. **To modify**: Always respect relationships and constraints

---

**Author:** Auto Shop Management System  
**Last Update:** March 2026  
**Schema Version:** 1.0### 6️⃣ ITENS DO ORÇAMENTO (`itens_orcamento`)

**Descrição:** Itens detalhados de cada orçamento (peças e serviços)

| Campo            | Tipo            | Restrições                  | Descrição                    |
| ---------------- | --------------- | --------------------------- | ---------------------------- |
| `id`             | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Identificador único          |
| `orcamento_id`   | BIGINT UNSIGNED | NOT NULL, FK                | Referência ao orçamento      |
| `tipo`           | ENUM            | NOT NULL                    | Tipo do item (peca, servico) |
| `descricao`      | VARCHAR(255)    | NOT NULL                    | Descrição da peça ou serviço |
| `quantidade`     | INT             | NOT NULL, DEFAULT 1         | Quantidade                   |
| `valor_unitario` | DECIMAL(10,2)   | NOT NULL                    | Valor unitário               |
| `valor_total`    | DECIMAL(10,2)   | NOT NULL                    | quantidade \* valor_unitario |
| `observacao`     | TEXT            | NULL                        | Observações                  |
| `created_at`     | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Data de criação              |
| `updated_at`     | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Data de atualização          |

**Tipos possíveis:**

- `peca` - Peças/componentes
- `servico` - Mão de obra/serviços

**Índices:**

- PRIMARY KEY: `id`
- INDEX: `idx_orcamento_id` (orcamento_id)
- INDEX: `idx_tipo` (tipo)

**Relacionamentos:**

- ⬅️ `N:1` com `orcamentos` (CASCADE ON DELETE)

---

### 7️⃣ SERVIÇOS EXECUTADOS (`servicos_executados`)

**Descrição:** Registro dos serviços executados em cada OS

| Campo                  | Tipo            | Restrições                  | Descrição                         |
| ---------------------- | --------------- | --------------------------- | --------------------------------- |
| `id`                   | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Identificador único               |
| `os_id`                | BIGINT UNSIGNED | NOT NULL, FK                | Referência à OS                   |
| `mecanico_responsavel` | VARCHAR(255)    | NULL                        | Nome do mecânico                  |
| `data_inicio`          | TIMESTAMP       | NULL                        | Data de início do serviço         |
| `data_conclusao`       | TIMESTAMP       | NULL                        | Data de conclusão do serviço      |
| `descricao_servico`    | TEXT            | NOT NULL                    | Detalhamento do serviço realizado |
| `tempo_estimado_horas` | DECIMAL(5,2)    | NULL                        | Tempo estimado em horas           |
| `tempo_real_horas`     | DECIMAL(5,2)    | NULL                        | Tempo real gasto                  |
| `status`               | ENUM            | DEFAULT 'pendente'          | Status do serviço                 |
| `created_at`           | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Data de criação                   |
| `updated_at`           | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Data de atualização               |

**Status possíveis:**

- `pendente`
- `em_andamento`
- `concluido`

**Índices:**

- PRIMARY KEY: `id`
- INDEX: `idx_os_id` (os_id)
- INDEX: `idx_status` (status)
- INDEX: `idx_mecanico` (mecanico_responsavel)

**Relacionamentos:**

- ⬅️ `N:1` com `ordens_servico` (CASCADE ON DELETE)

---

### 8️⃣ PAGAMENTOS (`pagamentos`)

**Descrição:** Registro de pagamentos das OS

| Campo               | Tipo            | Restrições                  | Descrição                      |
| ------------------- | --------------- | --------------------------- | ------------------------------ |
| `id`                | BIGINT UNSIGNED | PK, AUTO_INCREMENT          | Identificador único            |
| `os_id`             | BIGINT UNSIGNED | NOT NULL, FK                | Referência à OS                |
| `data_pagamento`    | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Data do pagamento              |
| `forma_pagamento`   | ENUM            | NOT NULL                    | Forma de pagamento             |
| `valor_pago`        | DECIMAL(10,2)   | NOT NULL                    | Valor pago                     |
| `desconto_aplicado` | DECIMAL(10,2)   | DEFAULT 0.00                | Desconto aplicado              |
| `valor_final`       | DECIMAL(10,2)   | NOT NULL                    | valor_pago - desconto_aplicado |
| `status`            | ENUM            | DEFAULT 'pendente'          | Status do pagamento            |
| `observacoes`       | TEXT            | NULL                        | Observações                    |
| `created_at`        | TIMESTAMP       | DEFAULT CURRENT_TIMESTAMP   | Data de criação                |
| `updated_at`        | TIMESTAMP       | ON UPDATE CURRENT_TIMESTAMP | Data de atualização            |

**Formas de pagamento:**

- `dinheiro`
- `cartao_debito`
- `cartao_credito`
- `pix`
- `boleto`
- `cheque`

**Status possíveis:**

- `pendente`
- `pago_parcial`
- `pago_total`

**Índices:**

- PRIMARY KEY: `id`
- INDEX: `idx_os_id` (os_id)
- INDEX: `idx_status` (status)
- INDEX: `idx_forma_pagamento` (forma_pagamento)
- INDEX: `idx_data_pagamento` (data_pagamento)

**Relacionamentos:**

- ⬅️ `N:1` com `ordens_servico` (RESTRICT ON DELETE)

---

## 🔐 Resumo dos Relacionamentos (Foreign Keys)

### Tabela `enderecos`

```sql
FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE
```

### Tabela `veiculos`

```sql
FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE
```

### Tabela `ordens_servico`

```sql
FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE RESTRICT
FOREIGN KEY (veiculo_id) REFERENCES veiculos(id) ON DELETE RESTRICT
```

### Tabela `orcamentos`

```sql
FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE
```

### Tabela `itens_orcamento`

```sql
FOREIGN KEY (orcamento_id) REFERENCES orcamentos(id) ON DELETE CASCADE
```

### Tabela `servicos_executados`

```sql
FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE
```

### Tabela `pagamentos`

```sql
FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE RESTRICT
```

---

## ⚙️ Triggers Automáticos

### 1. Cálculo automático do valor total dos itens

- **`trg_calcula_valor_item_insert`**: Calcula `valor_total = quantidade * valor_unitario` ao inserir item
- **`trg_calcula_valor_item_update`**: Calcula `valor_total = quantidade * valor_unitario` ao atualizar item

### 2. Atualização automática dos totais do orçamento

- **`trg_atualiza_orcamento_after_item_insert`**: Após inserir um item, atualiza:
  - `valor_pecas` (soma dos itens do tipo 'peca')
  - `valor_mao_obra` (soma dos itens do tipo 'servico')
  - `valor_total` (valor_pecas + valor_mao_obra)
  - `valor_final` (valor_total - desconto)

---

## 📊 Views Úteis

### 1. `vw_os_completas`

Exibe ordens de serviço com dados completos do cliente e veículo.

```sql
SELECT * FROM vw_os_completas;
```

**Campos retornados:**

- id, numero_os, data_entrada, status
- cliente_nome, cliente_cpf, cliente_telefone
- veiculo_placa, veiculo_marca, veiculo_modelo
- descricao_problema, data_prevista_conclusao

### 2. `vw_resumo_financeiro_os`

Resumo financeiro por ordem de serviço.

```sql
SELECT * FROM vw_resumo_financeiro_os;
```

**Campos retornados:**

- os_id, numero_os
- valor_orcamento, valor_pago, saldo_pendente
- situacao_pagamento (pago/parcial/pendente)

---

## 🔧 Procedures Úteis

### 1. `sp_historico_cliente`

Busca o histórico completo de um cliente.

```sql
CALL sp_historico_cliente(1); -- ID do cliente
```

### 2. `sp_os_pendentes_orcamento`

Lista ordens de serviço aguardando orçamento.

```sql
CALL sp_os_pendentes_orcamento();
```

---

## 📝 Ordem de Criação das Tabelas

Para criar o banco manualmente, siga esta ordem (respeite as dependências):

1. ✅ `clientes` (sem dependências)
2. ✅ `enderecos` (depende de `clientes`)
3. ✅ `veiculos` (depende de `clientes`)
4. ✅ `ordens_servico` (depende de `clientes` e `veiculos`)
5. ✅ `orcamentos` (depende de `ordens_servico`)
6. ✅ `itens_orcamento` (depende de `orcamentos`)
7. ✅ `servicos_executados` (depende de `ordens_servico`)
8. ✅ `pagamentos` (depende de `ordens_servico`)

---

## 🎯 Regras de Negócio Importantes

1. **Soft Delete**: Tabelas `clientes` e `veiculos` usam soft delete (campo `deleted_at`)
2. **Cascade Delete**:
   - Deletar cliente → deleta endereços e veículos
   - Deletar OS → deleta orçamentos, serviços executados
3. **Restrict Delete**:
   - Não pode deletar cliente ou veículo com OS ativa
   - Não pode deletar OS com pagamentos registrados
4. **Cálculos Automáticos**:
   - Valores dos orçamentos são calculados automaticamente via triggers
5. **Validações**:
   - CPF deve ser único
   - Placa do veículo deve ser única
   - Número da OS deve ser único
   - Número do orçamento deve ser único

---

## 🚀 Como Usar Este Documento

1. **Para criar o banco**: Execute o arquivo `create_database.sql`
2. **Para entender a estrutura**: Use este documento como referência
3. **Para fazer consultas**: Utilize as views e procedures já criadas
4. **Para modificar**: Sempre respeite os relacionamentos e constraints

---

**Autor:** Sistema de Mecânica  
**Última Atualização:** Março 2026  
**Versão do Schema:** 1.0
