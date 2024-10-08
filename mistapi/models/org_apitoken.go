package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OrgApitoken represents a OrgApitoken struct.
// Org API Token
type OrgApitoken struct {
    // email of the token creator / null if creator is deleted
    CreatedBy            Optional[string]  `json:"created_by"`
    CreatedTime          *float64          `json:"created_time,omitempty"`
    Id                   *uuid.UUID        `json:"id,omitempty"`
    Key                  *string           `json:"key,omitempty"`
    LastUsed             Optional[float64] `json:"last_used"`
    // name of the token
    Name                 Optional[string]  `json:"name"`
    OrgId                *uuid.UUID        `json:"org_id,omitempty"`
    // list of privileges the token has on the orgs/sites
    Privileges           []PrivilegeOrg    `json:"privileges,omitempty"`
    // to restrict where the API can be called from
    SrcIps               []string          `json:"src_ips,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgApitoken.
// It customizes the JSON marshaling process for OrgApitoken objects.
func (o OrgApitoken) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgApitoken object to a map representation for JSON marshaling.
func (o OrgApitoken) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CreatedBy.IsValueSet() {
        if o.CreatedBy.Value() != nil {
            structMap["created_by"] = o.CreatedBy.Value()
        } else {
            structMap["created_by"] = nil
        }
    }
    if o.CreatedTime != nil {
        structMap["created_time"] = o.CreatedTime
    }
    if o.Id != nil {
        structMap["id"] = o.Id
    }
    if o.Key != nil {
        structMap["key"] = o.Key
    }
    if o.LastUsed.IsValueSet() {
        if o.LastUsed.Value() != nil {
            structMap["last_used"] = o.LastUsed.Value()
        } else {
            structMap["last_used"] = nil
        }
    }
    if o.Name.IsValueSet() {
        if o.Name.Value() != nil {
            structMap["name"] = o.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.Privileges != nil {
        structMap["privileges"] = o.Privileges
    }
    if o.SrcIps != nil {
        structMap["src_ips"] = o.SrcIps
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgApitoken.
// It customizes the JSON unmarshaling process for OrgApitoken objects.
func (o *OrgApitoken) UnmarshalJSON(input []byte) error {
    var temp orgApitoken
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_by", "created_time", "id", "key", "last_used", "name", "org_id", "privileges", "src_ips")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.CreatedBy = temp.CreatedBy
    o.CreatedTime = temp.CreatedTime
    o.Id = temp.Id
    o.Key = temp.Key
    o.LastUsed = temp.LastUsed
    o.Name = temp.Name
    o.OrgId = temp.OrgId
    o.Privileges = temp.Privileges
    o.SrcIps = temp.SrcIps
    return nil
}

// orgApitoken is a temporary struct used for validating the fields of OrgApitoken.
type orgApitoken  struct {
    CreatedBy   Optional[string]  `json:"created_by"`
    CreatedTime *float64          `json:"created_time,omitempty"`
    Id          *uuid.UUID        `json:"id,omitempty"`
    Key         *string           `json:"key,omitempty"`
    LastUsed    Optional[float64] `json:"last_used"`
    Name        Optional[string]  `json:"name"`
    OrgId       *uuid.UUID        `json:"org_id,omitempty"`
    Privileges  []PrivilegeOrg    `json:"privileges,omitempty"`
    SrcIps      []string          `json:"src_ips,omitempty"`
}
