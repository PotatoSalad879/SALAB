ENG23 3031 System Analysis and Design

การบ้าน/ปฏิบัติการ (Assignment / Lab) — ระดับข้อสอบ (Advanced)

หัวข้อ: Relational Database Mapping ด้วย Go Structs & GORM Tags — ระบบสำรองที่นั่งสายการบิน (SkyLink Airline Reservation System)

ผลลัพธ์การเรียนรู้: CLO2 — ออกแบบและ implement การ mapping ระหว่างแบบจำลองข้อมูลเชิงแนวคิดกับฐานข้อมูลเชิงสัมพันธ์ (สอดคล้อง PLO2)

---

## วัตถุประสงค์ (Objectives)

1. วิเคราะห์ข้อกำหนดของระบบและ Database Schema Specification แล้วแปลงเป็นโครงสร้าง Go Structs
2. กำหนด Primary Key, Foreign Key, Composite Key, Composite Foreign Key, Natural Key, Unique Constraint, Nullability, Soft Delete และความสัมพันธ์ (1:N, M:N, Self-Reference, Polymorphic) ด้วย GORM Tags ให้ถูกต้องตามหลักการออกแบบฐานข้อมูลเชิงสัมพันธ์

---

## ขั้นตอนการทำงาน (Instructions)

1. อ่านและทำความเข้าใจ **Database Schema Specification** และ **ความสัมพันธ์ระหว่างเอนทิตี** ด้านล่างให้ครบถ้วน
2. แก้ไขไฟล์ในโฟลเดอร์ `internal/models/` เท่านั้น โดยปรับ Go Structs และ GORM Tags ให้ตรงตามข้อกำหนด
   - โค้ดตั้งต้นยัง map ความสัมพันธ์ไม่ครบ/ไม่ถูกต้อง และ **ไม่มีคอมเมนต์ระบุจุดที่ต้องแก้**
   - ห้ามแก้ไฟล์ในโฟลเดอร์ `tests/`
3. รันชุดทดสอบจนผ่านครบทุกกรณี (`PASS` ทั้ง 10 ข้อ)

```bash
go test -v ./tests/... -count=1
```

รันทีละข้อ:

```bash
go test -v ./tests/ -run TestFlightAirportForeignKeys -count=1
```

รัน migration จริงลง PostgreSQL (หลังแก้ครบ ไม่บังคับ):

```bash
docker compose up -d
go run ./cmd/server
```

---

## สถานการณ์ระบบ (System Scenario)

ระบบสำรองที่นั่งสายการบิน "SkyLink" ให้บริการจัดการสนามบิน อากาศยาน เที่ยวบิน ลูกเรือ ผู้โดยสาร การจอง และบัตรโดยสาร โดยมีข้อกำหนดทางธุรกิจดังนี้

