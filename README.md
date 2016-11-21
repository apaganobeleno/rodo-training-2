### Challenge #2
1. Create a we app that handles http requests for the folllowing things:
```
/ip/details 
```

Should take the IP from the request and using freegeoip.net should return:

- Country
- Region
- City
- Zip code
- Time zone
- Link to google maps with geolocation.

```
/ip/details?for=X.X.X.X
```

Should return the same details as the first one but for the passed "for" attribute.

- Application should return 404 for no existing routes, we would like it to answer with :
```
{error: "Page not found"}
```

And http status 404 as mentioned earlier.
