
POST /rate-limit

Configura los límites de tasa del proxy.
Cuerpo de la solicitud:

JSON
{
  "rate": 1000, // Solicitudes máximas por minuto
  "burst": 5 // Ráfaga máxima de solicitudes
}


Respuesta:

200 OK: Configuración exitosa.
400 Bad Request: Cuerpo de la solicitud inválido.

Métricas

GET /metrics

Exporta métricas en formato Prometheus.
Proxy

GET /{path}
Reenvía la solicitud a https://api.mercadolibre.com/{path}.

Manejo de Errores

429 Too Many Requests: Se devuelve cuando se excede el límite de tasa.
500 Internal Server Error: Se devuelve en caso de errores internos o fallos de comunicación con la API de MercadoLibre.

Dependencias

Gorilla Mux: Para el enrutamiento de las solicitudes.
Prometheus Go Client: Para exportar métricas.
Gorilla Handlers: Para configurar CORS.
Rate Limiter: Para implementar la limitación de tasa.