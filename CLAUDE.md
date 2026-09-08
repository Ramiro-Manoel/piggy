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
│   ├── transaction/     # struct comum `transaction` (não exportada: ID, Ref, Description, Amount, Date, CategoryID *string) embutida em AccountTransaction (AccountID string) e CardTransaction (InvoiceID string, InstallmentNumber, TotalInstallments int). Repository/Service hoje cobrem só AccountTransaction
│   ├── category/        # Category (ParentID *string, suporta um nível de subcategoria), Repository, Service
│   ├── account/         # Account (ID, Number, Name, Owner, Balance int64 centavos), Repository, Service
│   ├── card/             # [planejado] Card (ID, Ref, Name, Brand, CreditLimit, AvailableLimit, ClosingDay, DueDay), Repository, Service
│   ├── invoice/          # [planejado] Invoice (ID, CardID, ClosingDate, DueDate, TotalAmount, Status), Repository, Service — decide a qual fatura cada CardTransaction pertence
│   └── adapters/
│       ├── storage/
│       │   ├── memory/  # implementações em memória (transaction, category, account)
│       │   └── postgres/ # implementações Postgres — pgx/v5 (account, transaction, category implementados)
│       ├── finance_provider/pluggy/ # implementa financeProvider (auth, FetchAccounts, FetchTransactions)
│       ├── ai/          # implementa Categorizer
│       └── whatsapp/    # bot whatsmeow
├── migrations/          # golang-migrate, up/down SQL por contexto
│   ├── 001_create_accounts.up.sql / down.sql
│   ├── 002_create_categories.up.sql / down.sql
│   ├── 003_create_transactions.up.sql / down.sql
│   └── [planejado] 004_create_cards, 005_create_invoices, 006_create_card_transactions
│       (card_transactions referencia invoices, não cards diretamente — sem FK redundante,
│       já que Invoice.CardID já resolve isso)
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

Fase 1 concluída. Fase 2 (Postgres) concluída para account, transaction (AccountTransaction) e category. Integração Pluggy iniciada (client + mapper, autenticação e fetch de contas/transações via sandbox).

Contextos implementados:
- transaction: struct comum não-exportada `transaction` (ID, Ref, Description, Amount, Date, CategoryID *string) embutida em `AccountTransaction` (AccountID string) e `CardTransaction` (InvoiceID string, InstallmentNumber, TotalInstallments int — struct existe, ainda sem Repository/Service/adapter). Repository, Service (Create, List, Read, Sync) cobrem hoje só AccountTransaction. Implementação em memória e Postgres.
- category: struct Category (ParentID *string), interface Repository, Service (Create, List, Read). Implementação em memória.
- account: struct Account (Balance int64), interface Repository, Service (Create, List, Read, Sync). Implementação em memória E Postgres (pgx/v5).
- handler: Handler com rotas GET/POST para /transactions, /categories, /accounts, e POST /accounts/sync, /transactions/sync/{accountID}. Interfaces locais por contexto em interfaces.go.
- cmd/server/main.go: injeção de dependências, conexão Postgres via pgx, autenticação Pluggy, carregamento de .env via godotenv.
- adapters/finance_provider/pluggy: client HTTP (Authenticate, FetchAccounts, FetchTransactions) e mapper pra account.Account / transaction.AccountTransaction.

Pendente (Fase 2):
- Testes do service (bloqueado por política de AV corporativo — aguardando TI liberar C:\SAPDevelop)

Pendente (cartão — desenho já decidido em conversa, implementação não iniciada):
- internal/card/: Card, Repository, Service, financeProvider (FetchCards) — mapeado do Pluggy via campo `type: CREDIT` em /accounts (hoje o client só busca contas BANK)
- internal/invoice/: Invoice, Repository, Service (FindOrCreate, Close, List) — calcula a fatura de cada CardTransaction a partir do ClosingDay do cartão; não confiar no `billId`/`billForecastDate` do Pluggy (vêm nulos na maioria das compras não parceladas em dados reais de sandbox)
- transaction.Repository/Service viram genéricos (Repository[T], helper saveAll[T identifiable]) pra cobrir CardTransaction sem duplicar a parte de CRUD; CardTransactionService.Sync fica separado do de AccountTransaction porque depende de um invoiceFinder pra resolver InvoiceID antes de salvar
- migrations 004_create_cards, 005_create_invoices, 006_create_card_transactions

Pendente (fases futuras):
- transaction: método CategorizeManual (categorização manual/IA)
- cmd/cli/ — Cobra CLI
