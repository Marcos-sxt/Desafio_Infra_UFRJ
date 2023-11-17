# Usa a imagem oficial do Go como imagem base
FROM golang:latest AS builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia apenas os arquivos de dependências necessários para o container antes de compilar
COPY ./go.mod ./go.sum ./

# Fazendo download das dependências necessárias antes de compilar
RUN go mod download

# Copia os arquivos do projeto para o contêiner
COPY . .

# Compila o código Go com otimizações
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Segunda fase do build para criar uma imagem mínima
FROM scratch

# Copia o executável da fase anterior
COPY --from=builder /app/main /

# Expondo a porta 3000
EXPOSE 3000

# Comando para iniciar o servidor
CMD ["./main"]
