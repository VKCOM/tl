# Документация пакета RPC

Этот пакет реализует VK RPC для TCP/Unix и UDP транспорта.

Модификации этого пакета требуют глубоких знаний протокола VK RPC и синхронизации с командой движков.

### Cвойства, которые есть сейчас

И которые нельзя терять при любых рефакторах пакета.

1. Строгость - всё, что отклоняется от протокола должна приводить к ошибке.
2. Скорость - обрабатывать 1млн RPS и более на среднем ноутбуке (2024г), и не хотим замедления.
3. Отсутствие аллокаций при длительной работе соединения, аллокации только при установлении.
4. Минимальные внешние и внутренние зависимости - из vkgo нельзя использовать ничего вне папки пакета.
5. Гарантия лимита памяти rpc.Server и backpressure на клиентов - при любых внешних условиях не будет выделено больше памяти и горутин-работников, чем настроено. Из этого свойства часто следуют гарантии лимита памяти и сервисов, построенных на базе rpc.Server. 
5. Тестируемость - максимальное использование библиотек детерминистского тестирования. 

### Свойства, которые хотелось бы реализовать

1. RPC level Congestion Control для rpc.Client - если клиент видит, что запрос вероятно не будет обработан сервером за таймаут, запрос должен генерировать ошибку локально. 
2. Лучшее инструментирование rpc.Server - писать события и различную информацию (например версию протокола) в Statshouse.

# Как использовать rpc.Server

## Инициализация

В принципе, достаточно передать обработчик. Но если вы собираетесь общаться с другими машинами,
то также понадобится криптоключ и доверенные подсети. Без шифрования между подсетями данные передавать заперещено.

```
	server := rpc.NewServer(
		rpc.ServerWithHandler(myHandler),
		rpc.ServerWithTrustedSubnetGroups(subnet.DefaultTrustedSubnetGroups),
		rpc.ServerWithCryptoKeys([]string{string(aesKeyBytes)}),
	)

	go func() {
		if err := server.ListenAndServe("udp", "127.0.0.1:2401"); err != nil {
			log.Fatal(err)
		}
	}()

	if err := server.ListenAndServe("tcp", ":2401"); err != nil {
		log.Fatal(err)
	}
```

## Обработчик и контекст обработчика

Обработчик получает контекст обработчика `hctx`, где есть Request, RequestExtra и другие поля.
Обработчик должен заполнить Response и ResponseExtra и вернуть nil, либо вернуть ошибку.

Обработчик не имеет право запоминать `hctx` и потом обращаться к полям после выхода из функции обработчика,
так как контекст будет немедленно переиспользован сервером.

Перед вызовом у `ctx` будет установлен таймаут, в зависимости от пожеланий клиента и Default таймаута в сервере.

Также `ctx` будет отменён, если соединение, где сделан запос закроется. Так что если обработчик делает какие-то долгие
операции, он должен быть готов прервать их при отмене контекста, как и принято в го.  

```
	myHandler := func(ctx context.Context, hctx *rpc.HandlerContext) (err error) {
		fmt.Printf("function: %d\n", len(hctx.Request))
		hctx.Response = append(hctx.Response, 1, 2, 3)
		fmt.Printf("response: %d\n", len(hctx.Response))
		return nil
	}

	myHandler := func(ctx context.Context, hctx *rpc.HandlerContext) (err error) {
		return &rpc.Error{
			Code:        ErrBrokenLocationCode,
			Description: "Unable to pack/unpack location",
		}
	}
```

В качестве обработчика обычно используется код сгенерированный `tl2gen`.
Здесь за сериализацию и десериализацию отвечает сгенерированный код, так что в `hctx` обычно смотреть не нужно,
но если нужно, он доступен через `context.Context`.

```
	myHandler := tlmemcache.Handler{
		Get: func(ctx context.Context, args tlmemcache.Get) (tlmemcache.Value, error) {
			hctx := rpc.GetHandlerContext(ctx)
			log.Printf("transport %s request len = %d protocol version %d keyID %x", hctx.ProtocolTransport(), len(hctx.Request), hctx.ProtocolVersion(), hctx.KeyID())
			log.Printf("memcache request with key=%s", args.Key)
			strValue := tlmemcache.Strvalue{
				Value: "Hello " + args.Key + " from Go via " + hctx.ProtocolTransport(),
				Flags: 0xAA05,
			}
			return strValue.AsUnion(), nil // or rpc.NewError(-303, "dummy error")
		},
	}
```

## Цепочка обработчиков

Обработчики можно вызывать в цепочке, если первый вернул `rpc.ErrNoHandler`, то вызывается следующий и так далее.

