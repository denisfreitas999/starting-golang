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
3. [Convensões do Go](#03-convensões-do-go)

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


## 03. Convensões do Go

1. **Variáveis declaradas**

    Variáveis declaradas devem ser obrigatóriamente utilizadas, caso contrarário o haverá erro de compilação e execução do código. 

2. **Ponto e vígula**

    O uso do ponto e vírgula para indicar o término de expressões em Go é opcional, porém, é recomendado não utiliza-los. 

3. **Valores padrões**

    Por padrão, caso variáveis não sejam declaradas possuindo valores, o Go irá atribuir valores padrões para elas.

4. **Inferência de tipos**

    O Go consegue inferir os tipos das variáveis sem precisar declará-los, caso seja atribuido um valor a variável.
    ```go
    // var idade int = 27
    var idade = 27
    ```

5. **Inferência de variável**

    O Go consegue inferir uma variável sem precisar utilizar palavra reservada `var` antes do nome da variável. Basta alterar o sinal de atribuição `=` para `:=`

    ```go
    // var peso = 84
    peso := 84
    ```

5. **Inferência de Tipo com o `Scan`**

    Em Go, o compilador consegue inferir o tipo da variável na leitura de entrada, sem a necessidade de especificar explicitamente o tipo da variável a ser lida, desde que o tipo já tenha sido declarado anteriormente na declaração da variável. 

    Exemplo:
    ```go
    var comando int
    // Se já declaramos a variável `comando` como um inteiro, podemos usar `Scan` sem especificar o tipo novamente
    // fmt.Scanf("%d", &comando)
    fmt.Scan(&comando)
    ```

6. **Controle de Fluxo com `Switch`**

    Em Go, não é necessário utilizar a palavra-chave `break` dentro dos blocos `case` para evitar a execução dos casos subsequentes. O Go automaticamente encerra a execução após o primeiro `case` correspondente, tornando o código mais limpo e intuitivo.

    Exemplo:
    ```go
    switch comando {
        case 1:
            fmt.Println("Monitorando...")
        case 2:
            fmt.Println("Exibindo Logs...")
        case 0:
            fmt.Println("Saindo do Programa...")
        default:
            fmt.Println("Comando inválido! Escolha entre os valores 0, 1 ou 2 para acessar as opções do menu.")
    }
    ```

7. **Parênteses Delimitadores**

    Em Go, não é necessário utilizar parênteses para delimitar expressões em estruturas de controle como `if`, `for`, etc. O escopo da lógica é delimitado automaticamente pelas chaves `{}` que envolvem o bloco de código.

    Exemplo:
    ```go
    if comando == 1 {
        fmt.Println("Monitorando...")
    } else if comando == 2 {
        fmt.Println("Exibindo Logs...")
    } else if comando == 0 {
        fmt.Println("Saindo do Programa...")
    } else {
        fmt.Println("Comando inválido! Escolha entre os valores 0, 1 ou 2 para acessar as opções do menu.")
    }
    ```
8. **Saídas com o `os.Exit`**

    Em Go, a função `os.Exit` é usada para terminar a execução de um programa imediatamente. Ela aceita um argumento inteiro que é retornado ao sistema operacional como código de saída, onde `0` geralmente indica sucesso e qualquer outro valor indica uma falha ou encerramento anormal.

    Exemplo:
    ```go
    switch comando {
        case 1:
            fmt.Println("Monitorando...")
        case 2:
            fmt.Println("Exibindo Logs...")
        case 0:
            fmt.Println("Saindo do Programa...")
            os.Exit(0)  // Encerra o programa com sucesso
        default:
            fmt.Println("Comando inválido! Escolha entre os valores 0, 1 ou 2 para acessar as opções do menu.")
            os.Exit(-1) // Encerra o programa indicando um erro
    }
    ```

    No exemplo acima, `os.Exit(0)` é chamado para encerrar o programa normalmente quando o comando é `0`, enquanto `os.Exit(-1)` é usado para sinalizar um erro ao usuário em caso de comando inválido.

9. **Funções com Múltiplos Retornos e Operador em Branco `_`**

    Em Go, as funções podem retornar múltiplos valores. Isso é particularmente útil quando uma função precisa retornar um valor de sucesso e um valor de erro, ou outras combinações de resultados. Quando você não precisa de um dos valores retornados, o operador em branco (`_`) pode ser usado para ignorá-lo.

    **Exemplo de Função com Retorno Composto:**
    ```go
    func dividir(numerador, denominador int) (int, error) {
        if denominador == 0 {
            return 0, fmt.Errorf("divisão por zero")
        }
        return numerador / denominador, nil
    }

    func main() {
        resultado, err := dividir(10, 2)
        if err != nil {
            fmt.Println("Erro:", err)
        } else {
            fmt.Println("Resultado:", resultado)
        }
    }
    ```

    Neste exemplo, a função `dividir` retorna dois valores: o resultado da divisão e um valor de erro, que será `nil` se a operação for bem-sucedida.

    **Exemplo de Função com Retorno Composto Usando o Operador em Branco:**
    ```go
    func dividir(numerador, denominador int) (int, error) {
        if denominador == 0 {
            return 0, fmt.Errorf("divisão por zero")
        }
        return numerador / denominador, nil
    }

    func main() {
        resultado, _ := dividir(10, 2)  // Ignorando o valor de erro
        fmt.Println("Resultado:", resultado)
    }
    ```

    Neste segundo exemplo, usamos o operador em branco (`_`) para ignorar o valor de erro retornado pela função `dividir`, caso o usuário não queira lidar com possíveis erros.

10. **Looping Infinito com `for` e Inexistência do `while`**

    Em Go, não existe uma estrutura de loop `while`. Em vez disso, a linguagem utiliza o `for` para todos os tipos de loops, incluindo loops infinitos. Um loop infinito pode ser criado simplesmente omitindo todas as condições do `for`, fazendo com que ele execute indefinidamente até que seja explicitamente interrompido.

    ```go
    func main() {
        for {
            fmt.Println("Executando indefinidamente...")
            // O loop pode ser interrompido com um break ou return
        }
    }
    ```

    No exemplo acima, o loop `for` não tem condições, portanto, ele continua a executar indefinidamente. Para sair desse loop, você pode usar uma instrução `break` ou `return` dentro do bloco de código.

    Go adota essa abordagem simplificada para loops, substituindo a necessidade de um comando `while` separado.

11. **Array vs. Slices**

    Em Go, um **array** é uma estrutura de dados com um tamanho fixo, definido no momento da sua criação. Já o **slice** é uma abstração sobre o array, que permite ao desenvolvedor trabalhar com coleções de elementos de forma mais flexível, sem a necessidade de se preocupar com a capacidade subjacente.

    Um slice não armazena diretamente os dados, mas sim uma referência a um array. A principal vantagem do slice é que ele pode crescer ou diminuir dinamicamente, o que o torna muito mais conveniente de usar do que arrays fixos.

    **Exemplo de Uso de Slice:**
    ```go
    func exibeNomes() {
        nomes := []string{"Douglas", "Daniel", "Bernardo"}
        fmt.Println("O meu slice tem", len(nomes), "itens")
        fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

        nomes = append(nomes, "Aparecida")
        fmt.Println("O meu slice tem", len(nomes), "itens")
        fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
    }
    ```

    Neste exemplo, o slice `nomes` é inicializado com três elementos. Ao usar a função `append` para adicionar um novo nome, o slice é automaticamente redimensionado para acomodar o novo item. O comprimento (`len`) indica o número de elementos atuais no slice, enquanto a capacidade (`cap`) indica o tamanho máximo que o slice pode ter antes de precisar realocar a memória.

    Ao contrário dos arrays, os slices são muito mais flexíveis, permitindo ao desenvolvedor manipular coleções de dados sem precisar se preocupar manualmente com a alocação de memória ou o redimensionamento.

12. **Tratamento de Erros**

    Em Go, o tratamento de erros é feito de forma explícita, utilizando o valor de retorno das funções. A maioria das funções em Go retorna um valor adicional do tipo `error` para indicar se ocorreu algum problema durante a execução. É uma prática comum verificar imediatamente se ocorreu um erro e tratá-lo de maneira adequada.

    ```go
    func testaSite(site string) {
        resp, err := http.Get(site)
        
        // Tratamento de Erro
        if err != nil {
            fmt.Println("Erro detectado:", err)
            return  // Interrompe a execução da função em caso de erro
        }

        if resp.StatusCode == 200 {
            fmt.Println("Site:", site, "foi carregado com sucesso!")
        } else {
            fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
        }
    }
    ```

    Neste exemplo, a função `testaSite` faz uma requisição HTTP para o site fornecido e utiliza o padrão de tratamento de erros do Go. Se `http.Get` retornar um erro, ele é imediatamente tratado com uma mensagem informativa e a função é interrompida. Caso contrário, a função verifica o código de status da resposta (`StatusCode`) para determinar se o site foi carregado com sucesso ou se houve algum problema.

    Esse método de tratamento de erros, onde os erros são verificados e tratados logo após a ocorrência, ajuda a manter o código Go simples, claro e seguro.

13. **Leitura de Arquivos**

    Em Go, a leitura de arquivos pode ser feita utilizando funções da biblioteca padrão como `os.Open` para abrir o arquivo e `bufio.NewReader` para ler o conteúdo linha por linha. É importante sempre tratar erros durante a manipulação de arquivos para evitar falhas inesperadas.

    ```go
    func lerSitesDoArquivo() []string {
        var sites []string

        arquivo, err := os.Open("sites.txt")
        // Alternativamente, você pode usar: arquivo, err := os.ReadFile("sites.txt")

        if err != nil {
            fmt.Println("Erro detectado:", err)
            return sites // Retorna o slice vazio em caso de erro
        }

        leitor := bufio.NewReader(arquivo)
        
        for {
            linha, err := leitor.ReadString('\n')
            linha = strings.TrimSpace(linha)
            
            sites = append(sites, linha)

            if err == io.EOF {
                break
            }
        }

        arquivo.Close()

        return sites
    }
    ```

    Neste exemplo, a função `lerSitesDoArquivo` abre um arquivo chamado `sites.txt` e lê cada linha dele, armazenando os sites em um slice de strings. A leitura é feita utilizando `bufio.NewReader`, que permite ler o arquivo linha por linha até encontrar o final do arquivo (`io.EOF`). 

    Caso ocorra algum erro ao abrir o arquivo, ele é imediatamente tratado com uma mensagem de erro, e a função retorna um slice vazio. No final, o arquivo é fechado com `arquivo.Close()` para liberar os recursos do sistema.

    Essa abordagem é útil para processar arquivos grandes ou para quando se deseja manipular o conteúdo linha por linha.

14. **Escrita em Arquivos**

    Em Go, a escrita em arquivos pode ser realizada utilizando a função `os.OpenFile`, que permite abrir um arquivo com diferentes modos (leitura, escrita, criação, anexação, etc.). É importante sempre tratar possíveis erros ao abrir ou manipular arquivos e garantir que o arquivo seja fechado corretamente após a operação.

    ```go
    func registrarLog(site string, status bool) {
        arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

        if err != nil {    
            fmt.Println("Erro detectado:", err)
            return  // Encerra a função em caso de erro
        }

        // Escreve a data e hora atual, o site e o status online no arquivo
        _, err = arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + 
                                     " | Site: " + site + 
                                     " | Status Online: " + strconv.FormatBool(status) + "\n")

        if err != nil {
            fmt.Println("Erro ao escrever no arquivo:", err)
        }

        arquivo.Close()
    }
    ```

    Neste exemplo, a função `registrarLog` abre ou cria um arquivo chamado `logs.txt` no modo de leitura e escrita (`os.O_RDWR`), criando o arquivo se ele não existir (`os.O_CREATE`) e adicionando o conteúdo no final do arquivo (`os.O_APPEND`). O texto a ser escrito inclui a data e hora atual, o nome do site e o status de disponibilidade (`true` ou `false`).

    Se houver um erro ao abrir ou escrever no arquivo, o erro é tratado com uma mensagem informativa. Após a operação, o arquivo é fechado com `arquivo.Close()` para garantir que os dados sejam salvos corretamente e que os recursos do sistema sejam liberados.

    Esse método é útil para registrar logs ou qualquer tipo de histórico de operações em arquivos de texto.
