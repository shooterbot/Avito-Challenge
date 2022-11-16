# Avito-Challenge
Deadline 17.11.22 10:24  
Readme в процессе создания (будет готов 16.11 до 22:00)

## Запуск
В корневой папке выполнить команду  
```docker-compose build```   
После завершения ее работы выполнить  
```docker-compose up```   

Команда будет выполняться не менее 10 секунд. После удачного ее выполнения будет выведена строчка  
```Balance system server is running on 0.0.0.0:31337```   
Если ее нет, вероятно, контейнер с PostgreSQL не успел инициализировать БД и таблицы за отведенные 10 секунд. Попробуйте повторить ```docker-compose up```.   
К приложению можно обратиться по адресу 127.0.0.1:31337/api/v1/.

## Маршруты API
См. [документацию на swagger](https://github.com/shooterbot/Avito-Challenge/blob/main/swagger.yaml)  
Ссылка для удобства: https://editor.swagger.io/

## Пример работы
БД пустая. Пытаемся получить счет пользователя 1:  
![image](https://user-images.githubusercontent.com/73233230/202247016-45736152-c22c-4061-9685-34349cc3ce19.png)  

Получили 0, тк деньги пользователю еще не начислялись. Добавим ему 100 у.е.  
![image](https://user-images.githubusercontent.com/73233230/202247456-db71e4b1-a091-49b4-991f-4c15b3c606d7.png)  

Снова проверяем баланс:  
![image](https://user-images.githubusercontent.com/73233230/202247863-db4eb93e-5eaa-46e9-94e6-37e22999be40.png)  

Теперь у пользователя есть 100 у.е. Допустим, он попытался купить услугу, на которую денег не хватает:  
![image](https://user-images.githubusercontent.com/73233230/202248117-915b510f-c7c1-4cdf-a92d-48d963128a9c.png)  

В услуге отказано. Снова проверим счет:  
![image](https://user-images.githubusercontent.com/73233230/202248409-6ca7711e-6f34-4db3-8111-75b9c4de3eaa.png)  

Видим, что деньги не списались. Пробуем купить что-то по-дешевле:  
![image](https://user-images.githubusercontent.com/73233230/202248560-2adaf659-3fc5-4eae-9218-710c83ad7b21.png)  

В этот раз все получилось. Опять проверяем счет:  
![image](https://user-images.githubusercontent.com/73233230/202248667-749ad780-6e2a-4e34-ac96-96957bd1d609.png)  

Видим, что деньги списались. Но, допустим, услугу решили отменить:  
![image](https://user-images.githubusercontent.com/73233230/202248879-64237c70-f684-48af-bc2d-0cd81ec664cd.png)  

Проверим, что деньги вернулись на счет:  
![image](https://user-images.githubusercontent.com/73233230/202248993-8802c8d4-42d4-46ee-be9e-6033e9fc500c.png)  

Повторим покупку, но в этот раз подтвердим ее:  
![image](https://user-images.githubusercontent.com/73233230/202249119-c80d6b48-a6a7-4508-a2c0-6f036655a2eb.png)  
![image](https://user-images.githubusercontent.com/73233230/202249171-3c47ee31-caba-435e-907a-77f6d85d7663.png)  

Составим отчет для бухгалтерии:  
![image](https://user-images.githubusercontent.com/73233230/202249479-e05d4795-bad5-4e28-b3cb-fe186d0fc397.png)

В указанном файле содержится наша прибыль: по услуге с id 1 полчили 55.5 у.е. А больше ничего и не получали.
![image](https://user-images.githubusercontent.com/73233230/202249795-3f07540c-7f21-4cb7-a80e-c020db646853.png)

Напомню, что сейчас на счету пользоваля 44.50. Переведем 4.5 у.е. пользователю 2:  
![image](https://user-images.githubusercontent.com/73233230/202251426-4108f925-b130-4b34-83d4-59c251e8402f.png)  

Проверяем балансы пользователей:  
![image](https://user-images.githubusercontent.com/73233230/202251546-16857d2f-3cf0-4708-8d8f-97be89c599dc.png)  
![image](https://user-images.githubusercontent.com/73233230/202251589-a609ecbd-af1c-4099-be27-56bbf55dbb8e.png)  

Напоследок посмотрим выписку по транзакциям пользователя 1:  
![image](https://user-images.githubusercontent.com/73233230/202254999-8158761a-26b1-4449-9460-cafd458872d9.png)  
Получчаем следующий список:
```
[
    {
        "UserId": 1,
        "Other": "visa",
        "Reason": "Пополнение счета",
        "Date": "16.11.2022",
        "Amount": 100
    },
    {
        "UserId": 1,
        "Other": "User reservation bill",
        "Reason": "Reserved for a service",
        "Date": "16.11.2022",
        "Amount": -55.5
    },
    {
        "UserId": 1,
        "Other": "User reservation bill",
        "Reason": "Reservation has been canceled",
        "Date": "16.11.2022",
        "Amount": 55.5
    },
    {
        "UserId": 1,
        "Other": "User reservation bill",
        "Reason": "Reserved for a service",
        "Date": "16.11.2022",
        "Amount": -55.5
    },
    {
        "UserId": 1,
        "Other": "Another user",
        "Reason": "because",
        "Date": "16.11.2022",
        "Amount": -4.5
    }
]
```
