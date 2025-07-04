# redis-api

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


