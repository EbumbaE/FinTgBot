Телеграм бот для записи финансовых трат

Команды:\
**/setNote** дата категория сумма\
пример: **/setNote 27.09.2022 food 453.12**\
добавляет трату в заданный день по заданной категории, отвечает **Done** в случае успешной записи

**/getStatistic** период (week, month или year)\
example: **/getStatistic week**\
выводит статистику за заданный период, ответа на команду:\
**Statistic for the week:\
food: 245.12\
school: 85.01**

**/selectReportCurrency**\
дает выбор валюты для команд getStatistic, setNote, getMonthlyBudget

**/setBudget** дата сумма валюта\
example: **/setBudget 10.2022 254 USD**\
устанавливает бюджет на месяц

**/getBudget** дата\
example: **/getBudget 10.2022**\
выводит расход за месяц рабочей валюте

**/start**\
просто выводит *hellow*

**/help**\
выводит информацию о командах