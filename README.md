# BMI-manager
This is my project for the course "Programming with GO" in FMI. BMI manager - web application that I develop with Go. The application calculate BMI index of current user, show graphic and statistics.//English below


BMI manager е уеб-приложение, разработващо се на езика Gо.

Фукнционалности:

1.Регистрация в системата

2.Всеки регистриран потребител може да пресметне своя индекс на телесна маса. 

Системата пресмята BMI индекса на потребителя и показва графика за въведените BMI след въвеждане на кг, височина. Чрез графиката потребителя може да следи как се изменя неговия BMI. 

3.Потребителят може да въвежда веднъж на ден BMI, ако го направи повече от 1 път, се запазва последното за деня.

4.За потребителя е налична и страница (My profile), която съдържа данните, с които се е регистрирал.

5.Предоставя се и страница с обща статистика за всички потребители, която показва графика на средния BMI на всички регистрирани мъже и жени.

Системата се държи адекватно като изкарва необходимите съобщения за това. Например, при поискване на собствена статистика, а потребителя няма въведен нито един BMI, системата уведомява за  това като изкарва съобщение.

При невъведени потребителско име и парола, системата казва, че тези полета са задължителни.

Не се позволява на нерегистрирани потребители да достъпват страници, различни от index page и страницата с регистрация.

Потребителския интерфейс е интуитивен и лесен за използване.

Приложението е изцяло responsivness.

Инсталация:

1.Клонирате repository-то https://github.com/mpopova/BMI-manager.git .

2.Сваляте използваните пакети с go get. Това са: "github.com/gorilla/mux", "github.com/gorilla/securecookie", "github.com/go-sql-driver/mysql".

3.Инсталирате XAMPP Control Panel и стартирате модулите Apache и MySQL.

4.От конзолата влизате в папката, където сте клонирали проекта и изпълнявате go build server.go

5.Изпълнявате server

6.Достъпвате application-а на http://localhost:8080/ .


Лиценз:

MIT License - open source license.

BMI manager is web-application, written in Go.

//English
Functional requirements:

1.Registration

2.Registered user can calculate BMI.

The application calculate BMI of the user and show graphic after kilos and height had entered. The user can follow the changes of his BMI.

3.The user can save one BMI of day. If is saved multiple times, the system keep the latest record of the day.

4.The system support My profile page that contains all the data of user.

5.The system support page with average statistic - how is average BMI of males and females and show graphic for it.


Installation instructions:

1.Clone the repository https://github.com/mpopova/BMI-manager.git .

2.Download used packages with the command "go get". Тhey are: "github.com/gorilla/mux", "github.com/gorilla/securecookie", "github.com/go-sql-driver/mysql".

3.Install XAMPP Control Panel and start Apache and MySQL modules.

4.From cmd we can stay on the folder with the cloned project and run the following:
go build server.go

5.run:
server

6.Now, you can access the application here ->  http://localhost:8080/ .


Licence:

MIT License - open source license.

