### Kata: "The Ad-Fraud Detective"

**Business Context:**
You are a data engineer at "ClickFlow," a digital advertising network. A major client, "Nexus Motors," has reported a suspicious spike in clicks for their new electric car campaign. They suspect click fraud from a network of bots. Your task is to analyze one hour of raw clickstream data to identify potential fraud patterns.

**The Problem:**
Fraudulent activity often leaves a footprint. In this case, you are looking for two key indicators:

1.  **Click Farms:** A large number of clicks originating from the same IP address in a very short time.
2.  **Impossible Travel:** Clicks from the same User ID that originate from geographically distant locations within an impossibly short time frame (e.g., New York and London within 10 minutes).

**Your Mission:**
Design a high-level batch processing workflow to process the raw click data and generate two alert reports:

1.  A list of IP addresses that generated more than 100 clicks in the one-hour window, along with their total click count.
2.  A list of User IDs that have clicks from locations more than 500 km apart within a 10-minute window.

**Available Data (in a distributed filesystem):**

- `clicks/` directory: Thousands of files containing raw click events in JSON format.
  - Example record: `{"timestamp": "2023-10-27T14:05:32Z", "user_id": "user123", "ip_address": "192.168.1.1", "campaign_id": "nexus_ev", "location": {"country": "US", "city": "New York"}}`
- `geo/` directory: A static reference dataset that maps city names to their approximate latitude and longitude coordinates.
  - Example record: `{"city": "New York", "lat": 40.7128, "lon": -74.0060}`

**Kata Tasks (30-minute thought exercise):**

1.  **Workflow Design (10 mins):** Sketch the high-level stages of your batch processing job. Don't write code; describe the dataflow using concepts from the chapter (e.g., Map, Reduce, Join, Filter, Group). How many jobs will you need? In what order?

    > Since I need to generate two reports, a need a group of jobs for each report. For the first report I need a map Job from the clicks directory that gonna output as key the ip_address and as value the timestamp. In the reducer side, I'll have the ip_address and a list of timestamps. From the actual conditions, the dataset now have only a single hour of data. The reducer, then, should output the ip_address and the list size from these timestamps and output it to another job. This another job will have a map to invert the relation between the key and the value. Now the key is the list size and value is the ip_address. In the reducer, as the result gonna be sorted by key, I can discard all the keys there represents values less than 100, and output only keys and values that the key is more or equal than 100. Maybe this additional job with it's map and reducer can be discarded and output the results from the first reducer. For the second report, It gonna be more complex, because, first, I need to process the click directory, by mapping the user_id as key and city and timestamp as values. Then, I use a join from each city value to the geo directory city to get the lat and lon. I don't know exactly how to continue from here, maybe ordering the values that I have by timestamp and group by 10 minutes windows, and then group by user_id and enrich the data of the lat/lon to generate the deltas between cities to discover the distance and determine if they are 500km apart.

2.  **Joins Strategy (10 mins):** To calculate the distance between two cities for the "Impossible Travel" detection, you need to enrich the click data with geographic coordinates.

    - What type of join (Map-Side or Reduce-Side) would you use to combine the `clicks` data with the `geo` data? Justify your choice based on the size and nature of the datasets.
      > I think that reduce-side should be more appropriate here, because the question are about the users, and clicks is sufficiently big to use the reducer on map-side. In the reduce-side, the geolocalization table can be sufficiently small to fit on memory to make the join.

3.  **Algorithm & Shuffling (10 mins):** For the "Impossible Travel" report, you need to find pairs of clicks from the same user that are close in time but far in geography.
    - Describe the key(s) you would use in your Map and Reduce phases to group the data effectively.
      > I think that I would use the user_id as key, because it will group all the user click data into it's id. From here, we can sort the values of the timestamp to see the distance between clicks for each user. Maybe we can pass this data by inverting the relation to use the timestamp slice as key to see deltas between cities.
    - What computation would happen in your Reducer function to identify the "impossible" pairs? How would you avoid a performance pitfall when a single user has thousands of clicks?
      > The computation here should be a sort by timestamp, I think, because it'll put each click that remains short to each other. The sort should be made by value.

**Success Criteria:**
A successful kata response will clearly articulate a multi-stage batch workflow, make a reasoned choice for the join strategy, and outline a logical data grouping and computation plan that efficiently identifies the required fraud patterns.
