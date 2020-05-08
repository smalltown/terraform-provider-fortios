package fortios

import (
	"fmt"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceJSONGenericAPI() *schema.Resource {
	return &schema.Resource{
		Create: resourceJSONGenericAPICreateUpdate,
		Read:   resourceJSONGenericAPIRead,
		Update: resourceJSONGenericAPICreateUpdate,
		Delete: resourceJSONGenericAPIDelete,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"specialparams": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"json": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceJSONGenericAPICreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Build input data by sdk
	i := &forticlient.JSONJSONGenericAPI{
		Path:          d.Get("path").(string),
		Method:        d.Get("method").(string),
		Specialparams: d.Get("specialparams").(string),
		Json:          d.Get("json").(string),
	}

	//Call process by sdk
	res, err := c.CreateJSONGenericAPI(i)
	if err != nil {
		return fmt.Errorf("Error creating json generic api: %s", err)
	}

	//Set index for d
	d.SetId("JsonGenericApi")
	d.Set("response", res)

	return nil
}

func resourceJSONGenericAPIDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceJSONGenericAPIRead(d *schema.ResourceData, m interface{}) error {
	return nil
}