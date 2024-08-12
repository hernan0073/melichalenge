Descripción del Diagrama AWS

Voy a contarte cómo me imagino el setup en AWS:

Amazon API Gateway: Este es el punto de entrada para todas las solicitudes de los clientes. Se encarga de redirigir esas peticiones a los servicios backend adecuados.

Amazon RDS o Amazon DynamoDB: Para almacenar métricas y logs, uso RDS si necesito una base de datos SQL o DynamoDB si me inclino por NoSQL.

Amazon CloudWatch: Es el servicio que utilizo para recolectar y visualizar logs y métricas. Es fundamental para monitorear lo que está pasando.

Amazon EKS: Aquí es donde corre el proxy y otros servicios en un clúster de Kubernetes administrado por AWS. Me da flexibilidad y control sobre los contenedores.

Amazon Elastic Load Balancer (ELB): Este distribuye el tráfico de red entre varias instancias de EC2, ayudando a balancear la carga y mejorar la disponibilidad.

Amazon CloudFront: Uso esta CDN para mejorar el rendimiento de la entrega de contenido y reducir la latencia.

Observabilidad: Para mantener todo bajo control, aplico un módulo de APM como Datadog o NewRelic y monitoreo todo desde allí.