- **สนามบิน** ถูกระบุด้วยรหัส IATA 3 ตัวอักษร (เช่น `BKK`, `HKT`, `CNX`) ซึ่งถือเป็นกุญแจหลักตามธรรมชาติ (natural key) ของสนามบิน ระบบนี้ไม่ใช้เลขรัน (surrogate id) เป็นกุญแจหลักของสนามบิน
- **ที่นั่ง** บนอากาศยานถูกระบุตัวตนด้วย "อากาศยานลำใด" ร่วมกับ "หมายเลขที่นั่งใด" หมายเลขที่นั่ง (เช่น `12A`) สามารถซ้ำกันได้เมื่ออยู่คนละลำ แต่ห้ามซ้ำภายในอากาศยานลำเดียวกัน
- **เที่ยวบิน** หนึ่งเที่ยวถูกระบุด้วย "หมายเลขเที่ยวบิน" ร่วมกับ "วันที่ให้บริการ" (เที่ยวบินหมายเลขเดียวกันแต่คนละวันถือเป็นคนละเที่ยวบิน) แต่ละเที่ยวบินออกเดินทางจากสนามบินต้นทาง 1 แห่ง ไปยังสนามบินปลายทาง 1 แห่ง โดยอ้างอิงด้วยรหัส IATA ของสนามบิน และใช้อากาศยาน 1 ลำ
- **ลูกเรือ** ทุกคนถูกจัดเก็บในตารางเดียว (Single Table Inheritance) จำแนกบทบาทด้วยประเภทลูกเรือ ได้แก่ นักบิน (pilot) และพนักงานต้อนรับ (attendant) ข้อมูลเฉพาะของแต่ละบทบาทต้องเว้นว่างได้เมื่อบันทึกลูกเรืออีกบทบาทหนึ่ง
- ระบบมอบหมายลูกเรือเข้าประจำเที่ยวบินผ่านรายการ **การมอบหมายลูกเรือ** ซึ่งเป็นความสัมพันธ์ระหว่างเที่ยวบินกับลูกเรือแบบกลุ่มต่อกลุ่ม (M:N) และต้องบันทึกหน้าที่ของลูกเรือในเที่ยวบินนั้นเพิ่มเติม กุญแจหลักของรายการมอบหมายคือ "เที่ยวบิน (หมายเลข + วันที่)" ร่วมกับ "ลูกเรือ" และห้ามมอบหมายลูกเรือคนเดิมเข้าเที่ยวบินเดิมซ้ำซ้อน
- **ผู้โดยสาร** ถูกระบุด้วยเลขหนังสือเดินทาง ผู้โดยสารสามารถระบุผู้ร่วมเดินทาง (companions) ซึ่งเป็นผู้โดยสารในระบบเช่นกันได้หลายคน และผู้โดยสารหนึ่งคนก็เป็นผู้ร่วมเดินทางของผู้อื่นได้ ความสัมพันธ์นี้เป็นแบบกลุ่มต่อกลุ่มที่อ้างอิงตารางผู้โดยสารเอง (self-referencing M:N) ผ่านตารางเชื่อม `passenger_companions` ที่ประกอบด้วยคอลัมน์ `passenger_id` และ `companion_id`
- เจ้าหน้าที่สามารถแนบ **บันทึกภายใน** เข้ากับ "ผู้โดยสาร" หรือ "การจอง" ก็ได้ (polymorphic association) โดยเก็บชนิดของเจ้าของไว้ที่คอลัมน์ `owner_type` และรหัสของเจ้าของไว้ที่คอลัมน์ `owner_id`
- **การจอง** หนึ่งรายการเป็นของผู้โดยสาร 1 คน และประกอบด้วยบัตรโดยสารได้หลายใบ รหัสการจอง (`ref`) 6 หลักต้องไม่ซ้ำกันในระบบ เมื่อการจองถูกยกเลิก ระบบต้องไม่ลบข้อมูลออกจริง แต่ทำเป็นการลบแบบนุ่มนวล (soft delete) เพื่อเก็บประวัติไว้ตรวจสอบย้อนหลัง
- **บัตรโดยสาร** หนึ่งใบอยู่ภายใต้การจอง 1 รายการ และผูกกับเที่ยวบิน 1 เที่ยว การอ้างอิงเที่ยวบินต้องใช้ "หมายเลขเที่ยวบิน" ร่วมกับ "วันที่ให้บริการ" เป็น Foreign Key แบบหลายคอลัมน์ ให้ตรงกับกุญแจหลักของเที่ยวบิน

---

## Database Schema Specification

> รายการด้านล่างคือ "แอตทริบิวต์ประจำเอนทิตี" ที่ต้องมี — ส่วนกุญแจหลัก, Foreign Key, ความสัมพันธ์
> และ Constraint ต่าง ๆ ให้พิจารณาจากสถานการณ์ระบบด้านบนและกรณีทดสอบ

### สนามบิน (Airport)
- รหัสสนามบิน IATA (code)
- ชื่อสนามบิน (name)
- เมืองที่ตั้ง (city)

### อากาศยาน (Aircraft) — โครงสร้างถูกต้องแล้ว
- ทะเบียนอากาศยาน (registration)
- รุ่นอากาศยาน (model)

### ที่นั่ง (Seat)
- หมายเลขที่นั่ง (seat_no)
- ชั้นโดยสาร (class)
- อ้างอิงถึงอากาศยานที่ที่นั่งนั้นอยู่ (aircraft)

