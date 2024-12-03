




	// var son1, son2 int
	// var operator string

	// fmt.Println("Введите первое число: ")
	// fmt.Scanln(&son1)

	// fmt.Println("Введите второе число: ")
	// fmt.Scanln(&son2)

	// fmt.Println("Введите оператор: ")
	// fmt.Scanln(&operator)

	// if operator == "+" {
	// 	fmt.Println("Результат:", son1 + son2)
	// } else if operator == "-" {
	// 	fmt.Println("Результат:", son1 - son2)
	// } else if operator == "*" {
	// 	fmt.Println("Результат:", son1 * son2)
	// } else if operator == "/" {
	// 	if son2 != 0 {
	// 		fmt.Println("Результат:", son1 / son2)
	// 	} else {
	// 		fmt.Println("Ошибка: деление на ноль")
	// 	}
	// } else {
	// 	fmt.Println("Неправильный оператор")
	// }


	// var fasl string
	// fmt.Println("Ведите время года")
	// fmt.Scanln(&fasl)

	// switch fasl{
	//     case "весна": 
	//    		fmt.Println("Надеваем куртку и что тёплое")

	// 	case "лето": 
	// 		fmt.Println("Надеваем шорты и купаться")

	// 	case "осень": 
	// 		fmt.Println("Начинается листопад время фоткаться")

	// 	case "зима": 
	// 		fmt.Println("время бросать снежки")
	// 	default:
	// 		fmt.Println("неизвестное время года")
	// }




    // a, b := 0, 1

    // fmt.Println("Числа Фибоначчи до 100:")
    // for a <= 100 {
    //     fmt.Println(a)
    //     a, b = b, a+b
    // }



	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("fizzbuzz")
	// 	} else if i%3 == 0 {
	// 		fmt.Println("fizz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }







//     for {
//         math() 
        
//         fmt.Print("Введите команду: ")
//         var input string
//         fmt.Scanln(&input) 
        
//         if input == "exit" {
//             fmt.Println("Завершение программы...")
//             break
//         }
//     }
// }

// func math() {
//     var son1, son2 int
//     var operator string

//     fmt.Println("Введите первое число: ")
//     fmt.Scanln(&son1)

//     fmt.Println("Введите второе число: ")
//     fmt.Scanln(&son2)

//     fmt.Println("Введите оператор (+, -, *, /): ")
//     fmt.Scanln(&operator)

//     switch operator {
//     case "+":
//         fmt.Println("Результат:", son1 + son2)
//     case "-":
//         fmt.Println("Результат:", son1 - son2)
//     case "/":
//         if son2 == 0 {
//             fmt.Println("Ошибка: деление на ноль.")
//         } else {
//             fmt.Println("Результат:", son1 / son2)
//         }
//     case "*":
//         fmt.Println("Результат:", son1 * son2)
//     default:
//         fmt.Println("Нет такого оператора, пожалуйста, повторите.")
//     }


//  var i int = 0
// 	for {
// 		i++
// 		switch i {
// 			case 1:
// 				fmt.Println("yanvar")
// 			case 2:
// 				fmt.Println("fevral")
// 			case 3:
// 				fmt.Println("mart")
// 			case 4:
// 				fmt.Println("aprel")
// 			case 5:
// 				fmt.Println("may")
// 			case 6:
// 				fmt.Println("iyun")
// 			case 7:
// 				fmt.Println("iyul")
// 			case 8:
// 				fmt.Println("avgust")
// 			case 9:
// 				fmt.Println("sentabr")
// 			case 10:
// 				fmt.Println("oktabr")
// 			case 11:
// 				fmt.Println("noyabr")
// 			case 12:
// 				fmt.Println("dekabr")
// 		}

// 		if i >= 12 {
// 			fmt.Println("Цикл завершен")
//             break
// 		}

// }



// var n int
//     fmt.Print("Введите число: ")
//     fmt.Scanln(&n)

// 	chot := 0
// 	nechot := 0


