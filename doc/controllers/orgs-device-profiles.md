# Orgs Device Profiles

```go
orgsDeviceProfiles := client.OrgsDeviceProfiles()
```

## Class Name

`OrgsDeviceProfiles`

## Methods

* [Assign Org Device Profile](../../doc/controllers/orgs-device-profiles.md#assign-org-device-profile)
* [Create Org Device Profiles](../../doc/controllers/orgs-device-profiles.md#create-org-device-profiles)
* [Delete Org Device Profile](../../doc/controllers/orgs-device-profiles.md#delete-org-device-profile)
* [Get Org Device Profile](../../doc/controllers/orgs-device-profiles.md#get-org-device-profile)
* [List Org Device Profiles](../../doc/controllers/orgs-device-profiles.md#list-org-device-profiles)
* [Unassign Org Device Profile](../../doc/controllers/orgs-device-profiles.md#unassign-org-device-profile)
* [Update Org Device Profile](../../doc/controllers/orgs-device-profiles.md#update-org-device-profile)


# Assign Org Device Profile

Assign Org Device Profile to Devices

```go
AssignOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseAssignSuccess`](../../doc/models/response-assign-success.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "5c5b350e0001",
        "5c5b350e0003",
    },
}

apiResponse, err := orgsDeviceProfiles.AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
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
{
  "success": [
    "5c5b350e0001"
  ]
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


# Create Org Device Profiles

Create Device Profile

```go
CreateOrgDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Deviceprofile2) (
    models.ApiResponse[models.CreateOrgDeviceProfilesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Deviceprofile2`](../../doc/models/containers/deviceprofile-2.md) | Body, Optional | Request Body |

## Response Type

[`models.CreateOrgDeviceProfilesResponse`](../../doc/models/containers/create-org-device-profiles-response.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.DeviceprofileContainer.FromDeviceprofileAp(models.DeviceprofileAp{
    Aeroscout:        models.ToPointer(models.ApAeroscout{
        Enabled:         models.ToPointer(false),
        Host:            models.NewOptional(models.ToPointer("aero.pvt.net")),
        LocateConnected: models.ToPointer(true),
    }),
    Led:              models.ToPointer(models.ApLed{
        Brightness: models.ToPointer(255),
        Enabled:    models.ToPointer(true),
    }),
    Name:             models.NewOptional(models.ToPointer("string")),
    NtpServers:       []string{
        "10.10.10.10",
    },
    Type:             models.ToPointer(models.DeviceTypeApEnum("ap")),
    UsbConfig:        models.ToPointer(models.ApUsb{
        Cacert:     models.NewOptional(models.ToPointer("string")),
        Channel:    models.ToPointer(3),
        Enabled:    models.ToPointer(true),
        Host:       models.ToPointer("1.1.1.1"),
        Port:       models.ToPointer(0),
        Type:       models.ToPointer(models.ApUsbTypeEnum("imagotag")),
        VerifyCert: models.ToPointer(true),
        VlanId:     models.ToPointer(1),
    }),
})

apiResponse, err := orgsDeviceProfiles.CreateOrgDeviceProfiles(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceprofileAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *r)
    } else if r, ok := responseBody.AsDeviceprofileGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "aeroscout": {
    "enabled": true,
    "host": "string"
  },
  "ble_config": {
    "beacon_enabled": true,
    "beacon_rate": 0,
    "beacon_rate_mode": "default",
    "beam_disabled": [
      0
    ],
    "eddystone_uid_adv_power": -100,
    "eddystone_uid_beams": "string",
    "eddystone_uid_enabled": true,
    "eddystone_uid_freq_msec": 0,
    "eddystone_uid_instance": "string",
    "eddystone_uid_namespace": "string",
    "eddystone_url_adv_power": 0,
    "eddystone_url_beams": "string",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 0,
    "eddystone_url_url": "string",
    "ibeacon_adv_power": -100,
    "ibeacon_beams": "string",
    "ibeacon_enabled": true,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 0,
    "ibeacon_minor": 0,
    "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "power": 1,
    "power_mode": "string"
  },
  "created_time": 0,
  "disable_eth1": true,
  "disable_module": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ip_config": {
    "dns": [
      "string"
    ],
    "dns_suffix": [
      "string"
    ],
    "gateway": "192.168.0.1",
    "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "ip": "192.168.0.1",
    "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "mtu": 0,
    "netmask": "192.168.0.1",
    "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "mesh": {
    "enabled": true,
    "group": 0,
    "role": "base"
  },
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "poe_passthrough": true,
  "radio_config": {
    "ant_gain_24": 0,
    "ant_gain_5": 0,
    "band_24": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "band_24_usage": "24",
    "band_5": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "band_5_on_24_radio": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "scanning_enabled": true
  },
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "switch_config": {
    "enabled": true,
    "eth0": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth1": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth2": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth3": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "module": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "wds": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    }
  },
  "usb_config": {
    "cacert": "string",
    "channel": 0,
    "enabled": true,
    "host": "string",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true
  }
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


# Delete Org Device Profile

Delete Org Device Profile

```go
DeleteOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsDeviceProfiles.DeleteOrgDeviceProfile(ctx, orgId, deviceprofileId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
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


# Get Org Device Profile

Get Org device Profile Details

```go
GetOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID) (
    models.ApiResponse[models.GetOrgDeviceProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.GetOrgDeviceProfileResponse`](../../doc/models/containers/get-org-device-profile-response.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDeviceProfiles.GetOrgDeviceProfile(ctx, orgId, deviceprofileId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceprofileAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *r)
    } else if r, ok := responseBody.AsDeviceprofileGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "aeroscout": {
    "enabled": true,
    "host": "string"
  },
  "ble_config": {
    "beacon_enabled": true,
    "beacon_rate": 0,
    "beacon_rate_mode": "default",
    "beam_disabled": [
      0
    ],
    "eddystone_uid_adv_power": -100,
    "eddystone_uid_beams": "string",
    "eddystone_uid_enabled": true,
    "eddystone_uid_freq_msec": 0,
    "eddystone_uid_instance": "string",
    "eddystone_uid_namespace": "string",
    "eddystone_url_adv_power": 0,
    "eddystone_url_beams": "string",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 0,
    "eddystone_url_url": "string",
    "ibeacon_adv_power": -100,
    "ibeacon_beams": "string",
    "ibeacon_enabled": true,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 0,
    "ibeacon_minor": 0,
    "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "power": 1,
    "power_mode": "string"
  },
  "created_time": 0,
  "disable_eth1": true,
  "disable_module": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ip_config": {
    "dns": [
      "string"
    ],
    "dns_suffix": [
      "string"
    ],
    "gateway": "192.168.0.1",
    "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "ip": "192.168.0.1",
    "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "mtu": 0,
    "netmask": "192.168.0.1",
    "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "mesh": {
    "enabled": true,
    "group": 0,
    "role": "base"
  },
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "poe_passthrough": true,
  "radio_config": {
    "ant_gain_24": 0,
    "ant_gain_5": 0,
    "band_24": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "band_24_usage": "24",
    "band_5": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "band_5_on_24_radio": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "scanning_enabled": true
  },
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "switch_config": {
    "enabled": true,
    "eth0": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth1": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth2": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth3": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "module": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "wds": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    }
  },
  "usb_config": {
    "cacert": "string",
    "channel": 0,
    "enabled": true,
    "host": "string",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true
  }
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


# List Org Device Profiles

Get List of Org Device Profiles

```go
ListOrgDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.ListOrgDeviceProfilesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`[]models.ListOrgDeviceProfilesResponse`](../../doc/models/containers/list-org-device-profiles-response.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeEnum("ap")

limit := 100

page := 1

apiResponse, err := orgsDeviceProfiles.ListOrgDeviceProfiles(ctx, orgId, &mType, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceprofileAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *i)
        } else if i, ok := item.AsDeviceprofileGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "aeroscout": {
      "enabled": false,
      "host": "aero.pvt.net",
      "locate_connected": true
    },
    "ble_config": {
      "beacon_enabled": false,
      "beacon_rate": 3,
      "beacon_rate_mode": "custom",
      "beam_disabled": [
        1,
        3,
        6
      ],
      "custom_ble_packet_enabled": false,
      "custom_ble_packet_frame": "0x........",
      "custom_ble_packet_freq_msec": 300,
      "eddystone_uid_adv_power": -65,
      "eddystone_uid_beams": "2-4,7",
      "eddystone_uid_enabled": false,
      "eddystone_uid_freq_msec": 200,
      "eddystone_uid_instance": "5c5b35000001",
      "eddystone_uid_namespace": "2818e3868dec25629ede",
      "eddystone_url_adv_power": -65,
      "eddystone_url_beams": "2-4,7",
      "eddystone_url_enabled": true,
      "eddystone_url_freq_msec": 1000,
      "eddystone_url_url": "https://www.abc.com",
      "ibeacon_adv_power": -65,
      "ibeacon_beams": "2-4,7",
      "ibeacon_enabled": false,
      "ibeacon_freq_msec": 0,
      "ibeacon_major": 13,
      "ibeacon_minor": 138,
      "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
      "power": 10,
      "power_mode": "custom"
    },
    "created_time": 0,
    "disable_eth1": false,
    "disable_eth2": false,
    "disable_eth3": false,
    "disable_module": false,
    "for_site": true,
    "height": 0,
    "id": "497f6eca-6276-4993-bfeb-53cbbbba6108",
    "ip_config": {
      "dns": [
        "8.8.8.8",
        "4.4.4.4"
      ],
      "dns_suffix": [
        ".mist.local",
        ".mist.com"
      ],
      "gateway": "10.2.1.254",
      "gateway6": "2607:f8b0:4005:808::1",
      "ip": "10.2.1.1",
      "ip6": "2607:f8b0:4005:808::2004",
      "mtu": 0,
      "netmask": "255.255.255.0",
      "netmask6": "/32",
      "type": "static",
      "type6": "static",
      "vlan_id": 1
    },
    "led": {
      "brightness": 255,
      "enabled": true
    },
    "map_id": "09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1",
    "mesh": {
      "enabled": false,
      "group": 1,
      "role": "base"
    },
    "modified_time": 0,
    "name": "string",
    "notes": "string",
    "ntp_servers": [
      "string"
    ],
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "orientation": 0,
    "orientation_overwrite": true,
    "poe_passthrough": false,
    "port_config": {
      "property1": {
        "additional_vlan_ids": [
          55,
          66
        ],
        "authentication_protocol": "pap",
        "disabled": true,
        "dynamic_vlan": {
          "default_vlan_id": 999,
          "enabled": true,
          "type": "string",
          "vlans": {
            "1-10": null,
            "user": null
          }
        },
        "enable_mac_auth": false,
        "forwarding": "all",
        "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
        "mxtunnel_name": "string",
        "port_auth": "none",
        "port_vlan_id": 1,
        "radius_config": {
          "acct_interim_interval": 0,
          "acct_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1813,
              "secret": "testing123"
            }
          ],
          "auth_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1812,
              "secret": "testing123"
            }
          ],
          "auth_servers_retries": 3,
          "auth_servers_timeout": 5,
          "coa_enabled": false,
          "coa_port": 3799,
          "network": "string",
          "source_ip": "string"
        },
        "radsec": {
          "enabled": true,
          "idle_timeout": 60,
          "mxcluster_ids": [
            "572586b7-f97b-a22b-526c-8b97a3f609c4"
          ],
          "proxy_hosts": [
            "mxedge1.local"
          ],
          "server_name": "radsec.abc.com",
          "servers": [
            {
              "host": "1.1.1.1",
              "port": 1812
            }
          ],
          "use_mxedge": true,
          "use_site_mxedge": false
        },
        "vlan_id": 9,
        "vland_ids": [
          1,
          10,
          50
        ],
        "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
        "wxtunnel_remote_id": "wifiguest"
      },
      "property2": {
        "additional_vlan_ids": [
          55,
          66
        ],
        "authentication_protocol": "pap",
        "disabled": true,
        "dynamic_vlan": {
          "default_vlan_id": 999,
          "enabled": true,
          "type": "string",
          "vlans": {
            "1-10": null,
            "user": null
          }
        },
        "enable_mac_auth": false,
        "forwarding": "all",
        "mx_tunnel_id": "08cd7499-5841-51c8-e663-fb16b6f3b45e",
        "mxtunnel_name": "string",
        "port_auth": "none",
        "port_vlan_id": 1,
        "radius_config": {
          "acct_interim_interval": 0,
          "acct_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1813,
              "secret": "testing123"
            }
          ],
          "auth_servers": [
            {
              "host": "1.2.3.4",
              "keywrap_enabled": true,
              "keywrap_format": "hex",
              "keywrap_kek": "1122334455",
              "keywrap_mack": "1122334455",
              "port": 1812,
              "secret": "testing123"
            }
          ],
          "auth_servers_retries": 3,
          "auth_servers_timeout": 5,
          "coa_enabled": false,
          "coa_port": 3799,
          "network": "string",
          "source_ip": "string"
        },
        "radsec": {
          "enabled": true,
          "idle_timeout": 60,
          "mxcluster_ids": [
            "572586b7-f97b-a22b-526c-8b97a3f609c4"
          ],
          "proxy_hosts": [
            "mxedge1.local"
          ],
          "server_name": "radsec.abc.com",
          "servers": [
            {
              "host": "1.1.1.1",
              "port": 1812
            }
          ],
          "use_mxedge": true,
          "use_site_mxedge": false
        },
        "vlan_id": 9,
        "vland_ids": [
          1,
          10,
          50
        ],
        "wxtunnel_id": "7dae216d-7c98-a51b-e068-dd7d477b7216",
        "wxtunnel_remote_id": "wifiguest"
      }
    },
    "pwr_config": {
      "base": 0
    },
    "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
    "type": "ap",
    "x": 0,
    "y": 0
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


