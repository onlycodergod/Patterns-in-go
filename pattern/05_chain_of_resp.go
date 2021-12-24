package pattern

import "fmt"

/*
	Реализовать паттерн "цепочка вызовов".
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного паттерна на практике
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern


В качестве примера рассмотрим госпиталь, имеющий помещения "Регистратура", "Доктор", "Комната медикаментов", "Кассир"
Используя пример больницы, пациент сперва попадает в Регистратуру. Затем, зависимо от его состояния,
Регистратура отправляет его к следующему исполнителю в цепи.
*/

// Интерфейс обработчика
// Обработчик определяет общий для всех конкретных обработчиков интерфейс. Обычно достаточно описать единственный метод обработки запросов,
// но иногда здесь может быть объявлен и метод выставления следующего обработчика
type department interface {
	execute(*patient)
	setNext(department)
}

// Конкретные обработчики
// Содержат код обработки запросов. При получении запроса каждый обработчик решает, может ли он обработать запрос, а также стоит ли передать его следующему объекту.
//В большинстве случаев обработчики могут работать сами по себе и быть неизменяемыми, получив все нужные детали через параметры конструктора.

// Конкретный обработчик - Регистратура
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// Конкретный обработчик - Доктор
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// Конкретный обработчик - Комната медикаментов
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// Конкретный обработчик - Кассир
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

// Клиент может либо сформировать цепочку обработчиков единожды, либо перестраивать её динамически, в зависимости от логики программы.
// Клиент может отправлять запросы любому из объектов цепочки, не обязательно первому из них.
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func ChainOfRespPatternRun() {
	patient := &patient{name: "Ivan"}
	reception := &reception{}
	doctor := &doctor{}
	medical := &medical{}
	cashier := &cashier{}

	//Set next for departments
	reception.setNext(doctor)
	doctor.setNext(medical)
	medical.setNext(cashier)

	//Patient visiting
	reception.execute(patient)
}