// 	for i := 1; i <= n; i++ {
// 		switch {
// 		case i%2 == 0:
// 			chot = chot + i
// 		case i%2 != 0:
// 			nechot = nechot + i
// 		}
// 	}


// 	fmt.Println("Сумма всех чётных чисел:", chot)
// 	fmt.Println("Сумма всех нечётных чисел:", nechot)





// var a [6]string

// a = [6]string{"Muhamad Rasul, Abubakr", "bilolidin;", "sumaya","odil bobo"}
// fmt.Println(a)
// fmt.Println(a[2])

// var s []string = a[1:]
// fmt.Println(s)










// a :=[6]int {6, 7, 9, 12, 47, 32}
// var sum int;
// for _, v := range a {
// 	fmt.Println(v, "index")
// 	sum = sum+v
// }
// fmt.Println(sum)



// slice := a[0:6]
// var max int
// for _, v := range slice {
// if v >= max {
// 	max = v
// }
// }
// fmt.Println(max, "Максимально большое число")


// var numbersNot []int

//   var n int
//   fmt.Println("Nechta son kiritmoqchisiz? ")
//   fmt.Scanln(&n)

//   for i := 0; i <= n; i++ {
//     var number int
//     fmt.Println("Sonni kiriting : ", i+1)
//     fmt.Scanln(&number)
//     numbersNot = append(numbersNot, number)
//   }

//   oddCount := 0
//   evenCount := 0

//   for _, number := range numbersNot {
//     if number%2 == 0 {
//       evenCount++
//     } else {
//       oddCount++
//     }
//   }

//   fmt.Println("Juft sonlar soni:", evenCount)
//   fmt.Println("Toq sonlar soni:", oddCount)

// var input string

// fmt.Scanln(&input)

// m:= make(map[string]int)

// for _, v := range input {
// 	m[string(v)]++
// }

// fmt.Println(m)















// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var students []map[string]interface{}

// 	for {
// 		var name string
// 		fmt.Println("Введите имя студента (или 'exit' для выхода):")
// 		fmt.Scanln(&name)
// 		if name == "exit" {
// 			break
// 		}

// 		var id int
// 		fmt.Println("Введите ID студента:")
// 		fmt.Scanln(&id)

// 		var class int
// 		fmt.Println("Введите класс студента:")
// 		fmt.Scanln(&class)

// 		subjects := make(map[string]int)
// 		for {
// 			var subject string
// 			fmt.Println("Введите предмет и оценку (например, 'Math 5') или 'done' для завершения:")
// 			fmt.Scanln(&subject)
// 			if subject == "done" {
// 				break
// 			}

// 			var grade int
// 			fmt.Scanln(&grade)
// 			subjects[subject] = grade
// 		}

// 		student := map[string]interface{}{
// 			"name":     name,
// 			"id":       id,
// 			"class":    class,
// 			"subjects": subjects,
// 		}

// 		students = append(students, student)
// 	}