# Unassign Org Device Profile

Unassign Org Device Profile from Devices

```go
UnassignOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseAssignSuccess`](../../doc/models/response-assign-success.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs: []string{
        "5c5b350e0001",
        "5c5b350e0003",
    },
}

apiResponse, err := orgsDeviceProfiles.UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
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
{
  "success": [
    "5c5b350e0001"
  ]
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


# Update Org Device Profile

Update Org Device Profile

```go
UpdateOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.Deviceprofile2) (
    models.ApiResponse[models.UpdateOrgDeviceProfileResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `deviceprofileId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Deviceprofile2`](../../doc/models/containers/deviceprofile-2.md) | Body, Optional | Request Body |

## Response Type

[`models.UpdateOrgDeviceProfileResponse`](../../doc/models/containers/update-org-device-profile-response.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceprofileId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.DeviceprofileContainer.FromDeviceprofileGateway(models.DeviceprofileGateway{
    Name:                  "string",
})

apiResponse, err := orgsDeviceProfiles.UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceprofileAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileAp: ", *r)
    } else if r, ok := responseBody.AsDeviceprofileGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceprofileGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "aeroscout": {
    "enabled": true,
    "host": "string"
  },
  "ble_config": {
    "beacon_enabled": true,
    "beacon_rate": 0,
    "beacon_rate_mode": "default",
    "beam_disabled": [
      0
    ],
    "eddystone_uid_adv_power": -100,
    "eddystone_uid_beams": "string",
    "eddystone_uid_enabled": true,
    "eddystone_uid_freq_msec": 0,
    "eddystone_uid_instance": "string",
    "eddystone_uid_namespace": "string",
    "eddystone_url_adv_power": 0,
    "eddystone_url_beams": "string",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 0,
    "eddystone_url_url": "string",
    "ibeacon_adv_power": -100,
    "ibeacon_beams": "string",
    "ibeacon_enabled": true,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 0,
    "ibeacon_minor": 0,
    "ibeacon_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "power": 1,
    "power_mode": "string"
  },
  "created_time": 0,
  "disable_eth1": true,
  "disable_module": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "ip_config": {
    "dns": [
      "string"
    ],
    "dns_suffix": [
      "string"
    ],
    "gateway": "192.168.0.1",
    "gateway6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "ip": "192.168.0.1",
    "ip6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "mtu": 0,
    "netmask": "192.168.0.1",
    "netmask6": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "mesh": {
    "enabled": true,
    "group": 0,
    "role": "base"
  },
  "modified_time": 0,
  "name": "string",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "poe_passthrough": true,
  "radio_config": {
    "ant_gain_24": 0,
    "ant_gain_5": 0,
    "band_24": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "band_24_usage": "24",
    "band_5": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "band_5_on_24_radio": {
      "allow_rrm_disable": true,
      "antenna_mode": "default",
      "bandwidth": 20,
      "channel": 0,
      "disabled": true,
      "power": 0,
      "power_max": 0,
      "power_min": 0,
      "preamble": "auto",
      "usage": "24"
    },
    "scanning_enabled": true
  },
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "switch_config": {
    "enabled": true,
    "eth0": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth1": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth2": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "eth3": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "module": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    },
    "wds": {
      "port_vlan_id": 1,
      "vlan_ids": [
        0
      ]
    }
  },
  "usb_config": {
    "cacert": "string",
    "channel": 0,
    "enabled": true,
    "host": "string",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true
  }
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

