package admin

import (
	"html/template"

	"github.com/qor/admin"
)

type MyStruct struct {
	Name string
	Age  int64
}

func initFuncMap() {
	Admin.RegisterFuncMap("render_latest_order", renderLatestOrder)
	Admin.RegisterFuncMap("test_banner", renderTest)

}

func renderLatestOrder(context *admin.Context) template.HTML {
	var orderContext = context.NewResourceContext("Order")
	orderContext.Searcher.Pagination.PerPage = 5
	// orderContext.SetDB(orderContext.GetDB().Where("state in (?)", []string{"paid"}))

	if orders, err := orderContext.FindMany(); err == nil {
		return orderContext.Render("index/table", orders)
	}
	return template.HTML("")
}

func renderTest(context *admin.Context) template.HTML {
	// context = NewContext(map[string]interface{}{})
	// result := context.Result{"oleg": "oleg"}

	res := MyStruct{
		Name: "oleg",
	}

	return context.Render("jbo", res)
	// var orderContext = context.NewResourceContext("Order")
	// orderContext.Searcher.Pagination.PerPage = 5
	// // orderContext.SetDB(orderContext.GetDB().Where("state in (?)", []string{"paid"}))

	// if orders, err := orderContext.FindMany(); err == nil {
	// 	return orderContext.Render("index/table", orders)
	// }
	// return template.HTML("")
}
