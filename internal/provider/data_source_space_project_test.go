package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	spaceProjectDataSourceTestResourceId = "3U65GM4NKw79"
)

func TestAccDataSourceSpaceProject(t *testing.T) {
	//t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSpaceProject(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						fmt.Sprintf("data.%s.test", spaceProjectDataSourceResourceName), "id", regexp.MustCompile("")),
				),
			},
		},
	})
}

func testAccDataSourceSpaceProject() string {
	return fmt.Sprintf(`
data %s "test" {
  id = "%s"
}
`,
		spaceProjectDataSourceResourceName,
		spaceProjectDataSourceTestResourceId,
	)
}
