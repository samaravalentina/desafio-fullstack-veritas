Desafio Fullstack – Mini Kanban de Tarefas

Aplicação fullstack simples utilizando React no frontend e Go no backend. Sistema Kanban com três colunas fixas (A Fazer, Em Progresso e Concluídas).

## 📋 Requisitos

Antes de começar, certifique-se de ter instalado:

- **Go** (versão 1.21 ou superior): [Download](https://golang.org/dl/)
  - Após instalar, verifique com: `go version`
- **Node.js** (versão 18 ou superior): [Download](https://nodejs.org/)
  - O npm vem incluído com o Node.js
  - **Nota para Windows/PowerShell**: Se encontrar erro de política de execução, execute:
    ```powershell
    Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
    ```

Para mais detalhes, consulte o arquivo [REQUISITOS.md](REQUISITOS.md).

## 🚀 Como Rodar

### Backend (Go)

1. Navegue até a pasta do backend:
```bash
cd backend
```

2. Instale as dependências (se ainda não fez):
```bash
go mod init backend
go mod tidy
```

**Nota**: O `go.mod` já está criado. Se necessário, apenas execute `go mod tidy` para verificar dependências.

3. Execute o servidor:
```bash
go run ./cmd/server
```

O servidor estará rodando em `http://localhost:8080`

### Frontend (React)

1. Navegue até a pasta do frontend:
```bash
cd frontend
```

2. Instale as dependências:
```bash
npm install
```

**Nota**: Na primeira execução, isso pode levar alguns minutos para baixar todas as dependências do React.

3. Execute a aplicação:
```bash
npm start
```

A aplicação estará disponível em `http://localhost:3000`

## 📁 Estrutura do Projeto

```
/backend
  cmd/
    server/             # Configuração do servidor e rotas
  internal/
    handlers/           # Handlers das requisições HTTP
    middleware/         # Middleware (CORS)
    storage/            # Persistência de dados
    scripts/
      start.bat         # Script para iniciar no Windows
      start.sh          # Script para iniciar no Linux/macOS
      test-api.http     # Teste de API (VS Code/Insomnia)
      test-backend.ps1  # Script de teste em PowerShell
  go.mod                # Dependências do Go
  go.sum                # Checksum das dependências

/frontend
  package.json          # Dependências do React
  src/                  # Código fonte do React
    App.js
    components/
    services/

/docs
  user-flow.png         # Diagrama de fluxo do usuário (obrigatório)
```

## 🛠️ Decisões Técnicas

### Backend
- **Framework**: Utilizado `net/http` nativo do Go para simplicidade
- **Armazenamento**: Inicialmente em memória, com suporte a persistência em JSON
- **CORS**: Configurado para permitir acesso do frontend React
- **Validações**: Validação básica de campos obrigatórios e status válidos

### Frontend
- **Framework**: React com Create React App
- **Estilização**: CSS puro para manter simplicidade
- **Gerenciamento de Estado**: useState e useEffect do React
- **Requisições HTTP**: Fetch API nativa

## ⚠️ Limitações Conhecidas

- Armazenamento em memória: dados são perdidos ao reiniciar o servidor
- Sem autenticação/autorização
- Sem validação avançada de dados no frontend
- Interface básica sem bibliotecas de UI

## 🔮 Melhorias Futuras

- [ ] Persistência em banco de dados (PostgreSQL, SQLite)
- [ ] Drag and drop para mover tarefas entre colunas
- [ ] Autenticação de usuários
- [ ] Testes unitários e de integração
- [ ] Docker para containerização
- [ ] Interface mais polida com biblioteca de componentes (Material-UI, Ant Design)
- [ ] Validações mais robustas
- [ ] Tratamento de erros mais detalhado
- [ ] Logging estruturado

## 📊 Documentação

- **User Flow**: Ver `docs/user-flow.png`

## 📝 Endpoints da API

### GET /tasks
Retorna todas as tarefas

### POST /tasks
Cria uma nova tarefa
```json
{
  "title": "Título da tarefa",
  "description": "Descrição opcional",
  "status": "todo"
}
```

### PUT /tasks/:id
Atualiza uma tarefa existente
```json
{
  "title": "Título atualizado",
  "description": "Descrição atualizada",
  "status": "in_progress"
}
```

### DELETE /tasks/:id
Remove uma tarefa

## 🔒 Status Válidos

- `todo` - A Fazer
- `in_progress` - Em Progresso
- `done` - Concluídas
=======
# Desafio-fullstack-veritas
Desafio Fullstack – Veritas
