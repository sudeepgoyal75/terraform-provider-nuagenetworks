package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceFirewallAcl() *schema.Resource {
    return &schema.Resource{
        Create: resourceFirewallAclCreate,
        Read:   resourceFirewallAclRead,
        Update: resourceFirewallAclUpdate,
        Delete: resourceFirewallAclDelete,
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
            "active": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "default_allow_ip": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "default_allow_non_ip": &schema.Schema{
                Type:     schema.TypeBool,
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
            "rule_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "auto_generate_priority": &schema.Schema{
                Type:     schema.TypeBool,
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

func resourceFirewallAclCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize FirewallAcl object
    o := &vspk.FirewallAcl{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_ip"); ok {
        o.DefaultAllowIP = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_non_ip"); ok {
        o.DefaultAllowNonIP = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("rule_ids"); ok {
        o.RuleIds = attr.([]interface{})
    }
    if attr, ok := d.GetOk("auto_generate_priority"); ok {
        o.AutoGeneratePriority = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateFirewallAcl(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceFirewallAclRead(d, m)
}

func resourceFirewallAclRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FirewallAcl{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("active", o.Active)
    d.Set("default_allow_ip", o.DefaultAllowIP)
    d.Set("default_allow_non_ip", o.DefaultAllowNonIP)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("rule_ids", o.RuleIds)
    d.Set("auto_generate_priority", o.AutoGeneratePriority)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceFirewallAclUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FirewallAcl{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_ip"); ok {
        o.DefaultAllowIP = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_non_ip"); ok {
        o.DefaultAllowNonIP = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("rule_ids"); ok {
        o.RuleIds = attr.([]interface{})
    }
    if attr, ok := d.GetOk("auto_generate_priority"); ok {
        o.AutoGeneratePriority = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceFirewallAclDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FirewallAcl{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}