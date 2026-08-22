Organizador financeiro pessoal — contexto do projeto
Objetivo

Projeto pessoal (uso real + prática de Go): organizador financeiro que:

Busca dados via Open Finance (agregador Pluggy)
Responde perguntas financeiras via WhatsApp
Categoriza transações automaticamente com IA
Tem duas interfaces consumindo a mesma lógica: Web (API REST) e CLI

Autor é iniciante em Go — priorizar código idiomático, explicado, sem abstrações prematuras nem "mágica" desnecessária.

Arquitetura

Clean Architecture / Ports & Adapters (Hexagonal).

Regra de dependência: o núcleo (domínio + casos de uso) não importa nenhuma lib externa — nada de net/http, driver de banco, ou SDK de terceiros ali dentro. O núcleo define interfaces ("ports") em internal/usecase/port. Adapters (HTTP, CLI, WhatsApp, Postgres, Pluggy, IA) é que implementam ou consomem essas interfaces. Nunca inverter essa direção.

SOLID aplicado:

SRP: cada usecase faz uma coisa só (sync não categoriza, categorizar não sincroniza).
OCP: trocar um adapter (ex: Pluggy por Belvo) não deve exigir mudança nos usecases.
LSP: qualquer implementação de uma port é substituível por outra sem quebrar o usecase (repo em memória vs Postgres).
ISP: interfaces pequenas e específicas — nunca uma "mega interface" de repositório.
DIP: usecases dependem de abstrações que eles mesmos definem, nunca de pacotes concretos.
Estrutura de pastas
meu-organizador/
├── cmd/
│   ├── server/          # main.go: API HTTP + webhook do WhatsApp
│   └── cli/             # main.go: comandos Cobra
├── internal/
│   ├── domain/          # Transaction, Account, Category, Budget — Go puro
│   ├── usecase/
│   │   ├── port/        # interfaces: repository.go, openfinance.go, categorizer.go, messenger.go
│   │   ├── sync_transactions.go
│   │   ├── categorize_transactions.go
│   │   ├── get_summary.go
│   │   └── answer_question.go
│   ├── adapter/
│   │   ├── in/
│   │   │   ├── http/      # handlers REST
│   │   │   ├── cli/       # comandos que chamam os usecases
│   │   │   └── whatsapp/  # webhook / listener
│   │   └── out/
│   │       ├── postgres/  # implementa port.TransactionRepository
│   │       ├── pluggy/    # implementa port.OpenFinanceProvider
│   │       ├── ai/        # implementa port.Categorizer
│   │       └── whatsapp/  # implementa port.Messenger
│   └── config/
├── migrations/
└── go.mod
Stack escolhida
Open Finance: Pluggy, via fluxo "Meu Pluggy" (gratuito e sem prazo de expiração para uso pessoal, desde que as contas conectadas sejam do próprio usuário). Sem SDK oficial em Go — consumir a API REST direto via net/http.
WhatsApp: whatsmeow (go.mau.fi/whatsmeow) — não oficial, protocolo WhatsApp Web multidevice, ainda pré-1.0. Fixar versão exata no go.mod.
Banco: PostgreSQL + pgx ou sqlc.
CLI: cobra.
HTTP: net/http da stdlib pra começar; migrar pra chi só se precisar de mais middleware.
Categorização: chamada a uma API de LLM pedindo categoria estruturada (JSON) a partir da descrição/estabelecimento da transação.
Convenções de código
Erros explícitos sempre (if err != nil), nunca engolir silenciosamente; empacotar com fmt.Errorf("contexto: %w", err).
Sem variáveis globais mutáveis; dependências entram via injeção no construtor (New...).
Cada usecase precisa de teste unitário usando fakes manuais das ports (interfaces pequenas tornam isso fácil, sem precisar de framework de mock).
Comentários estilo godoc em identificadores exportados.
Roadmap (ordem de construção — não pular fases)
Domínio + casos de uso com repositório em memória (sem infra externa ainda)
Persistência real em Postgres
Integração Pluggy (sandbox primeiro, depois Meu Pluggy com conta real)
Categorização com IA
API HTTP (Web)
CLI
Bot de WhatsApp
Status atual

Nenhum código escrito ainda. Começando pela Fase 1.