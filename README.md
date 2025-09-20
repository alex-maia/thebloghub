# TheBlogHub

O TheBlogHub é um projeto desenvolvido apenas para fins de estudo, utilizando Go no backend e tecnologias modernas no frontend. Serve como exemplo de estrutura e organização para quem pretende aprender ou experimentar desenvolvimento web com Go.

## Funcionalidades

- Publicação e gestão de artigos
- Estrutura modular (Go, SCSS, JS)
- Layout responsivo
- Minificação de assets para melhor performance

## Estrutura do Projeto

```
assets/        # Ficheiros estáticos (CSS, JS, imagens, fontes)
config/        # Configurações da aplicação
controllers/   # Lógica de controlo (Go)
database/      # Conexão e gestão da base de dados
models/        # Modelos de dados
routes/        # Definição de rotas
views/         # Templates HTML
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

- Em modo `production`, todas as variáveis de base de dados são obrigatórias
- A `SECRET_KEY` deve ter pelo menos 32 caracteres
- O `APP_ENV` deve ser um dos valores válidos

## Execução

1. Compile e execute o backend:
   ```bash
   go run main.go
   ```
2. Para compilar/minificar assets:
   ```bash
   npm run build
   ```

## Como contribuir

1. Faça um fork do projeto
2. Crie uma branch para a sua funcionalidade/correção
3. Envie um pull request

## Licença

Este projeto está sob a licença MIT.
