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
