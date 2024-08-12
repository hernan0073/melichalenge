Proxy API Documentation
Este proyecto implementa un proxy que reenvía solicitudes al dominio api.mercadolibre.com. El proxy maneja solicitudes HTTP y las redirige al API de MercadoLibre, respetando límites de tasa (rate limiting) configurables.

Endpoints
POST /rate-limit
Configura los límites de tasa para el proxy.

Request Body:

json
Copiar código
{
  "rate": 1000,
  "burst": 5
}
rate: Número máximo de solicitudes permitidas por minuto.
burst: Número máximo de solicitudes permitidas en ráfaga.
Response:

Status 200 OK: Confirmación de que el límite de tasa ha sido configurado.
Status 400 Bad Request: Si el formato del cuerpo de la solicitud es inválido.
Example:

bash
Copiar código
curl -X POST http://localhost:8080/rate-limit -H "Content-Type: application/json" -d '{"rate": 1000, "burst": 5}'
GET /metrics
Exposición de métricas para Prometheus.

Response:

Status 200 OK: Devuelve métricas en formato Prometheus.
Example:

bash
Copiar código
curl http://localhost:8080/metrics
GET /{path}
Reenvía solicitudes al dominio https://api.mercadolibre.com.

Path Parameters:

{path}: El endpoint de la API de MercadoLibre que se desea consultar.
Response:

Devuelve la respuesta del API de MercadoLibre sin redirección ni caché.
Example:

bash
Copiar código
curl http://localhost:8080/categories/MLA97994
Rate Limiting
El proxy implementa un sistema básico de rate limiting que se configura por IP y por ruta:

IP 152.152.152.152: Se permite un máximo de 1000 solicitudes por minuto.
Path /categories/*: Se permite un máximo de 10,000 solicitudes por minuto.
IP 152.152.152.152 y Path /items/*: Se permite un máximo de 10 solicitudes por minuto.
CORS Configuration
El proxy está configurado para permitir solicitudes desde http://localhost:3000 con los siguientes métodos HTTP:

GET
POST
OPTIONS
Y los siguientes encabezados permitidos:

Content-Type
Authorization
Error Handling
Rate Limit Exceeded: Si se supera el límite de tasa configurado, se devuelve un error 429 Too Many Requests.
Internal Server Error: Si ocurre un error al procesar la solicitud o al comunicarse con el API de MercadoLibre, se devuelve un error 500 Internal Server Error.
Dependencies
Gorilla Mux: Para el enrutamiento de solicitudes.
Prometheus Go Client: Para la exposición de métricas.
Gorilla Handlers: Para la configuración de CORS.
Rate Limiter: Para la implementación del control de tasa.
Running the Application
Asegúrate de tener Go instalado.
Clona el repositorio.
Instala las dependencias con go mod tidy.
Ejecuta la aplicación con go run main.go.
La aplicación se ejecutará en el puerto 8080.