# easyFlight-airline-authority-service

**Search Operation Logic Documentation**

**1. Overview:**

The flight search operation in the flight booking project involves a detailed process orchestrated through several functions. The process begins when a gRPC request is received by the `SearchFlightInitial` function. The search details are then extracted and passed to the `SearchFlight` function, initiating the flight search.

**2. SearchFlightInitial:**

- **Input:** gRPC request containing search details.
- **Output:** Directs the request to the `SearchFlight` function.

**3. SearchFlight:**

- **Input:** dom.SearchDetails struct.
- **Output:** Calls either `OneWayFlightSearch` or `ReturnFlightSearch` based on whether it is a one-way or return flight search.

**4. OneWayFlightSearch:**

- **Input:** Search parameters from the `SearchFlight` function.
- **Output:** A list of flight paths from the departure to arrival airport with specified stops.

    - **Process:**
        - Creates three slices of type `FlightChart` from the database.
        - Collects routes from the departure airport to the destination airport on the departure date.
        - Collects routes from the arrival airport as the new departure airport, limiting stops to 2.
        - Combines the two sets of routes into a slice of type `FlightDetails`.

**5. Findallpaths:**

- **Input:** Flight details and max stops.
- **Output:** Graph traversal result - a list of flights grouped by departure airport.

    - **Process:**
        - Constructs a graph with departure airport as the key and a list of flights originating from that airport.
        - Conducts a graph traversal through flights originating from the original departure airport.
        - Returns a list of flights in the form of `[][]model.FlightDetails`.

**6. PathResponse:**

- **Input:** List of flight paths from `Findallpaths`.
- **Output:** Filtered paths based on the number of stops.

    - **Process:**
        - Checks if the number of flights in a path is greater than the specified number of stops.
        - Appends valid paths to a slice of type `Path`.

**7. Path:**

- A struct representing a flight path with attributes such as `PathId`, `Flights`, `NumberOfStops`, `TotalTravelTime`, `DepartureAirport`, and `ArrivalAirport`.

**8. Final Result:**

- The final response includes a list of flight paths, each containing a unique path ID, detailed flight information, the number of stops, total travel time, departure airport, and arrival airport.

This meticulous search operation ensures a comprehensive exploration of available flight paths, providing users with detailed and relevant information based on their search criteria. The combination of database queries, graph traversal, and path filtering results in a robust and efficient flight search mechanism in the flight booking project.

**9. Caching and Search Token Generation:**

- **Process:**
    - Upon receiving the flight search response, the values are stored in Redis and other caching mechanisms.
    - A unique search token is generated for the specific search, providing a reference identifier for subsequent requests related to the same search.
    - This ensures that the results of a particular search are cached and retrievable for a specific duration.

**10. Search Token:**

- **Definition:**
    - A unique identifier representing a specific flight search.

- **Attributes:**
    - **Token Value:** A unique string or identifier associated with a particular search.
    - **Expiration Time:** The duration for which the search results are considered valid.
    - **Cached Data:** The stored information related to the specific search, such as flight paths, fares, and other relevant details.

**11. Subsequent Search Requests:**

- **Usage:**
    - For every subsequent search request within the specified time frame, the search token generated in the initial search is required.
    - Users must include the search token in their requests to retrieve cached results, ensuring consistency and preventing changes in fare during the specified period.

**12. Benefits:**

- **Efficiency:**
    - Subsequent searches leveraging the search token are faster, as they retrieve cached data without the need for a complete search operation.

- **Consistency:**
    - Users receive consistent results for a particular search, minimizing discrepancies caused by fluctuating fares during the specified duration.

**13. Considerations:**

- **Token Expiry:**
    - It is crucial to manage the expiration time of search tokens to balance efficiency and accuracy. Tokens should expire after a reasonable duration to accommodate potential fare changes.

- **Cache Management:**
    - Regularly review and clear outdated or unnecessary cached data to maintain system efficiency and storage optimization.

**14. Conclusion:**

- The integration of caching and search token generation enhances the overall efficiency and consistency of the flight search system. Users benefit from quicker responses and a stabilized fare environment within the specified time frame, contributing to a seamless and reliable user experience in the flight booking project.

**To generate mock file using mockgen**

- mockgen -source=source.go -destination=destination.go -package=package_no