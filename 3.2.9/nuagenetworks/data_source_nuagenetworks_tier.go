package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceTier() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceTierRead,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_floating_ip_pool_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_macro_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_network_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_app": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceTierRead(d *schema.ResourceData, m interface{}) error {
    filteredTiers := vspk.TiersList{}
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
    parent := &vspk.App{ID: d.Get("parent_app").(string)}
    filteredTiers, err = parent.Tiers(fetchFilter)
    if err != nil {
        return err
    }

    Tier := &vspk.Tier{}

    if len(filteredTiers) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredTiers) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Tier = filteredTiers[0]

    d.Set("name", Tier.Name)
    d.Set("last_updated_by", Tier.LastUpdatedBy)
    d.Set("gateway", Tier.Gateway)
    d.Set("address", Tier.Address)
    d.Set("description", Tier.Description)
    d.Set("metadata", Tier.Metadata)
    d.Set("netmask", Tier.Netmask)
    d.Set("entity_scope", Tier.EntityScope)
    d.Set("associated_application_id", Tier.AssociatedApplicationID)
    d.Set("associated_floating_ip_pool_id", Tier.AssociatedFloatingIPPoolID)
    d.Set("associated_network_macro_id", Tier.AssociatedNetworkMacroID)
    d.Set("associated_network_object_id", Tier.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", Tier.AssociatedNetworkObjectType)
    d.Set("external_id", Tier.ExternalID)
    d.Set("type", Tier.Type)
    
    d.Set("id", Tier.Identifier())
    d.Set("parent_id", Tier.ParentID)
    d.Set("parent_type", Tier.ParentType)
    d.Set("owner", Tier.Owner)

    d.SetId(Tier.Identifier())
    
    return nil
}