// 	fmt.Println("Список студентов:")
// 	for _, student := range students {
// 		fmt.Println("Имя:", student["name"])
// 		fmt.Println("ID:", student["id"])
// 		fmt.Println("Класс:", student["class"])
// 		fmt.Println("Оценки:")
// 		for subject, grade := range student["subjects"].(map[string]int) {
// 			fmt.Println("  ", subject, ":", grade)
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println("Средняя оценка по каждому студенту:")
// 	for _, student := range students {
// 		total := 0
// 		count := 0
// 		for _, grade := range student["subjects"].(map[string]int) {
// 			total += grade
// 			count++
// 		}
// 		average := float64(total) / float64(count)
// 		fmt.Println("Имя:", student["name"], "Средняя оценка:", average)
// 	}
// }









	// m := make(map[string]int)
	// s, i := "asdf", 1
	// m[s] = i

	// v := m[s]
	// package main

	// import (
	// 	"fmt"
	// )
	
	// type Student struct {
	// 	Name  string
	// 	Id    string
	// 	Class string
	// 	Grade map[string]int
	// }
	
	// func main() {
	// 	students := make([]Student, 0, 10)
	
	// 	for {
	// 		var komand int
	// 		fmt.Println("Введите команду (1 - добавить студента, 2 - показать всех студентов, 3 - найти студента по ID, 0 - выход):")
	// 		fmt.Scanln(&komand)
	
	// 		switch komand {
	// 		case 1:
	// 			var student Student
	// 			student.Grade = make(map[string]int)
	
	// 			fmt.Println("Введите имя ученика:")
	// 			fmt.Scanln(&student.Name)
	
	// 			fmt.Println("Введите ID студента:")
	// 			fmt.Scanln(&student.Id)
	
	// 			fmt.Println("Введите класс студента:")
	// 			fmt.Scanln(&student.Class)
	
	// 			for {
	// 				var subject string
	// 				var grade int
	// 				fmt.Println("Введите предмет и оценку или 'done' для завершения:")
	// 				fmt.Scanln(&subject)
	// 				if subject == "done" {
	// 					break
	// 				}
	// 				fmt.Scanln(&grade)
	// 				student.Grade[subject] = grade
	// 			}
	
	// 			students = append(students, student)
	// 			fmt.Println("Ученик добавлен.")
	
	// 		case 2:
	// 			for _, student := range students {
	// 				fmt.Println("Имя:", student.Name, "ID:", student.Id, "Класс:", student.Class)
	// 				fmt.Println("Оценки:", student.Grade)
	// 			}
	
	// 		case 3:
	// 			var searchId string
	// 			fmt.Println("Введите ID студента для поиска:")
	// 			fmt.Scanln(&searchId)
	
	// 			found := false
	// 			for _, student := range students {
	// 				if student.Id == searchId {
	// 					fmt.Println("Имя:", student.Name, "ID:", student.Id, "Класс:", student.Class)
	// 					fmt.Println("Оценки:", student.Grade)
	
	// 					total := 0
	// 					count := 0
	// 					for _, grade := range student.Grade {
	// 						total += grade
	// 						count++
	// 					}
	
	// 					if count > 0 {
	// 						average := float64(total) / float64(count)
	// 						fmt.Printf("Средняя оценка: %.2f\n", average)
	// 					} else {
	// 						fmt.Println("Нет оценок для расчета средней.")
	// 					}
	// 					found = true
	// 					break
	// 				}
	// 			}
	
	// 			if !found {
	// 				fmt.Println("Студент с таким ID не найден.")
	// 			}
	
	// 		case 0:
	// 			fmt.Println("Выход из программы.")
	// 			return
	
	// 		default:
	// 			fmt.Println("Неизвестная команда")
	// 		}
	// 	}
	// }
	




	// var son int
	// a := [7]int{7, 14, 32, 8, 9, 6, 3}
	// fmt.Println(a)
	// for _, v := range a {
		// 	son = son + v
		// }
		// fmt.Println(son)
		
		// knife := a[0:5]
		// summa := 0
		// for _, v := range knife {
			// 		if summa <= v {
				// 			summa =  v
				// 		}
				// }
				// fmt.Println(summa, "Самое большое число")
				
				// BoshArray := []int {42, 12, 9, 6}
				// var el int
				// fmt.Scanln(&el)
				// BoshArray = append(BoshArray, el)
				// fmt.Println(BoshArray)
				
				// var toq int
				// var juft int
				// for _, v := range BoshArray {
					// 	if v%2 == 0 {
						// 		juft = juft +1
						// 	} else {
							// 		toq = toq +1 
							// 	}
							// }
							// fmt.Println("Toq sonlar", toq, "juft sonlar", juft)
							
							
							
	// package main
	// import "fmt"
	// 	func main() {
	// 	var input string
	// 	fmt.Scanln(&input)

	// 	 letterCont := make(map[rune]int)
	// 	for _, v := range input {
	// 		letterCont[v]++

	// 	}

	// 	for k, v := range letterCont{
	// 		fmt.Println("Буквы:", string(k), "повтор:", v)
	// 	}


		
								
	// }


// 	package main

// import (
// 	"fmt"
// )

