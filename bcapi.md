# API 

## GET /item

* Content-Type: "application/json" 
* ?keyword="2120250"
* ?token="877387487387"

===
### RESPONSE
```json
{"status":"success",
  "data":{
    "item_code":"2120250","item_name":"น้ำยาเชื่อมท่อ PVC 250 กรัม","unit_code":"กระป๋อง",
    "stock_qty":989,
    "stock_list":
      [
        {"qty":16,"unit_code":"กระป๋อง","wh_code":"S1-A","shelf_code":"-"},
        {"qty":240,"unit_code":"กระป๋อง","wh_code":"S1-B","shelf_code":"-"},
        {"qty":1,"unit_code":"กระป๋อง","wh_code":"S1-OFS","shelf_code":"-"},
        {"qty":0,"unit_code":"กระป๋อง","wh_code":"S1-SPO","shelf_code":"-"},
        {"qty":732,"unit_code":"กระป๋อง","wh_code":"S2-A","shelf_code":"-"}],
    "units":[
      {"unit_id":6475949,"unit_code":"กระป๋อง","unit_name":"กระป๋อง","packing_rate":1,"price":116},
      {"unit_id":6475953,"unit_code":"กล่อง","unit_name":"กล่อง","packing_rate":20,"price":2313}],
    "so_qty":52,"ro_qty":0,"po_qty":180,"my_grade":"A","comm":0,
    "img_profile":"http://qserver.nopadol.com/pictures/item/f8d71e76487b48b1e22280cce3fde5e3.jpg"
    }
}
```
* Status: 200
* Content-Type: "application/json; charset=utf-8"

===
## GET /items

* Content-Type: "application/json" 
* ?keyword="2120250"
* ?token="877387487387"

===
### RESPONSE
```json
{"status":"success","data":[{"item_code":"2120250","item_name":"น้ำยาเชื่อมท่อ PVC 250 กรัม","unit_code":"กระป๋อง","stock_qty":989,"stock_list":[{"qty":16,"unit_code":"กระป๋อง","wh_code":"S1-A","shelf_code":"-"},{"qty":240,"unit_code":"กระป๋อง","wh_code":"S1-B","shelf_code":"-"},{"qty":1,"unit_code":"กระป๋อง","wh_code":"S1-OFS","shelf_code":"-"},{"qty":0,"unit_code":"กระป๋อง","wh_code":"S1-SPO","shelf_code":"-"},{"qty":732,"unit_code":"กระป๋อง","wh_code":"S2-A","shelf_code":"-"}],"units":[{"unit_id":6475949,"unit_code":"กระป๋อง","unit_name":"กระป๋อง","packing_rate":1,"price":116},{"unit_id":6475953,"unit_code":"กล่อง","unit_name":"กล่อง","packing_rate":20,"price":2313}],"so_qty":52,"ro_qty":0,"po_qty":180,"my_grade":"A","comm":0,"img_profile":"http://qserver.nopadol.com/pictures/item/f8d71e76487b48b1e22280cce3fde5e3.jpg"}]}
```
* Status: 200
* Content-Type: "application/json; charset=utf-8"