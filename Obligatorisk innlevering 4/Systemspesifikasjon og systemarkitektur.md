Beskrivelse:
Det er tiltenkt at denne web-tjenesten skal fungere som et planleggingsverktøy for innbyggere i Kristiansand. Tjenestens funksjonalitet går ut på at et Golang-program henter værdata fra https://openweathermap.org via API, og parser dette gjennom til en webside. 

OpenWeatherMap samler data fra sensorer og andre datakilder, både fra profesjonelle og private aktører. Flesteparten av sensorene er installert på flyplasser, i store byer osv., og det finnes værstasjoner på totalt 200 000 lokasjoner rundt om i verden.
En underliggende logikk forteller brukeren hva slags type klær som egner seg basert på temperaturen i Kristiansand på nåværende tidspunkt, og kommer i tillegg med forslag til steder hvor aktiviteter kan bedrives, også basert på været/temperaturen.
Disse forslagene vil presenteres som pins på et Google Maps-kart på samme side. 
Tjenesten vil i utgangspunktet komme med relativt enkle anbefalinger, men kan enkelt utvides til å passe mer avanserte værforhold. Tjenesten utvikles som en web-tjeneste, som gjør at den også enkelt kan porteres over til for eksempel mobile operativsystemer (iOS og Android) via såkalte "web views".

Arkitektur:
Nodene i vår systemarkitektur er en klient som kjører applikasjonen vår. Denne applikasjonen starter en webserver som lytter på port ":XXXX". Webserveren lager og sender en API-request til openweathermap.org for å hente data i JSON-form.
