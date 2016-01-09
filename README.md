# BMI-manager
This is my project for the course "Programming with GO" in FMI. BMI manager - web application that I develop with Go. The application calculate BMI index of current user, show graphic and statistics.

BMI manager е уеб-приложение, разработващо се на езика Go. 

С BMI manager всеки регистриран потребител може да пресметне своя индекс на телесна маса чрез разработения калкулатор. 

Системата показва графика и потребителя може да следи как се изменя неговия BMI. 

За потребителя е налична и страница (My profile), която съдържа данните, с които се е регистрирал.

Предоставя се и страница с обща статистика за BMI-я на всички регистрирани потребители.

Инсталация:

1.Клонирате repository-то https://github.com/mpopova/BMI-manager.git .

2.Сваляте използваните пакети с go get. Това са: "fmt", "html/template", "github.com/gorilla/mux", "github.com/gorilla/securecookie", "net/http", "database/sql", "github.com/go-sql-driver/mysql", "log".

3.Инсталирате XAMPP Control Panel и стартирате модулите Apache и MySQL.

4.От конзолата влизате в папката, където сте клонирали проекта и изпълнявате go build server.go

5.Изпълнявате server.go

6.Достъпвате application-а на http://localhost:8080/ .


Лиценз:

MIT License - open source license.
