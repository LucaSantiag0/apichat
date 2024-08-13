Este projeto envolve a criação de uma API robusta utilizando a linguagem Go, conhecida por sua eficiência e desempenho. 
A API está equipada para lidar com diversas funcionalidades essenciais para aplicações modernas, como roteamento, middleware, suporte a WebSockets, gerenciamento de UUIDs, e interação com banco de dados PostgreSQL.
O projeto envolve uma sala de com ip independete e interativo, onde você pode abrir e fazer uma pergunta por exemplo, e ser respondido igualmente em tempo real, além de, compartilhar a sala atraves do link.


Dependências Principais
go-chi/chi: Utilizado para roteamento, chi é um pacote leve e expressivo que facilita a criação de rotas HTTP em Go.

go-chi/middleware: Este pacote oferece middlewares pré-construídos que podem ser facilmente integrados ao projeto para tarefas comuns, como logging, recuperação de pânico e compressão de resposta.

go-chi/cors: Pacote para configuração e gerenciamento de políticas de CORS (Cross-Origin Resource Sharing), essencial para a segurança e comunicação entre diferentes origens.

google/uuid: Utilizado para gerar UUIDs, este pacote é fundamental para criar identificadores únicos em várias partes da aplicação.

gorilla/websocket: Implementa WebSockets, permitindo comunicação em tempo real entre o cliente e o servidor.

jackc/pgx: Um driver PostgreSQL de alta performance e interface de conexão que é usado para a interação com o banco de dados.
