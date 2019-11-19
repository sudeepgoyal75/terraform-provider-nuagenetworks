package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceGatewaysLocation() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceGatewaysLocationRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "latitude": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ignore_geocode": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "time_zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "locality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "longitude": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "country": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_entity_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceGatewaysLocationRead(d *schema.ResourceData, m interface{}) error {
    filteredGatewaysLocations := vspk.GatewaysLocationsList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredGatewaysLocations, err = parent.GatewaysLocations(fetchFilter)
    if err != nil {
        return err
    }

    GatewaysLocation := &vspk.GatewaysLocation{}

    if len(filteredGatewaysLocations) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredGatewaysLocations) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    GatewaysLocation = filteredGatewaysLocations[0]

    d.Set("last_updated_by", GatewaysLocation.LastUpdatedBy)
    d.Set("latitude", GatewaysLocation.Latitude)
    d.Set("address", GatewaysLocation.Address)
    d.Set("ignore_geocode", GatewaysLocation.IgnoreGeocode)
    d.Set("time_zone_id", GatewaysLocation.TimeZoneID)
    d.Set("entity_scope", GatewaysLocation.EntityScope)
    d.Set("locality", GatewaysLocation.Locality)
    d.Set("longitude", GatewaysLocation.Longitude)
    d.Set("country", GatewaysLocation.Country)
    d.Set("associated_entity_name", GatewaysLocation.AssociatedEntityName)
    d.Set("associated_entity_type", GatewaysLocation.AssociatedEntityType)
    d.Set("state", GatewaysLocation.State)
    d.Set("external_id", GatewaysLocation.ExternalID)
    
    d.Set("id", GatewaysLocation.Identifier())
    d.Set("parent_id", GatewaysLocation.ParentID)
    d.Set("parent_type", GatewaysLocation.ParentType)
    d.Set("owner", GatewaysLocation.Owner)

    d.SetId(GatewaysLocation.Identifier())
    
    return nil
}