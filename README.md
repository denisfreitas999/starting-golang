# Liguagem Go

A linguagem Go, também conhecida como GoLang, é uma linguagem open source que foi criada pelo Google em 2007, e desde então é utilizada para a construção de produtos e serviços de grande escala. Atualmente a linguagem é utilizada por diversas empresas, como Uber, Twitch, Medium e Mercado livre.

Go é uma linguagem simples e produtiva de se utilizar, com foco no desenvolvimento de aplicações que necessitam de alta performance. Embora tenha sido criada para lidar com sistemas de redes e infraestrutura, Go também é bastante utilizada no mercado para:

Desenvolvimento de aplicações server-side e hospedadas em ambientes cloud;
Construção de scripts e ferramentas de automações utilizadas por times DevOps;
Construção de ferramentas de linha de comando;
Soluções de inteligência artifical e data science.

Fonte: [Formação Go](https://cursos.alura.com.br/formacao-go)

## Cursos
#### Curso 01 - [Go: a linguagem do Google](https://www.alura.com.br/curso-online-golang)

## Indices
1. [As 25 palavras reservadas do Go](#01-as-25-palavras-reservadas-do-go)
2. [Comandos de execução](#02-comandos-de-execução)


## 01. [As 25 palavras reservadas do Go](https://vinniciusgomes.dev/articles/as-25-palavras-reservadas-do-go)

### Declaração

As palavras reservadas nesta categoria são usadas para declarar diferentes elementos de código.

1. **const**

    A palavra-chave `const` é usada para declarar constantes, que são valores imutáveis conhecidos em tempo de compilação.
    ```go
    const Pi = 3.14
    ```
2. **var**

    A palavra-chave `var` é usada para declarar variáveis.
    ```go
    var age int
    ```
3. **func**

    A palavra-chave `func` é usada para declarar funções.
    ```go
    func add(a int, b int) int {
        return a + b
    }
    ```
4. **type**

    A palavra-chave `type` é usada para introduzir novos tipos, como structs.
    ```go
    type Person struct {
        Name string
        Age  int
    }
    ```
5. **import**

    A palavra-chave `import` é usada para importar pacotes.

    ```go
    import "fmt"
    ```
6. **package**

    A palavra-chave `package` define o pacote ao qual o código pertence.

    ```go
    package main
    ```
### Tipos compostos
Essas palavras-chave são usadas para definir tipos compostos.

7. **chan**

    A palavra-chave `chan` é usada para definir canais, que permitem comunicação entre goroutines.
    ```go
    ch := make(chan int)
    ```
8. **interface**

    A palavra-chave `interface` é usada para definir um conjunto de métodos que um tipo deve implementar.
    ```go
    type Speaker interface {
        Speak() string
    }
    ```
9. **map**

    A palavra-chave map define um tipo de mapa, uma coleção não ordenada de pares chave-valor.
    ```go
    m := make(map[string]int)
    ```
10. **struct**
    A palavra-chave `struct` é uma coleção de campos. Podemos usar a palavra-chave `struct` seguida pelas declarações de campo.
    ```go
    type Rectangle struct {
        Width, Height int
    }
    ```
### Controle de fluxo

Essas palavras-chave controlam o fluxo do código.

11. **break**

    A palavra-chave break interrompe a execução de loops ou switch.
    ```go
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
    }
    ```
12. **case**

    A palavra-chave `case` é usada em instruções switch.

    ```go
    switch day {
        case "Monday":
            ...
        }
    }
    ```
13. **continue**

    A palavra-chave continue pula para a próxima iteração do loop.

    ```go
    switch day {
        case "Monday":
            ...
        }
    }
    ```

14. **default**

    A palavra-chave default especifica o bloco de código a ser executado se nenhum case for correspondido.

    ```go
    switch day {
    default:
        fmt.Println("Dia desconhecido")
    }
    ```

15. **if**

    A palavra-chave if é usada para executar código condicionalmente.
    ```go
    if x > 0 {
        fmt.Println("Positivo")
    }
    ```
16. **else**
    
    A palavra-chave `else` define um bloco alternativo para uma instrução if.
    ```go
    if x < 0 {
        fmt.Println("Número negativo")
    } else {
        fmt.Println("Número não negativo")
    }
    ```

17. **fallthrough**

    A palavra-chave fallthrough transfere a execução para o próximo case em uma instrução switch.
    ```go
    switch day {
    case "Monday":
        fmt.Println("Segunda-feira")
        fallthrough
    case "Tuesday":
        fmt.Println("Terça-feira")
    }
    ```

18. **for**

    A palavra-chave for é usada para iniciar um loop.

    ```go
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    ```

19. **goto**

    A palavra-chave `goto` permite saltar para um rótulo especificado.

    ```go
    func example() {
    if x > 0 {
        goto Label
    }
        Label:
        fmt.Println("Label reached")
    }
    ```

20. **range**

    A palavra-chave `range` é usada para iterar sobre arrays, slices, mapas e canais.

    ```go
    for index, value := range slice {
        fmt.Println(index, value)
    }
    ```
21. **select**

    A palavra-chave `select` permite que uma goroutine espere em várias operações de comunicação.

    ```go
    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
    ```
22. **switch**

    A palavra-chave switch permite executar um bloco de código entre várias opções.

    ```go
    switch day {
    case "Monday":
        fmt.Println("Segunda-feira")
    }
    ```
### Modificador de função

Essas palavras-chave controlam a execução de chamadas de função de maneiras específicas.

23. **defer**

    A palavra-chave `defer` adia a execução de uma função até que a função circundante retorne.

    ```go
    func main() {
        defer fmt.Println("Executado por último")
        fmt.Println("Executado primeiro")
    }
    ```
24. **go**

    A palavra-chave `go` inicia uma nova goroutine.

    ```go
    go func() {
        fmt.Println("Executando em uma goroutine")
    }()
    ```

## 02. Comandos de execução

### `go build nome_do_arquivo.go`

O comando `go build nome_do_arquivo.go` é utilizado para compilar o código fonte de um arquivo Go. Ele gera um executável a partir do código especificado no arquivo `nome_do_arquivo.go`. Este executável pode ser executado de forma independente, sem a necessidade do ambiente Go instalado. Se nenhum erro for encontrado durante a compilação, o arquivo executável será criado no diretório atual.

### `go run nome_do_arquivo.go`

O comando `go run nome_do_arquivo.go` compila e executa o código fonte Go em uma única etapa. Ao contrário de `go build`, o `go run` não gera um arquivo executável no disco; em vez disso, ele compila o código em tempo real e executa o programa imediatamente. Esse comando é útil durante o desenvolvimento para testar rapidamente o código sem precisar gerar um executável separado.




