package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.3"
)

func resourcePATNATPool() *schema.Resource {
    return &schema.Resource{
        Create: resourcePATNATPoolCreate,
        Read:   resourcePATNATPoolRead,
        Update: resourcePATNATPoolUpdate,
        Delete: resourcePATNATPoolDelete,
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
            "address_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "default_patip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "end_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "end_source_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "translation_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "start_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "start_source_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dynamic_source_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
        },
    }
}

func resourcePATNATPoolCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PATNATPool object
    o := &vspk.PATNATPool{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("address_range"); ok {
        o.AddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("default_patip"); ok {
        o.DefaultPATIP = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("end_address_range"); ok {
        o.EndAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("end_source_address"); ok {
        o.EndSourceAddress = attr.(string)
    }
    if attr, ok := d.GetOk("translation_timeout"); ok {
        o.TranslationTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_gateway_type"); ok {
        o.AssociatedGatewayType = attr.(string)
    }
    if attr, ok := d.GetOk("start_address_range"); ok {
        o.StartAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("start_source_address"); ok {
        o.StartSourceAddress = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_source_enabled"); ok {
        o.DynamicSourceEnabled = attr.(bool)
    }
    parent := m.(*vspk.Me)
    err := parent.CreatePATNATPool(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourcePATNATPoolRead(d, m)
}

func resourcePATNATPoolRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PATNATPool{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("address_range", o.AddressRange)
    d.Set("default_patip", o.DefaultPATIP)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("end_address_range", o.EndAddressRange)
    d.Set("end_source_address", o.EndSourceAddress)
    d.Set("entity_scope", o.EntityScope)
    d.Set("translation_timeout", o.TranslationTimeout)
    d.Set("associated_gateway_id", o.AssociatedGatewayId)
    d.Set("associated_gateway_type", o.AssociatedGatewayType)
    d.Set("start_address_range", o.StartAddressRange)
    d.Set("start_source_address", o.StartSourceAddress)
    d.Set("external_id", o.ExternalID)
    d.Set("dynamic_source_enabled", o.DynamicSourceEnabled)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePATNATPoolUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PATNATPool{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("address_range"); ok {
        o.AddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("default_patip"); ok {
        o.DefaultPATIP = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("end_address_range"); ok {
        o.EndAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("end_source_address"); ok {
        o.EndSourceAddress = attr.(string)
    }
    if attr, ok := d.GetOk("translation_timeout"); ok {
        o.TranslationTimeout = attr.(int)
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_gateway_type"); ok {
        o.AssociatedGatewayType = attr.(string)
    }
    if attr, ok := d.GetOk("start_address_range"); ok {
        o.StartAddressRange = attr.(string)
    }
    if attr, ok := d.GetOk("start_source_address"); ok {
        o.StartSourceAddress = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_source_enabled"); ok {
        o.DynamicSourceEnabled = attr.(bool)
    }

    o.Save()

    return nil
}

func resourcePATNATPoolDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PATNATPool{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}