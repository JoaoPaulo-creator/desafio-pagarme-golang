## transactions

- infos da compra
- dados do cartao
- valor

## playable

- recebiveis pagos ao cliente

# Um exemplo de transacao

```json
{
  "value": 2342,
  "productDescription": "Iphone 14 128gb",
  "paymentMethod": "debit_card ou credit_card",
  "cardNumber": "123456789098765",
  "nameInCard": "JOAO P S MOREIRA",
  "cardExpirationDate": "04-27",
  "cvv": "123"
}
```

## Pre requisitos

- O numero do cartao deve ter uma mascara, e retornar somente os ultimos 4 digitos como visivel

A transacao devera criar um payable, e o payable devera ter as seguintes propriedades e requisitos

## payable

```json
{
  "status": "paid se for pago com cartao de debito - waiting_funds se for com credito",
  "payment_date": "01-01-2024", //(payment_date) = data da criação da transação (D+0). caso credito, senao payment_date) = data da criação da transação + 30 dias (D+30). para credito.
  "fee": "taxa de processamento. Desconta 3% para compras em debito, 5% para credito",
  "transactions": [
    {
      "value": 2342,
      "productDescription": "Iphone 14 128gb",
      "paymentMethod": "debit_card ou credit_card",
      "cardNumber": "123456789098765",
      "nameInCard": "JOAO P S MOREIRA",
      "cardExpirationDate": "04-27",
      "cvv": "123"
    }
  ]
}
```

"O serviço deve prover um meio de consulta para que o cliente visualize seu saldo com as seguintes informações:
Saldo available (disponível): tudo que o cliente já recebeu (payables paid)
Saldo waiting_funds (a receber): tudo que o cliente tem a receber (payables waiting_funds)"

no caso acima, acredito ser possivel fazer via query, onde em um dado endpoint seja possivel trazer essas infos
