package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
)

func resourceApp() *schema.Resource {
    return &schema.Resource{
        Create: resourceAppCreate,
        Read:   resourceAppRead,
        Update: resourceAppUpdate,
        Delete: resourceAppDelete,
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
                Required: true,
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
            "assoc_egress_acl_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "assoc_ingress_acl_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "associated_domain_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "associated_network_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_network_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceAppCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize App object
    o := &vspk.App{
        Name: d.Get("name").(string),
        AssociatedDomainID: d.Get("associated_domain_id").(string),
        AssociatedDomainType: d.Get("associated_domain_type").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_egress_acl_template_id"); ok {
        o.AssocEgressACLTemplateId = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_ingress_acl_template_id"); ok {
        o.AssocIngressACLTemplateId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_id"); ok {
        o.AssociatedNetworkObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_type"); ok {
        o.AssociatedNetworkObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateApp(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceAppRead(d, m)
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.App{
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
    d.Set("assoc_egress_acl_template_id", o.AssocEgressACLTemplateId)
    d.Set("assoc_ingress_acl_template_id", o.AssocIngressACLTemplateId)
    d.Set("associated_domain_id", o.AssociatedDomainID)
    d.Set("associated_domain_type", o.AssociatedDomainType)
    d.Set("associated_network_object_id", o.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", o.AssociatedNetworkObjectType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.App{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.AssociatedDomainID = d.Get("associated_domain_id").(string)
    o.AssociatedDomainType = d.Get("associated_domain_type").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_egress_acl_template_id"); ok {
        o.AssocEgressACLTemplateId = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_ingress_acl_template_id"); ok {
        o.AssocIngressACLTemplateId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_id"); ok {
        o.AssociatedNetworkObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_network_object_type"); ok {
        o.AssociatedNetworkObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.App{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}