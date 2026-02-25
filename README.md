# Country Info Service



**Course: PROG2005 - Cloud Technologies, 2026**

**Author: Isak Olai Varjord – isakov@stud.ntnu.no**



This service provides information about countries and currency exchange rates for neighboring countries. It fetches data from two external APIs (Rest Countries and Currency API) and combines them into a simple interface. You get status information about the external services, general country information, and exchange rates from a country's base currency to the currencies of its neighboring countries.



#### Endpoints:



#### 1\. Status – GET /countryinfo/v1/status/

**Returns HTTP status codes for the two external APIs this service depends on, as well as how long the service has been running.**



Example:

http://localhost:8080/countryinfo/v1/status/



Response:



json

{

&nbsp; "restcountriesapi": 200,

&nbsp; "currenciesapi": 200,

&nbsp; "version": "v1",

&nbsp; "uptime": 42

}

uptime is in seconds since the last restart.



#### 2\. Country Info – GET /countryinfo/v1/info/{code}

**Retrieves general information about a country based on its two-letter country code (ISO 3166-1 alpha-2).**



Example:

http://localhost:8080/countryinfo/v1/info/no

Render link



Response:



json

{

&nbsp; "name": "Norway",

&nbsp; "capital": "Oslo",

&nbsp; "population": 5379475,

&nbsp; "area": 323802,

&nbsp; "continents": \["Europe"],

&nbsp; "languages": {

&nbsp;   "nno": "Norwegian Nynorsk",

&nbsp;   "nob": "Norwegian Bokmål",

&nbsp;   "smi": "Sami"

&nbsp; },

&nbsp; "borders": \["FIN", "SWE", "RUS"],

&nbsp; "flag": "https://flagcdn.com/w320/no.png"

}

languages is a map with language codes and full names.

borders is a list of neighboring countries (same two-letter format).



##### **3. Exchange Rates to Neighboring Countries – GET /countryinfo/v1/exchange/{code}**

**Returns the country's base currency and exchange rates to all neighboring countries' currencies.**



Example:

http://localhost:8080/countryinfo/v1/exchange/no

Render link



Response:



json

{

&nbsp; "country": "Norway",

&nbsp; "base-currency": "NOK",

&nbsp; "exchange-rates": {

&nbsp;   "EUR": 0.088973,

&nbsp;   "SEK": 0.950225,

&nbsp;   "RUB": 8.033862

&nbsp; }

}

base-currency is the ISO 4217 code for the country's own currency.

exchange-rates contains rates for each neighboring country's currency (1 base unit = X foreign). If the country has no neighbors, the object is empty.





**How to Run Locally**

**Clone the repository:**



bash

git clone https://github.com/your-username/cloud-assignment-1.git

cd cloud-assignment-1

Run the server:



bash

go run cmd/server/main.go

Open http://localhost:8080 in your browser or use curl:



bash

curl http://localhost:8080/countryinfo/v1/status/

curl http://localhost:8080/countryinfo/v1/info/no

curl http://localhost:8080/countryinfo/v1/exchange/no



Dependencies (External APIs):

Rest Countries API (self-hosted): http://129.241.150.113:8080/v3.1/



Currency API (self-hosted): http://129.241.150.113:9090/currency/



The service uses only Go's standard library – no third-party packages.



