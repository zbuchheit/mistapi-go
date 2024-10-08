
# Privilege Msp View Enum

Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.

You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.

Below are the list of supported UI views. Note that this is UI only feature
Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.

You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.

Below are the list of supported UI views. Note that this is UI only feature

| UI View | Description |
| --- | --- |
| `reporting` | full access to all analytics tools |
| `marketing` | can view analytics and location maps |
| `location` | can view and manage location maps |
| `security` | can view and manage WLAN, rogues and authentication |
| `switch_admin` | can view and manage Switch ports |
| `mxedge_admin` | can view and manage Mist edges and Mist tunnels |
| `lobby_admin` | full access to Org and Site Pre-shared keys |

## Enumeration

`PrivilegeMspViewEnum`

## Fields

| Name |
|  --- |
| `lobby_admin` |
| `location` |
| `marketing` |
| `mxedge_admin` |
| `reporting` |
| `security` |
| `switch_admin` |

