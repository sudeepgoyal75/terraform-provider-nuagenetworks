package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePatch() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePatchRead,
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
            "patch_build_number": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "patch_summary": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "patch_tag": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "patch_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "supports_deletion": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "supports_network_acceleration": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourcePatchRead(d *schema.ResourceData, m interface{}) error {
    filteredPatchs := vspk.PatchsList{}
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
    parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
    filteredPatchs, err = parent.Patchs(fetchFilter)
    if err != nil {
        return err
    }

    Patch := &vspk.Patch{}

    if len(filteredPatchs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPatchs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Patch = filteredPatchs[0]

    d.Set("name", Patch.Name)
    d.Set("last_updated_by", Patch.LastUpdatedBy)
    d.Set("patch_build_number", Patch.PatchBuildNumber)
    d.Set("patch_summary", Patch.PatchSummary)
    d.Set("patch_tag", Patch.PatchTag)
    d.Set("patch_version", Patch.PatchVersion)
    d.Set("description", Patch.Description)
    d.Set("entity_scope", Patch.EntityScope)
    d.Set("supports_deletion", Patch.SupportsDeletion)
    d.Set("supports_network_acceleration", Patch.SupportsNetworkAcceleration)
    d.Set("external_id", Patch.ExternalID)
    
    d.Set("id", Patch.Identifier())
    d.Set("parent_id", Patch.ParentID)
    d.Set("parent_type", Patch.ParentType)
    d.Set("owner", Patch.Owner)

    d.SetId(Patch.Identifier())
    
    return nil
}