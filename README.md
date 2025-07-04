# redis-api

Aufgabe: Erstelle zwei namespaces in kubernetes. In dem einem namespace soll ein
Deployment von redis mit 3 replicas vorhanden sein. Dabei sollen die Daten FLÜCHTIG gespeichert
werden. In einem anderen Namespace soll ebenfalls ein Deployment erstellt werden
mit drei Replicas basierend auf dem image `docker.io/naivary/redis-api:latest`. Das Deployment muss die folgenden Punkte enthalten:
1. Konfiguration einer liveness probe
2. Konfiguration einer readiness probe
3. Konfiguration der redis-url unter welchen der redis server in dem ersten
   Namespace erreichbar ist (für das Format der URL siehe [env-vars](#env-vars))

Nach dem Deployment sollte es möglich sein über einen NodePort Service die
redis-api anzusprechen und die abfrage
`http://<node-api>:<node-port>?key=teccle` erfolgreich durchzuführen.

## livez

Der Endpoint livez ist erreichbar mit einer GET-Methode unter dem root pfad e.g.
`localhost:8080/livez`

## readyz

Der Endpoint readyz ist erreichbar mit einer GET-Methode unter dem root pfad e.g.
`localhost:8080/readyz`


## get

get erlaubt es einen gespeichertern Key aus dem Redis-Store zu erhalten. Es
werden zwei keys gespeichert: `etomer` und `teccle`. Diese kann man wie folgt
erhalten: `localhost:8080?key=etomer` oder `localhost:8080?key=teccle`.

## env-vars

Die folgenden env variables müssen gesetzt werden

1. APP_PORT: port des redis-apis
2. APP_HOST: host des redis-apis
3. REDIS_URL: url, welche genutzt werden kann um sich mit den redis server zu
   verbinden e.g. redis://@<adresse>:6379


