-- =====================================================
-- SISTEMA DE GERENCIAMENTO DE MECÂNICA
-- Script de Criação do Banco de Dados MySQL
-- Versão: 1.0
-- Data: Março 2026
-- =====================================================

-- Remover banco se já existir (cuidado em produção!)
-- DROP DATABASE IF EXISTS mecanica_db;

-- Criar banco de dados
CREATE DATABASE IF NOT EXISTS mecanica_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

-- Usar o banco
USE mecanica_db;

-- =====================================================
-- TABLE: CLIENTS
-- =====================================================
CREATE TABLE clients (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE COMMENT 'Format: 000.000.000-00',
    primary_phone VARCHAR(20) NOT NULL,
    secondary_phone VARCHAR(20),
    email VARCHAR(255),
    status ENUM('active', 'inactive') DEFAULT 'active',
    zip_code VARCHAR(10) COMMENT 'Format: 00000-000',
    street VARCHAR(255),
    number VARCHAR(20),
    complement VARCHAR(255),
    neighborhood VARCHAR(100),
    city VARCHAR(100),
    state CHAR(2) COMMENT 'State abbreviation: SP, RJ, etc',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL COMMENT 'Soft delete',
    
    INDEX idx_cpf (cpf),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Client registration';

-- =====================================================
-- TABLE: VEHICLES
-- =====================================================
CREATE TABLE vehicles (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    client_id BIGINT UNSIGNED NOT NULL,
    plate VARCHAR(10) NOT NULL UNIQUE COMMENT 'Formats: ABC-1234 or ABC1D23',
    brand VARCHAR(100) NOT NULL COMMENT 'Ex: Fiat, Volkswagen, Ford',
    model VARCHAR(100) NOT NULL COMMENT 'Ex: Uno, Gol, Fiesta',
    manufacture_year INT NOT NULL,
    model_year INT NOT NULL,
    color VARCHAR(50),
    chassis VARCHAR(50) COMMENT 'Chassis number',
    current_mileage INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL COMMENT 'Soft delete',
    
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    INDEX idx_client_id (client_id),
    INDEX idx_plate (plate),
    INDEX idx_brand_model (brand, model),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Registered vehicles (one client can have multiple vehicles)';

-- =====================================================
-- TABLE: SERVICE_ORDERS
-- =====================================================
CREATE TABLE service_orders (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_number VARCHAR(20) NOT NULL UNIQUE COMMENT 'Ex: SO-2024-0001',
    client_id BIGINT UNSIGNED NOT NULL,
    vehicle_id BIGINT UNSIGNED NOT NULL,
    entry_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    entry_mileage INT COMMENT 'Vehicle mileage at entry',
    problem_description TEXT NOT NULL COMMENT 'Problem reported by client',
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
    notes TEXT COMMENT 'General service order notes',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE RESTRICT,
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE RESTRICT,
    INDEX idx_order_number (order_number),
    INDEX idx_client_id (client_id),
    INDEX idx_vehicle_id (vehicle_id),
    INDEX idx_status (status),
    INDEX idx_entry_date (entry_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Service orders - vehicle entry control';

-- =====================================================
-- TABLE: BUDGETS
-- =====================================================
CREATE TABLE budgets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_order_id BIGINT UNSIGNED NOT NULL,
    budget_number VARCHAR(20) NOT NULL UNIQUE COMMENT 'Ex: BUD-2024-0001',
    creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parts_amount DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'Sum of parts values',
    labor_amount DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'Sum of service values',
    total_amount DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'parts_amount + labor_amount',
    discount DECIMAL(10, 2) DEFAULT 0.00,
    final_amount DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'total_amount - discount',
    expiration_date DATE COMMENT 'Budget expiration date',
    status ENUM('pending', 'approved', 'rejected', 'expired') DEFAULT 'pending',
    approval_rejection_date TIMESTAMP NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE CASCADE,
    INDEX idx_service_order_id (service_order_id),
    INDEX idx_status (status),
    INDEX idx_budget_number (budget_number),
    INDEX idx_expiration_date (expiration_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Budgets generated for each service order';

-- =====================================================
-- TABLE: BUDGET_ITEMS
-- =====================================================
CREATE TABLE budget_items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    budget_id BIGINT UNSIGNED NOT NULL,
    type ENUM('part', 'service') NOT NULL COMMENT 'Item type',
    description VARCHAR(255) NOT NULL COMMENT 'Part or service description',
    quantity INT NOT NULL DEFAULT 1,
    unit_price DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL COMMENT 'quantity * unit_price',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (budget_id) REFERENCES budgets(id) ON DELETE CASCADE,
    INDEX idx_budget_id (budget_id),
    INDEX idx_type (type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Detailed items of each budget (parts and services)';

-- =====================================================
-- TABLE: EXECUTED_SERVICES
-- =====================================================
CREATE TABLE executed_services (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_order_id BIGINT UNSIGNED NOT NULL,
    mechanic_name VARCHAR(255) COMMENT 'Mechanic name',
    start_date TIMESTAMP NULL,
    completion_date TIMESTAMP NULL,
    service_description TEXT NOT NULL COMMENT 'Service details',
    estimated_hours DECIMAL(5, 2) COMMENT 'Estimated time in hours',
    actual_hours DECIMAL(5, 2) COMMENT 'Actual time spent',
    status ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE CASCADE,
    INDEX idx_service_order_id (service_order_id),
    INDEX idx_status (status),
    INDEX idx_mechanic (mechanic_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Record of services executed for each service order';

-- =====================================================
-- TABLE: PAYMENTS
-- =====================================================
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
    final_amount DECIMAL(10, 2) NOT NULL COMMENT 'amount_paid - discount_applied',
    status ENUM('pending', 'partially_paid', 'fully_paid') DEFAULT 'pending',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (service_order_id) REFERENCES service_orders(id) ON DELETE RESTRICT,
    INDEX idx_service_order_id (service_order_id),
    INDEX idx_status (status),
    INDEX idx_payment_method (payment_method),
    INDEX idx_payment_date (payment_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Service order payment records';

-- =====================================================
-- TEST DATA (OPTIONAL)
-- =====================================================

-- Insert test clients
INSERT INTO clients (name, cpf, primary_phone, email, status) VALUES
('João da Silva', '123.456.789-00', '(11) 98765-4321', 'joao.silva@email.com', 'active'),
('Maria Oliveira', '987.654.321-00', '(11) 91234-5678', 'maria.oliveira@email.com', 'active');

-- Insert test vehicles
INSERT INTO vehicles (client_id, plate, brand, model, manufacture_year, model_year, color, current_mileage) VALUES
(1, 'ABC-1234', 'Fiat', 'Uno', 2020, 2020, 'White', 50000),
(1, 'DEF-5678', 'Volkswagen', 'Gol', 2019, 2019, 'Silver', 80000),
(2, 'GHI-9012', 'Chevrolet', 'Onix', 2021, 2022, 'Black', 30000);

-- Insert test service orders
INSERT INTO service_orders (order_number, client_id, vehicle_id, entry_mileage, problem_description, status) VALUES
('SO-2024-0001', 1, 1, 50000, 'Strange engine noise and dashboard light on', 'awaiting_budget'),
('SO-2024-0002', 2, 3, 30000, 'Oil change and 30k km service', 'approved');

-- =====================================================
-- USEFUL VIEWS
-- =====================================================

-- View: Service orders with complete client and vehicle data
CREATE OR REPLACE VIEW vw_complete_service_orders AS
SELECT 
    so.id,
    so.order_number,
    so.entry_date,
    so.status,
    c.name AS client_name,
    c.cpf AS client_cpf,
    c.primary_phone AS client_phone,
    v.plate AS vehicle_plate,
    v.brand AS vehicle_brand,
    v.model AS vehicle_model,
    so.problem_description,
    so.expected_completion_date
FROM service_orders so
INNER JOIN clients c ON so.client_id = c.id
INNER JOIN vehicles v ON so.vehicle_id = v.id
WHERE so.deleted_at IS NULL;

-- View: Financial summary per service order
CREATE OR REPLACE VIEW vw_financial_summary AS
SELECT 
    so.id AS service_order_id,
    so.order_number,
    b.final_amount AS budget_amount,
    COALESCE(SUM(p.final_amount), 0) AS amount_paid,
    (b.final_amount - COALESCE(SUM(p.final_amount), 0)) AS pending_balance,
    CASE 
        WHEN COALESCE(SUM(p.final_amount), 0) >= b.final_amount THEN 'paid'
        WHEN COALESCE(SUM(p.final_amount), 0) > 0 THEN 'partial'
        ELSE 'pending'
    END AS payment_status
FROM service_orders so
LEFT JOIN budgets b ON so.id = b.service_order_id AND b.status = 'approved'
LEFT JOIN payments p ON so.id = p.service_order_id
GROUP BY so.id, so.order_number, b.final_amount;

-- =====================================================
-- TRIGGERS
-- =====================================================

-- Trigger: Calculate total_price of item on insert
DELIMITER $$
CREATE TRIGGER trg_calculate_item_price_insert
BEFORE INSERT ON budget_items
FOR EACH ROW
BEGIN
    SET NEW.total_price = NEW.quantity * NEW.unit_price;
END$$
DELIMITER ;

-- Trigger: Calculate total_price of item on update
DELIMITER $$
CREATE TRIGGER trg_calculate_item_price_update
BEFORE UPDATE ON budget_items
FOR EACH ROW
BEGIN
    SET NEW.total_price = NEW.quantity * NEW.unit_price;
END$$
DELIMITER ;

-- Trigger: Update budget totals when adding item
DELIMITER $$
CREATE TRIGGER trg_update_budget_after_item_insert
AFTER INSERT ON budget_items
FOR EACH ROW
BEGIN
    UPDATE budgets
    SET 
        parts_amount = (
            SELECT COALESCE(SUM(total_price), 0) 
            FROM budget_items 
            WHERE budget_id = NEW.budget_id AND type = 'part'
        ),
        labor_amount = (
            SELECT COALESCE(SUM(total_price), 0) 
            FROM budget_items 
            WHERE budget_id = NEW.budget_id AND type = 'service'
        )
    WHERE id = NEW.budget_id;
    
    UPDATE budgets
    SET 
        total_amount = parts_amount + labor_amount,
        final_amount = parts_amount + labor_amount - discount
    WHERE id = NEW.budget_id;
END$$
DELIMITER ;

-- =====================================================
-- USEFUL PROCEDURES
-- =====================================================

-- Procedure: Get client history
DELIMITER $$
CREATE PROCEDURE sp_client_history(IN p_client_id BIGINT)
BEGIN
    SELECT 
        so.order_number,
        so.entry_date,
        so.status,
        v.plate,
        v.brand,
        v.model,
        so.problem_description,
        b.final_amount AS budget_amount
    FROM service_orders so
    INNER JOIN vehicles v ON so.vehicle_id = v.id
    LEFT JOIN budgets b ON so.id = b.service_order_id AND b.status = 'approved'
    WHERE so.client_id = p_client_id
    ORDER BY so.entry_date DESC;
END$$
DELIMITER ;

-- Procedure: Get service orders awaiting budget
DELIMITER $$
CREATE PROCEDURE sp_pending_budget_orders()
BEGIN
    SELECT 
        so.order_number,
        c.name AS client_name,
        v.plate,
        v.model,
        so.problem_description,
        so.entry_date,
        DATEDIFF(CURRENT_DATE, DATE(so.entry_date)) AS days_waiting
    FROM service_orders so
    INNER JOIN clients c ON so.client_id = c.id
    INNER JOIN vehicles v ON so.vehicle_id = v.id
    WHERE so.status = 'awaiting_budget'
    ORDER BY so.entry_date ASC;
END$$
DELIMITER ;

-- =====================================================
-- SUCCESS MESSAGE
-- =====================================================
SELECT 
    '✅ Database created successfully!' AS status,
    DATABASE() AS current_database,
    (SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE()) AS total_tables;

-- List all created tables
SHOW TABLES;

-- =====================================================
-- END OF SCRIPT
-- =====================================================
