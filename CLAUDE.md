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
│   ├── transaction/     # Transaction, Repository, Service
│   ├── category/        # Category, Repository, Categorizer, Service
│   ├── account/         # Account, Repository, Service
│   └── adapters/
│       ├── storage/
│       │   ├── memory/  # implementações em memória
│       │   └── postgres/
│       ├── pluggy/      # implementa FinanceProvider
│       ├── ai/          # implementa Categorizer
│       └── whatsapp/    # bot whatsmeow
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

Fase 1 em andamento. Contexto transaction implementado com struct, interface Repository e Service básico (Save, List). Implementação em memória em adapters/storage/memory. Sem infraestrutura externa ainda.