// // Функция для управления банковским счетом
// func BankAccountManager(initialBalance float64) (
// 	func(amount float64) string,   // Функция пополнения
// 	func(amount float64) string,   // Функция снятия
// 	func() string,                 // Функция отображения баланса
// 	func() string,                 // Функция обнуления
// ) {
// 	balance := initialBalance // Локальная переменная баланса

// 	// Анонимная функция для пополнения счета
// 	deposit := func(amount float64) string {
// 		balance += amount
// 		return fmt.Sprintf("Счет пополнен на %.2f. Текущий баланс: %.2f.", amount, balance)
// 	}

// 	// Анонимная функция для снятия денег
// 	withdraw := func(amount float64) string {
// 		if amount > balance {
// 			return fmt.Sprintf("Недостаточно средств. Текущий баланс: %.2f.", balance)
// 		}
// 		balance -= amount
// 		return fmt.Sprintf("Со счета снято %.2f. Остаток: %.2f.", amount, balance)
// 	}

// 	// Анонимная функция для отображения баланса
// 	showBalance := func() string {
// 		return fmt.Sprintf("Текущий баланс: %.2f.", balance)
// 	}

// 	// Анонимная функция для обнуления счета с использованием defer
// 	resetBalance := func() string {
// 		defer fmt.Println("Счет был обнулен.")
// 		balance = 0
// 		return "Баланс успешно обнулен."
// 	}

// 	// Возвращаем все функции
// 	return deposit, withdraw, showBalance, resetBalance
// }

// func main() {
// 	// Инициализация управления счетом
// 	deposit, withdraw, showBalance, resetBalance := BankAccountManager(100.0)

// 	// Вызов различных операций
// 	fmt.Println(deposit(50.0))        // Пополнение
// 	fmt.Println(showBalance())        // Просмотр баланса
// 	fmt.Println(withdraw(30.0))       // Снятие денег
// 	fmt.Println(showBalance())        // Просмотр баланса
// 	fmt.Println(resetBalance())       // Обнуление счета
// 	fmt.Println(showBalance())        // Просмотр баланса после обнуления
// }

		


// Код для управления библиотекой
// package main

// import (
// 	"fmt"
// )


// type Book struct {
// 	Title           string
// 	Author          string
// 	ISBN            string
// 	AvailableCopies int
// }


// type Library struct {
// 	collection map[string]*Book 
// }


// type LibraryOperations interface {
// 	AddBook(book Book)               
// 	CheckOutBook(isbn string) error  
// 	ReturnBook(isbn string) error    
// 	ListAvailableBooks()             
// }


// type BookError struct {
// 	Message string
// }

// func (e *BookError) Error() string {
// 	return e.Message
// }


// func (lib *Library) AddBook(book Book) {
// 	if _, exists := lib.collection[book.ISBN]; exists {
// 		fmt.Printf("Книга с ISBN %s уже существует.\n", book.ISBN)
// 	} else {
// 		lib.collection[book.ISBN] = &book
// 		fmt.Printf("Книга \"%s\" добавлена в библиотеку.\n", book.Title)
// 	}
// }


// func (lib *Library) CheckOutBook(isbn string) error {
// 	book, exists := lib.collection[isbn]
// 	if !exists {
// 		return &BookError{Message: "Книга не найдена."}
// 	}
// 	if book.AvailableCopies <= 0 {
// 		return &BookError{Message: "Нет доступных экземпляров."}
// 	}
// 	book.AvailableCopies--
// 	fmt.Printf("Книга \"%s\" успешно взята.\n", book.Title)
// 	return nil
// }


// func (lib *Library) ReturnBook(isbn string) error {
// 	book, exists := lib.collection[isbn]
// 	if !exists {
// 		return &BookError{Message: "Книга не найдена."}
// 	}
// 	book.AvailableCopies++
// 	fmt.Printf("Книга \"%s\" успешно возвращена.\n", book.Title)
// 	return nil
// }