```
    myHandler := rpc.ChainHandler(router1.Handle, router2.Handle)
```

## Синхронный обработчик

Обработчик обычно вызывается из пула горутин-работников, для того, чтобы освободить сетевую горутину для ответов
на пинги, а также чтобы параллельно обрабатывать запросы даже и от одного соединения.

Иногда нужна низкая latency ответа, так что время на перекладывание запроса между горутинами становится заметно.

Тогда можно попросить вызывать обработчик прямо из сетевой горутины, определив `SyncHandler`.

```
	server := rpc.NewServer(
		rpc.ServerWithSyncHandler(mySyncHandler),
		rpc.ServerWithHandler(myHandler),
		rpc.ServerWithTrustedSubnetGroups(subnet.DefaultTrustedSubnetGroups),
		rpc.ServerWithCryptoKeys([]string{string(aesKeyBytes)}),
	)
```

По своей структуре SyncHandler не отличается от обычного, но не должен делать никаких долгих операций, иначе
заблокируется сетевые горутины. Если `SyncHandler` вернул `rpc.ErrNoHandler` (безусловно для функции, либо в зависимости от её аргументов), 
то обычный обработчик вызовется через некоторое время из горутины-работника.

## Поддержка Long poll

Long poll является популярным механизмом доставки событий до клиентов в момент, когда они происходят.

С точки зрения сервера, long poll это запрос, обработка которого откладывается на значительное время,
до наступления некоторого условия в бизнес-логике.

Наивная поддержка бы потребовала блокировки горутины-работника для каждого клиента на значительно время,
если клиентов 30000+, как в Barsic Master или Statshouse, то это бы потребовало 30000 висящих горутин с
контекстом вызова hctx, что было бы дико неэффективно.

Поэтому используется специальное API. Обрабатывать long poll можно только в `SyncHandler`, так как там
нужна детерминистская отмена, а значит отмена должна происходить строго после постановки long poll, но
отмена может придти в момент, когда long poll ждет горутины-обработчика.

Итак, в классическом случае SyncHandler берёт блокировку на структуру данных бизнес-логики и смотрит,
может ли он вернуть ответ сразу, если может, то возвращает ответ. Если же он решает, что этот запрос
будет ждать определенного события, он обязан не отпуская блокировку на структуру данных бизнес логики,
взять дополнительно блокировку на список ожидающих запросов, вызвать `hctx.StartLongPoll()`, если вернулась
ошибка, вернуть её из обработчика, если же запрос успешен, то добавить `rpc.LongpollHandle` в список ожидающих.

Смысл этой операции в том, что тяжеловесный `hctx` (500+ байтов плюс байты запроса и байты Extra) обменивается 
на легковесный `rpc.LongpollHandle` (100 байтов), так что на каждые десять тысяч ожидающих запросов нужно всего
около мегабайт памяти.

```
func (ms *JournalFast) HandleGetMetrics3(args tlstatshouse.GetMetrics3, hctx *rpc.HandlerContext) error {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	var ret tlmetadata.GetJournalResponsenew
	ms.getJournalDiffLocked3(args.From, &ret)
	if len(ret.Events) != 0 {
		hctx.Response, err = args.WriteResultTL1(hctx.Response, ret)
		return err
	}
	ms.clientsMu.Lock()
	defer ms.clientsMu.Unlock()
	lh, err := hctx.StartLongpoll(ms)
	if err != nil {
		return err
	}
	ms.metricsVersionClients3[lh] = args
	return nil
}
```

long poll запрос в любой момент может быть отменен клиентом, либо закрыться соединение.
Поэтому чтоы очистить структуру данных ожидающих запросов, в `StartLongpoll` передается интерфейс отмены,
который обязан быть в простейшем случае реализован таким образом:

```
func (ms *JournalFast) CancelLongpoll(lh rpc.LongpollHandle) {
	ms.clientsMu.Lock()
	defer ms.clientsMu.Unlock()
	delete(ms.metricsVersionClients3, lh)
}

func (ms *JournalFast) WriteEmptyResponse(lh rpc.LongpollHandle, hctx *rpc.HandlerContext) error {
	ms.CancelLongpoll(lh)
	return rpc.ErrLongpollNoEmptyResponse
}
```

При такой реализации если клиент использует таймаут (что обязательно при использовании UDP транспорта
на любом сегменте путешествия запроса от клиента к серверу, по причине того, что в UDP нет детерминисткой
отмены при закрытии соединений), то мы будем видеть фон ошибок `rpc.ErrLongpollNoEmptyResponse`.

