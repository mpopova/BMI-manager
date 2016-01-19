# BMI-manager
This is my project for the course "Programming with GO" in FMI. BMI manager - web application that I develop with Go. The application calculate BMI index of current user, show graphic and statistics.

BMI manager е уеб-приложение, разработващо се на езика Go. 

Фукнционалности:

1.Регистрация в системата

2.Всеки регистриран потребител може да пресметне своя индекс на телесна маса. 

Системата пресмята BMI индекса на потребителя и показва графика за въведените BMI след въвеждане на кг, височина, години и 

пол. Чрез графиката потребителя може да следи как се изменя неговия BMI. 

3.Потребителят може да въвежда веднъж на ден BMI, ако го направи повече от 1 път, се запазва последното за деня.

4.За потребителя е налична и страница (My profile), която съдържа данните, с които се е регистрирал.

5.Предоставя се и страница с обща статистика за всички потребители, например колко мъже имат BMI>30 (наднормено тегло).


Инсталация:

1.Клонирате repository-то https://github.com/mpopova/BMI-manager.git .

2.Сваляте използваните пакети с go get. Това са: "fmt", "html/template", "github.com/gorilla/mux", "github.com/gorilla/securecookie", "net/http", "database/sql", "github.com/go-sql-driver/mysql", "log".

3.Инсталирате XAMPP Control Panel и стартирате модулите Apache и MySQL.

4.От конзолата влизате в папката, където сте клонирали проекта и изпълнявате go build server.go

5.Изпълнявате server

6.Достъпвате application-а на http://localhost:8080/ .


Лиценз:

MIT License - open source license.
