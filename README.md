# Desafio Final - Proposta de arquitetura utilizando golang
Pós Arquitetura de Software e Soluções com IA - Desafio Final


*estrutura de pastas seguindo o padrão MVC*
```bash
product-api/
│
├── main.go                        // Ponto de entrada da aplicação
│
├── config/
│   └── database.go                // Inicialização do banco de dados
│
├── controller/
│   └── product_controller.go      // Controllers: manipulam as requisições
│
├── service/
│   └── product_service.go         // Regras de negócio
│
├── model/
│   └── product.go                 // Modelo de dados (entidade Produto)
│
├── repository/
│   └── product_repository.go      // Acesso a dados (CRUD com o banco)
│
└── router/
    └── router.go                  // Define os endpoints e rotas
```


*UML Class Diagram*

* Product                   *--> como entidade de domínio (Model)*
* ProductRepository         *--> como camada de persistência*
* ProductService            *--> com regras de negócio*
* ProductController         *--> responsável pelas requisições*
* product_router e main.go  *-->como bootstrapping da aplicação*


![alt](./imgs/UML.png)

*Testes*

* Criar um produto `*(POST /products)*`

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mouse Gamer",
    "description": "Mouse com 8000 DPI",
    "price": 149.99
}'
```

* Buscar todos os produtos `*(GET /products)*` 

```bash
curl http://localhost:8080/products
```

* Buscar produto por ID `*(GET /products/{id})*`

```bash
curl http://localhost:8080/products/1
```

* Buscar produtos por nome `*(GET /products/nome/{name})*`

```bash
curl http://localhost:8080/products/name/Mouse
```

* Atualizar um produto `*(PUT /products/{id})*`

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mouse Gamer RGB",
    "description": "Mouse com RGB e 12000 DPI",
    "price": 199.99
}'
```

* Contar produtos `*(GET /products/count)*`

```bash
curl http://localhost:8080/products/count
```

* Deletar um produto `*(DELETE /products/{id})*`

```bash
curl -X DELETE http://localhost:8080/products/1
```