Поэтому рекомендуется поддерживать отправку семантически пустого события, при этом так как `WriteEmptyResponse`
заранее, таймаут клиента обычно не успевает сработать, и в метриках мы не видим никаких ошибок.

```
func (ms *JournalFast) WriteEmptyResponse(lh rpc.LongpollHandle, hctx *rpc.HandlerContext) error {
	ms.clientsMu.Lock()
	defer ms.clientsMu.Unlock()
	args, ok := ms.metricsVersionClients3[lh]
	if !ok {
		return nil
	}
	delete(ms.metricsVersionClients3, lh)
	var ret tlmetadata.GetJournalResponsenew
	hctx.Response, err = args.WriteResultTL1(hctx.Response, ret)
	return err
}
```

Фактически здесь при запросе событий журнала с таймаутом 60 секунд, если их нет сразу, то либо ответ вернётся,
когда событие появится, либо через (примерно) 55 секунд вернётся пустой ответ.

Итак, задача кода обработчика когда появляются события, найти все ожидающие запросы, на которые теперь нужно
сгенерировать ответ и скаать серверу отправить его. Для этого берётся блокировка на структуру бизнес логики, затем
блокировка на список ожидающих клиентов, затем для каждого запроса на который будет сгенерирован ответ, 
`rpc.LongpollHandle` обменивается на hctx, ответ записывается и отправляется.

```
func (ms *JournalFast) broadcastJournal() {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	ms.clientsMu.Lock()
	defer ms.clientsMu.Unlock()
	var ret tlmetadata.GetJournalResponsenew
	for lh, args := range ms.metricsVersionClients3 {
		ms.getJournalDiffLocked3(args.From, &ret)
		if len(ret.Events) == 0 {
			continue
		}
		delete(ms.metricsVersionClients3, lh)
		if hctx, _ := lh.FinishLongpoll(); hctx != nil {
			hctx.Response, err = args.WriteResultTL1(hctx.Response, ret)
			hctx.SendLongpollResponse(err)
		}
	}
}
```

Если ответы большие, то поскольку `hctx.SendLongpollResponse` не блокируется на семафоре 
по памяти ответа (используется `sem.ForceAcquire()`), то обьем памяти может значительно увеличиться.

Поэтому рекомендуется блокироваться на семафоре явно, но понимая, что и весь метод `broadcast` заблокируется.

```
...
		if hctx, _ := lh.FinishLongpoll(); hctx != nil {
			var err error
			hctx.Response, err = args.WriteResultTL1(hctx.Response, ret)
			hctx.AccountResponseMem(len(hctx.Response))
			hctx.SendLongpollResponse(err)
		}
...
```

Представим, что `broadcast` вызывается в результате проигрывания бинлога, тогда оно подтормозит ровно на время,
пока не отправятся события клиентам. И всё это будет происходить в пределах лимита памяти, установленного
серверу на ответы, как и должно быть.


# Как использовать rpc.Client

## Инициализация

## Поддержка Long poll

С точки зрения клиента, long poll это обычный вызов. Рекомендуется строго ограничивать их количество,
например по 1 из каждого инстанса сервиса, и вставлять защитные паузы.

```
func (ms *JournalFast) goUpdateMetrics(aggLog AggLog) {
	backoffTimeout := time.Duration(0)
	for {
		err := ms.updateJournal(ms.ctxParent)
		if err == nil {
			backoffTimeout = 0
			time.Sleep(ms.journalRequestDelay) // if aggregator invariants are broken and they reply immediately forever
			continue
		}
		backoffTimeout = data_model.NextBackoffDuration(backoffTimeout)
		log.Printf("Failed to update metrics, will retry: %v", err)
		time.Sleep(backoffTimeout)
	}
}

func (ms *JournalFast) updateJournal(ctxParent context.Context) error {
    ctx, cancel := context.WithTimeout(ctxParent, ms.journalTimeout)
    defer cancel()
	
	extra := rpc.InvokeReqExtra{FailIfNoConnection: true}
	args := tlstatshouse.GetMetrics3{
		From: ms.getCurrentVersion(),
	}
	var ret tlmetadata.GetJournalResponsenew
	err := ms.getJournalClient().GetMetrics3(ctxParent, args, &extra, &ret)
	if err != nil {
		return fmt.Errorf("cannot load meta journal - %w", err)
	}
	ms.updateEvents(ret.Events, ret.CurrentVersion)
	return nil
}
```

Из интересного здесь также флаг `FailIfNoConnection: true`, который позволяет переключаться на реплики
до таймаута, если соединенеи до еплики разорвано (либо не получилось создать первый раз).