// func (lib *Library) ListAvailableBooks() {
// 	fmt.Println("Доступные книги:")
// 	for _, book := range lib.collection {
// 		fmt.Printf("Название: %s, Автор: %s, ISBN: %s, Доступных экземпляров: %d\n",
// 			book.Title, book.Author, book.ISBN, book.AvailableCopies)
// 	}
// }


// func main() {
// 	library := Library{collection: make(map[string]*Book)}

	
// 	for {
// 		var choice int
// 		fmt.Println("\nСистема управления библиотекой")
// 		fmt.Println("1. Добавить книгу")
// 		fmt.Println("2. Взять книгу")
// 		fmt.Println("3. Вернуть книгу")
// 		fmt.Println("4. Просмотреть доступные книги")
// 		fmt.Println("5. Выход")
// 		fmt.Print("Выберите опцию: ")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
			
// 			var title, author, isbn string
// 			var copies int
// 			fmt.Print("Введите название книги: ")
// 			fmt.Scanln(&title)
// 			fmt.Print("Введите автора книги: ")
// 			fmt.Scanln(&author)
// 			fmt.Print("Введите ISBN книги: ")
// 			fmt.Scanln(&isbn)
// 			fmt.Print("Введите количество доступных экземпляров: ")
// 			fmt.Scanln(&copies)

// 			library.AddBook(Book{Title: title, Author: author, ISBN: isbn, AvailableCopies: copies})

// 		case 2:
			
// 			var isbn string
// 			fmt.Print("Введите ISBN книги для взятия: ")
// 			fmt.Scanln(&isbn)

// 			if err := library.CheckOutBook(isbn); err != nil {
// 				fmt.Println("Ошибка:", err)
// 			}

// 		case 3:
			
// 			var isbn string
// 			fmt.Print("Введите ISBN книги для возврата: ")
// 			fmt.Scanln(&isbn)

// 			if err := library.ReturnBook(isbn); err != nil {
// 				fmt.Println("Ошибка:", err)
// 			}

// 		case 4:
		
// 			library.ListAvailableBooks()

// 		case 5:
			
// 			fmt.Println("Выход из программы.")
// 			return

// 		default:
// 			fmt.Println("Неверная опция, попробуйте снова.")
// 		}
// 	}
// }



// package main 
// import (
// 	"fmt"
// )


// type Car struct {
// 	Brand string
// 	Id int
// 	Rented bool 
// }

// type Bike struct {
// 	Brand string
// 	Id int
// 	Rented bool
// }

// type Skate struct {
// 	Brand string
// 	Id int
// 	Rented bool
// }

// type  Vehicle interface {
// 	Rent()error
// 	Return()error
// 	Details()error
// }


// func identifyType (a []Vehicle, t int) []Vehicle {
// 	if t == 1 {
		 
// 	}
// }


// func (c *Car) Rent() {
	
// }


// func main() {
// 	var aray  []Vehicle 
// 	var input int
// 	for {
// 	fmt.Println("Выберите какое транспортное средство хотите добавить")
// 	fmt.Println("1)Car, 2)Bike, 3)Skate, 0) Exit")
// 	fmt.Scanln(&input)

// 		if input == 0 {
// 			fmt.Println("Выход из программы.")
// 			break
// 		}

// 	switch input {
// 		case 1:
// 			var brand string
// 			var id int
// 			fmt.Println("Ведите бренд машины")
// 			fmt.Scanln(&brand)
// 			fmt.Println("Ведите id машины")
// 			fmt.Scanln(&id)
// 			c := Car{
// 				Brand: brand,
// 				Id: id
// 			}
// 			ar = append(ar, c)
// 		case 2:
// 			var brand string
// 			var id int
// 			fmt.Println("Ведите бренд мотоцикла")
// 			fmt.Scanln(&brand)
// 			fmt.Println("Ведите id мотоцикла")
// 			fmt.Scanln(&id)
// 			b := Car{
// 				Brand: brand,
// 				Id: id
// 			}
// 			ar = append(ar, b)
	
