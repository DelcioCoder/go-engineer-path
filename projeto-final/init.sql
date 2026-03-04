-- Inicialização do banco de dados
CREATE TABLE IF NOT EXISTS events (
    id          SERIAL PRIMARY KEY,
    tipo        VARCHAR(100) NOT NULL,
    destinatario VARCHAR(255) NOT NULL,
    dados       JSONB NOT NULL DEFAULT '{}',
    status      VARCHAR(20) NOT NULL DEFAULT 'pendente',
    criado_em   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    processado_em TIMESTAMP WITH TIME ZONE,

    CONSTRAINT chk_status CHECK (status IN ('pendente', 'processado', 'falha'))
);

-- Índice para busca por status
CREATE INDEX idx_events_status ON events (status);

-- Índice para busca por tipo
CREATE INDEX idx_events_tipo ON events (tipo);
