# 📚 DevBook API

    DevBook API é uma aplicação backend em Go desenvolvida como parte do curso "Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!" na Udemy.
    O sistema simula uma rede social para desenvolvedores, permitindo o cadastro de usuários, a criação de publicações e a interação entre seguidores.


## 📌 Funcionalidades da API

    🔑 POST /login → Autenticar usuário e gerar token JWT

    👤 POST /usuarios → Criar usuário
    👤 GET /usuarios → Listar todos os usuários
    👤 GET /usuarios/{id} → Buscar usuário por ID
    👤 PUT /usuarios/{id} → Atualizar usuário
    👤 DELETE /usuarios/{id} → Deletar usuário
    👤 POST /usuarios/{id}/atualizar-senha → Atualizar senha

    📝 POST /publicacoes → Criar publicação
    📝 GET /publicacoes → Listar todas as publicações
    📝 GET /publicacoes/{id} → Buscar publicação por ID
    📝 PUT /publicacoes/{id} → Atualizar publicação
    📝 DELETE /publicacoes/{id} → Remover publicação
    📝 POST /publicacoes/{id}/curtir → Curtir publicação
    📝 POST /publicacoes/{id}/descurtir → Descurtir publicação
    📝 GET /usuarios/{id}/publicacoes → Listar publicações de um usuário

    🤝 POST /usuarios/{id}/seguir → Seguir um usuário
    🤝 POST /usuarios/{id}/deixar-de-seguir → Deixar de seguir um usuário
    🤝 GET /usuarios/{id}/seguidores → Listar seguidores de um usuário
    🤝 GET /usuarios/{id}/seguindo → Listar quem o usuário está seguindo

## 📦 Tecnologias e Pacotes Utilizados

    [Go] → Linguagem principal da API
    [Mux] → Router HTTP 
    [MySQL] → Banco de dados relacional
    [GoDotEnv] → Gerencimanto de variáveis de ambiente
    [Bcrypt] → Criptografia de senhas
    [JWT Go] → Implementação de JSON Web Tokens
    [CheckMail] → Validação de e-mail dos usuários

## 📂 Estrutura do Projeto

    devbook-api/                        # Raiz do projeto
    │
    │── internal/                       # Código privado do projeto
    │   │
    │   │── auth/                       # Módulo de autenticação JWT
    │   │   └── auth.go                 # Funções principais de autenticação (validação de token)
    │   │
    │   │── config/                     # Configurações específicas
    │   │   └── config.go               # Carregamento de configs (.env)
    │   │
    │   │── data/                       # Conexão com o Banco de Dados
    │   │   └── data.go                 # Função principal de conexão
    │   │
    │   │── middlewares/                # Middlewares internos
    │   │   └── middlewares.go          # Definição de middlewares globais
    │   │
    │   │── repository/                 # Regras de persistência (DAO)
    │   │   │── publicacoes.go          # Consulta e operações de banco para publicações.
    │   │   └── usuarios.go             # Consulta e operações de banco para usuários.
    │   │
    │   └── security/                   # Módelo de segurança bcrypt
    │       └── security.go             # Funções para hashing de senha
    │
    │── pkg/                            # Código público do projeto
    │   │
    │   │── controllers/                # Controladores: lógica de entrada (camada HTTP)
    │   │   │── login.go                # Controller responsável por login/autenticação
    │   │   │── publicacoes.go          # Controller para operações com publicações
    │   │   └── usuarios.go             # Controller para operações com usuários
    │   │
    │   │── models/                     # Estruturas e modelos de dados
    │   │   │── Publicacoes.go          # Struct de "Publicações"
    │   │   │── Senha.go                # Struct de "Senha"
    │   │   └── Usuario.go              # Struct de "Usuário"
    │   │
    │   │── routes/                     # Definição das rotas
    │   │   │── rotas/                  # Agrupamento de rotas por domínio
    │   │   │   │── login.go            # Rota de Login
    │   │   │   │── publicacoes.go      # Rotas relacionadas a publicações
    │   │   │   │── rotas.go            # Registro central de rotas
    │   │   │   └── usuarios.go         # Rotas relacionadas a usuários
    │   │   └── router.go               # Inicialização e configuração do roteador principal
    │   │
    │   └── utils/                      # Funções utilitárias
    │       └── response.go             # Helper para padronizar respostas HTTP
    │
    │── sql/                            # Scripts SQL
    │   └── data.sql                    # Script de criação e população inicial do banco
    │
    │── .env                            # Variáveis de ambiente
    │── go.mod                          # Dependências do Go
    │── go.sum                          # Hash das dependências
    └── main.go                         # Ponto de entrada da aplicação

## 📎 Autor

    Este projeto foi desenvolvido com base em um curso de Go para APIs REST,
    adaptado e expandido por Pedro Henrique Leite para aprendizado e prática.

## 📄 Licença

    Este projeto é de uso educacional, sem fins comerciais.  
    Sinta-se à vontade para utilizar como referência em seus estudos 🚀