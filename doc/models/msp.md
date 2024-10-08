
# Msp

## Structure

`Msp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowMist` | `*bool` | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `LogoUrl` | `*string` | Optional | For advanced tier (uMSPs) only |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Tier` | [`*models.MspTierEnum`](../../doc/models/msp-tier-enum.md) | Optional | enum: `advanced`, `base`<br>**Default**: `"base"` |
| `Url` | `*string` | Optional | For advanced tier (uMSPs) only |

## Example (as JSON)

```json
{
  "tier": "base",
  "allow_mist": false,
  "created_time": 254.96,
  "id": "00000682-0000-0000-0000-000000000000",
  "logo_url": "logo_url4",
  "modified_time": 80.0
}
```

