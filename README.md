# Loja

Esta é uma aplicação de exemplo chamada "Loja" desenvolvida em Go (Golang) para gerenciar produtos em um banco de dados PostgreSQL utilizando Docker.

## Pré-requisitos

Antes de executar a aplicação, certifique-se de ter os seguintes pré-requisitos instalados em seu sistema:

- Docker: https://docs.docker.com/get-docker/
- Docker Compose: https://docs.docker.com/compose/install/

## Executando a aplicação

Siga os passos abaixo para executar a aplicação:

1. Clone este repositório: https://github.com/DjonatasBorges/loja

2. Abra o projeto em sua raiz

3. Inicie os containers do Docker: ```docker-compose up -d```

4. Acesse a aplicação em seu navegador: http://localhost:8000/produtos

Pronto agora é só usar


## Funcionalidades

- A aplicação permite visualizar a lista de produtos cadastrados.
- É possível criar novos produtos com nome, descrição, preço e quantidade.
- Os produtos podem ser editados, e suas informações podem ser atualizadas.
- Os produtos também podem ser deletados do banco de dados.

## Estrutura do Projeto

- `controllers`: Contém os controladores (handlers) HTTP que processam as requisições e gerenciam as respostas.
- `db`: Responsável por configurar e conectar-se ao banco de dados PostgreSQL.
- `models`: Contém os modelos de dados da aplicação.
- `templates`: Diretório para armazenar os arquivos de template HTML utilizados para renderizar as páginas.

## Autores

- Autor: Djonatas Camillo Borges
- GitHub: https://github.com/DjonatasBorges

`Observação:` Este é um projeto de adaptação da Alura Loja, onde foram adicionadas funcionalidades extras e melhorias no código. Além disso, a parte de integração com Docker foi criada exclusivamente por mim.

