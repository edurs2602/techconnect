# TechConnect API

API de Rede Social voltada para jovens desenvolvedores e estudantes de tecnologia que buscam um espaço focado e performático para compartilhar atualizações técnicas e interagir com outros programadores.

Diferente de redes sociais genéricas, o TechConnect foca em uma arquitetura de alta performance construída em Go com Design Orientado a Domínio (DDD), preparada para futuras consultas complexas e integrações via GraphQL.

## Funcionalidades do MVP

- Criação e gerenciamento de perfis de usuário
- Autenticação segura (JWT)
- CRUD de publicações (posts de até 280 caracteres)
- Sistema de seguir/deixar de seguir outros usuários
- Feed cronológico baseado nos usuários seguidos

## Stack Tecnológica

| Camada | Tecnologia | Justificativa |
|---|---|---|
| Linguagem & Roteamento | Go + Chi | Alta concorrência, execução leve e injeção de dependências para arquitetura hexagonal |
| Arquitetura | Hexagonal + DDD + TDD | Domínio isolado no centro, ports & adapters permitem trocar REST por GraphQL sem alterar lógica de negócio |
| Banco de Dados | PostgreSQL + sqlc | Consistência relacional para grafos de conexões, type-safety nas queries Go |
| Infraestrutura | Docker + GitHub Actions | Deploy conteinerizado e CI com execução automática de testes |
| Deploy | AWS/Azure (EKS/AKS) | Build via GitHub Actions, orquestração com Kubernetes |

## Arquitetura

```
cmd/api/                          → Entrypoint da aplicação
config/                           → Configuração (env vars, conexão com banco)
internal/
  domain/
    user/                         → Entidade, serviço e porta do repositório de usuários
    post/                         → Entidade, serviço e porta do repositório de posts
  application/usecase/            → Casos de uso (register, login, manage post)
  adapter/
    in/http/                      → Handlers HTTP, router, middleware
    out/postgres/                 → Implementação dos repositórios (PostgreSQL)
    out/security/                 → JWT e hashing de senhas
migrations/                       → Migrações SQL
```

Fluxo de dependências: `handler → usecase → domain service → repository port ← postgres adapter`

## Como Executar

```bash
# Subir o banco de dados
docker-compose up -d

# Instalar dependências
go mod tidy

# Rodar a aplicação
go run ./cmd/api

# Rodar testes
go test ./...
```

## Equipe

- **Luis Eduardo Ribeiro** — Backend, implementação e testes
- **Ian Lucas Melo Trindade** — Backend, implementação e testes
- **Ian Mikahel Dionisio** — Backend, implementação e testes
- **Adriano Nobrega Filho** — Backend, implementação e testes

> Projeto da disciplina de Desenvolvimento Web II — UFRN
