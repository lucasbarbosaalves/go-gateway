# Go Gateway

Este projeto de uma API que simula um gateway de pagamentos desenvolvida em Go, com o objetivo de servir como estudo e aprendizado da linguagem.

# Configuração do Ambiente

1. **Pré-requisitos**:

- Certifique-se de ter o Go instalado em sua máquina. Você pode baixá-lo em [golang.org](https://golang.org/).
- Tenha o `git` instalado para clonar o repositório.

2. **Clone o repositório**:

```bash
git clone https://github.com/lucasbarbosaalves/go-gateway.git
cd go-gateway
```

3. **Instale as dependências**:
   Execute o comando abaixo para instalar as dependências do projeto:

```bash
go mod tidy
```

Antes de executar o projeto, é necessário configurar o arquivo `.env` e preparar o ambiente Docker.

### Alterando o arquivo `.env`

1. Localize o arquivo `.env` na raiz do projeto.
2. Edite o arquivo para incluir as variáveis de ambiente necessárias, como credenciais de banco de dados, configurações de Kafka e outras informações relevantes. Um exemplo de configuração pode ser:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
KAFKA_BROKER=localhost:9092
```

Certifique-se de ajustar os valores conforme o ambiente que você está utilizando.

### Subindo o ambiente Docker

O projeto inclui um arquivo `docker-compose.yml` para facilitar a configuração de serviços como banco de dados e Kafka. Para subir o ambiente, siga os passos abaixo:

1. Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.
2. Execute o comando para iniciar os serviços:

```bash
docker-compose up -d
```

3. Verifique se os containers estão em execução:

```bash
docker ps
```

Os serviços necessários, como banco de dados e Kafka, devem estar listados.

### Configurações do Kafka

O projeto utiliza o Kafka para gerenciar mensagens. Certifique-se de que o broker Kafka está configurado corretamente no arquivo `.env`. Além disso, você pode utilizar ferramentas como o `kafka-topics.sh` para criar tópicos necessários, caso ainda não existam:

```bash
kafka-topics.sh --create --topic seu_topico --bootstrap-server localhost:9092
```

Com o ambiente configurado e os serviços em execução, você pode seguir os passos anteriores para iniciar o servidor e testar o gateway.

4. **Execute o projeto**:
   Inicie o servidor com o comando:

```bash
go run cmd/app/main.go
```
