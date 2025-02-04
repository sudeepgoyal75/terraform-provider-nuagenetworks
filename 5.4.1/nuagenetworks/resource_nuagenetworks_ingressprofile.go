package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceIngressProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceIngressProfileCreate,
        Read:   resourceIngressProfileRead,
        Update: resourceIngressProfileUpdate,
        Delete: resourceIngressProfileDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "assoc_entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ip_filter_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ip_filter_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ipv6_filter_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ipv6_filter_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_mac_filter_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_mac_filter_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_sap_ingress_qo_s_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_sap_ingress_qo_s_profile_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_gateway"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redundancy_group"},
            },
        },
    }
}

func resourceIngressProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IngressProfile object
    o := &vspk.IngressProfile{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ip_filter_profile_id"); ok {
        o.AssociatedIPFilterProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ipv6_filter_profile_id"); ok {
        o.AssociatedIPv6FilterProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_mac_filter_profile_id"); ok {
        o.AssociatedMACFilterProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_sap_ingress_qo_s_profile_id"); ok {
        o.AssociatedSAPIngressQoSProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        err := parent.CreateIngressProfile(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        err := parent.CreateIngressProfile(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceIngressProfileRead(d, m)
}

func resourceIngressProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressProfile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("assoc_entity_type", o.AssocEntityType)
    d.Set("associated_ip_filter_profile_id", o.AssociatedIPFilterProfileID)
    d.Set("associated_ip_filter_profile_name", o.AssociatedIPFilterProfileName)
    d.Set("associated_ipv6_filter_profile_id", o.AssociatedIPv6FilterProfileID)
    d.Set("associated_ipv6_filter_profile_name", o.AssociatedIPv6FilterProfileName)
    d.Set("associated_mac_filter_profile_id", o.AssociatedMACFilterProfileID)
    d.Set("associated_mac_filter_profile_name", o.AssociatedMACFilterProfileName)
    d.Set("associated_sap_ingress_qo_s_profile_id", o.AssociatedSAPIngressQoSProfileID)
    d.Set("associated_sap_ingress_qo_s_profile_name", o.AssociatedSAPIngressQoSProfileName)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIngressProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ip_filter_profile_id"); ok {
        o.AssociatedIPFilterProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ipv6_filter_profile_id"); ok {
        o.AssociatedIPv6FilterProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_mac_filter_profile_id"); ok {
        o.AssociatedMACFilterProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_sap_ingress_qo_s_profile_id"); ok {
        o.AssociatedSAPIngressQoSProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIngressProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}