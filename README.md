# EuLevo-Delivery

Eu-Levo Delivery applicativo de entregas

- [Tecnologias](#Tecnologias)
- [Autor](#Autor)
- [Rodando o app](#Rodando-app)
    - [Simulator](#Simulator)
    - [Kafka](#kafka)




## Tecnologias

- Simulator: 
    - Golang 
    - Docker
- Backend:
    - Nest.js
    - Mongo
    - Docker
- Frontend: 
    - React
- Messageria & filas
    - kafka
    - kafka connect
- Container
    - Docker
    - Kubernets
- Rastreabilidade & Relatorios
    - Isto
    - Kiali
    - Prometheus
    - Grafana

    
- Banco de dados
    - Mongo
    - Elastic


## Rodando app

### Simulator

No terminal rodar o comando do docker, lembrando que a maquina terá que ter o docker instalado

```bash
docker exec -it simulator bash

docker-compose up -d

go run main.go
```

### kafka

```bash
Consumindo msg no topico

Entra no container do kafka

docker exec -it kafkaconfig_kafka_1 bash

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-direction

Produzinho msg para o topico
...

kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-position

Consumidor que fica ouvindo as mensagens do produtor

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal

```
## Autor

- [@Samuel Jose Souza Oliveira](https://www.linkedin.com/in/samuel-jose-souza-oliveira-81b1177b/)