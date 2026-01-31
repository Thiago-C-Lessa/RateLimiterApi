*RequestLimiter*

Um rate limiter simples em Go baseado em Token Bucket, para limitar a quantidade de requests que uma API pode aceitar por segundo. Ideal para proteger endpoints de sobrecarga e evitar abuso de clientes.

Compatível com Chi Router.

**⚡ Funcionalidades**

- Limita a taxa de requests por segundo (RPS)

- Permite burst controlado (bucket cheio no início)

- Middleware HTTP pronto para usar

- Thread-safe (suporta múltiplas goroutines)

- Fácil de integrar com Chi Router ou outros routers


**🔄 Fluxo de Requests**

Fluxo visual do que acontece quando um request chega:

```txt
Request HTTP
     |
     v
[Chi Router]
     |
     v
[RateLimiter Middleware]
     |
     |---(Tokens disponíveis?)---YES---> Passa para handler -> 200 OK
     |
     |---NO---> Retorna 429 Too Many Requests


```

**Explicação:**

1. O request chega no Chi Router.

2. O middleware verifica se há tokens disponíveis no bucket.

3. Se houver token: consome 1 token e passa para o endpoint → 200 OK

4. Se não houver: request é rejeitado → 429 Too Many Requests

5. Os tokens são reabastecidos a cada segundo conforme a taxa configurada (RefillRate).

*Melhorias*
- Limitar requests por IP 