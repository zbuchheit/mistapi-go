# Sites Stats-Clients Wireless

```go
sitesStatsClientsWireless := client.SitesStatsClientsWireless()
```

## Class Name

`SitesStatsClientsWireless`

## Methods

* [Get Site Wireless Client Stats](../../doc/controllers/sites-stats-clients-wireless.md#get-site-wireless-client-stats)
* [Get Site Wireless Clients Stats by Map](../../doc/controllers/sites-stats-clients-wireless.md#get-site-wireless-clients-stats-by-map)
* [List Site Unconnected Client Stats](../../doc/controllers/sites-stats-clients-wireless.md#list-site-unconnected-client-stats)
* [List Site Wireless Clients Stats](../../doc/controllers/sites-stats-clients-wireless.md#list-site-wireless-clients-stats)


# Get Site Wireless Client Stats

Get Site Client Stats Details

```go
GetSiteWirelessClientStats(
    ctx context.Context,
    siteId uuid.UUID,
    clientMac string,
    wired *bool) (
    models.ApiResponse[[]models.StatsClientAnyOf],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `clientMac` | `string` | Template, Required | - |
| `wired` | `*bool` | Query, Optional | - |

## Response Type

[`[]models.StatsClientAnyOf`](../../doc/models/containers/stats-client-any-of.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

clientMac := "0000000000ab"

wired := false

apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientStats(ctx, siteId, clientMac, &wired)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsArrayOfClientWirelessStats(); ok {
            fmt.Println("Value narrowed down to []models.ClientWirelessStats: ", *i)
        } else if i, ok := item.AsStatsWiredClient(); ok {
            fmt.Println("Value narrowed down to models.StatsWiredClient: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Wireless Clients Stats by Map

Get Site Clients Stats By Map

```go
GetSiteWirelessClientsStatsByMap(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[][]models.ClientWirelessStats],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`[][]models.ClientWirelessStats`](../../doc/models/client-wireless-stats.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesStatsClientsWireless.GetSiteWirelessClientsStatsByMap(ctx, siteId, mapId, &page, &limit, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Unconnected Client Stats

Get List of Site Unconnected Client Location

```go
ListSiteUnconnectedClientStats(
    ctx context.Context,
    siteId uuid.UUID,
    mapId uuid.UUID) (
    models.ApiResponse[[]models.UnconnectedClientStat],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mapId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.UnconnectedClientStat`](../../doc/models/unconnected-client-stat.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mapId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsClientsWireless.ListSiteUnconnectedClientStats(ctx, siteId, mapId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "ap_mac": "5c5b350e0410",
    "last_seen": 1428939600,
    "mac": "5684dae9ac8b",
    "manufacture": "Apple",
    "rssi": -75,
    "x": 60,
    "y": 80
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Wireless Clients Stats

Get List of Site All Clients Stats Details

```go
ListSiteWirelessClientsStats(
    ctx context.Context,
    siteId uuid.UUID,
    wired *bool,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[[]models.StatsClientAnyOf],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wired` | `*bool` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`[]models.StatsClientAnyOf`](../../doc/models/containers/stats-client-any-of.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wired := false

limit := 100





duration := "10m"

apiResponse, err := sitesStatsClientsWireless.ListSiteWirelessClientsStats(ctx, siteId, &wired, &limit, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsArrayOfClientWirelessStats(); ok {
            fmt.Println("Value narrowed down to []models.ClientWirelessStats: ", *i)
        } else if i, ok := item.AsStatsWiredClient(); ok {
            fmt.Println("Value narrowed down to models.StatsWiredClient: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

