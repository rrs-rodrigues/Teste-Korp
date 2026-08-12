# Use uma versão específica em vez de 'latest' para maior estabilidade (ex: 1.21 ou a que você usa)
FROM golang:latest

# O WORKDIR já cria a pasta /app e entra nela
WORKDIR /app

# Copie APENAS os arquivos de módulo primeiro. 
# Isso faz com que o Docker faça cache das dependências, acelerando builds futuros.
COPY go.mod ./
# Se você tiver um arquivo go.sum, descomente a linha abaixo:
# COPY go.sum ./

# Baixe as dependências (se houver)
RUN go mod download

# Agora copie o resto do código da sua máquina para o container
COPY . .

# Compile o aplicativo
RUN go build -o main .

# Exponha a porta que seu app usa
EXPOSE 8080

# Execute o binário
CMD ["./main"]
