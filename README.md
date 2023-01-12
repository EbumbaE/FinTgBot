### Телеграм бот для записи финансовых трат

## Архитектура
https://miro.com/app/board/uXjVPLZy6mg=/

## Функционал

**/start**\
выводит информацию о командах

**/setNote** дата категория сумма\
пример: **/setNote 27.09.2022 food 453.12**\
добавляет трату в заданный день по заданной категории, отвечает в случае успешной записи **Done** или **Over budget by _ RUB**

**/getStatistic** период (week, month или year)\
example: **/getStatistic week**\
выводит статистику за заданный период, ответа на команду:\
**Statistic for the week in RUB::\
food: 245.12\
school: 85.01**

**/setBudget** дата сумма валюта\
example: **/setBudget 10.2022 254 USD**\
устанавливает бюджет на месяц

**/getBudget** дата\
example: **/getBudget 10.2022**\
выводит расход за месяц рабочей валюте **Budget for the month: 175.27/71.08 CNY**

**/selectCurrency**\
дает выбор валюты для команд getStatistic, setNote, getMonthlyBudget, setBudget

## Logs

Graylog: http://127.0.0.1:7555/ (admin/admin)

## Metrics

Prometheus: http://127.0.0.1:9090/
Grafana: http://127.0.0.1:3000/ (admin/admin)

## Tracing

Jaeger: http://127.0.0.1:16686/
