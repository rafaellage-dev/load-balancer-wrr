# Utiliza uma imagem leve do Go baseada em Alpine Linux
FROM golang:1.21-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para dentro do container
COPY main.go .

# Compila o binário de forma otimizada para produção
RUN go build -o backend-service main.go

# Expõe a porta interna utilizada pela aplicação
EXPOSE 8080

# Executa o binário compilado
CMD ["./backend-service"]
