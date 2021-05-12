# mtb-trails-api-golang
Scrapes the local mountain bike trail status page https://www.trianglemtb.com/mobiletrailstatus.php into JSON

Try it out yourself! 

On Linux and Mac, `curl https://api.nashluffman.com/trails` 

On Windows, switch to linux :) 

Example response

```json
{
  "metadata": {
    "query_time": "2021.01.01 17:35:19"
  },
  "data": [
    {
      "trail": "Beaver Dam",
      "status": "closed",
      "last_updated": "12/31 8:19 am",
      "trail_info": "https://www.trianglemtb.com/beaverdam.php"
    },
    {
      "trail": "Briar Chapel",
      "status": "closed",
      "last_updated": "12/30 5:00 pm",
      "trail_info": "https://www.trianglemtb.com/briarchapel.php"
    },
    {
      "trail": "Skills Area",
      "status": "closed",
      "last_updated": "12/27 3:21 pm",
      "trail_info": "https://www.trianglemtb.com/briarchapel.php"
    },
    {
      "trail": "Herndon Loop",
      "status": "closed",
      "last_updated": "12/27 3:21 pm",
      "trail_info": "https://www.trianglemtb.com/briarchapel.php"
    },
    {
      "trail": "Brumley Forest",
      "status": "closed",
      "last_updated": "12/30 9:30 pm",
      "trail_info": "https://www.trianglemtb.com/brumley.php"
    },
    {
      "trail": "CNF (airport)",
      "status": "closed",
      "last_updated": "12/31 8:47 am",
      "trail_info": "https://www.trianglemtb.com/chapelhill.php"
    },
    {
      "trail": "CNF (school)",
      "status": "closed",
      "last_updated": "12/31 8:47 am",
      "trail_info": "https://www.trianglemtb.com/chapelhill.php"
    },
    {
      "trail": "Crabtree",
      "status": "closed",
      "last_updated": "12/28 2:38 pm",
      "trail_info": "https://www.trianglemtb.com/crabtree.php"
    },
    {
      "trail": "Forest Ridge",
      "status": "closed",
      "last_updated": "01/01 10:42 am",
      "trail_info": "https://www.trianglemtb.com/forest_ridge.php"
    },
    {
      "trail": "Green Hills",
      "status": "closed",
      "last_updated": "12/31 10:26 am",
      "trail_info": "https://www.trianglemtb.com/green_hills.php"
    },
    {
      "trail": "Harris",
      "status": "closed",
      "last_updated": "12/30 11:07 am",
      "trail_info": "https://www.trianglemtb.com/harris.php"
    },
    {
      "trail": "Legend (Lower)",
      "status": "closed",
      "last_updated": "12/13 4:02 pm",
      "trail_info": "https://www.trianglemtb.com/legend.php"
    },
    {
      "trail": "Legend (Upper)",
      "status": "closed",
      "last_updated": "12/24 11:24 am",
      "trail_info": "https://www.trianglemtb.com/legend.php"
    },
    {
      "trail": "Little River",
      "status": "closed",
      "last_updated": "12/31 2:35 pm",
      "trail_info": "https://www.trianglemtb.com/littleriver.php"
    },
    {
      "trail": "New Light",
      "status": "closed",
      "last_updated": "12/28 5:00 am",
      "trail_info": "https://www.trianglemtb.com/newlight.php"
    },
    {
      "trail": "RTP Trails",
      "status": "closed",
      "last_updated": "12/30 8:39 pm",
      "trail_info": "https://www.trianglemtb.com/othertrails.php#rtp"
    },
    {
      "trail": "San Lee Gravity Park",
      "status": "closed",
      "last_updated": "12/24 10:32 am",
      "trail_info": "https://www.trianglemtb.com/sanlee.php"
    },
    {
      "trail": "San Lee Singletrack",
      "status": "closed",
      "last_updated": "12/24 10:32 am",
      "trail_info": "https://www.trianglemtb.com/sanlee.php"
    },
    {
      "trail": "Wendell Falls",
      "status": "closed",
      "last_updated": "12/14 11:40 am",
      "trail_info": "https://www.trianglemtb.com/othertrails.php#wendellfalls"
    },
    {
      "trail": "Williamson Preserve",
      "status": "closed",
      "last_updated": "01/01 11:06 am",
      "trail_info": "https://www.trianglemtb.com/williamson_preserve.php"
    }
  ]
}
```
