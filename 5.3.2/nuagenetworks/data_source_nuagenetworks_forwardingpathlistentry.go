package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceForwardingPathListEntry() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceForwardingPathListEntryRead,
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
            "fc_override": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "forwarding_action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uplink_preference": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_forwarding_path_list": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceForwardingPathListEntryRead(d *schema.ResourceData, m interface{}) error {
    filteredForwardingPathListEntries := vspk.ForwardingPathListEntriesList{}
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
    parent := &vspk.ForwardingPathList{ID: d.Get("parent_forwarding_path_list").(string)}
    filteredForwardingPathListEntries, err = parent.ForwardingPathListEntries(fetchFilter)
    if err != nil {
        return err
    }

    ForwardingPathListEntry := &vspk.ForwardingPathListEntry{}

    if len(filteredForwardingPathListEntries) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredForwardingPathListEntries) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ForwardingPathListEntry = filteredForwardingPathListEntries[0]

    d.Set("fc_override", ForwardingPathListEntry.FCOverride)
    d.Set("last_updated_by", ForwardingPathListEntry.LastUpdatedBy)
    d.Set("entity_scope", ForwardingPathListEntry.EntityScope)
    d.Set("forwarding_action", ForwardingPathListEntry.ForwardingAction)
    d.Set("uplink_preference", ForwardingPathListEntry.UplinkPreference)
    d.Set("priority", ForwardingPathListEntry.Priority)
    d.Set("external_id", ForwardingPathListEntry.ExternalID)
    
    d.Set("id", ForwardingPathListEntry.Identifier())
    d.Set("parent_id", ForwardingPathListEntry.ParentID)
    d.Set("parent_type", ForwardingPathListEntry.ParentType)
    d.Set("owner", ForwardingPathListEntry.Owner)

    d.SetId(ForwardingPathListEntry.Identifier())
    
    return nil
}