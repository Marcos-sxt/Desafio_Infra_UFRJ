# Projeto SIGA - Servidor de Desenvolvimento

Este cenário tem como objetivo permitir que você configure um ambiente de desenvolvimento local em sua máquina. 

Siga as etapas abaixo para configurar o ambiente de desenvolvimento local:


## Pré-Requisitos

Antes de começar, certifique-se de que você tenha as seguintes ferramentas instaladas em sua máquina:

- Go(Golang)
- Redis
- Docker
- Docker Compose
- Além das ferramentas básicas de programação (Git, Uma IDE, etc...)

## Instalação

1. Baixe o seguinte repositório:

      https://gitlab.com/Marcos-sxt/infra-2023/-/tree/cenario_de_desenvolvimento?ref_type=heads

   Ou

   Faça o clone deste repositório:

      git clone https://gitlab.com/Marcos-sxt/infra-2023/-/tree/cenario_de_desenvolvimento.git


2. Navegue até o diretório do projeto:

```powershell
cd infra-2023-cenario_de_desenvolvimento
```

3. Construa a imagem Docker do servidor:

(Não se esqueça de abrir o aplicativo do Docker Desktop!)

```powershell
docker-compose build
```

4. Inicie o servidor e o banco de dados Redis:

```powershell
docker-compose up
```
 
 Aguarde até que ambos os contêineres estejam em execução. O servidor estará disponível em http://localhost:3000/.

## Uso

Agora que o servidor está em execução, você pode testar as seguintes rotas:

Rota de Hello World:

```powershell
Rota principal: http://localhost:3000/
(ex: Invoke-WebRequest -Uri "http://localhost:3000")
```

Rota que cria uma sessão:

```powershell
Rota de login (POST): http://localhost:3000/login

(ex: Invoke-WebRequest -Uri "http://localhost:3000/login" -Method POST -Body "username=admin&password=123456")
no powershell do windows
```

Rota que verifica se a sessão existe:

```powershell
Rota de verificação de sessão (substitua SESSION_ID pelo ID da sessão): http://localhost:3000/check/SESSION_ID

(ex: Invoke-WebRequest -Uri "http://localhost:3000/check/session:SESSION_ID")

(Troque a string "SESSION_ID" para o id retornado no metodo de Login(POST))
```

E para listar todas as sessões ativas no servidor, você pode usar a seguinte rota:

```powershell
Rota para listar sessões ativas: http://localhost:3000/list-sessions
(ex: Invoke-WebRequest -Uri "http://localhost:3000/list-sessions")
```

## Contato

Contato
Se você tiver alguma dúvida ou precisar de suporte, entre em contato comigo:

- Nome: Marcos Morais
- Email: marcossantos7955@gmail.com
- GitHub: https://github.com/Marcos-sxt


- 
- 
- 
- 
- 

## Troubleshooting

Aqui estão algumas soluções para problemas comuns que você pode encontrar ao configurar ou usar este ambiente de desenvolvimento local:

### Problema 1: Erro ao construir a imagem Docker

Se você encontrar erros ao construir a imagem Docker, verifique o seguinte:

- Certifique-se de que o Docker e o Docker Compose estejam instalados corretamente em sua máquina.
- Verifique se o arquivo `Dockerfile` está configurado corretamente e se todas as dependências estão sendo instaladas.
- Verifique se você está no diretório correto do projeto ao executar `docker-compose build`.

### Problema 2: O servidor não está respondendo

Se o servidor não estiver respondendo corretamente, siga estas etapas:

- Verifique se o Docker Compose está em execução e se os contêineres do servidor e do Redis estão ativos.
- Certifique-se de que a porta configurada (geralmente a porta 3000) não esteja em uso por outro processo em sua máquina.
- Verifique se todas as variáveis de ambiente necessárias estão configuradas corretamente.

### Problema 3: Erros ao criar ou verificar sessões

Se você encontrar erros ao criar ou verificar sessões, revise o seguinte:

- Certifique-se de que o servidor Redis esteja em execução e acessível a partir do servidor.
- Verifique se as rotas `/login` e `/check/:session` estão configuradas corretamente no servidor.
- Verifique se os dados enviados ao criar uma sessão são válidos.

### Problema 4: Erros ao listar sessões ativas

Se você encontrar problemas ao listar sessões ativas, siga estas etapas:

- Verifique se a rota `/list-sessions` está configurada corretamente no servidor.
- Certifique-se de que a função `getAllSessionKeys` no código esteja funcionando conforme o esperado.
- Verifique se há dados no Redis para listar.

Se você ainda estiver enfrentando problemas após verificar essas etapas, entre em contato comigo para obter suporte direto
