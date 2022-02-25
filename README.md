# Аналитика над данными Электронного Университета и ВКОНТАКТЕ

Тема
---
Аналитика над данными ЭУ: попытка построить портрет неуспевающего студента (выбрать тематики групп из профилей из социальных сетей для студентов с низкими средними баллами).

Операционализация задачи:
---
Найти в социальной сети ВКонтакте студентов хотя бы с одним незачетом или одной неудовлетворительной оценкой за последние 2 семестра по данным ЭУ,  вывести статистику самых популярных категорий подписок на сообщества, визуализировать с помощью диаграмм. Представить данные на сайте.

Назначение
---
Для разработки образовательных программ и методологий для более масштабного проекта кафедры (образовательная интерактивная система ОН Каткова)

Целевая аудитория
---
Составители образовательных программ.

Описание
---
Выявляем, в рамках проекта, "ключевые точки" интересов неуспевающих студентов для дальнейшего их внедрения в процесс (давать проекты, рефераты, ЛР на актуальные для студентов темы которые интересны, на лекциях строить примеры на основе понятных и знакомых для обучающихся образах и тд). 
Они получают от нашей системы эти данные, а потом используют в своей работе, как конкретно - нас уже не касается. Мы предоставляем доступ к системе, которая автоматизирует процесс анализа интересов учащихся.
Ограничиваемся неуспевающими студентами, так как общая образовательная система предназначена именно для них. Под понятием “неуспевающий” понимаются студенты, имеющие хотя бы один незачет или неудовлетворительную оценку за последние 2 семестра (на момент анализа).

Цель посещения сайта
---
Информативно-научная, для ознакомления с визуализированными данными аналитики. Рекомендации для адаптации образовательных программ путем предоставления примерных направлений тематик исследований.

Используемые технологии:
---
Аналитика данных Электронного Университета и ВКОНТАКТЕ, получаемых с помощью парсинга, и предоставлением их через rest api и визуализацией. 
* Parsing - Python
* Rest API - Golang
* Front End - HTML/CSS/JavaScript
* Data Storage - MySQL

Тип результата
---
[Single Page Application](https://en.wikipedia.org/wiki/Single-page_application)


Используемые источники
---
* [Электронный университет](http://eun.bmstu.ru/products/progress/current/)
* [ВКонтакте](https://vk.com/)
* [Python](https://python-docx.readthedocs.io/en/latest/)
* [Golang](https://go.dev/doc/)
* [MySQL](https://dev.mysql.com/doc/)