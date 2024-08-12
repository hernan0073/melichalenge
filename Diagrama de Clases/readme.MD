Clases:
------------
RateLimiter
------------
Atributos:
limiter: rate.Limiter - Controla la tasa de solicitudes permitidas.
lastSeen: time.Time - Registra la última vez que se usó el limitador.

Métodos:
No tiene métodos explícitos definidos dentro de la clase, pero es utilizado en otros métodos como getRateLimiter.

----------------
RateLimitConfig
----------------

Atributos:
Rate: int - Define el número de solicitudes permitidas por minuto.
Burst: int - Define el número de solicitudes que pueden ser procesadas en ráfaga.
Métodos:
No tiene métodos explícitos definidos.

-------------------
Funciones Globales
-------------------

setRateLimit:
Parámetros: http.ResponseWriter, *http.Request
Actualiza los límites de tasa globales (defaultRateLimit, defaultBurstLimit) y reinicia los visitantes.

getRateLimiter:
Parámetros: ip string, path string
Retorna un rate.Limiter específico para una combinación de IP y path.

rateLimiterMiddleware:
Parámetros: next http.Handler
Middleware que aplica el limitador de tasa antes de procesar una solicitud.

proxyHandler:
Parámetros: http.ResponseWriter, *http.Request
Maneja las solicitudes proxy hacia la API de MercadoLibre.

**********
Relaciones:
**********

RateLimiter usa la librería rate para crear limitadores y gestionar las tasas de solicitud.
RateLimitConfig se usa para configurar los valores de tasa y ráfaga.
RateLimiter se almacena en un mapa visitors para rastrear las limitaciones por IP y path.
Funciones interactúan con estructuras de datos globales como visitors y client.