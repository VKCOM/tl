### TL Formal Model on Arend

В `src/Model.ard` описаны примитивы для описания типов на языке TL.
* В данной модели нет _Boxed_ на момент написания этого файла (13.11.24)
* Для описания типа создайте объект типа `TLType` (разница между union-type и обычными типа лишь в количестве конструкторов)

Пример для `pair {X: Type} {Y: Type} left:X right:Y = Pair;`:
```haskell
\func TLPair : TLType => \new TLType {
  | typeHeader => tl-new-type 1 (^Type :: (^Type :: nil))
  | constructorsCount => 1
  | constructors => (\new TLConstructor {
    | fields-count => 2
    | fields =>
      (\new TLFieldDefinition {
        | fieldProp => idp
        | fieldMask => nothing
        | fieldType => field-generic (\new TLGenericRef {
          | i => 0
          | isGeneric => idp
        })
      }) eac-::
      (\new TLFieldDefinition {
        | fieldProp => idp
        | fieldMask => nothing
        | fieldType => field-generic (\new TLGenericRef {
          | i => 1
          | isGeneric => idp
        })
      }) eac-::
      eac-nil
  }) eac-:: eac-nil
  | newType => is-new-type
  | nonEmpty => contradiction
}
```

Другие примеры лежат в `src/Types.ard`. Также там есть пример не выводимой в тле конструкции (которая не компилиться в Arend)