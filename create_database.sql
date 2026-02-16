-- =====================================================
-- SISTEMA DE GERENCIAMENTO DE MECÂNICA
-- Script de Criação do Banco de Dados MySQL
-- Versão: 1.0
-- Data: Fevereiro 2026
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
-- TABELA: CLIENTES
-- =====================================================
CREATE TABLE clientes (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE COMMENT 'Formato: 000.000.000-00',
    telefone_principal VARCHAR(20) NOT NULL,
    telefone_secundario VARCHAR(20),
    email VARCHAR(255),
    status ENUM('ativo', 'inativo') DEFAULT 'ativo',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL COMMENT 'Soft delete',
    
    INDEX idx_cpf (cpf),
    INDEX idx_status (status),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Cadastro de clientes da mecânica';

-- =====================================================
-- TABELA: ENDERECOS
-- =====================================================
CREATE TABLE enderecos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cliente_id BIGINT UNSIGNED NOT NULL,
    cep VARCHAR(10) NOT NULL COMMENT 'Formato: 00000-000',
    rua VARCHAR(255) NOT NULL,
    numero VARCHAR(20) NOT NULL,
    complemento VARCHAR(255),
    bairro VARCHAR(100) NOT NULL,
    cidade VARCHAR(100) NOT NULL,
    estado CHAR(2) NOT NULL COMMENT 'Sigla do estado: SP, RJ, etc',
    tipo ENUM('residencial', 'comercial') DEFAULT 'residencial',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE,
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_cidade_estado (cidade, estado)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Endereços dos clientes';

-- =====================================================
-- TABELA: VEICULOS
-- =====================================================
CREATE TABLE veiculos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cliente_id BIGINT UNSIGNED NOT NULL,
    placa VARCHAR(10) NOT NULL UNIQUE COMMENT 'Formatos: ABC-1234 ou ABC1D23',
    marca VARCHAR(100) NOT NULL COMMENT 'Ex: Fiat, Volkswagen, Ford',
    modelo VARCHAR(100) NOT NULL COMMENT 'Ex: Uno, Gol, Fiesta',
    ano_fabricacao INT NOT NULL,
    ano_modelo INT NOT NULL,
    cor VARCHAR(50),
    chassi VARCHAR(50) COMMENT 'Número do chassi',
    quilometragem_atual INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL COMMENT 'Soft delete',
    
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE CASCADE,
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_placa (placa),
    INDEX idx_marca_modelo (marca, modelo),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Veículos cadastrados (um cliente pode ter vários veículos)';

-- =====================================================
-- TABELA: ORDENS_SERVICO
-- =====================================================
CREATE TABLE ordens_servico (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    numero_os VARCHAR(20) NOT NULL UNIQUE COMMENT 'Ex: OS-2024-0001',
    cliente_id BIGINT UNSIGNED NOT NULL,
    veiculo_id BIGINT UNSIGNED NOT NULL,
    data_entrada TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    quilometragem_entrada INT COMMENT 'KM do veículo na entrada',
    descricao_problema TEXT NOT NULL COMMENT 'Problema relatado pelo cliente',
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
    observacoes TEXT COMMENT 'Observações gerais da OS',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (cliente_id) REFERENCES clientes(id) ON DELETE RESTRICT,
    FOREIGN KEY (veiculo_id) REFERENCES veiculos(id) ON DELETE RESTRICT,
    INDEX idx_numero_os (numero_os),
    INDEX idx_cliente_id (cliente_id),
    INDEX idx_veiculo_id (veiculo_id),
    INDEX idx_status (status),
    INDEX idx_data_entrada (data_entrada)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Ordens de serviço (OS) - controle de entrada de veículos';

-- =====================================================
-- TABELA: ORCAMENTOS
-- =====================================================
CREATE TABLE orcamentos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    numero_orcamento VARCHAR(20) NOT NULL UNIQUE COMMENT 'Ex: ORC-2024-0001',
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    valor_pecas DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'Soma dos valores das peças',
    valor_mao_obra DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'Soma dos valores dos serviços',
    valor_total DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'valor_pecas + valor_mao_obra',
    desconto DECIMAL(10, 2) DEFAULT 0.00,
    valor_final DECIMAL(10, 2) DEFAULT 0.00 COMMENT 'valor_total - desconto',
    data_validade DATE COMMENT 'Validade do orçamento',
    status ENUM('pendente', 'aprovado', 'recusado', 'expirado') DEFAULT 'pendente',
    data_aprovacao_recusa TIMESTAMP NULL,
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_numero_orcamento (numero_orcamento),
    INDEX idx_data_validade (data_validade)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Orçamentos gerados para cada OS';

-- =====================================================
-- TABELA: ITENS_ORCAMENTO
-- =====================================================
CREATE TABLE itens_orcamento (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    orcamento_id BIGINT UNSIGNED NOT NULL,
    tipo ENUM('peca', 'servico') NOT NULL COMMENT 'Tipo do item',
    descricao VARCHAR(255) NOT NULL COMMENT 'Descrição da peça ou serviço',
    quantidade INT NOT NULL DEFAULT 1,
    valor_unitario DECIMAL(10, 2) NOT NULL,
    valor_total DECIMAL(10, 2) NOT NULL COMMENT 'quantidade * valor_unitario',
    observacao TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (orcamento_id) REFERENCES orcamentos(id) ON DELETE CASCADE,
    INDEX idx_orcamento_id (orcamento_id),
    INDEX idx_tipo (tipo)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Itens detalhados de cada orçamento (peças e serviços)';

-- =====================================================
-- TABELA: SERVICOS_EXECUTADOS
-- =====================================================
CREATE TABLE servicos_executados (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    os_id BIGINT UNSIGNED NOT NULL,
    mecanico_responsavel VARCHAR(255) COMMENT 'Nome do mecânico',
    data_inicio TIMESTAMP NULL,
    data_conclusao TIMESTAMP NULL,
    descricao_servico TEXT NOT NULL COMMENT 'Detalhamento do serviço realizado',
    tempo_estimado_horas DECIMAL(5, 2) COMMENT 'Tempo estimado em horas',
    tempo_real_horas DECIMAL(5, 2) COMMENT 'Tempo real gasto',
    status ENUM('pendente', 'em_andamento', 'concluido') DEFAULT 'pendente',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE CASCADE,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_mecanico (mecanico_responsavel)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Registro dos serviços executados em cada OS';

-- =====================================================
-- TABELA: PAGAMENTOS
-- =====================================================
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
    valor_final DECIMAL(10, 2) NOT NULL COMMENT 'valor_pago - desconto_aplicado',
    status ENUM('pendente', 'pago_parcial', 'pago_total') DEFAULT 'pendente',
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (os_id) REFERENCES ordens_servico(id) ON DELETE RESTRICT,
    INDEX idx_os_id (os_id),
    INDEX idx_status (status),
    INDEX idx_forma_pagamento (forma_pagamento),
    INDEX idx_data_pagamento (data_pagamento)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
COMMENT='Registro de pagamentos das OS';

-- =====================================================
-- DADOS DE TESTE (OPCIONAL)
-- =====================================================

-- Inserir cliente de teste
INSERT INTO clientes (nome, cpf, telefone_principal, email, status) VALUES
('João da Silva', '123.456.789-00', '(11) 98765-4321', 'joao.silva@email.com', 'ativo'),
('Maria Oliveira', '987.654.321-00', '(11) 91234-5678', 'maria.oliveira@email.com', 'ativo');

-- Inserir endereço de teste
INSERT INTO enderecos (cliente_id, cep, rua, numero, bairro, cidade, estado) VALUES
(1, '01310-100', 'Avenida Paulista', '1000', 'Bela Vista', 'São Paulo', 'SP'),
(2, '01310-200', 'Rua Augusta', '500', 'Consolação', 'São Paulo', 'SP');

-- Inserir veículos de teste
INSERT INTO veiculos (cliente_id, placa, marca, modelo, ano_fabricacao, ano_modelo, cor, quilometragem_atual) VALUES
(1, 'ABC-1234', 'Fiat', 'Uno', 2020, 2020, 'Branco', 50000),
(1, 'DEF-5678', 'Volkswagen', 'Gol', 2019, 2019, 'Prata', 80000),
(2, 'GHI-9012', 'Chevrolet', 'Onix', 2021, 2022, 'Preto', 30000);

-- Inserir OS de teste
INSERT INTO ordens_servico (numero_os, cliente_id, veiculo_id, quilometragem_entrada, descricao_problema, status) VALUES
('OS-2024-0001', 1, 1, 50000, 'Barulho estranho no motor e luz do painel acesa', 'aguardando_orcamento'),
('OS-2024-0002', 2, 3, 30000, 'Troca de óleo e revisão dos 30 mil km', 'aprovado');

-- =====================================================
-- VIEWS ÚTEIS
-- =====================================================

-- View: OS com dados completos do cliente e veículo
CREATE OR REPLACE VIEW vw_os_completas AS
SELECT 
    os.id,
    os.numero_os,
    os.data_entrada,
    os.status,
    c.nome AS cliente_nome,
    c.cpf AS cliente_cpf,
    c.telefone_principal AS cliente_telefone,
    v.placa AS veiculo_placa,
    v.marca AS veiculo_marca,
    v.modelo AS veiculo_modelo,
    os.descricao_problema,
    os.data_prevista_conclusao
FROM ordens_servico os
INNER JOIN clientes c ON os.cliente_id = c.id
INNER JOIN veiculos v ON os.veiculo_id = v.id
WHERE os.deleted_at IS NULL;

-- View: Resumo financeiro por OS
CREATE OR REPLACE VIEW vw_resumo_financeiro_os AS
SELECT 
    os.id AS os_id,
    os.numero_os,
    orc.valor_final AS valor_orcamento,
    COALESCE(SUM(p.valor_final), 0) AS valor_pago,
    (orc.valor_final - COALESCE(SUM(p.valor_final), 0)) AS saldo_pendente,
    CASE 
        WHEN COALESCE(SUM(p.valor_final), 0) >= orc.valor_final THEN 'pago'
        WHEN COALESCE(SUM(p.valor_final), 0) > 0 THEN 'parcial'
        ELSE 'pendente'
    END AS situacao_pagamento
FROM ordens_servico os
LEFT JOIN orcamentos orc ON os.id = orc.os_id AND orc.status = 'aprovado'
LEFT JOIN pagamentos p ON os.id = p.os_id
GROUP BY os.id, os.numero_os, orc.valor_final;

-- =====================================================
-- TRIGGERS
-- =====================================================

-- Trigger: Calcular valor_total do item ao inserir
DELIMITER $$
CREATE TRIGGER trg_calcula_valor_item_insert
BEFORE INSERT ON itens_orcamento
FOR EACH ROW
BEGIN
    SET NEW.valor_total = NEW.quantidade * NEW.valor_unitario;
END$$
DELIMITER ;

-- Trigger: Calcular valor_total do item ao atualizar
DELIMITER $$
CREATE TRIGGER trg_calcula_valor_item_update
BEFORE UPDATE ON itens_orcamento
FOR EACH ROW
BEGIN
    SET NEW.valor_total = NEW.quantidade * NEW.valor_unitario;
END$$
DELIMITER ;

-- Trigger: Atualizar totais do orçamento quando adicionar item
DELIMITER $$
CREATE TRIGGER trg_atualiza_orcamento_after_item_insert
AFTER INSERT ON itens_orcamento
FOR EACH ROW
BEGIN
    UPDATE orcamentos
    SET 
        valor_pecas = (
            SELECT COALESCE(SUM(valor_total), 0) 
            FROM itens_orcamento 
            WHERE orcamento_id = NEW.orcamento_id AND tipo = 'peca'
        ),
        valor_mao_obra = (
            SELECT COALESCE(SUM(valor_total), 0) 
            FROM itens_orcamento 
            WHERE orcamento_id = NEW.orcamento_id AND tipo = 'servico'
        )
    WHERE id = NEW.orcamento_id;
    
    UPDATE orcamentos
    SET 
        valor_total = valor_pecas + valor_mao_obra,
        valor_final = valor_pecas + valor_mao_obra - desconto
    WHERE id = NEW.orcamento_id;
END$$
DELIMITER ;

-- =====================================================
-- PROCEDURES ÚTEIS
-- =====================================================

-- Procedure: Buscar histórico de um cliente
DELIMITER $$
CREATE PROCEDURE sp_historico_cliente(IN p_cliente_id BIGINT)
BEGIN
    SELECT 
        os.numero_os,
        os.data_entrada,
        os.status,
        v.placa,
        v.marca,
        v.modelo,
        os.descricao_problema,
        orc.valor_final AS valor_orcamento
    FROM ordens_servico os
    INNER JOIN veiculos v ON os.veiculo_id = v.id
    LEFT JOIN orcamentos orc ON os.id = orc.os_id AND orc.status = 'aprovado'
    WHERE os.cliente_id = p_cliente_id
    ORDER BY os.data_entrada DESC;
END$$
DELIMITER ;

-- Procedure: Buscar OS pendentes de orçamento
DELIMITER $$
CREATE PROCEDURE sp_os_pendentes_orcamento()
BEGIN
    SELECT 
        os.numero_os,
        c.nome AS cliente,
        v.placa,
        v.modelo,
        os.descricao_problema,
        os.data_entrada,
        DATEDIFF(CURRENT_DATE, DATE(os.data_entrada)) AS dias_aguardando
    FROM ordens_servico os
    INNER JOIN clientes c ON os.cliente_id = c.id
    INNER JOIN veiculos v ON os.veiculo_id = v.id
    WHERE os.status = 'aguardando_orcamento'
    ORDER BY os.data_entrada ASC;
END$$
DELIMITER ;

-- =====================================================
-- MENSAGEM DE SUCESSO
-- =====================================================
SELECT 
    '✅ Banco de dados criado com sucesso!' AS status,
    DATABASE() AS banco_atual,
    (SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE()) AS total_tabelas;

-- Listar todas as tabelas criadas
SHOW TABLES;

-- =====================================================
-- FIM DO SCRIPT
-- =====================================================
