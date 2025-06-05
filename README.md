# BEWELL-APP - ระบบประมวลผลคำสั่งซื้อ

## โครงสร้างโปรเจค

BEWELL-APP/
├── api/
│ └── order_handler.go # จัดการการรับ-ส่งคำสั่งซื้อ
├── config/
│ └── config.go # ตั้งค่าการทำงานของแอป
├── entities/
│ └── order.go # โมเดลข้อมูลคำสั่งซื้อ
├── routes/
│ └── routes.go # กำหนดเส้นทาง API
├── services/
│ └── order_service.go # โลจิกการประมวลผล
├── utils/
│ ├── format.go # ฟังก์ชันจัดรูปแบบข้อมูล
│ └── response.go # การตอบกลับ API
├── go.mod # ไฟล์ระบุ dependencies
├── go.sum # ไฟล์ตรวจสอบความถูกต้อง
└── main.go # จุดเริ่มต้นแอปพลิเคชัน


## วิธีการใช้งาน

1. **ติดตั้งและรันเซิร์ฟเวอร์**:
- run คำสั่ง go run main.go

2. **ทดสอบAPIด้วยPostman**:

- นำเข้าไฟล์ beWellAssign.postman_collection.json เข้า Postman
- เลือกกรณีทดสอบที่ต้องการจาก Collection
- ยิง method post จากแต่ละ case ได้เลย

2. **ทดสอบ ด้วย  UNITEST**:

- cd ไปที่ path \bewell-app\services ที่มี file services\order_service_test.go
- run คำสั่ง go test
- ถ้า test ผ่าน จะขึ้น ok ถ้า test ไม่ผ่าน จะมีรายละเอียด error แสดงใน terminal

NOTE => Case 7 : one product and one bundle product with wrong prefix and have / symbol and * ยังไม่ถูกต้อง