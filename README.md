<h1>Transactions</h1>
<p>O projeto <strong>Transactions</strong> é uma aplicação backend desenvolvida em Go (Golang) que permite manipular transações financeiras, calcular valores convertidos com base nas taxas de câmbio e realizar persistência de dados em um banco de dados em memória.</p>

<h2>Funcionalidades</h2>
<ul>
    <li><strong>Criar e recuperar dados de transações:</strong> A aplicação permite criar e consultar transações financeiras através de uma API RESTful.</li>
    <li><strong>Conversão de moeda:</strong> A aplicação usa a API do <a href="https://api.fiscaldata.treasury.gov" target="_blank">Treasury Reporting Rates API</a> para buscar taxas de câmbio e realizar a conversão de valores com base na moeda e data da transação.</li>
</ul>

<h2>Stack de Tecnologias</h2>
<ul>
    <li><strong>Go (Golang):</strong> Linguagem de programação principal utilizada para o desenvolvimento da aplicação.</li>
    <li><strong>Gin:</strong> Framework web para Go, utilizado para construção da API RESTful.</li>
    <li><strong>Gorm:</strong> ORM utilizado para interagir com o banco de dados (no caso, um banco de dados em memória com SQLMock para testes).</li>
    <li><strong>SQLMock:</strong> Biblioteca para testes unitários que simula interações com o banco de dados.</li>
    <li><strong>Godotenv:</strong> Para o gerenciamento de variáveis de ambiente.</li>
    <li><strong>SQLite:</strong> Banco de dados utilizado para persistência de dados em um ambiente de desenvolvimento simples.</li>
    <li><strong>Makefile:</strong> Para automatizar comandos de build, execução e testes da aplicação.</li>
</ul>

<h2>Funcionalidades da API</h2>

<h3>1. Criar Transação</h3>
<p>A API permite criar transações financeiras com os seguintes parâmetros:</p>
<ul>
    <li><strong>id</strong>: ID único da transação</li>
    <li><strong>description</strong>: Descrição da transação</li>
    <li><strong>value</strong>: Valor da transação</li>
    <li><strong>date</strong>: Data da transação</li>
    <li><strong>currency</strong>: Moeda da transação</li>
</ul>

<h3>2. Consultar Transação e Conversão</h3>
<p>A API permite consultar uma transação e obter o valor convertido para uma moeda diferente. O endpoint espera que a moeda seja fornecida como um parâmetro de consulta.</p>
<p><strong>Exemplo:</strong> <code>/transaction/:id/currency?currency=Brazil-Real</code></p>

<h2>Estrutura de Pastas</h2>
<p>A estrutura de pastas do projeto é organizada da seguinte forma:</p>
<pre>
/cmd
    /main.go                  
/models
    /transactions.go           
/handlers
    /transaction_handler.go    
/utils
    /utils.go                 
/worker
    /memory_worker.go          
/queue
    /queue.go                  
/tests
    /handler_test.go           
    /worker_test.go            
    /queue_test.go             
/Makefile                     
</pre>

<h2>Como Rodar o Projeto</h2>

<h3>Requisitos</h3>
<ul>
    <li>Go 1.18 ou superior</li>
    <li>Instalação de dependências do Go (executar <code>go mod tidy</code>)</li>
</ul>

<h3>Comandos</h3>
<ul>
    <li><strong>Build:</strong> Compila a aplicação.
        <pre>make build</pre>
    </li>
    <li><strong>Rodar a aplicação:</strong> Inicia a aplicação.
        <pre>make run</pre>
    </li>
    <li><strong>Testar manipuladores:</strong> Executa os testes dos manipuladores da API.
        <pre>make testHandler</pre>
    </li>
    <li><strong>Testar workers:</strong> Executa os testes para a lógica de workers.
        <pre>make testWorker</pre>
    </li>
    <li><strong>Testar fila:</strong> Executa os testes para a fila de memória.
        <pre>make testQueue</pre>
    </li>
</ul>

<h2>Testes</h2>
<p>A aplicação conta com testes unitários para os manipuladores, workers e fila. Você pode executá-los com os comandos listados acima, através do Makefile.</p>

</body>
</html>
