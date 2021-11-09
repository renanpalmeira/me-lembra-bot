# :robot: me-lembra-bot

Estudo de caso sobre chatbot para o Curso de Graduação em Engenharia de Computação (projeto de TCC)

URL de produção: https://me-lembra-bot.herokuapp.com/

## Requisitos

- [Golang 1.7](https://golang.org/doc/go1.7)
- Conta + chaves https://www.twilio.com/
- [Node 14.16.0](https://github.com/nodejs/node/blob/master/doc/changelogs/CHANGELOG_V14.md#14.16.0)

## CI/CD

- A cada nova alteração na main vamos um deploy no Heroku
- A cada novo commit/push realizamos os lints:
  - Linguagem Go https://golangci-lint.run/
  - Dockerfile https://hadolint.github.io/hadolint/
  - Scan de segurança das dependências https://github.com/sonatype-nexus-community/nancy

## Ciclo de desenvolvimento

### Local

```sh
cp .env.sample .env # Create .env file
go run cmd/api/main.go
```

## Algumas telas

### O chat
![print 1](.github/images/1.png?raw=true)

### SMS chegando no celular
![print 2](.github/images/2.png?raw=true)
