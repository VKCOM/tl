# Сериализация TL-типов в JSON и из JSON

При сохранении по возможности делается оптимизация, когда пустые значения
(нулевые числа, пустые строки, false, пустые массивы) не записываются.
 
При чтении отсутствующего поля оно считается пустым (0 для численных типов,
пустая строка для строк, пустой массив для массивов, объект без полей).

# Числа, строки, Bool

Превращаются в соответствующий тип JSON.  

Строки, которые содержат в символы, не поддерживаемые в кодировке UTF-8, или сырые байты, будут закодированы **base64** и сохранены в `json:object`, с одним полем с именем `base64` и закодированным телом. `{"base64": "<base64-encoded bytes>"}`

Например `"\xC5"` представляет символ Å в кодировке ISO-8859-1. Этот символ не представлен в UTF-8 и будет записан как

```json
{
    "base64":"xQ=="
}
```

или последовательность байт `0xf0 0xf1 0xf2 0xf3` будет записана как
```json
{
    "base64":"8PHy8w=="
}
```

обычный текст при этом остается обычной строкой

```
TL:
foo str:string bin:string = Foo;

JSON:
{
    "str": "string as always",
    "bin": { "base64" : "8PHy8w==" }
}
```

# Структуры

Превращаются в `json:object`, где имена полей соответствуют именам в TL, например

```
TL:
memcache.strvalue value:string flags:int = memcache.Value;

JSON:
{
    "value":"Hello",
    "flags":1
}
```

Поле будет сохранено только если оно имеет непустое значение.

При чтении если поле не зависит от бита `field_mask`, и не указано, оно будет установлено в пустое значение.

А если зависит, то все комбинации интерпретируются так

|  Бит `field_mask`  | Указано ли поле явно | Результат
|-----|----|----|
|  1  |  да  | поле прочитано |
|  1  |  нет | поле установлено в пустое значение |
|  0  |  да  | если `field_mask` явная, поле прочитано, бит `field_mask` установлен в 1, иначе ошибка |
|  0  |  нет | поле установлено в пустое значение, бит `field_mask` остался 0 |

Лишние поля при чтении считаются опечатками и не допускаются.


# Объединения

Превращаются в `json:object` с полями `type`, `value`, где тип содержит полный TL-тип с тегом через #
(независимо от того, указана ли # в описании TL) 

```
TL:
not_found#08309efe              = Value;
longvalue#40b8737a value:long   = Value;
strvalue#c265bec1  value:string = Value;

JSON, одно из:
{
    "type":"not_found#08309efe",
    "value":{}
}
{
    "type":"longvalue#40b8737a",
    "value":{
        "value":5
    }
}
{
    "type":"strvalue#c265bec1",
    "value":{
        "value":"Hello",
    }
}
```
 
При чтении можно указывать как полный TL-тип (`"strvalue#c265bec1"`), так и
только имя (`"strvalue"`) или только тэг с решёткой (`"#c265bec1"`). 

# Перечисления

Фактически являются разновидностью объединения.

Превращаются в строку, содержащую полный TL-тип константы с тегом через #
(независимо от того, указана ли # в описании TL)

```
TL:
red = Color; 
blue = Color;

JSON:
"red#ad537640"
"blue#b53d8932"
```

При чтении можно указывать как полный TL-тип (`"red#ad537640"`),
так и только имя (`"red"`) или только тэг с решёткой (`"#ad537640"`).

# Vector, Tuple

Превращаются в `json:array`

```
TL:
series numbers:%(Vector int) start_indx:int = Ыeries;

JSON:
{
    "numbers":[1,5,20],
    "start_indx":1
}

TL:
point X:int Y:int = Point;
polygon points:%(Vector %Point) = Polygon;

JSON:
{
    "points":[
        {"X":1,"Y":1},
        {"X":3,"Y":1},
        {"X":2,"Y":2}
    ]
}

TL:
withInnerArray counters: %(Vector %(Tuple int 8)) = WithInnerArray;

JSON:
{
    "counters":[
        [0,1,2,3,4,5,6,7],
        [10,11,12,13,14,15,16,17]
    ],
}
```

При чтении, если размер массива зависит от параметра, размер в json должен точно совпадать с параметром.

```
{
    "n":5,
    "data":[0,1,2,3,4]
}
```

Если размер `data` не будет равен 5, то произойдёт ошибка чтения. 

# Dictionary

```
TL:
dictExample type_name:string description: (%Dictionary string) = DictExample;

JSON:
{
    "type_name":"name",
    "description":{
        "a":"alpha",
        "b":"beta"
    }
}

TL:
intKeyDictExamplw
     longs:%(IntKeyDictionary %(IntKeyDictionary long))
     = IntKeyDictExample;

JSON:
{
    "longs":{
        "1":{"10":100,"11":101},
        "2":{"20":200,"21":201}
    }
}
```

# Maybe

Если значение установлено, пишется `json:object` с полями `value` и `ok:true`.
Если значение не установлено, пишется пустой `json:object` (подразумевается `ok:false`)

```
TL:
maybeExample
    isYes:(Maybe string)
    isNo:(Maybe int)
    = MaybeExample;

JSON:
{
    "isYes":{
        "value":"yes",
        "ok":true
    },
    "isNo":{}
}
```

При чтении все комбинации

|    `ok`    | Указано ли value явно |                    Результат      ,            |
|------------|-----------------------|------------------------------------------------|
|  true      |          да           | MaybeTrue, value прочитано                     |
|  true      |          нет          | MaybeTrue, value установлено в пустое значение |
|  false     |          да           | ошибка                                         |
|  false     |          нет          | MaybeFalse                                     |
|  не указан |          да           | MaybeTrue, value прочитано                     |
|  не указан |          нет          | MaybeFalse                                     |

# True

Сохраняется в JSON только если используется, как `Boxed type`, либо когда
зависит от `fields_mask`, находящейся в самом объекте. Тогда поле дублирует бит маски при сохранении,
и позволяет указывать поле вместо бита маски при чтении (взаимодействие с битами маски аналогично полям остальных типов).   

```
TL:
exampleTrueType
     fields_mask  : #
     sorted       : fields_mask.1? %True
     reversed     : fields_mask.2? %True
     = ExampleTrueType;

JSON:
{
    "fields_mask":4,
    "and_mask":48,
    "xor_mask":15,
    "reverse":true
}
```
