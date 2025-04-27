# C3PO  
**GoLang data validator – simples, flexível e extensível.**

C3PO é um validador de dados rápido e minimalista para Go.  
Ele é pensado para ser leve, intuitivo e fácil de extender — sem refletividade pesada ou configurações complexas.

## Features ✨
- 🔥 Validação por tags (`required`, `min`, `max`,`minlen`)
- ⚡ Sem mágica: fácil de entender e debuggar
- 🛠️ Extensível: adicione suas próprias validações
- 🏎️ Alto desempenho: ideal para aplicações críticas

## Instalação
```bash
go get 5tk.dev/c3po
```

## Exemplo rápido
```go
package main

import (
    "fmt"
    "5tk.dev/c3po"
)

type User struct {
    Name string `validate:"required"`
    Age  int    `validate:"min=18"`
}

func main() {
    user := &User{}
    sch := c3po.Validate(user,map[string]any{"name": "cleitu", "age": "15"})
    if sch.HasErrors() {
        panic(sch.Errors())
    }
    u := sch.Value().(*User)
    fmt.Println(u) 
}
```

## Validações suportadas
| Tag       | Descrição                  |
|-----------|----------------------------|
| `required`| Campo obrigatório          |
| `min`     | Valor mínimo (número)      |
| `max`     | Valor máximo (número)      |
| `minlen`  | Valor máximo (tamanho)     |
| `maxlen`  | Valor máximo (tamanho)     |
| `escape`  | Html Escape     |

## Extensões e validações customizadas
Crie novas tags facilmente:
```go
c3po.Register("now", func(field reflect.Value, param string) error {
    field.Set(reflect.ValueOf(time.Now()))
    return nil
})
```

## Roadmap 🚀
- [x] Sistema de validação básico (`required`, `min`, `max`)
- [ ] Middleware de validação para `http.Request`
- [ ] Diretório de exemplos
- [ ] Documentação completa
- [ ] Benchmarks

## Contribuindo
Pull requests são bem-vindos!  
Se encontrar algum bug ou tiver ideias de melhoria, abra uma [issue](https://github.com/5tk-dev/c3po/issues).