### เที่ยวบิน (Flight)
- หมายเลขเที่ยวบิน (flight_no)
- วันที่ให้บริการ (service_date)
- สนามบินต้นทาง (origin) — อ้างอิงด้วยรหัส IATA
- สนามบินปลายทาง (destination) — อ้างอิงด้วยรหัส IATA
- อากาศยานที่ใช้บิน (aircraft)

### ลูกเรือ (CrewMember)
- รหัสพนักงาน (staff_no)
- ชื่อ-นามสกุล (full_name)
- ประเภทลูกเรือ (crew_role)

เฉพาะนักบิน (pilot):
- เลขที่ใบอนุญาตนักบิน (license_no)
- วันหมดอายุใบรับรองแพทย์ (medical_expiry)

เฉพาะพนักงานต้อนรับ (attendant):
- ภาษาที่ให้บริการได้ (cabin_languages)
- ระดับการบริการ (service_grade)

### การมอบหมายลูกเรือเข้าเที่ยวบิน (FlightCrew)
- หน้าที่ในเที่ยวบิน (duty)
- อ้างอิงถึงเที่ยวบิน (หมายเลขเที่ยวบิน + วันที่ให้บริการ)
- อ้างอิงถึงลูกเรือ (crew_member)

### ผู้โดยสาร (Passenger)
- เลขหนังสือเดินทาง (passport_no)
- ชื่อ-นามสกุล (full_name)
- ผู้ร่วมเดินทาง (companions) — เป็นผู้โดยสารด้วยกัน

### บันทึกภายใน (Note)
- ชนิดของเจ้าของบันทึก (owner_type)
- รหัสของเจ้าของบันทึก (owner_id)
- เนื้อหาบันทึก (body)

### การจอง (Booking)
- รหัสการจอง (ref)
- สถานะการจอง (status)
- อ้างอิงถึงผู้โดยสารเจ้าของการจอง (passenger)
- ข้อมูลการลบแบบนุ่มนวล (deleted_at)

### บัตรโดยสาร (Ticket)
- ค่าโดยสาร (fare_amount)
- อ้างอิงถึงการจอง (booking)
- อ้างอิงถึงเที่ยวบิน (หมายเลขเที่ยวบิน + วันที่ให้บริการ)

---

## กรณีทดสอบ (Test Cases) — 10 ข้อ

| # | ไฟล์ | ประเด็นที่ตรวจ |
|---|---|---|
| 01 | `01_airport_natural_key_test.go` | Natural key เป็น string (`code`) และไม่มี surrogate `id` |
| 02 | `02_seat_composite_pk_test.go` | Composite Primary Key `(aircraft_id, seat_no)` |
| 03 | `03_flight_composite_pk_test.go` | Composite Primary Key `(flight_no, service_date)` |
| 04 | `04_flight_airport_fk_test.go` | Foreign Key ต้นทาง/ปลายทาง อ้างอิง `Airport.code` (`references:`) |
| 05 | `05_crew_sti_test.go` | Single Table Inheritance + Nullability ของ subtype |
| 06 | `06_flight_crew_join_test.go` | M:N Join + Composite PK 3 คอลัมน์ + คอลัมน์ `duty` |
| 07 | `07_passenger_selfref_m2m_test.go` | Self-Referencing Many-to-Many (`passenger_companions`) |
| 08 | `08_polymorphic_notes_test.go` | Polymorphic Association (`owner_id`, `owner_type`) |
| 09 | `09_booking_softdelete_test.go` | Soft Delete (`deleted_at`) + `ref` Unique |
| 10 | `10_ticket_composite_fk_test.go` | Composite Foreign Key `(flight_no, service_date)` → Flight |

---

## โครงสร้างโปรเจกต์

- `internal/models/` — Go Structs กำหนด Relational Schema **(ไฟล์ที่ต้องแก้)**
- `internal/db/` — Database Connection และ `AutoMigrate`
- `cmd/server/` — Entrypoint สำหรับรัน migration
- `tests/` — In-Memory Schema Validation (ไม่ต้องเชื่อมต่อฐานข้อมูลจริง)

> `go build ./...` ผ่านตั้งแต่สถานะเริ่มต้น — ใช้ `go test ./tests/...` เป็นเกณฑ์ตรวจงานหลัก
