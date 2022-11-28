EndPoints

GET "/service/city?limit= ?offset= " - get all cities

GET "/service/city/{id}" - get city by id

POST "/service/city" - add city

PUT "service/city" - update city

DELETE "service/city" - delete city


Model
{"Id":0,"Name":"","Сarbases":[]}
{"Location":"","Sort":0,"Automobiles":[]}
{"Model":"","Brand":""}

JSON
{
  "Id":0,
  "Name":"Town1",
  "Сarbases":[
    {
      "Location":"A1",
      "Sort":0,
      "Automobiles":[
        {
          "Model":"AX",
          "Brand":""
        },
        {
          "Model":"AZ",
          "Brand":""
        }
      ]
    },
    {
      "Location":"B1",
      "Sort":0,
      "Automobiles":[
        {
          "Model":"BX",
          "Brand":""
        },
        {
          "Model":"BZ",
          "Brand":""
        }
      ]
    }
  ]
}


Data
{"Name":"Town1","Сarbases":[
  {"Location":"A1","Sort":2,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"A1","Sort":1,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"A1","Sort":4,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"A1","Sort":8,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"A1","Sort":5,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"B1","Sort":1,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}
  ]}

{"Name":"Town2","Сarbases":[
  {"Location":"A2","Sort":6,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"A2","Sort":1,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"A2","Sort":4,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},
  {"Location":"B2","Sort":8,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}
  ]}

{"Name":"Town3","Сarbases":[]}

{"Name":"Town4","Сarbases":[{"Location":"A4","Sort":4,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B4","Sort":3,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}

{"Name":"Town5","Сarbases":[{"Location":"A5","Sort":1,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B5","Sort"2:,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}

{"Name":"Town6","Сarbases":[{"Location":"A6","Sort":6,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B6","Sort":6,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}

{"Name":"Town7","Сarbases":[{"Location":"A7","Sort":7,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B7","Sort":7,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}

{"Name":"Town8","Сarbases":[{"Location":"A8","Sort":8,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B8","Sort":8,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}

{"Name":"Town9","Сarbases":[{"Location":"A9","Sort":9,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B1","Sort":9,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}

{"Name":"Town10","Сarbases":[{"Location":"A10","Sort":10,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]},{"Location":"B10","Sort":10,"Automobiles":[{"Model":"","Brand":""},{"Model":"","Brand":""}]}]}



	// mod := cache.Instance.Get(fmt.Sprintf("%d", city.Id))
	// if mod != nil {
	// 	city = mod.(*entity.City)
	// 	return city, nil
	// }
