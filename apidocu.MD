POST /rate-limit

Configura los límites de tasa del proxy.
Cuerpo de la solicitud:
JSON
{
  "rate": 1000, // Solicitudes máximas por minuto
  "burst": 5 // Ráfaga máxima de solicitudes
}


¡Claro! Aquí tienes una versión del README en español, con algunas mejoras adicionales:

README.md en Español: Proxy API para MercadoLibre
Descripción
Este proyecto implementa un proxy que actúa como intermediario entre tus aplicaciones y la API de MercadoLibre (api.mercadolibre.com). El proxy se encarga de reenviar las solicitudes a la API de MercadoLibre y de administrar los límites de tasa, asegurando que no se excedan las restricciones impuestas por la API.

Funcionalidades
Limitación de Tasa: Permite configurar límites de tasa por IP y por ruta para evitar sobrecargar la API de MercadoLibre.
Métricas: Exporta métricas en formato Prometheus para monitorear el rendimiento del proxy.
CORS: Soporta solicitudes CORS desde dominios específicos.
Manejo de Errores: Maneja errores comunes como límites de tasa excedidos y errores internos del servidor.
Endpoints
Configuración
POST /rate-limit
Configura los límites de tasa del proxy.
Cuerpo de la solicitud:
JSON
{
  "rate": 1000, // Solicitudes máximas por minuto
  "burst": 5 // Ráfaga máxima de solicitudes
}
Usa el código con precaución.

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