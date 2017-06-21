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
* ?keyword="21202"
* ?token="877387487387"

===
### RESPONSE
```json
{"status":"success",
  "data":[
    {"item_code":"2120225","item_name":"ฟลอร์เดรน 4\" No.160","unit_code":"ตัว","stock_qty":0,"stock_list":null,"units":[{"unit_id":14311,"unit_code":"ตัว","unit_name":"ตัว","packing_rate":1,"price":595}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""},{"item_code":"2120226","item_name":"ฟลอร์เดรน 6\" KNACK No.232B","unit_code":"ตัว","stock_qty":0,"stock_list":null,"units":[{"unit_id":14315,"unit_code":"ตัว","unit_name":"ตัว","packing_rate":1,"price":0}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""},
    {"item_code":"2120227","item_name":"ฟลอร์เดรน 2\" กันกลิ่น KNACK No.224B","unit_code":"ตัว","stock_qty":0,"stock_list":[{"qty":0,"unit_code":"ตัว","wh_code":"S1-SPO","shelf_code":"-"}],"units":[{"unit_id":6127822,"unit_code":"ตัว","unit_name":"ตัว","packing_rate":1,"price":560}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""},
    {"item_code":"2120228","item_name":"ฟลอร์เดรน 2\" ไม่กันกลิ่น KNACK No.224P","unit_code":"ตัว","stock_qty":0,"stock_list":[{"qty":0,"unit_code":"ตัว","wh_code":"S1-SPO","shelf_code":"-"}],"units":[{"unit_id":6127826,"unit_code":"ตัว","unit_name":"ตัว","packing_rate":1,"price":510}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""},
    {"item_code":"2120250","item_name":"น้ำยาเชื่อมท่อ PVC 250 กรัม","unit_code":"กระป๋อง","stock_qty":989,"stock_list":[{"qty":16,"unit_code":"กระป๋อง","wh_code":"S1-A","shelf_code":"-"},{"qty":240,"unit_code":"กระป๋อง","wh_code":"S1-B","shelf_code":"-"},{"qty":1,"unit_code":"กระป๋อง","wh_code":"S1-OFS","shelf_code":"-"},{"qty":0,"unit_code":"กระป๋อง","wh_code":"S1-SPO","shelf_code":"-"},{"qty":732,"unit_code":"กระป๋อง","wh_code":"S2-A","shelf_code":"-"}],"units":[{"unit_id":6475949,"unit_code":"กระป๋อง","unit_name":"กระป๋อง","packing_rate":1,"price":116},{"unit_id":6475953,"unit_code":"กล่อง","unit_name":"กล่อง","packing_rate":20,"price":2313}],"so_qty":52,"ro_qty":0,"po_qty":180,"my_grade":"A","comm":0,"img_profile":"http://qserver.nopadol.com/pictures/item/f8d71e76487b48b1e22280cce3fde5e3.jpg"},{"item_code":"5521202","item_name":"สกรูเกลียว FASTANIC #12*2\"2  กลมx50ตัว","unit_code":"แพ็ค","stock_qty":0,"stock_list":null,"units":[{"unit_id":4637391,"unit_code":"แพ็ค","unit_name":"แพ็ค","packing_rate":1,"price":35}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"C","comm":0,"img_profile":""},{"item_code":"8851123212021","item_name":"เครื่องดื่มบำรุงกำลัง M-150","unit_code":"ขวด","stock_qty":5,"stock_list":[{"qty":5,"unit_code":"ขวด","wh_code":"S2-NIV","shelf_code":"-"}],"units":[{"unit_id":490155,"unit_code":"ขวด","unit_name":"ขวด","packing_rate":1,"price":9}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""},{"item_code":"8851952120269","item_name":"น้ำอัดลมเอส ส้ม ขนาด 455 มล.","unit_code":"ขวด","stock_qty":0,"stock_list":null,"units":[{"unit_id":4669276,"unit_code":"ขวด","unit_name":"ขวด","packing_rate":1,"price":12}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""},{"item_code":"9212020","item_name":"ที่ใส่กระดาษชำระแบบแขวน","unit_code":"ชิ้น","stock_qty":0,"stock_list":null,"units":[{"unit_id":96895,"unit_code":"ชิ้น","unit_name":"ชิ้น","packing_rate":1,"price":100}],"so_qty":0,"ro_qty":0,"po_qty":0,"my_grade":"-","comm":0,"img_profile":""}]}
```
* Status: 200
* Content-Type: "application/json; charset=utf-8"