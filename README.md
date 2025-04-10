
# c3po

**c3po** é uma biblioteca de validação de structs em Go inspirada no Pydantic, com foco em performance, flexibilidade e facilidade de uso. Ideal para validações robustas em APIs web, sistemas embarcados e mais.

---

## Instalação

```bash
go get github.com/5tkgarage/c3po
```

---

## Exemplo Básico

```go
type User struct {
    Name string `c3po:"required"`
    Age  int    `c3po:"min=18"`
}

data := &User{Name: "Luke", Age: 17}
schema := c3po.ParseSchema(data)
res := schema.Decode(data)

if res.HasError() {
    fmt.Println(res.Errors())
}
```

---

## Tags Suportadas

| Tag        | Descrição                                  |
|------------|--------------------------------------------|
| `required` | Campo obrigatório                          |
| `min`      | Valor mínimo (para números, strings, etc.) |
| `max`      | Valor máximo                               |
| `in`       | Lista de valores aceitos                   |

---

## Tags Personalizadas

Você pode definir suas próprias tags com:

```go
c3po.ParseSchemaWithTag("chat", struct)
```

Isso permite utilizar a lib com frameworks como Fiber, Gin, Echo e outros, mantendo liberdade total nas tags de validação.

---

## Retorno de Erros

`Decode()` retorna um struct com:

- `ValidData()`: dados validados (com defaults aplicados)
- `Errors()`: mapa com os erros de validação encontrados

---

## Exemplo com Valor Padrão

```go
type Food struct {
    Name string `c3po:"required=false"`
}

schema := c3po.ParseSchema(&Food{Name: "fries"})
res := schema.Decode(map[string]any{}) // Name será "fries"
```

---

## Ideias Futuras

- Suporte a `enum` e validações encadeadas
- Validações condicionais entre campos
- Integração com serialização JSON nativa
- Geração automática de documentação a partir dos schemas

---

## Contribuição

Contribuições são bem-vindas! Abra uma issue, envie um PR ou mande uma ideia maluca — a casa é sua.

---

## Licença

MIT — use, modifique, compartilhe.