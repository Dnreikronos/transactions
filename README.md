<h1>Transactions</h1>
<p>The <strong>Transactions</strong> project is a backend application developed in Go (Golang) that allows for the handling of financial transactions, currency value conversions based on exchange rates, and data persistence in an in-memory database.</p>

<h2>Features</h2>
<ul>
    <li><strong>Create and retrieve transaction data:</strong> The application allows creating and querying financial transactions through a RESTful API.</li>
    <li><strong>Currency conversion:</strong> The application uses the <a href="https://fiscaldata.treasury.gov/datasets/treasury-reporting-rates-exchange/treasury-reporting-rates-of-exchange" target="_blank">Treasury Reporting Rates API</a> to fetch exchange rates and perform value conversion based on the currency and transaction date.</li>
</ul>

<h2>Technology Stack</h2>
<ul>
    <li><strong>Go (Golang):</strong> The main programming language used for application development.</li>
    <li><strong>Gin:</strong> A web framework for Go, used for building the RESTful API.</li>
    <li><strong>Gorm:</strong> An ORM used for interacting with the database (in this case, an in-memory database with SQLMock for testing).</li>
    <li><strong>SQLMock:</strong> A library for unit testing that simulates interactions with the database.</li>
    <li><strong>Godotenv:</strong> For managing environment variables.</li>
    <li><strong>SQLite:</strong> The database used for data persistence in a simple development environment.</li>
    <li><strong>Makefile:</strong> For automating build, run, and test commands for the application.</li>
</ul>

<h2>API Features</h2>

<h3>1. Create Transaction</h3>
<p>The API allows creating financial transactions with the following parameters:</p>
<ul>
    <li><strong>id</strong>: Unique transaction ID</li>
    <li><strong>description</strong>: Transaction description</li>
    <li><strong>value</strong>: Transaction value</li>
    <li><strong>date</strong>: Transaction date</li>
    <li><strong>currency</strong>: Transaction currency</li>
</ul>

<h3>2. Query Transaction and Conversion</h3>
<p>The API allows querying a transaction and obtaining the value converted to a different currency. The endpoint expects the currency to be provided as a query parameter.</p>
<p><strong>Example:</strong> <code>/transaction/:id/currency?currency=Brazil-Real</code></p>

<h2>Folder Structure</h2>
<p>The project's folder structure is organized as follows:</p>
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

<h2>How to Run the Project</h2>

<h3>Requirements</h3>
<ul>
    <li>Go 1.18 or higher</li>
    <li>Install Go dependencies (run <code>go mod tidy</code>)</li>
</ul>

<h3>Commands</h3>
<ul>
    <li><strong>Build:</strong> Compiles the application.
        <pre>make build</pre>
    </li>
    <li><strong>Run the application:</strong> Starts the application.
        <pre>make run</pre>
    </li>
    <li><strong>Test handlers:</strong> Runs the tests for the API handlers.
        <pre>make testHandler</pre>
    </li>
    <li><strong>Test workers:</strong> Runs the tests for the worker logic.
        <pre>make testWorker</pre>
    </li>
    <li><strong>Test queue:</strong> Runs the tests for the memory queue.
        <pre>make testQueue</pre>
    </li>
</ul>

<h2>Tests</h2>
<p>The application includes unit tests for the handlers, workers, and queue. You can run them with the commands listed above, using the Makefile.</p>
