# Сделано
* Тупла некорректный код генерит для вложенных туплов, сделать функции записи и чтения
* uint32 только bare
* передача нат-параметров для union
* часть проверок на уникальность конструкторов, имён типов, тэгов, запретил 0 тэг
* поля с одинаковыми Title(name) // hren a:Bool A:int = Hren;
* постарался вынести VK-специфичную часть в отдельный файл
* сделал проверки-эвристики для встроенных типов, була.
* Починить рекурсивные типы - генерировать * поля-указатели
* убраны конфликты глобальных имён
* Решить проблему совпадения имён для Go-типов, вероятно придётся сделать мапку с использованными гошными именами
* Разбиение по файликам
* tlglobal namespace
* Починить nat-param в dictionary сделать как в векторе
* Сгенерировать код проверки fields_mask
* Енумы сделать
* Сделать чтобы Arith.Res был uint32 и не было промежуточных переполнений
* Генерация фабрики factory.go
* Анонимные поля в квадратных скобках, вообще корректность квадратных скобок
* Структы из 1 безымянного поля делать typedef
* Убрать HandlerContext из старого tlgen
* сделать чтобы тэги встроенных типов, була, maybe можно было задавать как угодно
* Сделать тип Tuple обыкновенным, использовать для реализации его и вектора встроенный тип n*[]
* Bytes-версии - `map[*]T -> []struct{K, V}`
* убрал зависимость от package tlrw
* разложил типы по файлам, пока что примерно
* привёл в порядок имена функций записи обёрток, туплов и векторов
* Тип True поддерживать в филд масках особым образом (посмотреть, как в старом было сделано)
* Отдельный код для Maybe (пока некорретный - проверить с аргументом int Int, Vector, в том числе с пробросом парамтеров)
* Код работы с JSON
* задавать желаемый package для кода

# TODO
* запретить совпадения имён между полями и аргументами шаблона
* Оставить в пакете tlrw только независящий от тэгов код
* etc.

```tlschema
```
# Замена *повторяющегося* типа на `Vector` tlgen 1.0

```
hren b:a*[int] a:# = Hren;
```

`Compile TL schema to TLO
nat_term: found type type_type
make: *** [Makefile:16: gen] Error 1`

---
```
hren a:# c:float b:a*[int] = Hren;
```

`replace arrays: replace in constructors: hren: b: expected vector multiplicity to be next to the array`

---
```
hren a:# b:a*[int] c:a*[int] = Hren;
```
`replace arrays: replace in constructors: hren: b: expected a single usage of multiplicity, found 2`

---
```
hren a:float b:a*[int] = Hren;
```
`Compile TL schema to TLO
nat_term: found type type_type`

---
```
hren a:int b:a*[int] = Hren;
```
`Compile TL schema to TLO
nat_term: found type type_type`

---

О нашёл ещё работающий вариант с анонимным полем
```
hren a:# b:[int] = Hren;
```
```go
// Hren описывает следующий комбинатор:
// hren b:(Vector %Int) = Hren
type Hren struct {
        B []int32
}
```
Но анонимные поля отдельная фишка

---
```
hren # = Hren;
```
```go
// hren # = Hren
type Hren uint32
```
---
```
hren int int = Hren;
```
```go
// hren %Int %Int = Hren
type Hren struct {
        Int int32
        Int1 int32
}
```

---
```
hren int # = Hren;
```
bad line: 1 uint32
так почему-то нельзя, лол
Ты пока не умеешь с анонимными

---
```
anon int = Anon;
```
```go
// TL: anon int = Anon;
type Anon struct { 
	int32
}
```

---
```
hren # b:[int] = Hren;
```
```go
// hren b:(Vector %Int) = Hren
type Hren struct {
	B []int32
}
```

---

```
hren # [int] = Hren;
```
```go
// Hren описывает следующий комбинатор:
// hren (Vector %Int) = Hren
type Hren []int32
```
---

ого неожиданно
```
hren a:# [int] = Hren;
```
```golang
// Hren описывает следующий комбинатор:
// hren (Vector %Int) = Hren
type Hren []int32
```
ну понятно, сначала срабатывает свёртка 2х полей в одно, дальше если это поле получилось анонимным то срабатывает механизм генерации алиаса вместо структа, нам второй пока не нужен

---
```
hren a:# b:int c:a*[int] = Hren;
```

`replace arrays: replace in constructors: hren: c: expected vector multiplicity to be next to the array`

---

```
hren # b:a*[int] = Hren;
```

`Compile TL schema to TLO
nat_term: found type type_type`


---

Значит все паттерны замены которые я пока нашёл
```
N:# B:N*[t] -> B: Vector[t]
N:# B:[t] -> B: Vector[t]
# B:[t] -> B: Vector[t]
# [t] -> "пустое имя": Vector[t] 
```

