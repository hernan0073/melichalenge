Voy a contarte lo que hice con la tecnología que puedo compartir en pantalla:

MiniKube, Grafana, Prometheus, una aplicación hecha en GO, y DockerHub como Container Registry.

Descripción del Proxy

 Implementé un proxy que reenvía solicitudes al dominio api.mercadolibre.com. El proxy está diseñado para permitir el acceso a los endpoints de MercadoLibre sin redirección ni almacenamiento en caché. Las solicitudes se envían directamente al API de MercadoLibre y la respuesta vuelve al cliente sin alteraciones.

Funcionalidad del Proxy

Reenvío de Solicitudes: El proxy se encarga de manejar las solicitudes que llegan y las dirige al dominio de MercadoLibre. No hago redirección ni almacenamiento en caché; las solicitudes y respuestas se transmiten tal como llegan.

Control de Llamadas (Rate Limiting)

Implementé un sistema básico de rate limiting para controlar cuántas solicitudes se permiten:

Configuración por IP y Path:
IP  Permite hasta 1000 solicitudes por minuto.
Path /categories/*: Permite hasta 10,000 solicitudes por minuto.
IP  y Path /items/*: Permite hasta 10 solicitudes por minuto.
Este sistema asegura que se respeten los límites de llamadas establecidos para distintas combinaciones de IP y rutas.

Escalabilidad (Script.js Use K6 para testear y simular)
Para manejar una carga de más de 50,000 solicitudes por segundo, hay que tener en cuenta lo siguiente:

Optimización del Código: el código no esta tan optimizado para manejar un gran volumen de solicitudes concurrentes, le pase el corrector de ChatGpt pero Si le asigno mas tiempo lo puedo mejorar.

Configuración del Clúster: Escalé el número de réplicas del proxy en el clúster de Kubernetes y ajusta los recursos cuando el HPA llego a 80%. Se tendria mas tiempo utilizaria custom-metrics.