// 		case 3:
// 			var brand string
// 			var id int
// 			fmt.Println("Ведите бренд самоката")
// 			fmt.Scanln(&brand)
// 			fmt.Println("Ведите id самоката")
// 			fmt.Scanln(&id)
// 			d := Car{
// 				Brand: brand,
// 				Id: id
// 			}
// 			ar = append(ar, d)
// 		default:
// 			fmt.Scanln("Неверный ввод, попробуйте снова")
// 		}
	
// 	}

// 	var click int
// 	fmt.Println("Какое действие вы хотите выбрать: 1)Арендавать 2)Возрашать тачки 3)Показать арендованые тачки")
// 	fmt.Scanln(&click)
// 	switch click {
// 		case 1:
// 			var check int
// 			fmt.Println("Какую транспортную средстку вы хотите арендовать 1)Car 2) Bike 3)Scate") 
// 			switch check {
// 				case 1:
					
					
// 			}

// 	}

// }


package main

import (
	"fmt"
)

type Car struct {
	Brand  string
	Id     int
	Rented bool
}

type Bike struct {
	Brand  string
	Id     int
	Rented bool
}

type Skate struct {
	Brand  string
	Id     int
	Rented bool
}

type Vehicle interface {
	Rent() error
	Return() error
	Details() error
}

func (c *Car) Rent() error {
	if c.Rented {
		return fmt.Errorf("машина с ID %d уже арендована", c.Id)
	}
	c.Rented = true
	return nil
}

func (c *Car) Return() error {
	if !c.Rented {
		return fmt.Errorf("машина с ID %d не была арендована", c.Id)
	}
	c.Rented = false
	return nil
}

func (c *Car) Details() error {
	status := "доступна"
	if c.Rented {
		status = "арендована"
	}
	fmt.Printf("Car Brand: %s, ID: %d, Статус: %s\n", c.Brand, c.Id, status)
	return nil
}

func (b *Bike) Rent() error {
	if b.Rented {
		return fmt.Errorf("мотоцикл с ID %d уже арендован", b.Id)
	}
	b.Rented = true
	return nil
}

func (b *Bike) Return() error {
	if !b.Rented {
		return fmt.Errorf("мотоцикл с ID %d не был арендован", b.Id)
	}
	b.Rented = false
	return nil
}

func (b *Bike) Details() error {
	status := "доступен"
	if b.Rented {
		status = "арендован"
	}
	fmt.Printf("Bike Brand: %s, ID: %d, Статус: %s\n", b.Brand, b.Id, status)
	return nil
}

func (s *Skate) Rent() error {
	if s.Rented {
		return fmt.Errorf("самокат с ID %d уже арендован", s.Id)
	}
	s.Rented = true
	return nil
}

func (s *Skate) Return() error {
	if !s.Rented {
		return fmt.Errorf("самокат с ID %d не был арендован", s.Id)
	}
	s.Rented = false
	return nil
}

func (s *Skate) Details() error {
	status := "доступен"
	if s.Rented {
		status = "арендован"
	}
	fmt.Printf("Skate Brand: %s, ID: %d, Статус: %s\n", s.Brand, s.Id, status)
	return nil
}

func showAvailableVehicles(ar []Vehicle) {
	fmt.Println("Доступные транспортные средства:")
	for _, v := range ar {
		if v != nil {
			switch v := v.(type) {
			case *Car:
				if !v.Rented {
					v.Details()
				}
			case *Bike:
				if !v.Rented {
					v.Details()
				}
			case *Skate:
				if !v.Rented {
					v.Details()
				}
			}
		}
	}
}

