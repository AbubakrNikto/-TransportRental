
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


