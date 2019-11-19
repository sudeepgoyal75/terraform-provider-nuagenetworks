package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
)

func resourceFlowSecurityPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceFlowSecurityPolicyCreate,
        Read:   resourceFlowSecurityPolicyRead,
        Update: resourceFlowSecurityPolicyUpdate,
        Delete: resourceFlowSecurityPolicyDelete,
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
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_address_overwrite": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "flow_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_address_overwrite": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "associated_application_service_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
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
            "parent_flow": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceFlowSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize FlowSecurityPolicy object
    o := &vspk.FlowSecurityPolicy{
    }
    if attr, ok := d.GetOk("action"); ok {
        o.Action = attr.(string)
    }
    if attr, ok := d.GetOk("destination_address_overwrite"); ok {
        o.DestinationAddressOverwrite = attr.(string)
    }
    if attr, ok := d.GetOk("flow_id"); ok {
        o.FlowID = attr.(string)
    }
    if attr, ok := d.GetOk("source_address_overwrite"); ok {
        o.SourceAddressOverwrite = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("associated_application_service_id"); ok {
        o.AssociatedApplicationServiceID = attr.(string)
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
    parent := &vspk.Flow{ID: d.Get("parent_flow").(string)}
    err := parent.CreateFlowSecurityPolicy(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceFlowSecurityPolicyRead(d, m)
}

func resourceFlowSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FlowSecurityPolicy{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("action", o.Action)
    d.Set("destination_address_overwrite", o.DestinationAddressOverwrite)
    d.Set("flow_id", o.FlowID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("source_address_overwrite", o.SourceAddressOverwrite)
    d.Set("priority", o.Priority)
    d.Set("associated_application_service_id", o.AssociatedApplicationServiceID)
    d.Set("associated_network_object_id", o.AssociatedNetworkObjectID)
    d.Set("associated_network_object_type", o.AssociatedNetworkObjectType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceFlowSecurityPolicyUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FlowSecurityPolicy{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("action"); ok {
        o.Action = attr.(string)
    }
    if attr, ok := d.GetOk("destination_address_overwrite"); ok {
        o.DestinationAddressOverwrite = attr.(string)
    }
    if attr, ok := d.GetOk("flow_id"); ok {
        o.FlowID = attr.(string)
    }
    if attr, ok := d.GetOk("source_address_overwrite"); ok {
        o.SourceAddressOverwrite = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("associated_application_service_id"); ok {
        o.AssociatedApplicationServiceID = attr.(string)
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

func resourceFlowSecurityPolicyDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.FlowSecurityPolicy{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}