func main() {
	var aray []Vehicle
	var input int

	for {
		fmt.Println("Выберите какое транспортное средство хотите добавить (или 0 для выхода):")
		fmt.Println("1) Car, 2) Bike, 3) Skate, 0) Exit")
		fmt.Scanln(&input)

		if input == 0 {
			fmt.Println("Выход из программы.")
			break
		}

		switch input {
		case 1:
			var brand string
			var id int
			fmt.Println("Введите бренд машины:")
			fmt.Scanln(&brand)
			fmt.Println("Введите ID машины:")
			fmt.Scanln(&id)
			c := &Car{Brand: brand, Id: id}
			aray = append(aray, c)
		case 2:
			var brand string
			var id int
			fmt.Println("Введите бренд мотоцикла:")
			fmt.Scanln(&brand)
			fmt.Println("Введите ID мотоцикла:")
			fmt.Scanln(&id)
			b := &Bike{Brand: brand, Id: id}
			aray = append(aray, b)
		case 3:
			var brand string
			var id int
			fmt.Println("Введите бренд самоката:")
			fmt.Scanln(&brand)
			fmt.Println("Введите ID самоката:")
			fmt.Scanln(&id)
			s := &Skate{Brand: brand, Id: id}
			aray = append(aray, s)
		default:
			fmt.Println("Неверный ввод, попробуйте снова.")
		}
	}

	
	var click int
	for {
		fmt.Println("\nКакое действие вы хотите выбрать:")
		fmt.Println("1) Арендовать, 2) Возвратить, 3) Показать арендованные транспортные средства, 0) Выход")
		fmt.Scanln(&click)

		if click == 0 {
			fmt.Println("Выход из программы.")
			break
		}

		switch click {
		case 1:
		
			showAvailableVehicles(aray)

			var check int
			fmt.Println("Какое транспортное средство вы хотите арендовать?")
			fmt.Println("1) Car, 2) Bike, 3) Skate")
			fmt.Scanln(&check)

			var vehicle Vehicle
			var id int
			switch check {
			case 1:
				fmt.Println("Введите ID машины для аренды:")
				fmt.Scanln(&id)
				for _, v := range aray {
					if car, ok := v.(*Car); ok && car.Id == id && !car.Rented {
						vehicle = car
						break
					}
				}
			case 2:
				fmt.Println("Введите ID мотоцикла для аренды:")
				fmt.Scanln(&id)
				for _, v := range aray {
					if bike, ok := v.(*Bike); ok && bike.Id == id && !bike.Rented {
						vehicle = bike
						break
					}
				}
			case 3:
				fmt.Println("Введите ID самоката для аренды:")
				fmt.Scanln(&id)
				for _, v := range aray {
					if skate, ok := v.(*Skate); ok && skate.Id == id && !skate.Rented {
						vehicle = skate
						break
					}
				}
			}

			if vehicle != nil {
				if err := vehicle.Rent(); err != nil {
					fmt.Println("Ошибка аренды:", err)
				} else {
					fmt.Println("Транспорт арендован успешно!")
				}
			} else {
				fmt.Println("Транспорт не найден или уже арендован.")
			}

		case 2:
			var check int
			fmt.Println("Какое транспортное средство вы хотите вернуть?")
			fmt.Println("1) Car, 2) Bike, 3) Skate")
			fmt.Scanln(&check)

			var vehicle Vehicle
			switch check {
			case 1:
				var id int
				fmt.Println("Введите ID машины для возврата:")
				fmt.Scanln(&id)
				for _, v := range aray {
					if car, ok := v.(*Car); ok && car.Id == id {
						vehicle = car
						break
					}
				}
			case 2:
				var id int
				fmt.Println("Введите ID мотоцикла для возврата:")
				fmt.Scanln(&id)
				for _, v := range aray {
					if bike, ok := v.(*Bike); ok && bike.Id == id {
						vehicle = bike
						break
					}
				}
			case 3:
				var id int
				fmt.Println("Введите ID самоката для возврата:")
				fmt.Scanln(&id)
				for _, v := range aray {
					if skate, ok := v.(*Skate); ok && skate.Id == id {
						vehicle = skate
						break
					}
				}
			}

			if vehicle != nil {
				if err := vehicle.Return(); err != nil {
					fmt.Println("Ошибка возврата:", err)
				} else {
					fmt.Println("Транспорт возвращен успешно!")
				}
			} else {
				fmt.Println("Транспорт не найден.")
			}

		case 3:
			fmt.Println("Арендованные транспортные средства:")
			for _, v := range aray {
				if v != nil {
					v.Details()
				}
			}

		default:
			fmt.Println("Неверный ввод.")
		}
	}
}


