# Magic The Gathering Card Reader
Ruzzie made an awesome Card Reader application. He exposes a URL to retrieve all kinds of MTG cards.

For ease of use, the data has been downloaded and stored in [this json file](allcards.json)

## Part 1
Read the json file, and show only cards that contain “Creature” in the “Types” field

Example result:
```
[
   {
       "MultiverseId": 12346,
       "Legality": 412,
       "BasicType": 4,
       "ColorIdentity": 31,
       "Cmc": 4,
       "Rating": 0,
       "Price": 0.05,
       "ManaCost": "{3}{R}",
       "ImageRelativeUrl": "/cards/normal/en/con/77.jpg?1496793377",
       "Types": "Creature — Phoenix",
       "Name": "Worldheart Phoenix"
   }
   ...
]
```

## Part 2
Export the results in a CSV file

## Finished?

Pick one (or more) from this list you can do to expand:

- Take arguments in the application to filter on `rating >= X`. (`/filterrating:1`)
- Output the number of records written to the console
- Output the average price and rating of the records written
- Instead of reading the file from disk, retrieve it from `https://ruzziemtgstorage.blob.core.windows.net/demo/allcards.json`

