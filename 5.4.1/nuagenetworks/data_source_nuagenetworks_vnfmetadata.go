package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVNFMetadata() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVNFMetadataRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "blob": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vnf"},
            },
            "parent_vnf": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourceVNFMetadataRead(d *schema.ResourceData, m interface{}) error {
    filteredVNFMetadatas := vspk.VNFMetadatasList{}
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
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredVNFMetadatas, err = parent.VNFMetadatas(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vnf"); ok {
        parent := &vspk.VNF{ID: attr.(string)}
        filteredVNFMetadatas, err = parent.VNFMetadatas(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredVNFMetadatas, err = parent.VNFMetadatas(fetchFilter)
        if err != nil {
            return err
        }
    }

    VNFMetadata := &vspk.VNFMetadata{}

    if len(filteredVNFMetadatas) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVNFMetadatas) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VNFMetadata = filteredVNFMetadatas[0]

    d.Set("name", VNFMetadata.Name)
    d.Set("last_updated_by", VNFMetadata.LastUpdatedBy)
    d.Set("description", VNFMetadata.Description)
    d.Set("blob", VNFMetadata.Blob)
    d.Set("entity_scope", VNFMetadata.EntityScope)
    d.Set("assoc_entity_type", VNFMetadata.AssocEntityType)
    d.Set("external_id", VNFMetadata.ExternalID)
    
    d.Set("id", VNFMetadata.Identifier())
    d.Set("parent_id", VNFMetadata.ParentID)
    d.Set("parent_type", VNFMetadata.ParentType)
    d.Set("owner", VNFMetadata.Owner)

    d.SetId(VNFMetadata.Identifier())
    
    return nil
}