# Practice Lab (Hard): Relational Mapping — LearnHub Online Course Platform

> ระดับยากขึ้น: มี **many-to-many แบบมี join table มาตรฐานของ GORM**, ความสัมพันธ์ **อ้างอิงตัวเอง (self-reference)**,
> และ **composite UNIQUE index**. โจทย์ไม่บอกว่าบั๊กอยู่ตรงไหน — ให้อ่าน Requirements + รัน Test แล้ววิเคราะห์เอง

## Scenario (สถานการณ์โจทย์)

แพลตฟอร์มคอร์สเรียนออนไลน์ "LearnHub" มีข้อกำหนดดังนี้:

- **ผู้ใช้ (User)** — เก็บผู้ใช้ทุกคนไว้ในตารางเดียว (Single Table Inheritance) แยกบทบาทด้วยคอลัมน์ `user_role`
  - **ผู้สอน (instructor)**: มีข้อมูลเฉพาะ ประวัติโดยย่อ (`bio`) และสาขาความเชี่ยวชาญ (`expertise`)
  - **ผู้เรียน (learner)**: มีข้อมูลเฉพาะ พาดหัวโปรไฟล์ (`headline`) และเป้าหมายการเรียน (`learning_goal`)
  - ข้อมูลเฉพาะของแต่ละบทบาทต้องเว้นว่างได้เมื่อบันทึกผู้ใช้อีกบทบาทหนึ่ง

- **หมวดหมู่ (Category)** — จัดเป็นลำดับชั้นแบบ **อ้างอิงตัวเอง**: หมวดหมู่หนึ่งอาจมี "หมวดหมู่แม่" 1 หมวด หรือเป็นหมวดระดับบนสุด (ไม่มีแม่) ก็ได้ ค่าหมวดแม่จึงต้องเว้นว่างได้ และหนึ่งหมวดหมู่มีคอร์สได้หลายคอร์ส (1:N)

- **คอร์ส (Course)** — แต่ละคอร์สอยู่ใน 1 หมวดหมู่ และมีผู้สอนหลัก 1 คน (อ้างไปตาราง `users`) โดยคอร์สเป็นผู้ถือ Foreign Key ทั้งสอง

- **แท็ก (Tag)** — คอร์สหนึ่งติดแท็กได้หลายอัน และแท็กหนึ่งใช้กับคอร์สได้หลายคอร์ส (**M:N**) ความสัมพันธ์นี้ไม่มีข้อมูลเพิ่มเติม จึงใช้ **ตารางเชื่อมมาตรฐานที่ GORM สร้างและจัดการให้เอง** โดยทั้งสองฝั่งต้องชี้ไปตารางเชื่อมชื่อเดียวกันคือ **`course_tags`**

- **การลงทะเบียนเรียน (Enrollment)** — ผู้เรียนลงทะเบียนได้หลายคอร์ส คอร์สมีผู้เรียนได้หลายคน (M:N) แต่ความสัมพันธ์นี้ **มีข้อมูลเพิ่ม** (ความคืบหน้าเป็น % , วันที่ลงทะเบียน) จึงทำเป็นตารางเชื่อมของตัวเอง โดยใช้ `(learner_id, course_id)` ร่วมกันเป็นคีย์หลัก และผู้เรียนลงทะเบียนคอร์สเดิมซ้ำไม่ได้

- **บทเรียน (Lesson)** — แต่ละคอร์สมีหลายบทเรียน (1:N) เรียงด้วย `order_index` โดย**ภายในคอร์สเดียวกันห้ามมี `order_index` ซ้ำกัน**

- **รีวิว (Review)** — ผู้เรียนรีวิวคอร์สได้ โดย**ผู้เรียน 1 คนรีวิวคอร์สเดิมได้ไม่เกิน 1 ครั้ง** ข้อความรีวิว (`comment`) จะไม่ระบุก็ได้

- **ตั้งค่าคอร์ส (CourseSetting)** — แต่ละคอร์สมีชุดตั้งค่าได้ไม่เกิน 1 ชุด (**1:1**) ข้อความประกาศ (`announcement`) จะไม่ระบุก็ได้

---

## ภารกิจ (Task)

Go structs ในโฟลเดอร์ `internal/models/` ยัง map ความสัมพันธ์เชิงสัมพันธ์ไม่ถูกต้องตาม Requirements ด้านบน
**ไม่มีคอมเมนต์ชี้จุดบั๊กให้** — วิเคราะห์เอง แก้เอง จนกว่าจะรัน Test ผ่านครบทุกข้อ

แก้เฉพาะไฟล์ใน `internal/models/` เท่านั้น (ห้ามแก้โฟลเดอร์ `tests/`)

### รันเทสต์ทุกข้อ
```bash
go test -v ./tests/... -count=1
```

### รันทีละข้อ
```bash
go test -v ./tests/ -run TestCourseTagManyToMany -count=1
```

### รันเซิร์ฟเวอร์ (หลังแก้ครบ)
```bash
docker compose up -d
go run ./cmd/server
```

---

## หัวข้อ Test (8 ไฟล์)

| ไฟล์ | ตรวจเรื่อง |
|---|---|
| `01_user_subtype_test.go` | Single Table Inheritance + subtype fields ของ users |
| `02_category_selfref_test.go` | ความสัมพันธ์อ้างอิงตัวเองของ Category (parent) |
| `03_course_fk_test.go` | Owning side ของ Course (FK ไป category และ instructor) |
| `04_course_tag_many2many_test.go` | many2many ระหว่าง Course กับ Tag ผ่าน join table `course_tags` |
| `05_enrollment_composite_pk_test.go` | Composite PK ของ Enrollment |
| `06_lesson_unique_index_test.go` | Composite UNIQUE index (course_id, order_index) |
| `07_review_test.go` | Composite UNIQUE (course_id, learner_id) + nullable comment |
| `08_course_setting_test.go` | 1:1 (unique course_id) + nullable announcement |

---

## โครงสร้างโปรเจกต์

- `internal/models/` — Go structs กำหนด Relational Schema **(ไฟล์ที่ต้องแก้)**
- `internal/{dto,repositories,services,handlers,routes}/` — application layers
- `internal/config/` — DB connection + `AutoMigrate`
- `tests/` — In-memory schema validation (ไม่ต้องต่อ DB)

> `go build ./...` ในสถานะเริ่มต้นอาจ error จนกว่าจะแก้ schema ให้สอดคล้องกับ layer อื่น —
> ใช้ `go test ./tests/...` เป็นเกณฑ์ตรวจงานหลัก
