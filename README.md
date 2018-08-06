# Holiday Offers

A simple API written in Go for filtering holiday offers retrieved from external API.

## Usage
You can run the API with:

    go run *.go
    
It runs on port 8080.

## Debug
If you want to view uncompressed response in pretty form, you can pass `debug=true` query param.


## Endpoints
The only endpoint retrieves data from external API. You can use query params for filtering results.
* **URL**
    
    `/v1/`
    
* **Method:**
    
    `GET`
    
* **URL Params**

    * `debug = [bool]`
    * `earliest_departure_time = [%H:%M]`
    * `earliest_return_time = [%H:%M]`
    * `max_price = [numeric]`
    * `min_price = [numeric]`
    * `star_rating = [numeric]`

* **Success Response:**
    * **Code:** 200 OK<br/>
    **Content:**
    
    ```json
    {
      "offers": [
        {
          "Hotelname": "Catalonia+Punta+Del+Rey", 
          "Inboundarr": "23/08/2018 00:20", 
          "Inbounddep": "22/08/2018 20:05", 
          "Inboundfltnum": "LS1664", 
          "Outboundfltnum": "LS1663", 
          "Sellprice": "510.62", 
          "Starrating": "4"
        }, 
        {
          "Hotelname": "Catalonia+Las+Vegas", 
          "Inboundarr": "23/08/2018 00:20", 
          "Inbounddep": "22/08/2018 20:05", 
          "Inboundfltnum": "LS1664", 
          "Outboundfltnum": "LS1663", 
          "Sellprice": "518.81", 
          "Starrating": "4"
        }
      ], 
      "summary": {
        "average_price": 514.71, 
        "cheapest_price": 510.62, 
        "most_expensive_price": 518.81
      }
    }
    ```
    
* **Error Response:**
    * **Code:** 400 Bad Request<br/>
    **Content**: "Unknown parameter: {key}"
    
        *OR*
        
    * **Code:** 400 Bad Request<br/>
    **Content**: "{value} is not valid value for {key}"

        *OR*
        
    * **Code:** 400 Bad Request<br/>
    **Content**: "Invalid time format for {key}"