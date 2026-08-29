Organizador financeiro pessoal — contexto do projeto
Objetivo

Projeto pessoal (uso real + prática de Go): organizador financeiro que:

Busca dados via Open Finance (agregador Pluggy)
Responde perguntas financeiras via WhatsApp
Categoriza transações automaticamente com IA
Tem duas interfaces consumindo a mesma lógica: Web (API REST) e CLI

Autor é iniciante em Go — priorizar código idiomático, explicado, sem abstrações prematuras nem "mágica" desnecessária.

Arquitetura

Organização por contexto de negócio (idiomático Go).

Cada contexto (transaction, category, account) é dono do seu domínio, interface de repositório e service. As interfaces são definidas dentro do próprio contexto, no lado do consumidor. Implementações concretas ficam em internal/adapters/.

Regra de dependência: os contextos de negócio não importam nada de adapters. Os adapters importam os contextos para implementar as interfaces. O main.go conecta tudo via injeção de dependência.

SOLID aplicado:

SRP: cada service faz operações de um único contexto.
OCP: trocar um adapter (ex: Pluggy por Belvo) não exige mudança no service.
LSP: qualquer implementação de Repository é substituível (memória vs Postgres).
ISP: interfaces pequenas e específicas, definidas pelo consumidor.
DIP: services dependem de interfaces que eles mesmos definem, nunca de pacotes concretos.

Estrutura de pastas
piggy/
├── cmd/
│   ├── server/          # main.go: API HTTP + webhook do WhatsApp
│   └── cli/             # main.go: comandos Cobra
├── internal/
│   ├── transaction/     # Transaction (Amount int64 centavos, CategoryID *string, AccountID *string), Repository, Service
│   ├── category/        # Category (ParentID *string, suporta um nível de subcategoria), Repository, Service
│   ├── account/         # Account (ID, Number, Name, Owner, Balance int64 centavos), Repository, Service
│   └── adapters/
│       ├── storage/
│       │   ├── memory/  # implementações em memória (transaction, category, account)
│       │   └── postgres/ # implementações Postgres — pgx/v5 (account implementado)
│       ├── pluggy/      # implementa FinanceProvider
│       ├── ai/          # implementa Categorizer
│       └── whatsapp/    # bot whatsmeow
├── migrations/          # golang-migrate, up/down SQL por contexto
│   ├── 001_create_accounts.up.sql / down.sql
│   ├── 002_create_categories.up.sql / down.sql
│   └── 003_create_transactions.up.sql / down.sql
└── go.mod
Stack escolhida
Open Finance: Pluggy, via fluxo "Meu Pluggy" (gratuito e sem prazo de expiração para uso pessoal, desde que as contas conectadas sejam do próprio usuário). Sem SDK oficial em Go — consumir a API REST direto via net/http.
WhatsApp: whatsmeow (go.mau.fi/whatsmeow) — não oficial, protocolo WhatsApp Web multidevice, ainda pré-1.0. Fixar versão exata no go.mod.
Banco: PostgreSQL (Neon — cloud gratuito) + pgx/v5. Migrations via golang-migrate.
CLI: cobra.
HTTP: net/http da stdlib pra começar; migrar pra chi só se precisar de mais middleware.
Categorização: chamada a uma API de LLM pedindo categoria estruturada (JSON) a partir da descrição/estabelecimento da transação.
Convenções de código
Erros explícitos sempre (if err != nil), nunca engolir silenciosamente; empacotar com fmt.Errorf("contexto: %w", err).
Sem variáveis globais mutáveis; dependências entram via injeção no construtor (New...).
Cada usecase precisa de teste unitário usando fakes manuais das ports (interfaces pequenas tornam isso fácil, sem precisar de framework de mock).
Comentários estilo godoc em identificadores exportados.
Valores monetários sempre em centavos (int64) — nunca float.
Conexão com o banco criada no main.go e injetada nos repositórios via construtor.
Variáveis de ambiente carregadas via godotenv (.env local, .env.example commitado sem valores).
Roadmap (ordem de construção — não pular fases)
Domínio + casos de uso com repositório em memória (sem infra externa ainda)
Persistência real em Postgres
Integração Pluggy (sandbox primeiro, depois Meu Pluggy com conta real)
Categorização com IA
API HTTP (Web)
CLI
Bot de WhatsApp
Status atual

Fase 1 concluída. Fase 2 (Postgres) em andamento.

Contextos implementados:
- transaction: struct Transaction (Amount int64, CategoryID *string), interface Repository, Service (Create, List, Read). Implementação em memória.
- category: struct Category (ParentID *string), interface Repository, Service (Create, List, Read). Implementação em memória.
- account: struct Account (Balance int64), interface Repository, Service (Create, List, Read). Implementação em memória E Postgres (pgx/v5).
- handler: Handler com rotas GET/POST para /transactions, /categories e /accounts. Interfaces locais por contexto em interfaces.go.
- cmd/server/main.go: injeção de dependências, conexão Postgres via pgx, carregamento de .env via godotenv.

Pendente (Fase 2):
- postgres/transaction.go e postgres/category.go — adapters Postgres para transaction e category
- transaction.Transaction: adicionar campo AccountID string
- Testes do service (bloqueado por política de AV corporativo — aguardando TI liberar C:\SAPDevelop)

Pendente (fases futuras):
- transaction.Service: método CategorizeManual e interface FinanceProvider (para Sync via Pluggy)
- Contexto card (cartão de crédito separado de conta)
- cmd/cli/ — Cobra CLI
