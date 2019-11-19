package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10.1"
)

func resourceExternalAppService() *schema.Resource {
    return &schema.Resource{
        Create: resourceExternalAppServiceCreate,
        Read:   resourceExternalAppServiceRead,
        Update: resourceExternalAppServiceUpdate,
        Delete: resourceExternalAppServiceDelete,
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
            "destination_nat_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_nat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "destination_nat_mask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "metadata": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "egress_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "virtual_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "virtual_ip_required": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ingress_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_nat_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_nat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "associated_service_egress_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_service_egress_redirect_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_service_ingress_group_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_service_ingress_redirect_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain"},
            },
        },
    }
}

func resourceExternalAppServiceCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ExternalAppService object
    o := &vspk.ExternalAppService{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_nat_address"); ok {
        o.DestinationNATAddress = attr.(string)
    }
    if attr, ok := d.GetOk("destination_nat_enabled"); ok {
        o.DestinationNATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("destination_nat_mask"); ok {
        o.DestinationNATMask = attr.(string)
    }
    if attr, ok := d.GetOk("metadata"); ok {
        o.Metadata = attr.(string)
    }
    if attr, ok := d.GetOk("egress_type"); ok {
        o.EgressType = attr.(string)
    }
    if attr, ok := d.GetOk("virtual_ip"); ok {
        o.VirtualIP = attr.(string)
    }
    if attr, ok := d.GetOk("virtual_ip_required"); ok {
        o.VirtualIPRequired = attr.(bool)
    }
    if attr, ok := d.GetOk("ingress_type"); ok {
        o.IngressType = attr.(string)
    }
    if attr, ok := d.GetOk("source_nat_address"); ok {
        o.SourceNATAddress = attr.(string)
    }
    if attr, ok := d.GetOk("source_nat_enabled"); ok {
        o.SourceNATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_service_egress_group_id"); ok {
        o.AssociatedServiceEgressGroupID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_service_egress_redirect_id"); ok {
        o.AssociatedServiceEgressRedirectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_service_ingress_group_id"); ok {
        o.AssociatedServiceIngressGroupID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_service_ingress_redirect_id"); ok {
        o.AssociatedServiceIngressRedirectID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateExternalAppService(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        err := parent.CreateExternalAppService(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateExternalAppService(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceExternalAppServiceRead(d, m)
}

func resourceExternalAppServiceRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ExternalAppService{
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
    d.Set("destination_nat_address", o.DestinationNATAddress)
    d.Set("destination_nat_enabled", o.DestinationNATEnabled)
    d.Set("destination_nat_mask", o.DestinationNATMask)
    d.Set("metadata", o.Metadata)
    d.Set("egress_type", o.EgressType)
    d.Set("virtual_ip", o.VirtualIP)
    d.Set("virtual_ip_required", o.VirtualIPRequired)
    d.Set("ingress_type", o.IngressType)
    d.Set("entity_scope", o.EntityScope)
    d.Set("source_nat_address", o.SourceNATAddress)
    d.Set("source_nat_enabled", o.SourceNATEnabled)
    d.Set("associated_service_egress_group_id", o.AssociatedServiceEgressGroupID)
    d.Set("associated_service_egress_redirect_id", o.AssociatedServiceEgressRedirectID)
    d.Set("associated_service_ingress_group_id", o.AssociatedServiceIngressGroupID)
    d.Set("associated_service_ingress_redirect_id", o.AssociatedServiceIngressRedirectID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceExternalAppServiceUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ExternalAppService{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_nat_address"); ok {
        o.DestinationNATAddress = attr.(string)
    }
    if attr, ok := d.GetOk("destination_nat_enabled"); ok {
        o.DestinationNATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("destination_nat_mask"); ok {
        o.DestinationNATMask = attr.(string)
    }
    if attr, ok := d.GetOk("metadata"); ok {
        o.Metadata = attr.(string)
    }
    if attr, ok := d.GetOk("egress_type"); ok {
        o.EgressType = attr.(string)
    }
    if attr, ok := d.GetOk("virtual_ip"); ok {
        o.VirtualIP = attr.(string)
    }
    if attr, ok := d.GetOk("virtual_ip_required"); ok {
        o.VirtualIPRequired = attr.(bool)
    }
    if attr, ok := d.GetOk("ingress_type"); ok {
        o.IngressType = attr.(string)
    }
    if attr, ok := d.GetOk("source_nat_address"); ok {
        o.SourceNATAddress = attr.(string)
    }
    if attr, ok := d.GetOk("source_nat_enabled"); ok {
        o.SourceNATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_service_egress_group_id"); ok {
        o.AssociatedServiceEgressGroupID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_service_egress_redirect_id"); ok {
        o.AssociatedServiceEgressRedirectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_service_ingress_group_id"); ok {
        o.AssociatedServiceIngressGroupID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_service_ingress_redirect_id"); ok {
        o.AssociatedServiceIngressRedirectID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceExternalAppServiceDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ExternalAppService{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}