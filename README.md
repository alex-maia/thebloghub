# TheBlogHub

O TheBlogHub é um projeto desenvolvido apenas para fins de estudo, utilizando Go no backend e tecnologias modernas no frontend. Serve como exemplo de estrutura e organização para quem pretende aprender ou experimentar desenvolvimento web com Go.

## Funcionalidades

- **Backend em Go** com arquitetura MVC
- **Validação robusta de configuração** por ambiente
- **Gestão de variáveis de ambiente** com valores padrão
- **Sistema de templates** HTML modular
- **Build system** com Gulp (SCSS, JS)
- **Layout responsivo** com design moderno
- **Minificação de assets** para melhor performance
- **Logs informativos** com diferentes níveis por ambiente

## Estrutura do Projeto

```
public/        # Ficheiros estáticos compilados (CSS, JS, imagens, fontes)
resources/     # Ficheiros fonte (SCSS, JS, templates HTML)
├── js/        # JavaScript fonte
├── scss/      # Estilos SCSS
└── views/     # Templates HTML
config/        # Configurações da aplicação com validação
controllers/   # Lógica de controlo (Go)
database/      # Conexão e gestão da base de dados
models/        # Modelos de dados
routes/        # Definição de rotas
.env.example   # Exemplo de variáveis de ambiente
gulpfile.js    # Tarefas de build frontend
main.go        # Entrada principal da aplicação
```

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/alex-maia/thebloghub.git
   cd thebloghub
   ```
2. Instale as dependências Go:
   ```bash
   go mod tidy
   ```
3. Instale as dependências do frontend:
   ```bash
   npm install
   ```
4. Configure as variáveis de ambiente:
   ```bash
   cp .env.example .env
   # Edite o arquivo .env com as suas configurações
   ```

## Configuração

### Variáveis de Ambiente

O projeto utiliza variáveis de ambiente para configuração. Copie o arquivo `.env.example` para `.env` e configure as seguintes variáveis:

- `APP_NAME`: Nome da aplicação (padrão: TheBlogHub)
- `APP_ENV`: Ambiente de execução (development/production/testing)
- `APP_PORT`: Porta do servidor (padrão: 8080)
- `DB_HOST`: Host da base de dados (padrão: localhost)
- `DB_PORT`: Porta da base de dados (padrão: 5432)
- `DB_USER`: Utilizador da base de dados
- `DB_PASS`: Password da base de dados
- `DB_NAME`: Nome da base de dados
- `SECRET_KEY`: Chave secreta (mínimo 32 caracteres)

### Validação de Configuração

A aplicação valida automaticamente as configurações ao iniciar:

- **Ambientes suportados:** `development`, `production`, `testing`
- **Modo produção:** Todas as variáveis de base de dados são obrigatórias
- **Segurança:** `SECRET_KEY` deve ter pelo menos 32 caracteres
- **Valores padrão:** Configurações sensatas para desenvolvimento
- **Tratamento de erros:** Falha rápida com mensagens claras

### Funções Utilitárias

O módulo `config` inclui funções para verificar o ambiente atual:

```go
config.IsDevelopment()  // true se APP_ENV=development
config.IsProduction()   // true se APP_ENV=production
config.IsTesting()      // true se APP_ENV=testing
```

## Execução

### Desenvolvimento

```bash
# Executar diretamente
go run main.go

# Ou compilar e executar
go build -o thebloghub main.go
./thebloghub
```

### Build de Assets

```bash
# Compilar SCSS (resources/scss → public/css)
npm run build-dev

# Compilar JS (resources/js → public/js)
npm run build-js

# Watch mode (recompila automaticamente)
npm run build-watch
```

### Estrutura de Build

- **Fonte:** `resources/` (SCSS, JS, templates)
- **Compilado:** `public/` (CSS, JS bundles, assets finais)
- **Separação clara** entre código fonte e assets públicos
- **URLs:** Assets servidos em `/public/` (ex: `/public/css/main.css`)

### Exemplos de Logs

```
Configuração carregada com sucesso - Ambiente: development, Porta: 8080
Servidor TheBlogHub iniciado em http://localhost:8080
Ambiente: development
```

## Como contribuir

1. Faça um fork do projeto
2. Crie uma branch para a sua funcionalidade/correção
3. Envie um pull request

## Licença

Este projeto está sob a licença MIT.
